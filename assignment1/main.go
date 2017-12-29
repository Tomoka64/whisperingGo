package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	s := bufio.NewScanner(strings.NewReader(text))
	for s.Scan() {
		fmt.Printf("{body: %v}\n", s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

//
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter text: ")
// 	input, _ := reader.ReadString('\n')
// 	s := bufio.NewScanner(strings.NewReader(input))
// 	fmt.Println(s.Text())
// }
