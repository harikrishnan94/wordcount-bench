package main

import (
	"fmt"
	"os"
	"time"
)

func isWhiteSpace(c rune) bool {
	return c == '\r' || c == '\n' || c == ' ' || c == '\t'
}

func WordCount(str string) {
	wc := 0
	is_prev_white_space := false

	for i := 0; i < len(str); i++ {
		is_cur_white_space := isWhiteSpace(rune(str[i]))

		if is_cur_white_space && !is_prev_white_space {
			wc += 1
		}

		is_prev_white_space = is_cur_white_space
	}

	if len(str) != 0 && !is_prev_white_space {
		wc += 1
	}

	fmt.Printf("Num Words = %d\n", wc)
}

func main() {
	start := time.Now()

	b, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	WordCount(str)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed = %d ms\n", elapsed.Milliseconds())
}
