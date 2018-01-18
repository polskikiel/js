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
			var num []byte

			for _, b := range f {
				//fmt.Println(b)
				if b > 47 && b < 58 {
					num = append(num, b - 48)

					if whiteSpace {
						if _, err := convert(num); err != nil {
							fmt.Errorf("can't conver byte[] to int")
						} else if flag {
							flag = false
						} else if !flag {
							flag = true
						}
					}

				} else {
					whiteSpace = true
					fmt.Println(num)
					num = []byte{}
				}
			}
		} else {
			fmt.Println("txt file empty")
		}

	}
	return blocks
}

func convert(data []byte) (uint32, error) {
	v, err := strconv.ParseUint(string(data), 10, 32)
	if err != nil {
		return 0, err
	}
	fmt.Println(uint32(v))
	return uint32(v), nil
}
