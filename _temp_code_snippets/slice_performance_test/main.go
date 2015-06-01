package main

import (
	"fmt"
	"strconv"
	"testing"
	// "log"
	// "time"
)

var strSlice []string
var intSlice []int

func main() {
	
	for i := 0; i < 50000; i++ {
		iStr := strconv.Itoa(i)
		strSlice = append(strSlice, iStr)
		intSlice = append(intSlice, i)
	}
	brStr := testing.Benchmark(BenchmarkFunctionStrSlice)
	brInt := testing.Benchmark(BenchmarkFunctionIntSlice)
	fmt.Println("brStr")
	fmt.Println(brStr)
	fmt.Println("brInt")
	fmt.Println(brInt)
}

func lookupSliceTargetStr(target string) (b *testing.B){
	// defer un(trace("SOME_ARBITRARY_STRING_SO_YOU_CAN_KEEP_TRACK"))
	//_ = m[key]
	for _, s := range strSlice{
		if s == target{
			return
		}
	}
	return
}

func lookupSliceTargetInt(target int) (b *testing.B){
	// defer un(trace("SOME_ARBITRARY_STRING_SO_YOU_CAN_KEEP_TRACK"))
	//_ = m[key]
	for _, i := range intSlice{
		if i == target{
			return
		}
	}
	return
}



func BenchmarkFunctionStrSlice(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = lookupSliceTargetStr("25000")
    }
}

func BenchmarkFunctionIntSlice(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = lookupSliceTargetInt(25000)
    }
}
