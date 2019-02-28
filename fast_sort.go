package main

import "fmt"

var sort_array [10]int


func main() {
	sort_array = [10]int{10,7,25,99,1,8,50,85,26,12}
	fmt.Println(sort_array)
	quick_sort(0, len(sort_array)-1)
	fmt.Println(sort_array)
}

func quick_sort(left int, right int) {
	var i,j,t,tmp int
	if (left > right) {
		return
	}
	tmp = sort_array[left]
	i = left
	j = right
	for i!=j {
		for sort_array[j] >= tmp && i< j {
			j--
		}
		for sort_array[i] <=tmp && i< j {
			i++
		}
		if i <j {
			t=sort_array[j]
			sort_array[j] = sort_array[i]
			sort_array[i] = t
		}
	}
	sort_array[left] = sort_array[i]
	sort_array[i] = tmp
	quick_sort(left, i-1)
	quick_sort(i+1, right)
}