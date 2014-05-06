package main

/*
 Samples a random number generator 100 times a second and outputs
 JSON specifying the count of each number, for the purposes of
 demonstration.
 */

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	metric := make(map[string]interface{})
	for {
		var counts [6]int

		for i := 0; i < 100; i++ {
			k := rand.Intn(6)
			counts[k] = counts[k] + 1
		}

		for i, c := range counts {
			metric["service"] = fmt.Sprintf("graph-%d", i+1)
			metric["value"] = c
			metric["time"] = time.Now().Unix()
			b, _ := json.Marshal(metric)
			fmt.Println(string(b))
		}
		time.Sleep(1 * time.Second)
	}
}