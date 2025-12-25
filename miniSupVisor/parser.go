package main

import(
	"errors"
	"regexp"
	"strconv"
)

var latencybot = regexp.MustCompile(`time=([\d.]+)`)

func parseLatency(line string) (float64, error){
	// search the line for latency data
	match := latencybot.FindStringSubmatch(line)

	// check if latency data exists
	if len(match) < 2{
		return 0, errors.New("no latency data found")
	}

	//now converting string latency to float data
	latencystr := match[1]

	value, err := strconv.ParseFloat(latencystr, 64)

	if err != nil {
		return 0, err
	}

	return value, err

}