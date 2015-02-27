package helpers

import
(
	"time"
	"log"
)

func Trace(s string) (string, time.Time) {
    log.Println("START:", s)
    return s, time.Now()
}

func Un(s string, startTime time.Time) {
    endTime := time.Now()
    log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}