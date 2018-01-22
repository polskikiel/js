package main

import (
	"fmt"
	"js/internal/wall"
	"js/internal/block"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("wynik2.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Błąd tworzenia pliku", err)
		os.Exit(1)
	}

	_, werr := file.Write([]byte(strconv.Itoa(wall.NewWall(block.GetBlocksFromFile("dane2.txt")))))
	if werr != nil {
		fmt.Println("Błąd zapisywania do pliku", werr)
		os.Exit(1)
	}
}
