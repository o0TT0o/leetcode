//"42"
//"   -42"
//"4193 with words"
//"words and 987"
//" "
//"3.14159"
//"21474836460"
//"00000-42a1234"
//"9223372036854775808"
//"-9223372036854775808"
//"  +  413"
//"-4193 with words"

func myAtoi(s string) int {
    neg := false
    result := 0
    ifNegSet := false
    ifValSet := false
    const maxVal = math.MaxInt32
    const minVal = math.MinInt32
    const throttleVal = math.MaxInt32/10
    const posThrottleRemainderVal = math.MaxInt32%10
    const negThrottleRemainderVal = posThrottleRemainderVal + 1

    //fmt.Println(throttleVal)
    //fmt.Println(posThrottleRemainderVal)
    //fmt.Println(negThrottleRemainderVal)
    
    for _,char := range s{
        if !ifValSet && !ifNegSet && char == '-'{
            neg = true
            ifNegSet = true
            continue
        }

        if !ifValSet && !ifNegSet && char == '+'{
            neg = false
            ifNegSet = true
            continue
        }
       
        // int char '0' = 48
        intChar := int(char)
        //fmt.Println(intChar)
        if intChar>=48 && intChar<=57{
            if !neg && ( result > throttleVal || result == throttleVal && intChar > (48 + posThrottleRemainderVal)) {
                return maxVal
            }
            
            if neg && ( result > throttleVal || result == throttleVal && intChar > (48 + negThrottleRemainderVal)) {
                return minVal
            }         
            
            result = result*10 + intChar - 48
            ifValSet = true
            //fmt.Println(result)
            continue
        }
        
        if !ifValSet && ifNegSet {
            return 0
        }
       
        // int char 'blank' = 32
        if !ifValSet && intChar != 32 {
            return 0
        }
        
       
        if ifValSet{
            break
        }

    }
   
    //fmt.Println(result)
    //fmt.Println(neg)
    if neg{
        result = 0-result
    }
   
    return result

}