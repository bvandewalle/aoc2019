package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := []string{
		"10 KVPH => 5 HPRK",
		"5 RSTBJ => 5 QKBQL",
		"2 GZWFN, 21 WBPFQ => 5 KMFWH",
		"5 JDJB, 1 FSWFT, 1 NKVSV => 6 MGKSL",
		"5 BCRHK => 9 KXFTL",
		"23 NKVSV, 2 RSTBJ => 9 QPBVD",
		"19 BKFVS, 7 JZBFT => 7 XWTQ",
		"14 JLXP, 4 LSCL => 8 FWLTD",
		"173 ORE => 5 TZSDV",
		"2 FPVH, 1 JDJB, 3 KHRW => 2 QLNJ",
		"1 HTGMX, 1 GVJVK, 2 RLRK => 2 HWBM",
		"1 GLVHT, 1 PBCT, 5 ZWKGV, 1 QSVJ, 2 FWLTD, 3 CNVPB, 1 QGNL => 8 RNLTX",
		"1 KXZTS => 2 BKFVS",
		"1 KVPH, 6 PVHPV, 2 TZSDV => 4 RLRK",
		"118 ORE => 1 VRVZ",
		"7 MGKSL, 4 HWBM => 2 GZWFN",
		"5 PVHPV => 7 HTGMX",
		"25 LSCL, 12 GVMFW => 6 ZWKGV",
		"1 CTPND, 1 KXZTS => 3 FRQH",
		"1 KXFTL => 3 PBCT",
		"1 CMPX => 4 KZNBL",
		"2 HDQVB, 1 QPBVD => 5 CTPND",
		"14 KVPH => 1 FCBQN",
		"3 XWTQ, 22 CTHM, 4 KVPH, 4 BZTV, 1 KMFWH, 12 NRFK => 7 CXVR",
		"1 GVJVK => 7 RSTBJ",
		"1 GVJVK => 4 NSQHW",
		"3 NKVSV => 8 KHRW",
		"8 HDQVB, 9 BCRHK => 6 GVMFW",
		"142 ORE => 7 KVPH",
		"4 TZSDV => 2 GVJVK",
		"4 KVPH, 10 HWBM => 3 NRFK",
		"47 PBCT, 15 CXVR, 45 GVJVK, 23 KZNBL, 1 WFPNP, 14 RNLTX => 1 FUEL",
		"1 PCBNG => 4 QLJXM",
		"1 SHTQF => 2 FNWBZ",
		"2 FCBQN, 1 BCRHK => 5 HVFBV",
		"1 BZTQ => 9 CTHM",
		"16 SHTQF => 3 BZTQ",
		"11 PBCT, 5 PCBNG, 2 CTPND => 1 WBPFQ",
		"3 KHRW => 4 FSWFT",
		"12 HDQVB, 1 PBCT, 9 NRFK => 9 VLWJL",
		"5 SHTQF, 8 HVFBV => 6 BZTV",
		"2 KZNBL, 7 NRFK => 3 DVFS",
		"18 HTLSF, 14 DVFS => 6 TLFNL",
		"1 RSTBJ => 1 NKVSV",
		"2 QLNJ, 7 BZTQ => 6 PCBNG",
		"1 HTLSF, 19 CMPX => 7 JDJB",
		"6 KZNBL, 3 QSVJ => 8 SHTQF",
		"3 HTLSF, 1 VRVZ => 6 CMPX",
		"1 MGKSL, 15 CTPND => 6 STNPH",
		"2 NKVSV, 7 JDJB => 4 KXZTS",
		"3 KVPH => 4 QSVJ",
		"1 HPRK, 9 PCBNG, 2 KXFTL => 9 CNVPB",
		"27 GZWFN, 1 VLWJL, 15 LSCL => 3 GLVHT",
		"162 ORE => 4 HTLSF",
		"193 ORE => 8 PVHPV",
		"9 TLFNL, 1 KHRW => 6 HDQVB",
		"6 QLJXM, 4 FCBQN => 7 JLXP",
		"3 HTLSF, 21 NSQHW, 18 GVJVK => 7 BCRHK",
		"1 HTGMX, 20 CMPX, 6 RSTBJ => 6 FPVH",
		"4 KXZTS, 7 CNVPB, 1 STNPH => 2 LSCL",
		"3 KXZTS, 1 PCBNG => 3 JZBFT",
		"22 WBPFQ, 22 FRQH, 1 QLNJ, 4 CTHM, 3 GVMFW, 1 KMFWH, 4 QKBQL => 4 WFPNP",
		"3 QLJXM, 11 FNWBZ, 3 WBPFQ => 5 QGNL",
	}

	second(input)
}

func first(input []string) {
	fmt.Printf("%+v", parse(input))

	got := map[string]int{}
	next(parse(input), "FUEL", got)

	fmt.Println(got["ORE"])
}

func second(input []string) {

	pinput := parse(input)
	available := 1000000000000

	got := map[string]int{}
	next(pinput, "FUEL", got)

	orerequired := got["ORE"]

	fuelstupid := available / orerequired
	oreleftover := available % orerequired

	for k, v := range got {
		got[k] = v * fuelstupid
	}
	got["ORE"] = oreleftover
	got["FUEL"] = 0

	fmt.Printf("GOT: %+v \t fuelstupid: %d\n", got, fuelstupid)

	for k := range got {
		invert(pinput, k, got)
	}
	fmt.Printf("GOT: %+v \t fuelstupid: %d\n", got, fuelstupid)

	steps := []int{1000, 100, 10, 1}

	for _, step := range steps {
		for next2(pinput, "FUEL", step, got) {
			fuelstupid += step
			got["FUEL"] = 0
			//fmt.Printf("GOT: %+v \t fuelstupid: %d\n", got, fuelstupid)
			fmt.Printf("ORE: %d\t FUEL: %d\n", got["ORE"], fuelstupid)
		}
		fmt.Printf("GOT: %+v \t fuelstupid: %d\n", got, fuelstupid)
		for k := range got {
			invert(pinput, k, got)
		}
	}
	fmt.Printf("GOT: %+v \t fuelstupid: %d\n", got, fuelstupid)

}

func invert(eqs map[string]map[string]int, current string, got map[string]int) {
	if current == "ORE" {
		return
	}
	//fmt.Println(current)

	requirements := eqs[current]

	for got[current] >= requirements[current] {

		a := got[current] / requirements[current]

		for req, amount := range requirements {
			if req == current {
				got[req] -= a * amount
				continue
			}

			got[req] += a * amount

			invert(eqs, req, got)
		}
	}
}

func next(eqs map[string]map[string]int, current string, got map[string]int) {
	//fmt.Printf("CURRENT: %s \t GOT: %+v \n", current, got)

	requirements := eqs[current]

	for req, amount := range requirements {
		if req == current {
			got[req] += amount
			continue
		}

		if req == "ORE" {
			got["ORE"] += amount
			continue
		}

		for got[req] < amount {
			next(eqs, req, got)
		}

		got[req] -= amount
	}
}

func next2(eqs map[string]map[string]int, current string, want int, got map[string]int) bool {

	requirements := eqs[current]

	needed := want - got[current]
	reactionNeeded := (needed / requirements[current])
	if needed%requirements[current] != 0 {
		reactionNeeded++
	}

	//fmt.Printf("CURRENT: %s \t GOT: %+v \t WANT %d \t REACTION %d\n", current, got, want, reactionNeeded)

	for req, amount := range requirements {
		if req == current {
			continue
		}

		if req == "ORE" {
			if got["ORE"] < reactionNeeded*amount {
				return false
			}
		}

		for got[req] < reactionNeeded*amount {
			if !next2(eqs, req, reactionNeeded*amount, got) {
				return false
			}
		}

	}

	for req, amount := range requirements {
		if req == current {
			continue
		}

		got[req] -= reactionNeeded * amount
	}

	got[current] += reactionNeeded * requirements[current]

	return true
}

func parse(input []string) map[string]map[string]int {
	parsed := map[string]map[string]int{}

	for _, eq := range input {
		entry := map[string]int{}
		sides := strings.Split(eq, " => ")

		for _, r := range strings.Split(sides[0], ", ") {
			n := strings.Split(r, " ")
			needed, _ := strconv.Atoi(n[0])
			entry[n[1]] = needed
		}

		n := strings.Split(sides[1], " ")
		needed, _ := strconv.Atoi(n[0])
		entry[n[1]] = needed

		parsed[n[1]] = entry
	}

	return parsed
}
