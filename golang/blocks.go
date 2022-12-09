package golang

import (
	"github.com/EmiPhil/gogen"
)

func (s *Script) WriteHeader() {
	s.Line(gogen.Raw).SetValue("package " + s.PackageName)
	s.Line(gogen.Raw)
}
