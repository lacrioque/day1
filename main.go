package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getData() string {
	file, err1 := os.Open("./coords.txt")
	if err1 != nil {
		fmt.Printf("error making http request: %s\n", err1)
		os.Exit(1)
	}

	scn := bufio.NewReader(file)

	cnt, err2 := io.ReadAll(scn)

	if err2 != nil {
		fmt.Printf("error reading http request: %s\n", err2)
		os.Exit(1)
	}

	fmt.Printf("File read, taken %d bytes", len(cnt))
	fmt.Println()
	return string(cnt)
}

func getnMapValue(s string) int32 {
	nMap := make(map[string]int)
	nMap[`nine`] = 9
	nMap[`9`] = 9
	nMap[`eight`] = 8
	nMap[`8`] = 8
	nMap[`seven`] = 7
	nMap[`7`] = 7
	nMap[`six`] = 6
	nMap[`6`] = 6
	nMap[`five`] = 5
	nMap[`5`] = 5
	nMap[`four`] = 4
	nMap[`4`] = 4
	nMap[`three`] = 3
	nMap[`3`] = 3
	nMap[`two`] = 2
	nMap[`2`] = 2
	nMap[`one`] = 1
	nMap[`1`] = 1

	return int32(nMap[s])
}

func getFirstNumberFromLine(line string) int32 {
	decimalRx := regexp.MustCompile(`\d|nine|eight|seven|six|five|four|three|two|one`)
	first := decimalRx.FindString(line)

	if first == "" {
		return 0
	}
	return getnMapValue(first)
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func getLastNumberFromLine(line string) int32 {
	decimalRx := regexp.MustCompile(`\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)
	last := decimalRx.FindString(Reverse(line))

	if last == "" {
		return 0
	}
	return getnMapValue(Reverse(last))
}

func lnToNumber(line string) int32 {
	tryn, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("[ERROR] error turning line to number: %s\n", err)
		fmt.Println()
		return 0
	}
	return int32(tryn)
}

func sum(values []int32) int32 {
	sum := int32(0)
	for _, v := range values {
		sum = sum + v
	}
	return sum
}

func main() {
	scrambledCoords := getData()
	lines := strings.Split(scrambledCoords, "\n")

	var numbers []int32
	for n, line := range lines {
		first := getFirstNumberFromLine(line)
		last := getLastNumberFromLine(line)
		coords := fmt.Sprintf("%d%d", first, last)
		fmt.Printf("Coords in line %d: (%s) -> %s \n", n, line, coords)
		fmt.Printf("Adding %s to %d\n", coords, sum(numbers))
		numbers = append(numbers, lnToNumber(coords))
	}

	fmt.Printf("Sum of all coords: %d", sum(numbers))
	fmt.Println()

}
