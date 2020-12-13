package main

import "fmt"

func main() {
	// when slice length less than capacity
	nums := make([]int, 3, 5)
	nums[0], nums[1], nums[2] = 1, 2, 3 
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))

	// append num 4
	nums = append(nums, 4)
	fmt.Println("append num 4: ")
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))

	// append num 5
	nums = append(nums, 5)

	fmt.Println("append num 5: ")
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("length: %d\n", len(nums))
	fmt.Printf("capacity: %d\n", cap(nums))


	// append function for string
	fmt.Println("----------------------")
	str := "holo"
	suffix := "lllslsl"

	res := append([]byte(str), suffix...)
	fmt.Printf("res: %v\n", res)

}