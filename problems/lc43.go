package problems

import "math/big"

func multiply(num1 string, num2 string) string {
	var b1, b2 big.Int
	b1.SetString(num1, 10)
	b2.SetString(num2, 10)
	var product big.Int
	product.Mul(&b1, &b2)
	return product.String()
}
