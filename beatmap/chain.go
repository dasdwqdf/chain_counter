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

func newChain(hitObjects []objects.HitObject, currentPosition int, beatLength float64, snap float64, minSize int) (int, *Chain) {
	if !(currentPosition+1 < len(hitObjects)) {
		return currentPosition + 1, nil
	}

	starTime := hitObjects[currentPosition].Time
	size := 1
	currentObject := hitObjects[currentPosition]
	nextObject := hitObjects[currentPosition+1]

	snapLength := beatLength / snap
	expectedTime := int64(math.Round(float64(currentObject.Time) + snapLength))

	for currentPosition < len(hitObjects) && (nextObject.Time >= expectedTime-2 && nextObject.Time <= expectedTime+2) {
		currentPosition++
		size++
		currentObject = nextObject
		if currentPosition+1 < len(hitObjects) {
			nextObject = hitObjects[currentPosition+1]
		}
		expectedTime = int64(math.Round(float64(currentObject.Time) + snapLength))
		// fmt.Println(currentObject.Time, nextObject.Time, expectedTime)
	}

	endTime := currentObject.Time

	if size >= minSize {
		// fmt.Println(starTime, endTime, size)
		return currentPosition + 1, &Chain{snapDivisor: uint8(snap), startTime: starTime, endTime: endTime, size: size}
	} else {
		return currentPosition + 1, nil
	}
}

func GetBeatmapChains(beatmap Beatmap, snap uint8, minSize int) []Chain {
	chains := []Chain{}

	if len(beatmap.timingPoints) == 1 {
		beatLength := beatmap.timingPoints[0].BeatLength

		for i := 0; i < len(beatmap.hitObjects); {
			next, chain := newChain(beatmap.hitObjects, i, beatLength, float64(snap), minSize)
			if chain != nil {
				chains = append(chains, *chain)
			}
			i = next
		}

		return chains

	} else {
		return nil
	}
}
