package gogen

import (
	"github.com/EmiPhil/gogen/texts"
	"sync"
)

func makeBlock(script *Script, indentationLevel uint) *Block {
	block := new(Block)
	block.script = script
	block.IndentationLevel = indentationLevel
	block.Code = []Code{}

	return block
}

// Block represents a collection of code
type Block struct {
	// script provides important settings for lines
	script *Script

	// IndentationLevel tracks how indented the current block is
	IndentationLevel uint

	// BracketPrefixes set extra values to add before/after the brackets
	LeftBracketPrefix   string
	RightBracketPostfix string

	// protects Code
	sync.Mutex
	// Code are the statements/expressions in the block, including other blocks
	Code []Code
}

func (b *Block) SetLeftBracketPrefix(prefix string) {
	b.Lock()
	defer b.Unlock()

	b.LeftBracketPrefix = prefix
}

func (b *Block) SetRightBracketPrefix(prefix string) {
	b.Lock()
	defer b.Unlock()

	b.RightBracketPostfix = prefix
}

// Line creates a Line with Comment toggling. We split it like this to make
// it easier for end users to get lines/comments without passing bools around
func (b *Block) Line(kind LineKind) *Line {
	b.Lock()
	defer b.Unlock()

	line := makeLine(b.script, kind, b.IndentationLevel+1)
	b.Code = append(b.Code, line)

	return line
}

// Expr returns a line which is already registered to this block. This is
// the recommended way of making lines
func (b *Block) Expr() *Line {
	return b.Line(Standard)
}

// Comment is like Line but for adding comments
func (b *Block) Comment() *Line {
	return b.Line(Comment)
}

// Raw is like Line but for adding Raw text
func (b *Block) Raw() *Line {
	return b.Line(Raw)
}

// Block creates a new Block with correct indentation applied
func (b *Block) Block() *Block {
	b.Lock()
	defer b.Unlock()

	block := makeBlock(b.script, b.IndentationLevel+1)
	b.Code = append(b.Code, block)

	return block
}

func (b *Block) bracket(right bool) (out string) {
	idx := 0
	if right {
		idx = 1
	}

	out = texts.Prefix(b.script.Brackets[idx], b.IndentationLevel, b.script.IndentationCharacter) + b.script.ReturnChar
	if right {
		out += b.RightBracketPostfix
	} else {
		out = b.LeftBracketPrefix + out
	}

	return out
}

// Render returns a canonical version of the block Code
func (b *Block) Render() (out string) {
	b.Lock()
	defer b.Unlock()

	out = b.bracket(false)

	for _, code := range b.Code {
		out += code.Render()
	}

	out += b.bracket(true)
	return out
}

// String provided as a convenience. Equal to calling Render
func (b *Block) String() string {
	return b.Render()
}
