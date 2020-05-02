package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var ln = 1
var n bool

func init() {
	flag.BoolVar(&n, "n", false, "with line number")
}

func main() {
	flag.Parse()
	files := flag.Args()
	for _, fn := range files {
		err := readFile(fn)
		if err != nil {
			return
		}
	}
}

func readFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if n {
			fmt.Printf("%d: %s\n", ln, scanner.Text())
			ln++
		} else {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}
	return nil
}
