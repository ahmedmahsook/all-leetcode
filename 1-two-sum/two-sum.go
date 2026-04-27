func twoSum(nums []int, target int) []int {
maps:=make(map[int]int)

for i,num:= range nums{
    complement:=target-num
    if idx,found:=maps[complement];found{
        return []int{idx,i}
    }
    maps[num]=i
}
return nil
}