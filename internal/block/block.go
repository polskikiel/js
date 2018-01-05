package block

import (
	"io/ioutil"
	"fmt"
)

type Block struct {
	Width, Height int
}

func GetBlocksFromFile(path string) ([]Block) {
	var blocks []Block
	lines := 0

	if f, err := ioutil.ReadFile(path); err != nil {
		fmt.Println("txt file not found", err)
	} else {
		if len(f) > 0 {
			flag := true
			for i, b := range f {
				if b == byte('\n') {
					lines++
					continue
				}
				if lines > 0 {
					var num []byte
					for _, bt := range f[i:] {
						if bt == byte('\n') || bt == byte(' ') {
							break
						}
						num = append(num, bt)
					}

					if flag {
						flag = false
					} else {
						flag = true
					}
				}

			}
		} else {
			fmt.Println("txt file empty")
		}

	}
	return blocks
}