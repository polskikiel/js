package wall

import "js/internal/block"

func NewWall(blocks []block.Block) int {
	return buildWall(blocks)
}

func buildWall(blocks []block.Block) int {
	blockHeightCache := 0
	nrOfPosters := 0

	for _, block := range blocks {
		if block.Height == blockHeightCache {
			continue
		}

		blockHeightCache = block.Height
		nrOfPosters++
	}
	return nrOfPosters
}






