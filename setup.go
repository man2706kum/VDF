package main
import (
	"crypto/rsa"
	"crypto/rand"
	"math/big"
	"fmt"
)

func setup( n int , T int ) *big.Int {
	key, _ := rsa.GenerateKey(rand.Reader, n)
	pub := key.PublicKey.N
	//fmt.Println("\nPrime0 ",key.Primes[0])

	//fmt.Println("\n\nPrime1 ",key.Primes[1])
	return pub
}

func eval(N *big.Int, x *big.Int , l *big.Int , T int) (big.Int, big.Int) {

	var result big.Int
	result = *x
	for i := 1; i <= T; i++ {
		result.Exp(&result, big.NewInt(2), N)
	}

	return result, result
}

func main()  {

	//Setup phase

	var lambda, T int
	lambda = 1024
	T = 50
	fmt.Println("Public Parameter: N")
	N := setup(lambda, T)
	fmt.Println(N)
	fmt.Println(eval(N, big.NewInt(3), big.NewInt(3), 50))
	
}