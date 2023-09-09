package beatmap

import (
	obj "chain_counter/beatmap/objects"
)

type Beatmap struct {
	timingPoints []obj.TimingPoint
	hitObjects   []obj.HitObject
}

func CreateTimingPoints(timingPointsTxt []string) []obj.TimingPoint {
	timingPoints := []obj.TimingPoint{}

	for _, timingSection := range timingPointsTxt {
		if obj.IsRedLine(timingSection) {
			timingPoints = append(timingPoints, obj.NewTimingPoint(timingSection))
		}
	}

	return timingPoints
}

func CreateHitObjects(hitObjectsTxt []string) []obj.HitObject {
	hitObjects := []obj.HitObject{}

	for _, hitObject := range hitObjectsTxt {
		hitObjects = append(hitObjects, obj.NewHitObject(hitObject))
	}

	return hitObjects
}
