
package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    for i, num := range nums {
        diff := target - num

        if j, found := numMap[diff]; found {
            return []int{j, i} // Return the indices of the two numbers
        }

        numMap[num] = i
    }

    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9

    result := twoSum(nums, target)
    if len(result) > 0 {
        fmt.Printf("Indices: %v\n", result)
    } else {
        fmt.Println("No two sum solution found.")
    }
}
