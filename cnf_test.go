package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	verbose bool
	cert    bool
	mus     bool
	count   bool
	help    bool
)

func testCnf(path string) {
	if pb, printFn, err := parse(path); err != nil {
		fmt.Fprintf(os.Stderr, "could not parse problem: %v\n", err)
		os.Exit(1)
	} else {
		solve(pb, verbose, cert, printFn)
	}
}

func TestCnf(t *testing.T) {
	// verbose = true
	for i := 1; i <= 500; i++ {
		ufile := fmt.Sprintf("/Users/jpbirdy/Workspaces/go/src/github.com/DoOR-Team/door/uuf100/uuf100-0%d.cnf", i)
		testCnf(ufile)
	}
}
