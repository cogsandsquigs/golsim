package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.oxyl")
	if err != nil {
		fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}
