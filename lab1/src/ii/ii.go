package main

import (
	"fmt"
	"mapreduce"
	"os"
	"sort"
	"strconv"
	"strings"
)

// The mapping function is called once for each piece of the input.
// In this framework, the key is the name of the file that is being processed,
// and the value is the file's contents. The return value should be a slice of
// key/value pairs, each represented by a mapreduce.KeyValue.
func mapF(document string, value string) (res []mapreduce.KeyValue) {
	// TODO: you should complete this to do the inverted index challenge
	words := strings.FieldsFunc(value, contentSplit)
	for _, word := range words {
		res = append(res, mapreduce.KeyValue{word, document})
	}
	return
}

func contentSplit(r rune) bool {
	if r >= rune('A') && r <= rune('Z') ||
		r >= rune('a') && r <= rune('z') ||
		r >= rune('0') && r <= rune('9') {
		return false
	}
	return true
}

// The reduce function is called once for each key generated by Map, with a
// list of that key's string value (merged across all inputs). The return value
// should be a single output value for that key.
func reduceF(key string, values []string) string {
	// TODO: you should complete this to do the inverted index challenge
	uniqMap := make(map[string]bool, 128)
	for _, val := range values {
		uniqMap[val] = true
	}
	var uniqVal []string
	for val, _ := range uniqMap {
		uniqVal = append(uniqVal, val)
	}
	sort.Strings(uniqVal)
	return strconv.Itoa(len(uniqVal)) + " " + strings.Join(uniqVal, ",")
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master sequential x1.txt .. xN.txt)
// 2) Master (e.g., go run wc.go master localhost:7777 x1.txt .. xN.txt)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	if len(os.Args) < 4 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	} else if os.Args[1] == "master" {
		var mr *mapreduce.Master
		if os.Args[2] == "sequential" {
			mr = mapreduce.Sequential("iiseq", os.Args[3:], 3, mapF, reduceF)
		} else {
			mr = mapreduce.Distributed("iiseq", os.Args[3:], 3, os.Args[2])
		}
		mr.Wait()
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], mapF, reduceF, 100)
	}
}
