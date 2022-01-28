package main

import (
//	"math/rand"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func unsort(list []string) []string {
	newlist := []string{}
	
	return newlist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a space-delimeted list: ")
	scanner.Scan()
	liststr := scanner.Text()
	list := strings.Split(liststr, ",")
	fmt.Printf("%v\n", unsort(list))
}
