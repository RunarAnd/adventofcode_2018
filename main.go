package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var Freqs = make(map[int]int)

func getFreqValue(freqString string) (n int) {
	i, err := strconv.Atoi(freqString[1:])
	if err != nil {
		panic(err)
	}
	return i
}

func addFreq(freq int) (count int) {
	i, _ := Freqs[freq]
	Freqs[freq] = i + 1

	return (i + 1)
}

func printHashMap(m map[int]int) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

func main() {
	var frequency = 0
	Freqs[0] = 1
	dat, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	for {

		for _, line := range lines {
			if len(line) > 0 {
				if string(line[0]) == "+" {
					frequency += getFreqValue(line)
				}
				if string(line[0]) == "-" {
					frequency -= getFreqValue(line)
				}
				var nFreq = addFreq(frequency)
				if nFreq > 1 {
					fmt.Println("FOUNT IT: %s", frequency)
					panic(frequency)
				}
				fmt.Println(frequency, line, nFreq)
			}

		}
	}
	printHashMap(Freqs)

}
