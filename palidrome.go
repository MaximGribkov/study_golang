package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.Replace(text, " ", "", -1)
	rs := []rune(text)
	count := 0
	for j := range rs {
		if rs[j] == rs[len(rs)-j-1] {
			count++
		} else {
			count--
		}
	}
	if count == len(rs) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
