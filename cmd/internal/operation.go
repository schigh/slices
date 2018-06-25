package internal

// Operation describes a function placeholder
type Operation struct {
	Template     SliceTemplate
	ByValue      bool
	Copy         bool
	Tests        bool
	PackageName  string // added on at template execute time
	GenDate      string // added on at template execute time
	SourceStruct string // added on at template execute time
}
