package main

import (
	beatmap "chain_counter/beatmap"
	"fmt"
)

func main() {
	beatmap := beatmap.ImportBeatmapData("test.osu")
	fmt.Println(beatmap)
}
