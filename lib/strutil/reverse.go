// Package strutil contains utility functions for working with string.
// Go's convention is that the package name is the last element of the import path.
// There is no requirement that package names be unique across all packages linked into a single binary, only that the import paths (their full file names) be unique.
package strutil

// Reverse function returns its argument string reversed rune-wise left to right.
// In Go, a name is exported if it begins with a capital letter.
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
