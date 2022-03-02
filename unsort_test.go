package main

import (
	"fmt"
	"math"
	"encoding/json"
	"flag"
	"testing"
)

type unsortfuncpair struct {
	name string
	Run func([]string) []string
}

var ufindex = flag.Int("func", 0, "Specify the unsort function to be used by its index. 0=BasicUnsort, 1=OneSwapIndexUnsort")

func TestUnsort(t *testing.T) {
	unsortoptions := []unsortfuncpair {
		{"Basic unsort", BasicUnsort},
		{"One swap index unsort", OneSwapIndexUnsort},
	}
	sorted := []string{"0","1","2","3","4","5","6","7","8","9"}
	counter := map[string]int{
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
	if *ufindex < 0 || *ufindex > 1 {
		t.Errorf("Invalid unsort function index")
		return
	}
	unsortfunc := unsortoptions[*ufindex]
	fmt.Printf("--- %s ---\n", unsortfunc.name)
	for i := 100000; i > 0; i-- {
		unsorted := unsortfunc.Run(sorted)
		counter[unsorted[0]]++
	}
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
		t.Errorf("%s (most common @ %d) is more common than %s (least common @ %d) by over 10 percent", mostvalue, mostcount, leastvalue, int(leastcount))
	}
}
