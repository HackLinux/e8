// Package reporter defines compiler error report interface and
// provides several handy interface implementations.
package reporter

type Interface interface {
	Report(line int, col int, e error)
}
