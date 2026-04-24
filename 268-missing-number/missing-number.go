func missingNumber(nums []int) int {
    set := make(map[int]struct{})

    // store all numbers
    for _, num := range nums {
        set[num] = struct{}{}
    }

    // check from 0 to n
    for i := 0; i <= len(nums); i++ {
        if _, exists := set[i]; !exists {
            return i
        }
    }

    return -1
}