package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// pickRandom generate []int for 1~n, and take m elements from it
func pickRandom(n, m int) []int {
	res := make([]int, m)
	for i := 1; i <= n; i++ {
		if j := rand.Intn(i); j < m {
			res[i%m] = i
		}
	}
	return res
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var num, dice int
	if len(os.Args) < 2 {
		num = 1
		dice = 6
	} else {
		args := os.Args[1]
		fmt.Sscanf(args, "%dd%d", &num, &dice)
	}
	result := pickRandom(dice, num)
	fmt.Printf("You roll %d dice and %v\n", num, result)

}
