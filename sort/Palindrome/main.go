package main

import (
	"fmt"
	"sort"
)

type Palindrome []byte

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}


func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "hey"
	fmt.Printf("%s is a palindrome? %v\n", s, IsPalindrome(Palindrome([]byte(s))))

	s = "civic"
	fmt.Printf("%s is a palindrome? %v\n", s, IsPalindrome(Palindrome([]byte(s))))

	s = "anna"
	fmt.Printf("%s is a palindrome? %v\n", s, IsPalindrome(Palindrome([]byte(s))))

}
