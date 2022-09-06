package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main(){

}

func sortRoutine(arr []int, c chan []int) {
	quickSort(arr)
	c <- arr
}

func mergeRoutine(arr1 []int, arr2 []int, c chan []int) {
	arr := merge(arr1, arr2)
	c <- arr
}

func merge(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))
	i := 0
	j := 0
	for k := 0; k < len(arr); k++ {
		if i >= len(arr1) {
			arr[k] = arr2[j]
			j++
		} else if j >= len(arr2) {
			arr[k] = arr1[i]
			i++
		} else if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
	}
	return arr
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	quickSort(left)
	quickSort(right)
	copy(arr, append(append(left, pivot), right...))
}