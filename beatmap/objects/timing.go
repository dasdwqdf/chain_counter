package objects

import (
	"chain_counter/util"
	"strconv"
	"strings"
)

type TimingPoint struct {
	time       int64
	beatLength float64
}

func NewTimingPoint(timingSection string) TimingPoint {
	args := strings.Split(timingSection, ",")

	time, err := strconv.ParseInt(args[0], 0, 64)
	util.CheckError(err)

	beatLength, err := strconv.ParseFloat(args[1], 64)
	util.CheckError(err)

	return TimingPoint{time: time, beatLength: beatLength}
}

func IsRedLine(timingSection string) bool {
	args := strings.Split(timingSection, ",")
	return args[6] == "1"
}
