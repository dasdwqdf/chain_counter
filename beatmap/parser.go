package beatmap

import (
	"bufio"
	util "chain_counter/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getVersion(version string) string {
	args := strings.Split(version, " ")
	return args[3]
}

func handleVersion(versionTxt, timingPointsTxt, hitObjectsTxt string) Beatmap {
	version, err := strconv.ParseInt(getVersion(versionTxt), 0, 8)
	util.CheckError(err)

	if version == 3 {
		return Beatmap{}
	}

	return Beatmap{}
}

func ImportBeatmapData(fileName string) Beatmap {
	f, err := os.Open(fileName)
	util.CheckError(err)

	scanner := bufio.NewScanner(f)
	timingPointsTxt := []string{}
	hitObjectsTxt := []string{}

	version := ""

	if scanner.Scan() {
		version = scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "[TimingPoints]" {
			for line != "" && scanner.Scan() {
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

	err = f.Close()
	util.CheckError(err)

	if len(timingPointsTxt) > 0 {
		timingPointsTxt = timingPointsTxt[:len(timingPointsTxt)-1]
	}

	fmt.Println(version, timingPointsTxt, hitObjectsTxt)

	return Beatmap{}

	// return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
	// 	hitObjects: CreateHitObjects(hitObjectsTxt)}
}
