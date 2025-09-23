package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	path := flag.String("file", "", "Path to file")
	flag.Parse()

	if *path == "" {
		fmt.Println("Usage: go run dars3.go -file=sample.txt")
		os.Exit(1)
	}

	f, err := os.Open(*path)
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines, words, bytes, runes := 0, 0, 0, 0
	
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		bytes += len(line) + 1 // '\n' ni ham hisoblash
		runes += utf8.RuneCountInString(line)
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Skanner xatolik:", err)
	}

	fmt.Printf("Lines: %d\nWords: %d\nBytes(approx): %d\nRunes: %d\n", lines, words, bytes, runes)
}
