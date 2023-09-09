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
				timingPointsTxt = append(timingPointsTxt, line)
				line = scanner.Text()
			}
		} else if line == "[HitObjects]" {
			for scanner.Scan() && line != "" {
				hitObjectsTxt = append(hitObjectsTxt, line)
				line = scanner.Text()
			}
		}
	}

	f.Close()

	timingPointsTxt = timingPointsTxt[1:]
	hitObjectsTxt = hitObjectsTxt[1:]

	return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
		hitObjects: CreateHitObjects(hitObjectsTxt)}
}
