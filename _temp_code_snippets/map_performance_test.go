package main

import (
	"fmt"
	"strconv"
	"testing"
	// "log"
	// "time"
)

var m map[string]interface{}

func main() {
	m = make(map[string]interface{})
	for i := 0; i < 50000; i++ {
		iStr := strconv.Itoa(i)
		m[iStr] = fmt.Sprintf("%d",i)
	}
	br := testing.Benchmark(BenchmarkFunction)
	fmt.Println(br)
}

func lookupMapKey(key string, m map[string]interface{}) (b *testing.B){
	// defer un(trace("SOME_ARBITRARY_STRING_SO_YOU_CAN_KEEP_TRACK"))
	//_ = m[key]
	return
}

func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = lookupMapKey("25000", m)
    }
}

// func trace(s string) (string, time.Time) {
//     log.Println("START:", s)
//     return s, time.Now()
// }

// func un(s string, startTime time.Time) {
//     endTime := time.Now()
//     log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
// }