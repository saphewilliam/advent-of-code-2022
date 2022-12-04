package day02

const Path = "./day02/input.txt"

func Calculate(input []string) (score1 int, score2 int) {
	for _, v := range input {
		a, b := int(v[0])-65, int(v[2])-88
		score1 += 1 + b + 3*((4+b-a)%3)
		score2 += 1 + (a+b+2)%3 + b*3
	}
	return
}
