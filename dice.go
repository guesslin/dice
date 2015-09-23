package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var num, dice int32
	var result = make([]int32, 0)
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) < 2 {
		num = 1
		dice = 6
	} else {
		args := os.Args[1]
		fmt.Sscanf(args, "%dd%d", &num, &dice)
	}
	for i := int32(0); i < num; i++ {
		result = append(result, rand.Int31n(dice)+1)
	}
	fmt.Printf("You roll %d dice and %v\n", num, result)

}
