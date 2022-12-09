package texts

// Prefix prefixes string y n times to string x
func Prefix(x string, n uint, y string) (out string) {
	out = x
	for i := uint(0); i < n; i++ {
		out = y + out
	}
	return out
}

// Postfix postfixes string y n times to string x
func Postfix(x string, n uint, y string) (out string) {
	out = x
	for i := uint(0); i < n; i++ {
		out = out + y
	}
	return out
}
