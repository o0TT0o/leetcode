//"aaccc"
//"ccc"
//"aaaa"
//"babad"
//"cbbd"
//"a"
//"ac"
//"bb"
func longestPalindrome(s string) string {
    result := s[0:1]
    resLen := 1
    length := len(s)
    charArr := strings.Split(s, "")
    
    for pos, char := range(s) {
        leftIndex := pos -1 
        rightIndex := pos +1
        seqFlag := true
        
        for {
            
            for seqFlag && leftIndex >= 0 && charArr[leftIndex] == string(char) {
                leftIndex--
            }
            seqFlag = false
            
            if leftIndex >= 0 && rightIndex <length && charArr[leftIndex] == charArr[rightIndex] {
                leftIndex--
                rightIndex++
            }else{
                leftIndex++
                break
            }
        }   
        
        if ( rightIndex - leftIndex ) > resLen {
            //fmt.Println("tmpStr: ", tmpStr)
            resLen = rightIndex - leftIndex
            result = s[leftIndex:rightIndex]
        }
    }
    
    return result
}