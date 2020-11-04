package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	entry, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}
	reverse, err := reverseMessage(entry)
	if err != nil {
		panic(err)
	}
	fmt.Println(reverse)
}

func reverseMessage(s string) (string, error) {
	rs := []rune(s)
	lastPos := len(rs) - 1
	for pos := 0; pos <= lastPos/2; pos++ {
		rs[pos], rs[lastPos-pos] = rs[lastPos-pos], rs[pos]
	}
	return string(rs), nil
}
