package module

import (
	"io/ioutil"
	"strings"
	"bytes"
	"fmt"
)

func readLines(path string) ([]string, error) {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	lines := strings.Split(string(bytes), "\n")
	ret := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ret = append(ret, line)
	}

	return ret, nil
}

func writeLines(path string, lines []string) error {
	buf := new(bytes.Buffer)
	for _, line := range lines {
		fmt.Fprintln(buf, line)
	}

	return ioutil.WriteFile(path, buf.Bytes(), 0644)
}