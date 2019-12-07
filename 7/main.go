package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{
		3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 55, 76, 97, 114, 195, 276, 357, 438, 99999, 3, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 102, 5, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 101, 4, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 102, 4, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99,
	}

	second(input)
}

func first(input []int) {

	choices := []int{0, 1, 2, 3, 4}
	result := make(chan int)
	go firstRec(input, choices, 0, result)

	max := 0
	for {
		select {

		case r, ok := <-result:
			if !ok {
				return
			}

			if r > max {
				max = r
				fmt.Println(max)
			}

		}
	}
}

func second(input []int) {
	allChoices := [][]int{}
	generateAllChoices([]int{5, 6, 7, 8, 9}, []int{}, &allChoices)

	max := 0

	for _, choice := range allChoices {

		result := secondHelper(input, choice)
		if result > max {
			max = result
		}

	}

	fmt.Println("RESULT ", max)
}

func secondHelper(input []int, choice []int) int {

	nextInput := make(chan int)
	firstInput := nextInput

	ended := make(chan bool)
	for i := 0; i < 5; i++ {
		b := make([]int, len(input))
		copy(b, input)

		softOutput := make(chan int)

		if i == 4 {
			go generate(0, nextInput, softOutput, b, ended)
		} else {
			go generate(0, nextInput, softOutput, b, nil)
		}
		nextInput <- choice[i]

		nextInput = softOutput
	}

	firstInput <- 0

	for {
		select {

		case r, ok := <-nextInput:
			if !ok {
			} else {

				select {
				case <-ended:
					fmt.Println("cloosed. Last: ", r)

					close(firstInput)
					close(ended)
					return r

				case firstInput <- r:
					fmt.Println("555. Last: ", r)
				}
			}
		}
	}
}

func generateAllChoices(choices []int, current []int, fullChoices *[][]int) {
	if len(choices) == 0 {
		*fullChoices = append(*fullChoices, current)
	}

	for _, i := range choices {

		generateAllChoices(removeChoice(choices, i), append(current, i), fullChoices)

	}
}

func firstRec(input []int, choices []int, softInput int, result chan int) {
	if len(choices) == 0 {
		result <- softInput
	}
	for _, i := range choices {
		inputi := make(chan int, 2)
		inputi <- i
		inputi <- softInput
		outputi := make(chan int, 1)
		generate(0, inputi, outputi, input, nil)

		firstRec(input, removeChoice(choices, i), <-outputi, result)
	}

	if len(choices) == 5 {
		close(result)
	}
}

func removeChoice(choices []int, i int) []int {
	newChoice := []int{}
	for _, j := range choices {
		if i != j {
			newChoice = append(newChoice, j)
		}
	}
	return newChoice
}

func generate(id int, softInput chan int, softOutput chan int, input []int, ended chan bool) {
	//lastOutput := 0
	pos := 0

	for pos < len(input) {
		opCode, pos1, pos2, _ := parseOpcode(input[pos])

		//fmt.Printf("%d , POS %d, INPUT %d, OPCODE %d\n", id, pos, input[pos], opCode)

		if opCode == 0 {
			return
		}

		entry1 := 0
		entry2 := 0

		if pos1 {
			entry1 = input[pos+1]
		} else {
			if pos+1 < len(input) {
				if input[pos+1] < len(input) {
					entry1 = input[input[pos+1]]
				}
			}
		}
		if pos2 {
			entry2 = input[pos+2]
		} else {
			if pos+2 < len(input) {
				if input[pos+2] < len(input) {
					entry2 = input[input[pos+2]]
				}
			}
		}

		switch opCode {
		case 99:
			//			fmt.Printf("%d , END 99 \n", id)
			//			fmt.Printf("%d , Last Output %d\n", id, lastOutput)
			close(softOutput)
			if ended != nil {
				ended <- true
			}
			return
		case 1:
			input[input[pos+3]] = entry1 + entry2
			pos += 4

		case 2:
			input[input[pos+3]] = entry1 * entry2
			pos += 4

		case 3:
			//fmt.Printf("%d , INPUT tentative\n", id)
			input[input[pos+1]] = <-softInput
			//fmt.Printf("%d , INPUT successful %d\n", id, input[input[pos+1]])
			pos += 2

		case 4:
			//fmt.Printf("%d ,OUTPUT %d\n", id, entry1)
			//lastOutput = entry1
			softOutput <- entry1
			//fmt.Printf("%d ,OUTPUT successful %d\n", id, entry1)
			pos += 2

		case 5:
			if entry1 != 0 {
				pos = entry2
			} else {
				pos += 3
			}

		case 6:
			if entry1 == 0 {
				pos = entry2
			} else {
				pos += 3
			}

		case 7:
			if entry1 < entry2 {
				input[input[pos+3]] = 1
			} else {
				input[input[pos+3]] = 0
			}
			pos += 4

		case 8:
			if entry1 == entry2 {
				input[input[pos+3]] = 1
			} else {
				input[input[pos+3]] = 0
			}
			pos += 4
		}
	}

	return
}

func parseOpcode(a int) (int, bool, bool, bool) {
	opcode := 0
	pos1 := false
	pos2 := false
	pos3 := false

	if a < 10 {
		opcode = a
	} else {
		opcode = digit(a, 2)*10 + digit(a, 1)
	}

	if digit(a, 3) == 1 {
		pos1 = true
	}
	if digit(a, 4) == 1 {
		pos2 = true
	}
	if digit(a, 5) == 1 {
		pos3 = true
	}

	if opcode != 0 {
		//fmt.Println(opcode)
	}
	return opcode, pos1, pos2, pos3
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
