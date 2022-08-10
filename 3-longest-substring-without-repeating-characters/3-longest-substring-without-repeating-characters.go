func lengthOfLongestSubstring(s string) int {
    //test cases: 
    //  "dvdf" 
    longest := 0
    startIndex := 0 
    endIndex := 0
    collection := make(map[rune]int)
    for pos, char := range(s) {
        if val, ok := collection[char]; ok{
            //fmt.Println(collection)
            //fmt.Println(startIndex, endIndex)
            distance := endIndex - startIndex 
            //fmt.Println(distance)
            if distance > longest{
                //fmt.Println(longest , "------------>" , distance)
                longest = distance
            }
            
            newStart := val + 1
            if newStart > startIndex{
                startIndex = newStart
            }
            
            delete(collection,char)
        }
        
        collection[char] = pos
        endIndex++
    }
    distance := endIndex - startIndex 
    if distance > longest{
        //fmt.Println(longest , "+++++++>" , distance)
        longest = distance
    }
    return longest
}