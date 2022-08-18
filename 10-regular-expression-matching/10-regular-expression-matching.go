//"aa"
//"a"
//"aa"
//"a*"
//"ab"
//".*"
//"aabbcc"
//"a.*b.*c.*"
//"aabbccd"
//"a.*b.*c.*d"
//"aab"
//"c*a*b"
//"mississippi"
//"mis*is*p*."
//"ab"
//".*c"
//"a"
//"ab*"
//"aabcbcbcaccbcaabc"
//".*a*aa*.*b*.c*.*a*"
//"bbab"
//"b*a*"
//"bbbba"
//".*a*a"
//"aaaaaaaaaaaaab"
//"a*a*a*a*a*a*a*a*a*a*a*a*b"
func isMatch(s string, p string) bool {
    cache := make(map[[2]int]bool)
    sLen := len(s)
    pLen := len(p)
    var match func (sIndex int, pIndex int) bool
    match = func (sIndex int, pIndex int) bool {
        
        if val, ok := cache[[2]int{sIndex,pIndex}] ;ok{
            return val
        }
        
        //fmt.Println(s, p, sIndex, pIndex)
        if sIndex >= sLen && pIndex >= pLen{
            cache[[2]int{sIndex,pIndex}] = true
            return true
        }
        
        if pIndex >= pLen{
            cache[[2]int{sIndex,pIndex}] = false
            return false
        }
        
        paired := sIndex < sLen && (s[sIndex:sIndex+1] == p[pIndex:pIndex+1] || p[pIndex:pIndex+1]==".")
        
        if (pIndex+1 < pLen) && (p[pIndex+1:pIndex+2] == "*"){
            cache[[2]int{sIndex,pIndex}] = (paired && match(sIndex+1, pIndex)) || match(sIndex,pIndex+2)
            return cache[[2]int{sIndex,pIndex}]
        }
        
        if paired{
            cache[[2]int{sIndex,pIndex}] = match(sIndex+1, pIndex+1)
            return cache[[2]int{sIndex,pIndex}]
        }
        
        cache[[2]int{sIndex,pIndex}] = false
        return false
        
        
    }
    return match(0, 0)
}