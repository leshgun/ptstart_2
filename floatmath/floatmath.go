package floatmath

import (
	"errors"
	"log"
	"math/big"
	"strconv"
)

func SumStringFloat(x string, y string) (float64, error) {
	var x_f, y_f float64
	var err error

	x_f, err = strconv.ParseFloat(x, 64)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	y_f, err = strconv.ParseFloat(y, 64)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return x_f + y_f, nil
}

func SumStringArrayFloat(arr []string, fatal bool) (float64, error) {
	var sum float64

	if len(arr) == 0 {
		return 0, nil
	}

	sum = 0
	for _, str := range arr {
		cur, err := strconv.ParseFloat(str, 64)
		if err == nil {
			sum += cur
			continue
		}
		log.Fatal(err)
		if true {
			return sum, err
		}
	}

	return sum, nil
}

func SumStringArrayBigFloat(arr []string, precision int) (string, error) {
	var cur, sum *big.Float
	var err error

	if len(arr) == 0 {
		return "0", nil
	}
	if len(arr) == 1 {
		return arr[0], nil
	}

	sum = big.NewFloat(0.0).SetPrec(uint(precision * 4))
	for i := 0; i < len(arr); i++ {
		cur, err = ParseFloat(arr[i], precision)
		if err != nil {
			return BigFloatToString(sum, precision), err
		}
		sum.Add(sum, cur)
	}

	return BigFloatToString(sum, precision), nil
}

func SumStringBigFloat(x string, y string, precision int) (string, error) {
	var bx, by *big.Float
	var err error
	bx, err = ParseFloat(x, precision)
	if err != nil {
		return "", err
	}
	by, err = ParseFloat(y, precision)
	if err != nil {
		return "", err
	}
	bx.Add(bx, by)
	return bx.Text('f', precision), nil
}

func MulStringBigFloat(x string, y string, precision int) (string, error) {
	var bx, by *big.Float
	var err error
	bx, err = ParseFloat(x, precision)
	if err != nil {
		return "", err
	}
	by, err = ParseFloat(y, precision)
	if err != nil {
		return "", err
	}
	bx.Mul(bx, by)
	return BigFloatToString(bx, precision), nil
}

func DivStringBigFloat(x string, y string, precision int) (string, error) {
	var bx, by *big.Float
	var err error
	bx, err = ParseFloat(x, precision)
	if err != nil {
		return "", err
	}
	by, err = ParseFloat(y, precision)
	if err != nil {
		return "", err
	}
	bx.Quo(bx, by)
	return BigFloatToString(bx, precision), nil
}

func ParseFloat(x string, precision int) (*big.Float, error) {
	// For big.Float precision is the bits
	// Because 10^3 = 2^10, then:
	// precision_bits ~= precision * 4
	precision_bits := uint(precision * 4)
	bx, err := big.NewFloat(0.0).SetPrec(precision_bits).SetString(x)
	if !err {
		return big.NewFloat(0), errors.New("cant parse X")
	}
	return bx, nil
}

func BigFloatToString(bf *big.Float, precision int) string {
	return bf.Text('f', precision)
}
