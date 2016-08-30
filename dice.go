package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// pickRandom generate []int for 1~n, and take m elements from it
func pickRandom(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		if rand.Intn(i) == 0 {
			res = i
		}
	}
	return res
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var num, dice int
	var result []int
	if len(os.Args) < 2 {
		num = 1
		dice = 6
	} else {
		args := os.Args[1]
		fmt.Sscanf(args, "%dd%d", &num, &dice)
	}
	for i := 1; i <= num; i++ {
		result = append(result, pickRandom(dice))
	}
	// result := pickRandom(dice, num)
	fmt.Printf("You roll %d dice and %v\n", num, result)

}
