package block

import (
	"io/ioutil"
	"fmt"
	"strconv"
)

type Block struct {
	Width, Height int
}

func GetBlocksFromFile(path string) ([]Block) {
	var blocks []Block

	if f, err := ioutil.ReadFile(path); err != nil {
		fmt.Println("txt file not found", err)
	} else {
		if len(f) > 0 {
			flag := true
			whiteSpace := false
			wasWhiteSpace := false
			var num []byte
			var h, w int


			for _, b := range f {
				if b > 47 && b < 58 {
					num = append(num, b)
					if whiteSpace {
						whiteSpace = false
						if flag {
							flag = false
						} else {
							flag = true
						}
					}

				} else {
					if whiteSpace {
						wasWhiteSpace = true
					} else {
						wasWhiteSpace = false
					}
					whiteSpace = true

					if !wasWhiteSpace {
						if !flag {
							w, _ = strconv.Atoi(string(num))
							blocks = append(blocks, createBlock(h, w))
							fmt.Println(len(blocks))

						} else {
							h, _ = strconv.Atoi(string(num))
						}
						num = []byte{}
					}
				}
			}
		} else {
			fmt.Println("txt file empty")
		}
	}

	return blocks
}

func createBlock(width, height int) Block {
	return Block{Width: width, Height: height}
}

func (b *Block) String() string {
	return string(b.Height + b.Width)
}
