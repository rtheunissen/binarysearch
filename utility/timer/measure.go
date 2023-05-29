package timer

import "time"

func Measure(action func ()) time.Duration {
   start := time.Now()
   action()
   return time.Since(start)
}