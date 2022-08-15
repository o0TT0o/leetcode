func isPalindrome(x int) bool {
    input := x
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }
    
    rev_int :=0
    for x>=10{
        rev_int = rev_int*10 + x%10
        x = x/10
    }
    rev_int = rev_int*10 + x
    
    fmt.Println(rev_int)
    if rev_int == input {
        return true 
    } else {
        return false
    }
}