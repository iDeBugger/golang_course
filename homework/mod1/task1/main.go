package main

func ShiftSlice(s []int) []int {
	for i := 0; i < len(s) - 1; i++ {
		s[i], s[i + 1] = s[i + 1], s[i]
	}
	return s
}

func ReverseSlice(s []int) []int {
	for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
	return s
}

func main() {

}
