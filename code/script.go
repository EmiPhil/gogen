package code

import (
	"strings"
	"sync"
)

// Script represents a collection of code Blocks
type Script struct {
	// Line settings
	LineTerminator       string // Like ';'
	CommentString        string // Like '//'
	IndentationCharacter string // Like '\t'

	// Block settings
	IndentationFactor uint
	Brackets          [2]string // Like ["{", "}"]

	// Script settings
	EndWithNewLine bool
	ReturnChar     string // Like '\n'

	// Protects Blocks
	sync.Mutex
	Code []Code
}

// Normalize removes all whitespace and line terminators from the string
func (s *Script) Normalize(in string) (out string) {
	out = strings.TrimSpace(in)
	out = strings.TrimSuffix(out, s.LineTerminator)
	return out
}

// AsComment renders the given string like a Comment
func (s *Script) AsComment(in string) (out string) {
	prefix := strings.TrimSpace(s.CommentString)

	out = s.Normalize(in)
	if !strings.HasPrefix(out, prefix) {
		out = prefix + " " + out
	}

	return out
}

// Line creates a new line associated to this script
func (s *Script) Line(kind LineKind) *Line {
	s.Lock()
	defer s.Unlock()

	line := makeLine(s, kind, 0)
	s.Code = append(s.Code, line)
	return line
}

// Block creates a new block associated to this script.
func (s *Script) Block() *Block {
	s.Lock()
	defer s.Unlock()

	block := makeBlock(s, 0)
	s.Code = append(s.Code, block)
	return block
}

// Render outputs the entire script as a string ready to be written to a file
func (s *Script) Render() (out string) {
	s.Lock()
	defer s.Unlock()

	for _, code := range s.Code {
		out += code.Render()
	}

	out = strings.TrimSpace(out)
	if s.EndWithNewLine {
		out += s.ReturnChar
	}

	return out
}

// Script creates a new script based on the config of the current script
func (s *Script) Script() *Script {
	return MakeScript(
		s.LineTerminator,
		s.CommentString,
		s.IndentationCharacter,
		s.IndentationFactor,
		s.Brackets,
		s.EndWithNewLine,
		s.ReturnChar)
}

// MakeScript creates a new script based on an extensive list of config
func MakeScript(
	lineTerminator, commentString, indentationCharacter string,
	indentationLevel uint, brackets [2]string,
	endWithNewLine bool, returnCharacter string) *Script {
	script := new(Script)
	script.LineTerminator = lineTerminator
	script.CommentString = commentString
	script.IndentationCharacter = indentationCharacter
	script.IndentationFactor = indentationLevel
	script.Brackets = brackets
	script.EndWithNewLine = endWithNewLine
	script.ReturnChar = returnCharacter
	script.Code = []Code{}
	return script
}
