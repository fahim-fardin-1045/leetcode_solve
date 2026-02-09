package main

import (
	"fmt"
	"math"
)

func reverse(x int) (ans int) {
	for ; x != 0; x /= 10 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		ans = ans*10 + x%10
	}
	return
}

func main() {
	fmt.Println(reverse(123))        // 321
	fmt.Println(reverse(-123))       // -321
	fmt.Println(reverse(120))        // 21
	fmt.Println(reverse(1534236469)) // 0 (overflow)
}
