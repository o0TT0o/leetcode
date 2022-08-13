//"aaccc"
//"ccc"
//"aaaa"
//"babad"
//"cbbd"
//"a"
//"ac"
//"bb"
func longestPalindrome(s string) string {
    result := ""
    length := len(s)
    charArr := strings.Split(s, "")
    
    for pos, char := range(s) {
        tmpStr := string(char)
        leftIndex := pos -1 
        rightIndex := pos +1
        seqFlag := true
        
        for {
            
            for seqFlag && leftIndex >= 0 && charArr[leftIndex] == string(char) {
                tmpStr = string(charArr[leftIndex])+ tmpStr
                leftIndex--
            }
            seqFlag = false
            
            if leftIndex >= 0 && rightIndex <length && charArr[leftIndex] == charArr[rightIndex] {
                tmpStr = string(charArr[leftIndex])+ tmpStr+ string(charArr[rightIndex])
                leftIndex--
                rightIndex++
            }else{
                break
            }
        }   

        if len(tmpStr) > len(result) {
            //fmt.Println("tmpStr: ", tmpStr)
            result = tmpStr
        }
    }
    
    return result
}