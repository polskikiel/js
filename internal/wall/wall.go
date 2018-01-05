package wall

import "js/internal/block"

type Wall int

func NewWall(blocks []block.Block) Wall {
	return buildWall(blocks)
}

func buildWall(blocks []block.Block) Wall {
	blockHeightCache := 0
	nrOfPosters := 0

	for _, block := range blocks {
		if block.Height == blockHeightCache {
			continue
		}

		nrOfPosters++


		blockHeightCache = block.Height
	}
	return Wall(nrOfPosters)
}






