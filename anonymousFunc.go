package pkg

import (
	"fmt"
	"strconv"
)

//Функция на вход получает целое положительное число (uint),
//возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
//Если же получившееся число равно 0, то возвращается число 100.
//Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

func anomymousFunc() uint {
	var a uint
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	fn := func(num uint) uint {
		numStr := strconv.Itoa(int(num))
		newStr := ""
		for _, i := range numStr {
			if i%2 == 0 && i != 48 {
				newStr += string(i)
			}
		}
		if newStr != "" {
			numUint, err := strconv.ParseUint(newStr, 10, 64)
			if err != nil {
				panic(err)
			}
			if numUint == 0 {
				numUint = uint64(100)
			}
			return uint(numUint)
		} else {
			return uint(100)
		}
	}
	return fn(a)
}
