package reporter

type Interface interface {
	Report(line int, col int, e error)
}
