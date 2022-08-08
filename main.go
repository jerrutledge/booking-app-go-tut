package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {

	centered_palindrome := func(center int, s string) string {
		longest := s[center : center+1]
		// two types of palindromes, odd and even
		// ODD
		iterrange := math.Min(float64(center), float64(len(s)-center-1))
		for i := 1; i <= int(iterrange); i++ { // odd, this character is the center
			if s[center-i:center-i+1] == s[center+i:center+i+1] {
				// still a palindrome
				longest = s[center-i : center+i+1]
			} else {
				// not a palindrome anymore
				break
			}
		}
		// EVEN
		iterrange = math.Min(float64(center), float64(len(s)-center))
		for i := 1; i <= int(iterrange); i++ {
			if s[center-i:center-i+1] == s[center+i-1:center+i] {
				// still a palindrome
				if len(s[center-i:center+i]) > len(longest) {
					longest = s[center-i : center+i]
				}
			} else {
				// not a palindrome anymore
				break
			}
		}
		return longest
	}

	// use a pool of workers to concurrently calculate palindromes
	worker := func(index int, s string, jobs <-chan int, results chan<- string) {
		for j := range jobs {
			results <- centered_palindrome(j, s)
		}
	}
	numJobs := len(s)
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, s, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	longest := ""
	for a := 1; a <= numJobs; a++ {
		str := <-results
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}

func main() {

	s := "bb"
	fmt.Println(longestPalindrome(s))
}
