package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "59755896917240436883590128801944128314960209697748772345812613779993681653921392130717892227131006192013685880745266526841332344702777305618883690373009336723473576156891364433286347884341961199051928996407043083548530093856815242033836083385939123450194798886212218010265373470007419214532232070451413688761272161702869979111131739824016812416524959294631126604590525290614379571194343492489744116326306020911208862544356883420805148475867290136336455908593094711599372850605375386612760951870928631855149794159903638892258493374678363533942710253713596745816693277358122032544598918296670821584532099850685820371134731741105889842092969953797293495"

	//first(input)
	second(input)
}

func first(input string) {
	mem := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		mem = append(mem, n)
	}

	pattern := []int{0, 1, 0, -1}

	for phase := 1; phase <= 100; phase++ {

		next := make([]int, len(mem))
		for i := 0; i < len(mem); i++ {

			for j := 0; j < len(mem); j++ {
				next[i] += pattern[((j+1)/(i+1))%4] * mem[j]
			}

			next[i] = abs(next[i] % 10)
		}

		mem = next
		fmt.Printf("%d\t%+v\n", phase, mem)
	}

}

func second(input string) {
	offset, _ := strconv.Atoi(input[:7])
	offsetLength := len(input)*10000 - offset
	fmt.Println("Input size: ", len(input)*10000)
	fmt.Println("Input Offset: ", input[:7])
	fmt.Println("Input Work Length: ", offsetLength)
	mem := make([]int, offsetLength)

	for i := 0; i < offsetLength; i++ {
		n, _ := strconv.Atoi(string(input[abs((649 - (i % 650)))]))
		mem[i] = n
	}

	for phase := 1; phase <= 100; phase++ {
		var current int
		for i := 0; i < offsetLength; i++ {
			current += mem[i]
			current %= 10
			mem[i] = current
		}
	}

	fmt.Println("Message: ")
	for i := 1; i <= 8; i++ {
		fmt.Printf("%d", mem[offsetLength-i])
	}
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
