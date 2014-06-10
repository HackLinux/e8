package module

import (
	"fmt"
	"os"
	"time"
)

var zeroTime = time.Unix(0, 0)

func fileModTime(path string) (time.Time, error) {
	stat, e := os.Stat(path)
	if e != nil {
		return zeroTime, e
	}

	if !stat.Mode().IsRegular() {
		return zeroTime, fmt.Errorf("%q is not a regular file", path)
	}

	return stat.ModTime(), nil
}
