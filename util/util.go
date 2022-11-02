package util

import (
	"strconv"
)

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func RemoveArray(exluce []int, arr []int) int {
loop:
	for i := 0; i < len(arr); i++ {
		url := arr[i]
		for _, rem := range exluce {
			if url == rem {
				arr = append(arr[:i], arr[i+1:]...)
				i-- // Important: decrease index
				continue loop
			}
		}
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return SliceToInt(arr)

}

func SliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}

	return res
}

func RemoveMultipleElement(result []int) []int {
	keys := make(map[int]bool)

	list := []int{}

	for _, entry := range result {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ListToArray(nums int) []int {
	str := strconv.Itoa(nums)
	chars := []rune(str)
	result := []int{}
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		temp, _ := strconv.Atoi(char)
		result = append(result, temp)

	}
	return result
}
