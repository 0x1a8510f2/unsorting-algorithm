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

// Basic two swap index unsorting algorithm
func BasicUnsort(list []string) []string {
	for range list {
		swapindex1, swapindex2 := rand.Intn(len(list)), rand.Intn(len(list))
		list[swapindex1], list[swapindex2] = list[swapindex2], list[swapindex1]
	}
	return list
}

// Unsorting algorithm with one swap index
func OneSwapIndexUnsort(list []string) []string {
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
	fmt.Println("Choose an unsorting func:")
	for {
		fmt.Println("0=BasicUnsort, 1=OneSwapIndexUnsort")
		var unsortindex int
		fmt.Scanf("%d", &unsortindex)
		switch unsortindex {
		case 0:
			fmt.Printf("%v\n", BasicUnsort(list))
			return
		case 1:
			fmt.Printf("%v\n", OneSwapIndexUnsort(list))
			return
		default:
			fmt.Println("Please enter a valid index.")
		}
	}
}
