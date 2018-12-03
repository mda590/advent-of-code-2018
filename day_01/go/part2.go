package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")

	var FREQ int64 = 0
	FREQ_LIST := make([]string, 0)
	FREQ_LIST = append(FREQ_LIST, "0")
	var CHANGE_IND int = 0

	for {
		line, err := strconv.ParseInt(lines[CHANGE_IND], 10, 64)
		check(err)
		FREQ = FREQ + line

		if stringInSlice(string(FREQ), FREQ_LIST) {
			fmt.Println("The first frequency to appear twice is: ", FREQ)
			break
		}

		// Convert FREQ back to string to add to slice
		fmt.Println(FREQ)
		FREQ := strconv.FormatInt(FREQ, 16)
		fmt.Println(FREQ)
		FREQ_LIST = append(FREQ_LIST, FREQ)
		CHANGE_IND = CHANGE_IND + 1

		if CHANGE_IND == len(FREQ_LIST) {
			CHANGE_IND = 0
		}
	}
}