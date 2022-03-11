package main

import (
	"fmt"
	"math"
	"encoding/json"
	"testing"
)

var sorted = []string{"0","1","2","3","4","5","6","7","8","9"}
var counter = map[string]int{
	"0": 0, 
	"1": 0, 
	"2": 0,
	"3": 0, 
	"4": 0, 
	"5": 0, 
	"6": 0, 
	"7": 0, 
	"8": 0, 
	"9": 0,
}

func UnsortTestTemplate (b *testing.B, name string, Unsort func([]string)[]string) {
	for k, _ := range counter {
		counter[k] = 0
	}
	fmt.Println(name)
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < 100000; i++ {
			unsorted := Unsort(sorted)
			counter[unsorted[0]]++
		}
	}
	b.StopTimer()
	leastvalue := ""
	leastcount := math.Inf(1)
	mostvalue := ""
	mostcount := 0
	encoded, _ := json.Marshal(counter)
	fmt.Println(string(encoded))
	for value, count := range counter {
		if float64(count) < leastcount {
			leastcount = float64(count)
			leastvalue = value
		}
		if count > mostcount {
			mostcount = count
			mostvalue = value
		}
	}
	tenpercent := mostcount / 10
	if float64(mostcount - tenpercent) > leastcount {
		b.Errorf("%s (most common @ %d) is more common than %s (least common @ %d) by over 10 percent", mostvalue, mostcount, leastvalue, int(leastcount))
	}
}

func BenchmarkBasicUnsort(b *testing.B) {
	UnsortTestTemplate(b, "Basic unsort", BasicUnsort);
}

func BenchmarkOneSwapIndexUnsort(b *testing.B) {
	UnsortTestTemplate(b, "One swap index unsort", OneSwapIndexUnsort);
}

