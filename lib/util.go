package lib

import (
	"math"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func ParseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func ForEach[T any](arr []T, f func(T)) {
	for i := 0; i < len(arr); i++ {
		f(arr[i])
	}
}

func Map[T any, U any](arr []T, f func(T) U) (newArr []U) {
	newArr = make([]U, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = f(arr[i])
	}
	return newArr
}

func Filter[T any](arr []T, f func(T) bool) (newArr []T) {
	for i := 0; i < len(arr); i++ {
		if f(arr[i]) {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
