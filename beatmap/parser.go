package beatmap

import (
	"bufio"
	util "chain_counter/util"
	"errors"
	"os"
	"slices"
	"strings"
)

func getSupportedVersions() []string {
	return []string{
		"osu file format v3",
		"osu file format v4",
		"osu file format v5",
		"osu file format v6",
		"osu file format v7",
		"osu file format v8",
		"osu file format v9",
		"osu file format v10",
		"osu file format v11",
		"osu file format v12",
		"osu file format v13",
		"osu file format v14"}
}

func getVersion(version string) string {
	args := strings.Split(version, " ")
	return args[3]
}

func handleVersion(versionTxt string, timingPointsTxt, hitObjectsTxt []string) Beatmap {
	version := getVersion(versionTxt)

	if version == "v3" {
		for i, timingTxt := range timingPointsTxt {
			timingPointsTxt[i] = timingTxt + "1,1,1,1,1,1"
		}

		return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
			hitObjects: CreateHitObjects(hitObjectsTxt), version: version}

	} else if version == "v4" {
		for i, timingTxt := range timingPointsTxt {
			timingPointsTxt[i] = timingTxt + "1,1,1"
		}

		return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
			hitObjects: CreateHitObjects(hitObjectsTxt), version: version}

	} else {
		return Beatmap{timingPoints: CreateTimingPoints(timingPointsTxt),
			hitObjects: CreateHitObjects(hitObjectsTxt), version: version}
	}
}

func ImportBeatmapData(fileName string) Beatmap {
	f, err := os.Open(fileName)
	util.CheckError(err)

	scanner := bufio.NewScanner(f)
	timingPointsTxt := []string{}
	hitObjectsTxt := []string{}

	versionTxt := ""
	if scanner.Scan() {
		versionTxt = scanner.Text()
	} else {
		err = errors.New("empty file")
		util.CheckError(err)
	}

	if !slices.Contains(getSupportedVersions(), versionTxt) {
		err = errors.New("invalid version")
		util.CheckError(err)
	} else {
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
	}

	if len(timingPointsTxt) > 0 {
		timingPointsTxt = timingPointsTxt[:len(timingPointsTxt)-1]
	} else {
		err = errors.New("no timing points found")
		util.CheckError(err)
	}

	return handleVersion(versionTxt, timingPointsTxt, hitObjectsTxt)
}
