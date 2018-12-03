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

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")

	var FREQ int64
	FREQ = 0

	for _, line := range lines {
		line, err := strconv.ParseInt(line, 10, 64)
		check(err)
		FREQ = FREQ + line
	}
	fmt.Println("Final frequency after calibration is:", FREQ)
}