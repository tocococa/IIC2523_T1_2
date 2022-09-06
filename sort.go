package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Ingresa una lista de numeros separados por espacios")
	var input string
	fmt.Scanln(&input)
	//input fijo para el debugger:
	//input = "2 5 7 2 3 5 67 0 12 67 23 89 32 12 76"
	split_input := strings.Split(input, " ")
	arr := make([]int, len(split_input))
	for i := 0; i < len(split_input); i++ {
		arr[i], _ = strconv.Atoi(split_input[i])
	}
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)
	go sortRoutine(arr[:len(arr)/4], c1)
	go sortRoutine(arr[len(arr)/4:2*len(arr)/4], c2)
	go sortRoutine(arr[2*len(arr)/4:3*len(arr)/4], c3)
	go sortRoutine(arr[3*len(arr)/4:], c4)
	arr1 := <-c1
	arr2 := <-c2
	arr3 := <-c3
	arr4 := <-c4
	go mergeRoutine(arr1, arr2, c1)
	go mergeRoutine(arr3, arr4, c2)
	arr1 = <-c1
	arr2 = <-c2
	go mergeRoutine(arr1, arr2, c1)
	arr1 = <-c1
	fmt.Println(arr1)
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