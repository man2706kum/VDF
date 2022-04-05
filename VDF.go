package main
import (
	"crypto/rsa"
	"crypto/rand"
	"math/big"
	"fmt"
)
/*
	The setup function takes input a security parameter n & Time bound T and
	generates the Public Parameters. Here we are generating N which is the product of two large primes
*/
func setup( n int , T int ) *big.Int {
	key, _ := rsa.GenerateKey(rand.Reader, n)
	pub := key.PublicKey.N
	//fmt.Println("\nPrime0 ",key.Primes[0])
	//fmt.Println("\n\nPrime1 ",key.Primes[1])
	return pub
}
/*
	The eval function evaluates the VDF on the input x and returns the result and the proof of that result
*/

func eval(N *big.Int, x *big.Int , l *big.Int , T int) (big.Int, big.Int) {

	var result, q, r, product, proof big.Int
	result = *x
	//evaluation of VDF
	for i := 1; i <= T; i++ {
		result.Exp(&result, big.NewInt(2), N)	
	}
	//Generation of proof
	r.Exp(big.NewInt(2), big.NewInt(int64(T)), l)
	//fmt.Println("r:", r.String())
	product.Exp(big.NewInt(2), big.NewInt(int64(T)), big.NewInt(0))
	//fmt.Println("product:", product.String())
	q.Div(&product,l)
	//fmt.Println("q:", q.String())
	proof.Exp(x,&q,N)
	//fmt.Println("proof:", proof.String())


	return result, proof
}

/*
	The verify function takes input the result and the proof of it and gives a boolean output depending on whether the result of VDF is correct or not.
*/
func verify(proof big.Int, T int, l *big.Int, result *big.Int, x *big.Int, N *big.Int) bool {
	var val, r big.Int
	r.Exp(big.NewInt(2), big.NewInt(int64(T)), l)
	val.Exp(x,&r,N)
	proof.Exp(&proof, l, N)
	val.Mul(&val, &proof)
	val.Mod(&val, N)

	if val.Cmp(result)==0 {
		return true
	}
		
	return false
}

func main()  {

	//Setup phase
	var lambda, T int
	lambda = 1024	//security parameter
	T = 50	// time bound
	x := big.NewInt(5)	//input to VDF
	//fmt.Println("Public Parameter: N")
	N := setup(lambda, T)
	fmt.Println("-------------------------------Setup Phase---------------------------------------\n")
	fmt.Println("Public Parameter: N(=pq) ",N)

	//verifier chooses randomly a prime l
	l,_ := rand.Prime(rand.Reader, 25)
	fmt.Println("\nRandom Prime l chosen by verifier: ",l)
	//fmt.Println("result:")
	fmt.Println("\n---------------------------Evaluation phase--------------------------------------\n")
	//The evaluator generates the result & its proof
	result, proof := eval(N, x, l, 50)
	fmt.Println("result:", result.String())
	fmt.Println("\nproof:", proof.String())
	fmt.Println("--------------------------Verification Phase-------------------------------------")
	//The verifier checks the proof of result
	fmt.Println("\n",verify(proof,T,l,&result,x,N))

}