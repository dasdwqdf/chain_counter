package beatmap

type Chain struct {
	snapDivisor uint8
	size        uint64
}

func getBeatmapChains(beatmap Beatmap) []Chain {
	chains := []Chain{}

	if len(beatmap.timingPoints) == 1 {

	} else {
		return nil
	}
}
