package main

import (
	bmp "chain_counter/beatmap"
)

func main() {
	bmp.ImportBeatmapData("test.osu")
	// fmt.Println(bmp.GetBeatmapChains(beatmap, 4))
	// fmt.Println(bmp.GetBeatmapChains(beatmap, 4, 5))
}
