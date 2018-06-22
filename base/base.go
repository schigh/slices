package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type typeBucket struct {
	BaseType    string
	PackageType string
	ZeroValue   string
	BaseFile    string
	TestFile    string
}

var buckets []typeBucket

func init() {
	buckets = []typeBucket{
		{"int", "IntSlice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"int8", "Int8Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"int16", "Int16Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"int32", "Int32Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"int64", "Int64Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"uint", "UIntSlice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"uint8", "UInt8Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"uint16", "UInt16Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"uint32", "UInt32Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"uint64", "UInt64Slice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"float64", "Float64Slice", "0", "./tmpl/base.txt", "./tmpl/floattype_test.txt"},
		{"float32", "Float32Slice", "0", "./tmpl/base.txt", "./tmpl/floattype_test.txt"},
		{"uintptr", "UIntPtrSlice", "0", "./tmpl/base.txt", "./tmpl/integertype_test.txt"},
		{"string", "StringSlice", `""`, "./tmpl/base.txt", "./tmpl/stringtype_test.txt"},
	}
}

func main() {
	for _, bucket := range buckets {
		baseFileBytes, readBaseErr := ioutil.ReadFile(bucket.BaseFile)
		if readBaseErr != nil {
			fmt.Fprintf(os.Stderr, "Opening integertype base file failed with error: %v", readBaseErr)
			os.Exit(1)
		}

		testFileBytes, readTestErr := ioutil.ReadFile(bucket.TestFile)
		if readBaseErr != nil {
			fmt.Fprintf(os.Stderr, "Opening integertype test file failed with error: %v", readTestErr)
			os.Exit(1)
		}

		baseTmpl := template.Must(template.New("base").Parse(string(baseFileBytes)))
		testTmpl := template.Must(template.New("test").Parse(string(testFileBytes)))

		var execErr, writeErr error
		var baseBytes, testBytes []byte
		baseBuff := bytes.NewBuffer(baseBytes)
		testBuff := bytes.NewBuffer(testBytes)
		if execErr = baseTmpl.Execute(baseBuff, bucket); execErr != nil {
			fmt.Fprintf(os.Stderr, "executing template failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, execErr)
			os.Exit(1)
		}

		if execErr = testTmpl.Execute(testBuff, bucket); execErr != nil {
			fmt.Fprintf(os.Stderr, "executing template failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, execErr)
			os.Exit(1)
		}

		if writeErr = ioutil.WriteFile(fmt.Sprintf("../%s.go", bucket.BaseType), baseBuff.Bytes(), 0644); writeErr != nil {
			fmt.Fprintf(os.Stderr, "writing base file failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, writeErr)
			os.Exit(1)
		}

		if writeErr = ioutil.WriteFile(fmt.Sprintf("../%s_test.go", bucket.BaseType), testBuff.Bytes(), 0644); writeErr != nil {
			fmt.Fprintf(os.Stderr, "writing test file failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, writeErr)
			os.Exit(1)
		}
	}
}
