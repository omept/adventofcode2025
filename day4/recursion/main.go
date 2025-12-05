/*
--- Day 4: Printing Department ---
You ride the escalator down to the printing department. They're clearly getting ready for Christmas; they have lots of large rolls of paper everywhere, and there's even a massive printer in the corner (to handle the really big print jobs).

Decorating here will be easy: they can make their own decorations. What you really need is a way to get further into the North Pole base while the elevators are offline.

"Actually, maybe we can help with that," one of the Elves replies when you ask for help. "We're pretty sure there's a cafeteria on the other side of the back wall. If we could break through the wall, you'd be able to keep moving. It's too bad all of our forklifts are so busy moving those big rolls of paper around."

If you can optimize the work the forklifts are doing, maybe they would have time to spare to break through the wall.

The rolls of paper (@) are arranged on a large grid; the Elves even have a helpful diagram (your puzzle input) indicating where everything is located.

For example:

..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions. If you can figure out which rolls of paper the forklifts can access, they'll spend less time looking and more time breaking down the wall to the cafeteria.

In this example, there are 13 rolls of paper that can be accessed by a forklift (marked with x):

..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.
Consider your complete diagram of the paper roll locations. How many rolls of paper can be accessed by a forklift?

THIS EXAMPLE DOESN'T USE UNICODE-SAFE [][]rune grid. It's faster but it assume that the inputs are always ASCII
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, line)
	}

	h := len(grid)
	w := len(grid[0])

	memo := make([][]int, h)

	for i := range memo {
		memo[i] = make([]int, w)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// 8 directions
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	countNeighbors := func(r, c int) int {
		// memoized already?
		if memo[r][c] != -1 {
			return memo[r][c]
		}

		// only cells that are '@' matter
		if grid[r][c] != '@' {
			memo[r][c] = 0
			return 0
		}

		count := 0

		// recursively check each direction
		for _, d := range dirs {
			nr := r + d[0]
			nc := c + d[1]

			if nr < 0 || nr >= h || nc < 0 || nc >= w {
				continue
			}

			if grid[nr][nc] == '@' {
				// neighbor is a roll -- count it
				count++
			}
		}

		memo[r][c] = count
		return count
	}

	accessible := 0

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == '@' {
				if countNeighbors(r, c) < 4 {
					accessible++
				}
			}
		}
	}
	fmt.Println("accessible: ", accessible)
}
