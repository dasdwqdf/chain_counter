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

func newChain(hitObjects []objects.HitObject, currentPosition int, timingPoints []objects.TimingPoint, timingPointIndex int, snap float64, minSize int) (int, int, *Chain) {
	if !(currentPosition+1 < len(hitObjects)) {
		return currentPosition + 1, 0, nil
	}

	starTime := hitObjects[currentPosition].Time
	size := 1
	currentObject := hitObjects[currentPosition]
	nextObject := hitObjects[currentPosition+1]

	beatLength, timingPointIndex := getBeatLengthAt(currentObject.Time, timingPoints, timingPointIndex)
	snapLength := beatLength / snap
	expectedTime := int64(math.Round(float64(currentObject.Time) + snapLength))

	beatLength2, _ := getBeatLengthAt(nextObject.Time, timingPoints, timingPointIndex)
	snapLength2 := beatLength2 / snap
	expectedTime2 := int64(math.Round(float64(currentObject.Time) + snapLength2))

	isPartOfChain := (nextObject.Time >= expectedTime-2 && nextObject.Time <= expectedTime+2) ||
		(nextObject.Time >= expectedTime2-2 && nextObject.Time <= expectedTime2+2)

	for currentPosition < len(hitObjects) && isPartOfChain {
		currentPosition++
		size++
		currentObject = nextObject

		if currentPosition+1 < len(hitObjects) {
			nextObject = hitObjects[currentPosition+1]
		}

		beatLength, timingPointIndex = getBeatLengthAt(currentObject.Time, timingPoints, timingPointIndex)
		snapLength := beatLength / snap
		expectedTime = int64(math.Round(float64(currentObject.Time) + snapLength))

		beatLength2, _ = getBeatLengthAt(nextObject.Time, timingPoints, timingPointIndex)
		snapLength2 = beatLength2 / snap
		expectedTime2 = int64(math.Round(float64(currentObject.Time) + snapLength2))

		isPartOfChain = (nextObject.Time >= expectedTime-2 && nextObject.Time <= expectedTime+2) ||
			(nextObject.Time >= expectedTime2-2 && nextObject.Time <= expectedTime2+2)

		// if nextObject.Time == 163311 {
		// 	fmt.Println(expectedTime, expectedTime2)
		// }

		// fmt.Println(currentObject.Time, nextObject.Time, expectedTime)
	}

	endTime := currentObject.Time

	if size >= minSize {
		// fmt.Println(starTime, endTime, size)
		return currentPosition + 1, timingPointIndex, &Chain{snapDivisor: uint8(snap), startTime: starTime, endTime: endTime, size: size}
	} else {
		return currentPosition + 1, timingPointIndex, nil
	}
}

func GetBeatmapChains(beatmap Beatmap, snap uint8, minSize int) []Chain {
	chains := []Chain{}
	timingPointIndex := 0

	for i := 0; i < len(beatmap.hitObjects); {
		next, currentTimingPoint, chain := newChain(beatmap.hitObjects, i, beatmap.timingPoints, timingPointIndex, float64(snap), minSize)
		timingPointIndex = currentTimingPoint

		if chain != nil {
			chains = append(chains, *chain)
		}
		i = next
	}

	return chains
}
