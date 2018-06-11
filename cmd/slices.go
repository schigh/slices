package main

import (
	"fmt"
	"os"
	"path"

	"bytes"
	"github.com/schigh/slices"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"strings"
	"time"
)

const (
	FILTER = "filter"
	MAP    = "map"
)

type Tmpl struct {
	PackageName  string
	GenDate      string
	SourceStruct string
	Inclusions   []string
}

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
	allowed := slices.StringSlice([]string{FILTER, MAP})
	var includes, excludes []string
	for _, s := range scope {
		if len(s) > 0 && s[0] == '-' {
			ss := string(s[1:])
			if !allowed.Contains(ss) {
				printErr("illegal operation '%s' in file %s", ss, targetFile)
				os.Exit(1)
			}
			excludes = append(excludes, ss)
		} else {
			if !allowed.Contains(s) {
				printErr("illegal operation '%s' in file %s", s, targetFile)
				os.Exit(1)
			}
			includes = append(includes, s)
		}
	}

	all := len(includes) == 0
	if !all {
		// if we're not generating all funcs, then ignore any excludes
		excludes = nil
	} else {
		// otherwise, set includes to all allowed funcs
		includes = allowed.Value()
	}

	// filter out all excludes from includes, then sort
	includes = slices.StringSlice(includes).Filter(func(s string) bool {
		return !slices.StringSlice(excludes).Contains(s)
	}).SortAsc().Value()

	t := &Tmpl{
		PackageName:  targetPackage,
		GenDate:      time.Now().Format(time.RFC1123Z),
		SourceStruct: targetStruct,
		Inclusions:   includes,
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

func generateOutBytes(tmpl *Tmpl) ([]byte, error) {
	var b []byte
	bBuff := bytes.NewBuffer(b)

	var tmplErr error
	headTmpl, tmplErr := template.New("head").Parse(HEADTMPL)
	if tmplErr != nil {
		printErr("parsing template failed for HEADTMPL: %v", tmplErr)
		return nil, tmplErr
	}
	if headTmplErr := headTmpl.Execute(bBuff, tmpl); headTmplErr != nil {
		printErr("executing template failed for HEADTMPL: %v", headTmplErr)
		return nil, headTmplErr
	}
	bBuff.Write([]byte(NEWLINE))

	if slices.StringSlice(tmpl.Inclusions).Contains(MAP) {
		mapTmpl, tmplErr := template.New("map").Parse(MAPTMPL)
		if tmplErr != nil {
			printErr("parsing template failed for MAPTMPL: %v", tmplErr)
			return nil, tmplErr
		}
		if mapTmplErr := mapTmpl.Execute(bBuff, tmpl); mapTmplErr != nil {
			printErr("executing template failed for MAPTMPL: %v", mapTmplErr)
			return nil, mapTmplErr
		}
		bBuff.Write([]byte(NEWLINE))
	}

	if slices.StringSlice(tmpl.Inclusions).Contains(FILTER) {
		filterTmpl, tmplErr := template.New("package").Parse(FILTERTMPL)
		if tmplErr != nil {
			printErr("parsing template failed for FILTERTMPL: %v", tmplErr)
			return nil, tmplErr
		}
		if filterTmplErr := filterTmpl.Execute(bBuff, tmpl); filterTmplErr != nil {
			printErr("executing template failed for FILTERTMPL: %v", filterTmplErr)
			return nil, filterTmplErr
		}
		bBuff.Write([]byte(NEWLINE))
	}

	return bBuff.Bytes(), nil
}
