package golang

import (
	"github.com/EmiPhil/gogen/code"
	"strings"
)

func WriteComment(container code.Container, multilineComment string) {
	WriteComments(container, strings.Split(multilineComment, "\n"))
}

func WriteComments(container code.Container, comments []string) {
	for _, comment := range comments {
		container.Line(code.Comment).SetValue(comment)
	}
}
