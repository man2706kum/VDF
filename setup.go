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


func main()  {

	//Setup phase

	var lambda, T int
	lambda = 1024
	T = 50
	fmt.Println("\n\nN ",setup(lambda, T))
	
}