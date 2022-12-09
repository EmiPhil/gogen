package golang

import (
	"github.com/EmiPhil/gogen"
	"strings"
)

func WriteComment(container gogen.Container, multilineComment string) {
	WriteComments(container, strings.Split(multilineComment, "\n"))
}

func WriteComments(container gogen.Container, comments []string) {
	for _, comment := range comments {
		container.Line(gogen.Comment).SetValue(comment)
	}
}
