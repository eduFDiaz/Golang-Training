package main

import (
	"fmt"
	"io"
	"strings"
)

// Capper is a capepr
type Capper struct {
	wrt io.Writer
}

// Write
func (c *Capper) Write(p []byte) (n int, err error) {
	for _, char := range p {
		c.wrt.Write(strings.ToUpper(char))
	}
	integer := len(p)
	return integer, nil
}

func main() {
	fmt.Printf("Something")
}
