package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	inPath, outPath, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	in, err := os.Open(inPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open input:", err)
		os.Exit(1)
	}
	defer in.Close()

	raw, err := io.ReadAll(bufio.NewReader(in))
	if err != nil {
		fmt.Fprintln(os.Stderr, "read input:", err)
		os.Exit(1)
	}

	out := transform(string(raw))

	if outPath == "-" {
		fmt.Print(out)
		return
	}

	if err := os.WriteFile(outPath, []byte(out), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "write output:", err)
		os.Exit(1)
	}
}

func parseArgs(args []string) (string, string, error) {
	fs := flag.NewFlagSet("go-reloaded", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		return "", "", err
	}
	pos := fs.Args()
	if len(pos) < 2 {
		return "", "", errors.New("usage: go run . <input.txt> <output.txt|->")
	}
	return pos[0], pos[1], nil
}
