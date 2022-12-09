package golang

import (
	"github.com/EmiPhil/gogen/code"
)

func (s *Script) WriteHeader() {
	s.Line(code.Raw).SetValue("package " + s.PackageName)
	s.Line(code.Raw)
}
