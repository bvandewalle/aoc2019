package main

import (
	"fmt"
	"sort"
)

func main() {
	input1 := `#..#.#.#.######..#.#...##
##.#..#.#..##.#..######.#
.#.##.#..##..#.#.####.#..
.#..##.#.#..#.#...#...#.#
#...###.##.##..##...#..#.
##..#.#.#.###...#.##..#.#
###.###.#.##.##....#####.
.#####.#.#...#..#####..#.
.#.##...#.#...#####.##...
######.#..##.#..#.#.#....
###.##.#######....##.#..#
.####.##..#.##.#.#.##...#
##...##.######..##..#.###
...###...#..#...#.###..#.
.#####...##..#..#####.###
.#####..#.#######.###.##.
#...###.####.##.##.#.##.#
.#.#.#.#.#.##.#..#.#..###
##.#.####.###....###..##.
#..##.#....#..#..#.#..#.#
##..#..#...#..##..####..#
....#.....##..#.##.#...##
.##..#.#..##..##.#..##..#
.##..#####....#####.#.#.#
#..#..#..##...#..#.#.#.##
`
	//
	//	input2 := `.#..##.###...#######
	//##.############..##.
	//.#.######.########.#
	//.###.#######.####.#.
	//#####.##.#.##.###.##
	//..#####..#.#########
	//####################
	//#.####....###.#.#.##
	//##.#################
	//#####.##.###..####..
	//..######..##.#######
	//####.##.####...##..#
	//.#####..#.######.###
	//##...#.##########...
	//#.##########.#######
	//.####.#.###.###.#.##
	//....##.##.###..#####
	//.#.#.###########.###
	//#.#.#.#####.####.###
	//###.##.####.##.#..##
	//`

	//	input3 := `......#.#.
	//#..#.#....
	//..#######.
	//.#.#.###..
	//.#..#.....
	//..#....#.#
	//#..#....#.
	//.##.#..###
	//##...#..#.
	//.#....####
	//`

	second(input1)
}

func first(input string) {
	mem := parse(input)

	max := 0
	maxx := 0
	maxy := 0

	for j, vj := range mem {
		for i, vi := range vj {

			if vi == 1 {
				current, _ := count(mem, i, j)
				if current > max {
					max = current
					maxx = i
					maxy = j
				}
			}

		}
	}

	fmt.Println(max, maxx, maxy)
}

func second(input string) {
	bestx := 11
	besty := 19
	mem := parse(input)

	c, slopeMem := count(mem, bestx, besty)
	fmt.Println(c)

	slopes := []frac{}

	for k := range slopeMem {
		slopes = append(slopes, k)
	}

	sort.SliceStable(slopes, func(j, i int) bool {
		if slopes[i].inf {
			return !slopes[i].sign
		}
		if slopes[j].inf {
			return slopes[j].sign
		}
		if slopes[i].sign != slopes[j].sign {
			return slopes[i].sign
		}
		return float64(slopes[i].num)/float64(slopes[i].den) > (float64(slopes[j].num) / float64(slopes[j].den))
	})
	fmt.Printf("%+v\n", slopes)

	for k, v := range slopeMem {

		sort.SliceStable(v, func(i, j int) bool {
			disti := (v[i].x-bestx)*(v[i].x-bestx) + (v[i].y-besty)*(v[i].y-besty)
			distj := (v[j].x-bestx)*(v[j].x-bestx) + (v[j].y-besty)*(v[j].y-besty)
			return disti < distj
		})

		slopeMem[k] = v
	}

	for _, v := range slopes {
		for _, v2 := range slopeMem[v] {
			fmt.Printf("%d %d --", v2.x, v2.y)
		}
		fmt.Println()
	}

	round := 0
	complete := false
	count := 0

	for !complete {
		complete = true
		for _, v := range slopes {
			l := slopeMem[v]
			if len(l) > round {
				count++
				complete = false
				fmt.Printf("%d\t: %d %d \n", count, l[round].x, l[round].y)
				if count == 200 {
					fmt.Printf("200 %d %d --", l[round].x, l[round].y)
				}
			}
		}
		round++
		fmt.Println()
	}
	fmt.Println(round)

}

type frac struct {
	sign bool
	num  int
	den  int
	inf  bool
	//	neg  bool
}
type point struct {
	x int
	y int
}

func (f *frac) reduce() {
	//if f.num < 0 {
	//	f.neg = !f.neg
	//	f.num = -f.num
	//}
	if f.den < 0 {
		f.den = -f.den
		f.num = -f.num
	}

	a := f.num
	b := f.den

	if a == 0 {
		f.num = 0
		f.den = 1
		//f.neg = false
		return
	}
	if b == 0 {
		f.num = 0
		f.den = 1
		f.inf = true
		//f.neg = false
		return
	}

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	f.num /= a
	f.den /= a
}

func count(mem [][]int, x, y int) (int, map[frac][]point) {
	count := map[frac][]point{}

	for j, vj := range mem {
		for i, vi := range vj {
			if vi == 1 {
				if !((i == x) && (j == y)) {
					// (i,j) and (x,y)
					sign := false
					if i < x {
						sign = true
					}
					if i == x {
						if j < y {
							sign = true
						}
					}
					num := y - j
					den := x - i
					slope := frac{sign, num, den, false}
					slope.reduce()
					fmt.Printf("%d %d:  %+v", i, j, slope)
					if _, exists := count[slope]; exists {
						fmt.Printf(" >Exist")
						count[slope] = append(count[slope], point{i, j})
					} else {
						count[slope] = []point{
							point{i, j},
						}
					}
					fmt.Printf("\n")

				}
			}
		}
	}
	c := len(count)
	fmt.Println(c)

	return c, count
}

func parse(input string) [][]int {
	out := [][]int{}
	current := []int{}
	count := 0
	for _, c := range input {
		if c == '#' {
			current = append(current, 1)
			count++
		} else if c == '.' {
			current = append(current, 0)
		} else if c == '\n' {
			out = append(out, current)
			current = []int{}
		} else {
			fmt.Println("fail")
		}
	}

	fmt.Println(count)

	return out
}
