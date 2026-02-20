package main

import(
	"fmt"
)

func main(){
	nums := []int{2, 7, 11, 15}
	target := 9
	
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}


// func twoSum(nums []int, target int) []int {
// 	seen := make(map[int]int)

// 	for i, num := range nums {
// 		complement := target - num

// 		if index, found := seen[complement]; found {
// 			return []int{index, i}
// 		}

// 		seen[num] = i
// 	}
// 	return nil
// }


func twoSum(nums []int, target int) []int{
	for i := 0; i < len(nums); i++{
		for j := i+1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return nil
}