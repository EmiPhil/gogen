package golang

import (
	"github.com/EmiPhil/gogen"
)

// Script represents a go file
type Script struct {
	PackageName string
	*gogen.Script
}

func New(packageName string) *Script {
	script := new(Script)
	script.PackageName = packageName
	script.Script = gogen.MakeGoScript()

	script.WriteHeader()
	return script
}
