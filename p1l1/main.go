package main

import (
	"os"
	"fmt"
)

const ASSEM_SIMPLE = "listing_0037_single_register_mov"


func main() {
	instructions, _ := readData(ASSEM_SIMPLE)
	_ = instructions
}


func readData(filename string) ([16]int8, error) {
	instruction := [16]int8{}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("something aint right: %v", err)
		return instruction, err
	}
	for i, byte := range data  {
		fmt.Println(byte)
		for bit_pos := 7; bit_pos >= 0; bit_pos-- {
			if byte & (1 << bit_pos) != 0 {
				instruction[i*8 + bit_pos] = int8(1)	
			} else {
				instruction[i*8 + bit_pos] = int8(0)
			}
		}
	}
	fmt.Println(instruction)
	return instruction, nil
}


