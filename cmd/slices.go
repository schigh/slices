package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/schigh/slices"
	"github.com/schigh/slices/cmd/internal"
	"strconv"
)

var (
	targetPackage string
	targetStruct  string
	targetFile    string
)

func printErr(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31mslices: "+msg+"\n", args...)
}

func printSuccess(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[32mslices: "+msg+"\n", args...)
}

func main() {
	// make sure args are ok
	if len(os.Args) < 2 {
		printErr("usage: slices <struct> ...<funcs>")
		os.Exit(1)
	}

	// make sure env vars exist
	targetPackage = os.Getenv("GOPACKAGE")
	targetFile = os.Getenv("GOFILE")

	if targetPackage == "" || targetFile == "" {
		printErr("both GOPACKAGE and GOFILE environment variables are required")
	}
	targetStruct = os.Args[1]

	// assemble file path
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		printErr("unable to get pwd: %v", pwdErr)
		os.Exit(1)
	}
	targetFile = path.Join(pwd, targetFile)

	// make sure a struct by the provided name lives in this file
	found, parseErr := parseFileAndLocateStruct(targetFile, targetStruct)
	if parseErr != nil {
		printErr("error parsing file %s: %v", targetFile, parseErr)
		os.Exit(1)
	}

	if !found {
		printErr("unable to locate struct %s in file %s", targetStruct, targetFile)
		os.Exit(1)
	}

	// determine scope
	scope := slices.StringSlice(os.Args[2:]).Unique().Value()
	flags := strings.Join(scope, " ")
	operations, opsErr := internal.OperationsFromFlags(flags)
	if opsErr != nil {
		printErr(opsErr.Error())
		os.Exit(1)
	}

	// add the header operation
	header := []internal.Operation{
		{
			Template: internal.HeadTmpl,
		},
	}
	operations = append(header, operations...)

	t := &internal.Template{
		PackageName:  targetPackage,
		GenDate:      time.Now().Format(time.RFC1123Z),
		SourceStruct: targetStruct,
		Operations:   operations,
	}

	fileBytes, outErr := generateOutBytes(t)
	if outErr != nil {
		os.Exit(1)
	}

	outfile := path.Join(path.Dir(targetFile), fmt.Sprintf("%s_slices.go", strings.ToLower(targetStruct)))
	if writeFileErr := ioutil.WriteFile(outfile, fileBytes, 0644); writeFileErr != nil {
		printErr("error writing to file %s: %v", targetFile, writeFileErr)
		os.Exit(1)
	}

	printSuccess("wrote file %s", outfile)

}

func parseFileAndLocateStruct(filePath, targetStruct string) (bool, error) {
	fset := token.NewFileSet()
	node, parseErr := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if parseErr != nil {
		return false, parseErr
	}

	var found bool
	var structName string
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			structName = t.Name.String()
		case *ast.StructType:
			if structName == targetStruct {
				found = true
				return false
			}
		}
		return true
	})

	return found, nil
}

func generateOutBytes(tmpl *internal.Template) ([]byte, error) {
	var b []byte
	bBuff := bytes.NewBuffer(b)

	for _, op := range tmpl.Operations {
		op.SourceStruct = tmpl.SourceStruct
		op.GenDate = tmpl.GenDate
		op.PackageName = tmpl.PackageName
		tm, tmErr := template.New(strconv.FormatInt(time.Now().UnixNano(), 10)).Parse(op.Template.String())
		if tmErr != nil {
			printErr("parsing template failed")
			os.Exit(1)
		}
		if execErr := tm.Execute(bBuff, &op); execErr != nil {
			printErr("executing template failed")
		}
	}

	return bBuff.Bytes(), nil
}
