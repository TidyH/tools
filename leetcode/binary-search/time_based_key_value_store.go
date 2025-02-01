package main

import "fmt"

// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.
type TimeMap struct {
	timeline map[string][]TimestampValuePair // Time stamp is unique, value is not, original key is also unique
}

type TimestampValuePair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	TimeMap := TimeMap{
		timeline: make(map[string][]TimestampValuePair),
	}

	return TimeMap
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pair := TimestampValuePair{
		timestamp: timestamp,
		value:     value,
	}

	this.timeline[key] = append(this.timeline[key], pair)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := this.timeline[key]

	if !ok {
		return ""
	}

	left, right := 0, len(pairs)-1
	result := ""

	for left <= right {
		mid := left + (right-left)/2

		if timestamp >= pairs[mid].timestamp {
			result = pairs[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	fmt.Printf("981. Time Based Key-Value Store\n")

	TimeMap := Constructor()
	TimeMap.Set("foo", "bar", 1)
	fmt.Printf(TimeMap.Get("foo", 1))
}
