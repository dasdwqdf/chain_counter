package beatmap

import (
	"chain_counter/beatmap/objects"
	"math"
)

type Chain struct {
	snapDivisor uint8
	startTime   int64
	endTime     int64
	size        int
}

func getBeatLengthAt(time int64, timingPoints []objects.TimingPoint, currentTimingIndex int) (float64, int) {
	if len(timingPoints) == 1 {
		return timingPoints[0].BeatLength, 0

	} else if currentTimingIndex == len(timingPoints)-1 {
		return timingPoints[currentTimingIndex].BeatLength, currentTimingIndex

	} else {
		currentTiming := timingPoints[currentTimingIndex]
		nextTiming := timingPoints[currentTimingIndex+1]

		for currentTimingIndex < len(timingPoints) && time >= nextTiming.Time-2 {
			currentTimingIndex++
			currentTiming = nextTiming
			if currentTimingIndex+1 < len(timingPoints) {
				nextTiming = timingPoints[currentTimingIndex+1]
			}
		}

		if currentTimingIndex == len(timingPoints) {
			return currentTiming.BeatLength, currentTimingIndex - 1
		}

		return currentTiming.BeatLength, currentTimingIndex
	}
}

func getExpectedTime(currentObjectTime int64, timingPoints []objects.TimingPoint, timingPointIndex int, snap float64) (int64, int) {
	beatLength, timingPointIndex := getBeatLengthAt(currentObjectTime, timingPoints, timingPointIndex)
	snapLength := beatLength / snap

	return int64(math.Round(float64(currentObjectTime) + snapLength)), timingPointIndex
}

func isPartOfChain(objectTime, expectedTime1, expectedTime2 int64) bool {
	const error = 2

	case1 := objectTime >= expectedTime1-error && objectTime <= expectedTime1+error
	case2 := objectTime >= expectedTime2-error && objectTime <= expectedTime2+error

	return case1 || case2
}

func newChain(hitObjects []objects.HitObject, currentPosition int, timingPoints []objects.TimingPoint, timingPointIndex int, snap float64, minSize int) (int, int, *Chain) {
	if !(currentPosition+1 < len(hitObjects)) {
		return currentPosition + 1, 0, nil
	}

	starTime := hitObjects[currentPosition].Time
	size := 1
	currentObject := hitObjects[currentPosition]
	nextObject := hitObjects[currentPosition+1]

	expectedTime1, timingPointIndex := getExpectedTime(currentObject.Time, timingPoints, timingPointIndex, snap)
	expectedTime2, _ := getExpectedTime(nextObject.Time, timingPoints, timingPointIndex, snap)

	for currentPosition < len(hitObjects) && isPartOfChain(nextObject.Time, expectedTime1, expectedTime2) {
		currentPosition++
		size++
		currentObject = nextObject

		if currentPosition+1 < len(hitObjects) {
			nextObject = hitObjects[currentPosition+1]
		}

		expectedTime1, timingPointIndex = getExpectedTime(currentObject.Time, timingPoints, timingPointIndex, snap)
		expectedTime2, _ = getExpectedTime(nextObject.Time, timingPoints, timingPointIndex, snap)
	}

	endTime := currentObject.Time

	if size >= minSize {
		return currentPosition + 1, timingPointIndex, &Chain{snapDivisor: uint8(snap), startTime: starTime, endTime: endTime, size: size}
	} else {
		return currentPosition + 1, timingPointIndex, nil
	}
}

func GetBeatmapChains(beatmap Beatmap, snap int, minSize int) []Chain {
	chains := []Chain{}
	timingPointIndex := 0

	for i := 0; i < len(beatmap.hitObjects); {
		nextObjectIndex, currentTimingPointPos, chain := newChain(beatmap.hitObjects, i, beatmap.timingPoints, timingPointIndex, float64(snap), minSize)
		timingPointIndex = currentTimingPointPos

		if chain != nil {
			chains = append(chains, *chain)
		}

		i = nextObjectIndex
	}

	return chains
}
