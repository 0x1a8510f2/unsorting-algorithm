package main

import (
	"math/rand"
	"time"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Init PRNG
func init() { rand.Seed(time.Now().UnixNano()) }

// Unsorting algorithm
func Unsort(list []string) []string {
	for range list {
		swapindex1, swapindex2 := rand.Intn(len(list)), rand.Intn(len(list))
		list[swapindex1], list[swapindex2] = list[swapindex2], list[swapindex1]
	}
	return list
}

// Better unsorting algorithm
func BetterUnsort(list []string) []string {
	for i := len(list)-1; i >= 1; i-- {
		swapindex := rand.Intn(i+1)
		list[swapindex], list[i] = list[i], list[swapindex]
	}
	return list
}

// Unsort wrapper for demo purposes
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a space-delimeted list: ")
	scanner.Scan()
	liststr := scanner.Text()
	list := strings.Split(liststr, " ")
	fmt.Printf("%v\n", Unsort(list))
}
