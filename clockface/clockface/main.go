package main

import (
	"os"
	"time"

	"github.com/cabralfbenja/go-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
