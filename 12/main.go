package main

import "fmt"

func main() {
	input := [4][3]int{
		[3]int{10, 15, 7},
		[3]int{15, 10, 0},
		[3]int{20, 12, 3},
		[3]int{0, -3, 13},
	}

	//moons := [][]int{
	//	[]int{-1, 0, 2},
	//	[]int{2, -10, -7},
	//	[]int{4, -8, 8},
	//	[]int{3, 5, -1},
	//}

	second(input)
}

type moonState struct {
	moons    [4][3]int
	velocity [4][3]int
}

func copyState(in *moonState) *moonState {
	var m, v [4][3]int
	for i, m1 := range in.moons {
		for j, m2 := range m1 {
			m[i][j] = m2
			v[i][j] = in.velocity[i][j]
		}
	}

	return &moonState{
		moons:    m,
		velocity: v,
	}
}

func nextStep(in *moonState) *moonState {

	moons := in.moons
	velocity := in.velocity

	velocityChange := [4][3]int{}
	for i, m1 := range moons {
		for j, m2 := range moons[i:] {

			for k, v := range m1 {
				if v < m2[k] {
					velocityChange[i][k]++
					velocityChange[i+j][k]--
				} else if v > m2[k] {
					velocityChange[i][k]--
					velocityChange[i+j][k]++
				}
			}

		}
	}

	for i, m := range velocityChange {
		for j, v := range m {
			velocity[i][j] += v
		}
	}

	for i, m := range velocity {
		for j, v := range m {
			moons[i][j] += v
		}
	}

	return &moonState{
		moons:    moons,
		velocity: velocity,
	}
}

func first(moons [4][3]int) {

	velocity := [4][3]int{
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
	}

	count := 0

	moonState := &moonState{
		moons:    moons,
		velocity: velocity,
	}

	for count < 1000 {

		moonState = nextStep(moonState)
		count++

	}

	total := 0
	for i, m := range moonState.moons {
		k := 0
		p := 0
		for j, v := range m {
			if v < 0 {
				k += -v
			} else {
				k += v
			}

			if moonState.velocity[i][j] < 0 {
				p += -moonState.velocity[i][j]
			} else {
				p += moonState.velocity[i][j]
			}
		}
		total += k * p
	}

	fmt.Println(total)

}

type coordState struct {
	moons    [4]int
	velocity [4]int
}

func nextCoord(in coordState) coordState {
	moons := in.moons
	velocity := in.velocity

	velocityChange := [4]int{}
	for i, m1 := range moons {
		for j, m2 := range moons[i:] {

			if m1 < m2 {
				velocityChange[i]++
				velocityChange[i+j]--
			} else if m1 > m2 {
				velocityChange[i]--
				velocityChange[i+j]++
			}

		}
	}

	for i, v := range velocityChange {
		velocity[i] += v
	}

	for i, v := range velocity {
		moons[i] += v
	}

	return coordState{
		moons:    moons,
		velocity: velocity,
	}
}

func second(moons [4][3]int) {

	state := [3]coordState{
		coordState{
			moons:    [4]int{moons[0][0], moons[1][0], moons[2][0], moons[3][0]},
			velocity: [4]int{},
		},
		coordState{
			moons:    [4]int{moons[0][1], moons[1][1], moons[2][1], moons[3][1]},
			velocity: [4]int{},
		},
		coordState{
			moons:    [4]int{moons[0][2], moons[1][2], moons[2][2], moons[3][2]},
			velocity: [4]int{},
		},
	}

	mem := [3]map[coordState]int{
		map[coordState]int{},
		map[coordState]int{},
		map[coordState]int{},
	}

	result := [3]int{}

	for i := 0; i < 3; i++ {
		count := 0

		mem[i][state[i]] = 0

		for {

			count++
			state[i] = nextCoord(state[i])

			if match, exists := mem[i][state[i]]; exists {
				fmt.Println(i)
				fmt.Println(count)
				fmt.Println(match)
				fmt.Println("----")
				result[i] = count
				break
			}

			mem[i][state[i]] = count
		}
	}
	fmt.Println(LCM(result[0], result[1], result[2]))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// secondRainbow is a slow implementation using rainbow tables
func secondRainbow(moons [4][3]int) {
	velocity := [4][3]int{
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
	}

	count := 0

	state := &moonState{
		moons:    moons,
		velocity: velocity,
	}

	mem := map[moonState]int{}
	countmem := map[int]*moonState{}
	foundcount := 0
	matchcount := 0

	ratio := 1000000

	for {

		state = nextStep(state)
		count++

		if match, exists := mem[*state]; exists {
			fmt.Println(count)
			fmt.Println(match)

			foundcount = match
			matchcount = count

			break
		}

		if count%1000000 == 0 {
			//fmt.Println(count)
			mem[*state] = count
			countmem[count] = state
		}

	}

	fmt.Println("Stage2")

	count = foundcount - ratio
	state = countmem[count]

	for {
		state = nextStep(state)

		count++
		mem[*state] = count

		if count == foundcount {
			break
		}
	}

	fmt.Println("Filled")

	count = matchcount - (matchcount % ratio) - ratio
	state = countmem[count]

	for {
		state = nextStep(state)

		count++

		if match, exists := mem[*state]; exists {
			fmt.Println(count)
			fmt.Println(match)

			foundcount = match
			matchcount = count

			break
			return
		}
	}

}
