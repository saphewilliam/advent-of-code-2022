package lib

import (
	"math"
	"strconv"
)

type Point3D struct {
	X, Y, Z int
}

func NewP3D(x, y, z int) Point3D {
	return Point3D{X: x, Y: y, Z: z}
}

type Size3D struct {
	MinX, MaxX, MinY, MaxY, MinZ, MaxZ int
}

func NewS3D(minX, maxX, minY, maxY, minZ, maxZ int) Size3D {
	return Size3D{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY, MinZ: minZ, MaxZ: maxZ}
}

func (size *Size3D) SetAll(p Point3D) {
	size.MinX = Min(size.MinX, p.X)
	size.MaxX = Max(size.MaxX, p.X)
	size.MinY = Min(size.MinY, p.Y)
	size.MaxY = Max(size.MaxY, p.Y)
	size.MinZ = Min(size.MinZ, p.Z)
	size.MaxZ = Max(size.MaxZ, p.Z)
}

type Point2D struct {
	X, Y int
}

func NewP2D(x, y int) Point2D {
	return Point2D{X: x, Y: y}
}

type Size2D struct {
	MinX, MaxX, MinY, MaxY int
}

func newS2D(minX, maxX, minY, maxY int) Size2D {
	return Size2D{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
}

type Size struct {
	Min, Max int
}

func ParseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Max(numbers ...int) (max int) {
	for i, n := range numbers {
		if i == 0 {
			max = n
		} else {
			max = int(math.Max(float64(max), float64(n)))
		}
	}
	return
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
