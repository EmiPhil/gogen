package code

// Code represents a script, block, or line
type Code interface {
	Render() string
}

type Container interface {
	Line(kind LineKind) *Line
}
