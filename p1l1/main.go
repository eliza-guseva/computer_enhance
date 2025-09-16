package main

import (
	"os"
	"fmt"
)

const ASSEM_SIMPLE = "listing_0037_single_register_mov"
const ASSEM_LARGE = "listing_0038_many_register_mov"


var REGMAP = map[int][2]string{
	0: {"AX", "AL"},
	1: {"CX", "CL"},
	2: {"DX", "DL"},
	3: {"BX", "BL"},
	4: {"SP", "AH"},
	5: {"BP", "CH"},
	6: {"SI", "DH"},
	7: {"DI", "BH"},
}

func main() {
	data, _ := readData(ASSEM_LARGE)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Instructions ended")
		}
	}()
	isFirstByte := true
	isMov := false
	var w int
	var d bool
	output := ""
	for i, b := range data {
		fmt.Printf("Checking byte %d, %v\n", i, b)
		if isFirstByte {
			if b >> 2 == 34 {
				output += "MOV "
				isMov = true
				d = b << 6 & 128 != 0 
				fmt.Println(d)
				w = boolToInt(b << 7 & 128 == 0)
				fmt.Println(w)
				fmt.Println(isMov)
				isFirstByte = false
				continue
			} else {
				isMov = false
				isFirstByte = false
				continue
			}
		} else if ! isMov {
			isFirstByte = true
			continue
		} else {
			reg1 := b >> 3 & 7
			reg2 := b & 7
			mod  := b >> 6 & 3
			if mod != 3 {
				isFirstByte = true
				isMov = false
				continue
			} else if d {
				output += REGMAP[int(reg1)][w] + ", " + REGMAP[int(reg2)][w]
			} else {
				output += REGMAP[int(reg2)][w] + ", " + REGMAP[int(reg1)][w]
			}
		}
		isFirstByte = true
		isMov = false
		output += "\n"
		fmt.Println(output)
	}
}


func readData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("something aint right: %v", err)
		return []byte{}, err
	}
	return data, nil
}


func boolToInt(b bool) int {
      if b {
          return 1
      }
      return 0
  }
