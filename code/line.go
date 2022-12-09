package code

import (
	"github.com/EmiPhil/gogen/texts"
	"sync"
)

// lineState helps us render lines correctly depending on their status
type lineState uint8

const (
	user lineState = iota
	normal
	canonical
)

// LineKind helps distinguish between the kinds of lines that may be in a program
type LineKind uint8

const (
	Standard LineKind = iota
	Comment
	Raw
)

func makeLine(script *Script, kind LineKind, indentationLevel uint) *Line {
	line := new(Line)
	line.script = script
	line.kind = kind
	line.state = user
	line.IndentationLevel = indentationLevel
	line.Value = ""
	line.LineComment = ""

	return line
}

// Line is a single expression or statement
type Line struct {
	// script provides important settings for lines
	script *Script

	kind LineKind

	state lineState

	IndentationLevel uint

	sync.Mutex
	Value       string
	LineComment string
}

// SetValue is the thread safe way to set the line's value
func (l *Line) SetValue(value string) {
	l.Lock()
	defer l.Unlock()
	l.Value = value
}

// SetComment is the thread safe way to set the line's comment
func (l *Line) SetComment(value string) {
	l.Lock()
	defer l.Unlock()
	l.LineComment = value
}

// IsComment checks to see if this line is intended to be a Comment
func (l *Line) IsComment() bool {
	return l.kind == Comment
}

// IsRaw checks to see if this line is intended to be a Raw string
func (l *Line) IsRaw() bool {
	return l.kind == Raw
}

// IsNormal checks to see if the line has been normalized (whitespace & terminators
// removed)
func (l *Line) IsNormal() bool {
	return l.state == normal
}

// Normalize removes all whitespace and line terminators from the code. Lock before
// calling!
func (l *Line) Normalize() {
	if l.IsNormal() || l.IsRaw() {
		return
	}

	l.Value = l.script.Normalize(l.Value)
	l.state = normal
}

// IsCanonical checks to see if the line has been canonicalized (whitespace & terminators
// added)
func (l *Line) IsCanonical() bool {
	return l.state == canonical
}

// Indent applies an appropriate amount of indent prefixing to the line. Lock
// before calling!
func (l *Line) Indent() {
	l.Value = texts.Prefix(l.Value, l.IndentationLevel, l.script.IndentationCharacter)
}

// AddReturn adds the return char. Lock before calling!
func (l *Line) AddReturn() {
	l.Value += l.script.ReturnChar
}

// Canonicalize adds whitespace and line terminators to the code
func (l *Line) Canonicalize() {
	l.Lock()
	defer l.Unlock()

	defer l.Indent()
	defer l.AddReturn()
	if l.IsCanonical() || l.IsRaw() {
		return
	}

	l.Normalize()
	result := l.Value

	if l.IsComment() {
		result = l.script.AsComment(result)
	} else {
		// we don't need to add line terminators to comments
		result += l.script.LineTerminator
	}

	if len(l.LineComment) > 0 {
		result += l.script.AsComment(l.LineComment)
	}

	l.Value = result
	l.state = canonical
}

// Render returns a canonical string version of the line
func (l *Line) Render() string {
	l.Canonicalize()
	return l.Value
}

// String provided as a convenience. Equal to calling Render
func (l *Line) String() string {
	return l.Render()
}
