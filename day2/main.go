package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var pow10 = [...]uint64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
	10000000000000000000,
}

func main() {

	input := "61-71,12004923-12218173,907895-1086340,61083-74975,7676687127-7676868552,3328-4003,48-59,3826934-3859467,178-235,75491066-75643554,92-115,1487-1860,483139-586979,553489051-553589200,645895-722188,47720238-47818286,152157-192571,9797877401-9798014942,9326-11828,879837-904029,4347588-4499393,17-30,1-16,109218-145341,45794-60133,491-643,2155-2882,7576546102-7576769724,4104-5014,34-46,67594702-67751934,8541532888-8541668837,72-87,346340-480731,3358258808-3358456067,78265-98021,7969-9161,19293-27371,5143721-5316417,5641-7190,28793-36935,3232255123-3232366239,706-847,204915-242531,851-1135,790317-858666"
	ranges := parseInput(input)
	var total int64 = runRanges(ranges)

	fmt.Println(total)

}

func isRepeatedTwice(n int64) bool {
	if n < 10 {
		return false
	} else if n < 100 {
		x := n / 10
		y := n % 10
		return x == y
	}
	d := digitCount64(uint64(n))
	if d%2 != 0 {
		return false
	}
	newN := uint64(n)

	half := d / 2
	divisor := pow10[half]
	left := newN / divisor
	right := newN % divisor

	return left == right
}

func digitCount64(n uint64) int {
	switch {
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 1000:
		return 3
	case n < 10000:
		return 4
	case n < 100000:
		return 5
	case n < 1000000:
		return 6
	case n < 10000000:
		return 7
	case n < 100000000:
		return 8
	case n < 1000000000:
		return 9
	case n < 10000000000:
		return 10
	case n < 100000000000:
		return 11
	case n < 1000000000000:
		return 12
	case n < 10000000000000:
		return 13
	case n < 100000000000000:
		return 14
	case n < 1000000000000000:
		return 15
	case n < 10000000000000000:
		return 16
	case n < 100000000000000000:
		return 17
	case n < 1000000000000000000:
		return 18
	default:
		return 19
	}
}

func parseRange(r string) (int64, int64) {
	parts := strings.Split(r, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)
	return start, end
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}

func runRanges(ranges []string) int64 {
	sumChan := make(chan int64)
	var wg sync.WaitGroup

	for _, r := range ranges {
		if r == "" {
			continue
		}
		wg.Add(1)

		go func(r string) {
			defer wg.Done()
			start, end := parseRange(r)
			var total int64
			for n := start; n <= end; n++ {
				if isRepeated := isRepeatedTwice(n); isRepeated {
					total += n
				}
			}
			sumChan <- total
		}(r)
	}
	go func() {
		wg.Wait()
		close(sumChan)
	}()

	var total int64
	for s := range sumChan {
		total += s
	}

	return total
}
