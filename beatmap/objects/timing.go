package objects

import (
	"chain_counter/util"
	"strconv"
	"strings"
)

type TimingPoint struct {
	Time       int64
	BeatLength float64
}

func NewTimingPoint(timingSection string) TimingPoint {
	args := strings.Split(timingSection, ",")

	time, err := strconv.ParseInt(args[0], 0, 64)
	util.CheckError(err)

	beatLength, err := strconv.ParseFloat(args[1], 64)
	util.CheckError(err)

	return TimingPoint{Time: time, BeatLength: beatLength}
}

func IsRedLine(timingSection string) bool {
	args := strings.Split(timingSection, ",")
	return args[6] == "1"
}
