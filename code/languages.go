package code

import "github.com/EmiPhil/gogen/texts"

// Here we provide some default script styles for a few choice languages
func MakeGoScript() *Script {
	return MakeScript(
		"", "//", "\t",
		1, texts.CurlyBrackets(),
		true, "\n")
}
