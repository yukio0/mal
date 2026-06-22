package main

import (
	"bufio"
	"fmt"
	"os"
)

func READ(str string) string {
	return str
}

func EVAL(ast string, env string) string {
	return ast
}

func PRINT(exp string) string {
	return exp
}

func rep(str string) string {
	return PRINT(EVAL(READ(str), ""))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		fmt.Println(rep(input))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
