package internal

// Operation describes a function placeholder
type Operation struct {
	Template      SliceTemplate
	ByRef         bool
	Copy          bool
	Tests         bool
	Name          string
	PackageName   string // added on at template execute time
	GenDate       string // added on at template execute time
	SourceStruct  string // added on at template execute time
	PO            bool   // added on at template execute time
	SliceTypeName string // added on at template execute time
}
