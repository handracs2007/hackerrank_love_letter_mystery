// https://www.hackerrank.com/challenges/the-love-letter-mystery/problem

package main

import (
	"fmt"
	"math"
)

func theLoveLetterMystery(s string) int32 {
	var stepCount = int32(0)

	for i := 0; i < len(s)/2; i++ {
		var diff = int32(math.Abs(float64(int(s[i]) - int(s[len(s)-i-1]))))
		stepCount += diff
	}

	return stepCount
}

func main() {
	fmt.Println(theLoveLetterMystery("abc"))
	fmt.Println(theLoveLetterMystery("abcba"))
	fmt.Println(theLoveLetterMystery("abcd"))
	fmt.Println(theLoveLetterMystery("cba"))
}
