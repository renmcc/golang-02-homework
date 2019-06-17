package qiepian

func Reverse(slice []int) []int {
	start := 0
	end := len(slice) - 1
	for {
		if start > end {
			break
		}
		slice[start],slice[end] = slice[end],slice[start]
		start++
		end--
	}
	return slice
}