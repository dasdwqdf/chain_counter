package beatmap

import (
	"bufio"
	util "chain_counter/util"
	"os"
)

func ImportBeatmapData(fileName string) Beatmap {
	f, err := os.Open(fileName)
	util.CheckError(err)

	scanner := bufio.NewScanner(f)
	timingPointsTxt := []string{}
	hitObjectsTxt := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "[TimingPoints]" {
			for scanner.Scan() && line != "" {
				line = scanner.Text()
				timingPointsTxt = append(timingPointsTxt, line)
			}
		} else if line == "[HitObjects]" {
			for scanner.Scan() {
				line = scanner.Text()
				hitObjectsTxt = append(hitObjectsTxt, line)
			}
		}
	}

	f.Close()

	timingPointsTxt = timingPointsTxt[:len(timingPointsTxt)-1]

	return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
		hitObjects: CreateHitObjects(hitObjectsTxt)}
}
