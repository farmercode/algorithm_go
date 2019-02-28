//冒泡排序go的实现
//比较相邻的元素。 如果第一个比第二个大，就交换他们两个。
//对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。 在这一点，最后的元素应该会是最大的数。
//针对所有的元素重复以上的步骤，除了最后一个。
//持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

package main

import "fmt"

func main() {
	sort_array := [10]int{10,7,25,99,1,8,50,85,26,12}
	array_len := len(sort_array)
	for pos:=array_len-1; pos >= 0; pos--  {
		for i :=0; i < pos;  i++ {
			if sort_array[i] > sort_array[i+1] {
				tmp := sort_array[i+1]
				sort_array[i+1] = sort_array[i]
				sort_array[i] = tmp
			}
		}
	}
	fmt.Println(sort_array)
}
