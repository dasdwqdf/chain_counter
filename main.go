package main

import (
	bmp "chain_counter/beatmap"
	"fmt"
)

func main() {
	beatmap := bmp.ImportBeatmapData("test.osu")
	// fmt.Println(bmp.GetBeatmapChains(beatmap, 4))
	fmt.Println(bmp.GetBeatmapChains(beatmap, 4, 5))
}
