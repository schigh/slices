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
)

var (
	targetPackage string
	targetStruct  string
	targetFile    string
)

func main() {
	// make sure args are ok
	if len(os.Args) < 2 {
		internal.PrintErr("usage: slices <struct> ...<funcs>")
		os.Exit(1)
	}

	// make sure env vars exist
	targetPackage = os.Getenv("GOPACKAGE")
	targetFile = os.Getenv("GOFILE")

	if targetPackage == "" || targetFile == "" {
		internal.PrintErr("both GOPACKAGE and GOFILE environment variables are required")
	}

	targetStruct = os.Args[1]
	// test for pointer
	var pointerOverride bool
	if targetStruct[0] == '*' {
		pointerOverride = true
		targetStruct = targetStruct[1:]
	}

	// assemble file path
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		internal.PrintErr("unable to get pwd: %v", pwdErr)
		os.Exit(1)
	}
	targetFile = path.Join(pwd, targetFile)

	// inform user what is being created
	var pointerStr string
	if pointerOverride {
		pointerStr = "pointer "
	}
	internal.PrintInfo("generating %sslices for struct %s in %s", pointerStr, targetStruct, path.Base(targetFile))

	// make sure a struct by the provided name lives in this file
	found, parseErr := parseFileAndLocateStruct(targetFile, targetStruct)
	if parseErr != nil {
		internal.PrintErr("error parsing file %s: %v", targetFile, parseErr)
		os.Exit(1)
	}

	if !found {
		internal.PrintErr("unable to locate struct %s in file %s", targetStruct, targetFile)
		os.Exit(1)
	}

	// determine scope
	scope := slices.StringSlice(os.Args[2:]).Unique().Value()
	flags := strings.Join(scope, " ")
	operations, opsErr := internal.OperationsFromFlags(flags)
	if opsErr != nil {
		internal.PrintErr(opsErr.Error())
		os.Exit(1)
	}

	// add the header operation
	header := []internal.Operation{
		{
			Template: internal.HeadTmpl,
			Name:     "head",
		},
	}
	operations = append(header, operations...)

	t := &internal.Template{
		PO:           pointerOverride,
		PackageName:  targetPackage,
		GenDate:      time.Now().Format(time.RFC1123Z),
		SourceStruct: targetStruct,
		Operations:   operations,
	}

	fileBytes, outErr := generateOutBytes(t, pointerOverride)
	if outErr != nil {
		os.Exit(1)
	}

	// if it's a pointer slice, create a filename that indicates a pointer
	if pointerOverride {
		targetStruct = fmt.Sprintf("%s_ptr", targetStruct)
	}
	outfile := path.Join(path.Dir(targetFile), fmt.Sprintf("%s_slices.go", strings.ToLower(targetStruct)))
	if writeFileErr := ioutil.WriteFile(outfile, fileBytes, 0644); writeFileErr != nil {
		internal.PrintErr("error writing to file %s: %v", targetFile, writeFileErr)
		os.Exit(1)
	}

	internal.PrintSuccess("wrote file %s", outfile)

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

func generateOutBytes(tmpl *internal.Template, overridePtr bool) ([]byte, error) {
	var b []byte
	bBuff := bytes.NewBuffer(b)

	for _, op := range tmpl.Operations {
		op.SourceStruct = tmpl.SourceStruct
		op.SliceTypeName = tmpl.SourceStruct
		if overridePtr {
			op.SliceTypeName = fmt.Sprintf("%sPtr", op.SliceTypeName)
		}
		op.GenDate = tmpl.GenDate
		op.PackageName = tmpl.PackageName
		op.PO = tmpl.PO
		tm, tmErr := template.New(op.Name + op.SliceTypeName).Parse(op.Template.String())
		if tmErr != nil {
			internal.PrintErr("parsing template failed", tmErr)
			os.Exit(1)
		}
		if execErr := tm.Execute(bBuff, &op); execErr != nil {
			internal.PrintErr("executing template failed")
		}
	}

	return bBuff.Bytes(), nil
}
