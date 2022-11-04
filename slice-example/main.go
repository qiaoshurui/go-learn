package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	// 7 是容量到索引7(开区间,真正到索引6)

	fmt.Println(len(s2), cap(s2))

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)

	fmt.Println(s2)

	fmt.Println(slice)

	add(slice)
	fmt.Println(slice)
	tmp := slice[2:5:7]

	add(tmp)
	tmp = append(tmp, 100)
	add(tmp)
	fmt.Println(tmp[2:4])

	fmt.Println(slice)

}

func add(arr []int) {
	arr = append(arr, 1)
}
