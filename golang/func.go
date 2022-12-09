package golang

import (
	"fmt"
	"github.com/EmiPhil/gogen"
	"github.com/EmiPhil/gogen/texts"
)

// Params represents parameters to be sent to Funcs
type Params []string

func (p Params) String() string {
	return texts.Bracket(texts.RoundBrackets(), p...)
}

// Returns represents a list of return values to be sent to Funcs
type Returns []string

func (r Returns) String() string {
	switch len(r) {
	case 0:
		return ""
	case 1:
		return r[0]
	default:
		return texts.Bracket(texts.RoundBrackets(), r...)
	}
}

type Func struct {
	Name    string
	Params  Params
	Returns Returns

	*gogen.Block
}

func (fn *Func) SetLeftBracketPrefix() {
	fn.Block.SetLeftBracketPrefix(fmt.Sprintf("func %s%s %s", fn.Name, fn.Params, fn.Returns))
}

func (s *Script) WriteFunc(descriptor, name string, params []string, returns []string) Func {
	WriteComment(s, descriptor)

	fn := Func{}
	fn.Name = name
	fn.Params = params
	fn.Returns = returns

	fn.Block = s.Block()
	fn.SetLeftBracketPrefix()

	return fn
}
