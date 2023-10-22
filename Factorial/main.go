package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number for factorial value: ")
	scanner.Scan()
	s, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print(err, "\n")
	} else {
		recursionResult, err := recursion(s)
		if err != nil {
			fmt.Print("recursionResult error: ", err, "\n")
		}
		nonRecursionResult, err := nonrecursion(s)
		if err != nil {
			fmt.Print("nonRecursionResult error: ", err, "\n")
		}
		if nonRecursionResult != big.NewInt(0) && recursionResult != big.NewInt(0) && recursionResult.Cmp(nonRecursionResult) == 0 {
			fmt.Print("Result: ", recursionResult, "\n")
		} else {
			fmt.Print("Result are not equal", "\n")
		}
	}
}

func recursion(factorialNumber int) (*big.Int, error) {
	if factorialNumber < 0 {
		return big.NewInt(0), errors.New("undefined number")
	}
	if factorialNumber == 0 {
		return big.NewInt(0), nil
	}
	if factorialNumber == 1 {
		return big.NewInt(1), nil
	} else {
		recursionResult, err := recursion(factorialNumber - 1)
		if err != nil {
			fmt.Print("recursionResult error: ", err, "\n")
		}
		return recursionResult.Mul(recursionResult, big.NewInt(int64(factorialNumber))), nil
	}
}

func nonrecursion(factorialNumber int) (*big.Int, error) {
	if factorialNumber < 0 {
		return big.NewInt(0), errors.New("undefined number")
	}
	if factorialNumber == 0 {
		return big.NewInt(0), nil
	}
	var result = big.NewInt(1)
	for i := 1; i <= factorialNumber; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result, nil
}
