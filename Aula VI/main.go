package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := dataEntry("")
	if err != nil {
		panic(err)
	}
	s, err := generateFizzBuzz(n)
	fmt.Println(s)
}

func dataEntry(s string) (int, error) {
	if s == "" {
		fmt.Scanf("%s\n", &s)
	}
	var n int
	n, err := strconv.Atoi(s)
	if err != nil {
		return n, err
	}
	return n, nil
}

func generateFizzBuzz(n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("value must be bigger than one or bigger")
	}
	s := "1"
	for i := 2; i <= n; i++ {
		switch {
		case i%15 == 0:
			s = fmt.Sprintf("%s, %s", s, "Fizzbuzz")
		case i%5 == 0:
			s = fmt.Sprintf("%s, %s", s, "Buzz")
		case i%3 == 0:
			s = fmt.Sprintf("%s, %s", s, "Fizz")
		default:
			s = fmt.Sprintf("%s, %d", s, i)
		}
	}
	return s, nil
}
