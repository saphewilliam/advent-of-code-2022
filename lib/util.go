package lib

import "strconv"

func UnsafeParseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Map[T any, U any](arr []T, f func(T) U) []U {
	newArr := make([]U, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = f(arr[i])
	}
	return newArr
}
