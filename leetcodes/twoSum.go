import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	dictionary := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if _, found := dictionary[diff]; found {
			fmt.Println("inside here")
			fmt.Println(found)
			return []int{i, dictionary[diff]}
		} else {
			dictionary[n] = i
		}
	}
	return nil
}