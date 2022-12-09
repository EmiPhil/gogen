package golang

import "github.com/EmiPhil/gogen/code"

// Script represents a go file
type Script struct {
	PackageName string
	*code.Script
}

func New(packageName string) *Script {
	script := new(Script)
	script.PackageName = packageName
	script.Script = code.MakeGoScript()

	script.WriteHeader()
	return script
}
