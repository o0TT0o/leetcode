func twoSum(nums []int, target int) []int {
    tmp := make(map[int]int)
    result := []int{}
    for i,v := range nums {
        paired := target - v
        if val, ok := tmp[paired]; ok {
            result = append(result,val)
            result = append(result,i)
            return result
        }else{
            tmp[v] = i
        }
    }
    return result
}