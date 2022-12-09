package texts

import "strings"

// Bracket surrounds some values with brackets and returns them comma-separated
func Bracket(bracket [2]string, values ...string) string {
	return bracket[0] + strings.Join(values, ", ") + bracket[1]
}

func CurlyBrackets() [2]string {
	return [2]string{"{", "}"}
}

func RoundBrackets() [2]string {
	return [2]string{"(", ")"}
}

func SquareBrackets() [2]string {
	return [2]string{"[", "]"}
}

func Quotes() [2]string {
	return [2]string{"\"", "\""}
}

func SingleQuotes() [2]string {
	return [2]string{"'", "'"}
}
