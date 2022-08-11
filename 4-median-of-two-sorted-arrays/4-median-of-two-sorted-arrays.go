//  
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var result float64
    result = 0
    loopLen := len(nums1) + len(nums2)
    
    index1 := 0
    index2 := 0
    mergedArr := []int{}
    
    
    for i := 0 ; i< loopLen; i++ {
        
        if index1 >= len(nums1){
            for index2 < len(nums2){
                mergedArr = append(mergedArr, nums2[index2])
                index2++
            }
            break
        }
        
        if index2 >= len(nums2){
            for index1 < len(nums1){
                mergedArr = append(mergedArr, nums1[index1])
                index1++
            }       
            break
        }
        
        
        if nums1[index1] < nums2[index2] {
            mergedArr = append(mergedArr, nums1[index1])
            index1++
        }else{
            mergedArr = append(mergedArr, nums2[index2])
            index2++
        }
        
    }
    
    //fmt.Println(mergedArr) 
    //fmt.Println(math.Mod(float64(loopLen),float64(2)))
    //fmt.Println(loopLen/2)
    if  math.Mod(float64(loopLen),float64(2)) == 0{
        result = float64(mergedArr[loopLen/2-1] + mergedArr[loopLen/2]) /2
    }else{
        result = float64(mergedArr[loopLen/2])
    }
    
    
    return result
}