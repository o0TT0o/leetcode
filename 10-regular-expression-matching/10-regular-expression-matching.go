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
    sLen := len(s)
    pLen := len(p)
    var match func (sIndex int, pIndex int) bool
    match = func (sIndex int, pIndex int) bool {
        //fmt.Println(s, p, sIndex, pIndex)
        if sIndex >= sLen && pIndex >= pLen{
            return true
        }
        
        if pIndex >= pLen{
            return false
        }
        
        paired := sIndex < sLen && (s[sIndex:sIndex+1] == p[pIndex:pIndex+1] || p[pIndex:pIndex+1]==".")
        
        if (pIndex+1 < pLen) && (p[pIndex+1:pIndex+2] == "*"){
            return (paired && match(sIndex+1, pIndex)) || match(sIndex,pIndex+2)
        }
        
        if paired{
            return match(sIndex+1, pIndex+1)
        }
        
        return false
        
        
    }
    return match(0, 0)
}