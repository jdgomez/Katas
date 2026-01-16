package main

import (
	"os"
	"time"

	"github.com/jdgomez/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
