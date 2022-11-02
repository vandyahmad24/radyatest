package main

import (
	"errors"
	"fmt"
	"radya/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Soal1([]int{-1, 1})
	Soal1([]int{-1, -7, -5})
	Soal1([]int{1, 2, 1, 6})
	Soal2([]int{0, 1, 0}, 1259, 2601)
	Soal2([]int{0, 1, 0}, 1000, 1010)
	Soal3([]int{1, 2}, []int{1, 3})
	Soal3([]int{1, 2, 2}, []int{1, 2, 4})
	Soal4(1223334)
	Soal5([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})

}

func Soal1(arr []int) {
	tempPositive := []int{}
	tempNegative := []int{}
	result := 0
	for _, v := range arr {
		if v > 0 {
			tempPositive = append(tempPositive, v)
		} else {
			tempNegative = append(tempNegative, v)
		}
	}
	if len(tempNegative) > 0 && len(tempPositive) == 0 {
		result = -2
	}

	if len(tempPositive) > 0 && len(tempNegative) == 0 {
		result = 3
	}

	fmt.Println("Result nomer 1 :", result)

}

func Soal2(exclude []int, nums ...int) (int, error) {
	result := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		arrNum := util.IntToSlice(nums[i], []int{})
		resultRemove := util.RemoveArray(exclude, arrNum)
		if resultRemove != 0 {
			result = append(result, resultRemove)
		}

	}
	if len(result) == 0 {

		fmt.Println("Result Nomer 2 :", "not found")
		return 0, errors.New("not found")
	}

	a := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ""), "[]")
	hasil, _ := strconv.Atoi(a)
	fmt.Println("Result Nomer 2 :", hasil)
	return hasil, nil

}

func Soal3(numsA, numsB []int) []int {
	result := []int{}
	for _, v1 := range numsA {
		for _, v2 := range numsB {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}

	list := util.RemoveMultipleElement(result)

	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("Result Nomer 3 :", list)

	return list
}

func Soal4(num int) map[int]int {
	a := map[int]int{}

	listOfInt := util.ListToArray(num)

	for _, v := range listOfInt {
		if a[v] == 0 {
			a[v] = 1
		} else {
			a[v] = a[v] + 1
		}

	}
	fmt.Println("Result Nomer 4 :", a)

	return a

}

func Soal5(nums ...[]int) {
	// []int{1, 2, 3, 4, 5}
	for _, v := range nums {
		panjang := len(v) - 1

		for i, _ := range v {
			if i+1 > panjang {
				continue
			}

			fmt.Println(fmt.Sprintf("Perkalian %d * %d = %d", v[i], v[i+1], v[i]*v[i+1]))
			i += 1

			continue

		}
	}
}
