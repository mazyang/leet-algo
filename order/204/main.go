package main

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	res := 0
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			res++
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
