package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	input := []int64{
		3, 8, 1005, 8, 311, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 29, 1006, 0, 98, 2, 1005, 8, 10, 1, 1107, 11, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 62, 1006, 0, 27, 2, 1002, 12, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 90, 1, 1006, 1, 10, 2, 1, 20, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 102, 1, 8, 121, 1, 1003, 5, 10, 1, 1003, 12, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 151, 1006, 0, 17, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 175, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 197, 2, 6, 14, 10, 1006, 0, 92, 1006, 0, 4, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 229, 1006, 0, 21, 2, 102, 17, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 259, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 280, 1006, 0, 58, 1006, 0, 21, 2, 6, 11, 10, 101, 1, 9, 9, 1007, 9, 948, 10, 1005, 10, 15, 99, 109, 633, 104, 0, 104, 1, 21101, 937150919572, 0, 1, 21102, 328, 1, 0, 1105, 1, 432, 21101, 0, 387394675496, 1, 21102, 1, 339, 0, 1106, 0, 432, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 46325083283, 1, 1, 21102, 1, 386, 0, 1106, 0, 432, 21101, 0, 179519401051, 1, 21102, 397, 1, 0, 1106, 0, 432, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21102, 1, 868410348308, 1, 21102, 1, 420, 0, 1105, 1, 432, 21102, 718086501140, 1, 1, 21102, 1, 431, 0, 1105, 1, 432, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21101, 0, 463, 3, 21101, 453, 0, 0, 1106, 0, 496, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 458, 459, 474, 4, 0, 1001, 458, 1, 458, 108, 4, 458, 10, 1006, 10, 490, 1101, 0, 0, 458, 109, -2, 2105, 1, 0, 0, 109, 4, 2102, 1, -1, 495, 1207, -3, 0, 10, 1006, 10, 513, 21102, 0, 1, -3, 22102, 1, -3, 1, 22102, 1, -2, 2, 21102, 1, 1, 3, 21102, 1, 532, 0, 1105, 1, 537, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 560, 2207, -4, -2, 10, 1006, 10, 560, 22101, 0, -4, -4, 1105, 1, 628, 22102, 1, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 1, 579, 0, 1105, 1, 537, 22101, 0, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 598, 21102, 1, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 620, 22102, 1, -1, 1, 21102, 1, 620, 0, 105, 1, 495, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2106, 0, 0,
	}

	second(input)
}

type point struct {
	x int64
	y int64
}

func first(input []int64) {
	inputi := make(chan int64, 1)
	outputi := make(chan int64)
	go generate(inputi, outputi, input)

	var currentx, currenty int64
	var dir int64
	mem := map[point]int64{}
	count := 0

	for {
		inputi <- mem[point{currentx, currenty}]
		paint, ok := <-outputi
		if !ok {
			break
		}
		op, ok := <-outputi
		if !ok {
			break
		}
		fmt.Printf("%d\t: %d - %d\n", count, paint, op)
		fmt.Printf("Current: (%d,%d) Dir: %d\n", currentx, currenty, dir)

		mem[point{currentx, currenty}] = paint
		dir = calculateNextDir(dir, op)
		currentx, currenty = calculateNextPos(currentx, currenty, dir)
		count++
	}

	fmt.Println(len(mem))
}

func second(input []int64) {
	inputi := make(chan int64, 1)
	outputi := make(chan int64)
	go generate(inputi, outputi, input)

	var currentx, currenty int64
	var dir int64
	mem := map[point]int64{}
	count := 0
	mem[point{0, 0}] = 1

	var minx, miny, maxx, maxy int64

	for {
		inputi <- mem[point{currentx, currenty}]
		paint, ok := <-outputi
		if !ok {
			break
		}
		op, ok := <-outputi
		if !ok {
			break
		}
		fmt.Printf("%d\t: %d - %d\n", count, paint, op)
		fmt.Printf("Current: (%d,%d) Dir: %d\n", currentx, currenty, dir)

		mem[point{currentx, currenty}] = paint
		dir = calculateNextDir(dir, op)
		currentx, currenty = calculateNextPos(currentx, currenty, dir)
		count++
		if currentx < minx {
			minx = currentx
		}
		if currenty < miny {
			miny = currenty
		}
		if currentx > maxx {
			maxx = currentx
		}
		if currenty > maxy {
			maxy = currenty
		}
	}

	fmt.Println(len(mem))
	fmt.Println(minx, miny, maxx, maxy)

	img := image.NewRGBA(image.Rect(int(minx), 0, int(maxx)+1, int(-miny)+1))

	for p, v := range mem {
		if v == 0 {
			img.Set(int(p.x), int(-p.y), color.RGBA{255, 0, 0, 255})
		} else {
			img.Set(int(p.x), int(-p.y), color.RGBA{255, 255, 0, 255})
		}

	}

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func calculateNextDir(currentDir int64, op int64) int64 {
	if op == 0 {
		nextDir := []int64{3, 0, 1, 2}
		return nextDir[currentDir]
	} else {
		nextDir := []int64{1, 2, 3, 0}
		return nextDir[currentDir]
	}
}

func calculateNextPos(currentx, currenty, dir int64) (int64, int64) {
	nt := [][]int64{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	nextx := currentx + nt[dir][0]
	nexty := currenty + nt[dir][1]

	return nextx, nexty
}

func generate(softInput chan int64, softOutput chan int64, input []int64) {
	var pos, relative int64

	mem := map[int64]int64{}
	for k, v := range input {
		mem[int64(k)] = v
	}

	for pos < int64(len(input)) {
		opCode, mode1, mode2, mode3 := parseOpcode(input[pos])

		if opCode == 0 {
			return
		}

		var entry1 int64

		if opCode == 3 {
			entry1 = getMem(mem, int64(pos+1), mode1, relative)
		} else {
			entry1 = getParam(mem, int64(pos+1), mode1, relative)
		}
		entry2 := getParam(mem, int64(pos+2), mode2, relative)
		entry3 := getMem(mem, int64(pos+3), mode3, relative)

		//fmt.Printf("POS %d, INPUT %d, OPCODE %d, RELATIVE %d,ENTRY1: %d, ENTRY2: %d\n", pos, input[pos], opCode, relative, entry1, entry2)
		switch opCode {
		case 99:
			close(softOutput)
			return
		case 1:
			mem[entry3] = entry1 + entry2
			pos += 4

		case 2:
			mem[entry3] = entry1 * entry2
			pos += 4

		case 3:
			mem[entry1] = <-softInput
			pos += 2

		case 4:
			softOutput <- entry1
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
				mem[entry3] = 1
			} else {
				mem[entry3] = 0
			}
			pos += 4

		case 8:
			if entry1 == entry2 {
				mem[entry3] = 1
			} else {
				mem[entry3] = 0
			}
			pos += 4
		case 9:
			relative += entry1
			pos += 2
		}
	}

	return
}

func parseOpcode(a int64) (int64, int64, int64, int64) {

	var opcode int64
	if a < 10 {
		opcode = a
	} else {
		opcode = digit(a, 2)*10 + digit(a, 1)
	}

	return opcode, digit(a, 3), digit(a, 4), digit(a, 5)

}

func getParam(mem map[int64]int64, p int64, mode int64, relative int64) int64 {
	switch mode {
	case 0:
		return mem[mem[p]]
	case 1:
		return mem[p]
	case 2:
		return mem[relative+mem[p]]
	default:
		fmt.Println("ERROR: Unrecognised mode")
		return 0
	}
}

func getMem(mem map[int64]int64, p int64, mode int64, relative int64) int64 {
	switch mode {
	case 0:
		return mem[p]
	case 1:
		fmt.Println("ERROR: Unrecognised mode1 ")
		return 0
	case 2:
		return relative + mem[p]
	default:
		fmt.Println("ERROR: Unrecognised mode")
		return 0
	}
}

func digit(num, place int64) int64 {
	r := num % int64(math.Pow(10, float64(place)))
	return r / int64(math.Pow(10, float64(place-1)))
}
