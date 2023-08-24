package utils

import (
	"crypto/rand"
	"math/big"
)

// Generate a safe prime number of bit size n (i.e. p = 2q + 1, where p and q are primes)
func GenerateSafePrime(n int) (*big.Int, error) {
	var p, q *big.Int
	var err error
	for {
		q, err = rand.Prime(rand.Reader, n-1)
		if err != nil {
			return nil, err
		}
		p = new(big.Int).Mul(q, big.NewInt(2))
		p.Add(p, big.NewInt(1))
		//ProbablyPrime报告x是否可能是质数，应用带有n个伪随机选择的碱基的Miller-Rabin检验以及Baillie-PSW检验。
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}
