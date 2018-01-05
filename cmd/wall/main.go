package main

import (
	"fmt"
	"js/internal/wall"
	"js/internal/block"
)

func main() {
	fmt.Println(wall.NewWall(block.GetBlocksFromFile("plik1.txt")))
}