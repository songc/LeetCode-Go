package main

import "fmt"

func main() {
	s := "000011100"
	fmt.Println(minFlipsMonoIncr(s))

}

func minFlipsMonoIncr(s string) int {
	ans := 0
	num0, num1 := 0, 0
	for _, c := range s {
		if c == '0' && num1 > 0 {
			num0 += 1
			if num0 >= num1 {
				ans += num1
				num0, num1 = 0, 0
			}
		}
		if c == '1' {
			num1 += 1
		}
	}
	ans += num0
	return ans
}
