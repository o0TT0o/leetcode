//"PAYPALISHIRING"
//3
//"PAYPALISHIRING"
//4
//"A"
//1
//"A"
//2
//"A"
//3
//"ABC"
//2 
func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    
    result :=""
    steps := 0
    steps2 := 0
    
    for i := 0; i< numRows; i++{
        //result = result + " | "
        if i==0 || i==numRows-1 {
            steps = (numRows - 1)*2
            for j:=i ; j<len(s); j=j+steps{
                result = result + s[j:j+1]
            }
            
        }else{
            steps = (numRows - 1)*2
            steps2 = (numRows - 1)*2 - 2*i
            
            if i >= len(s){
                break 
            }
            
            result = result + s[i:i+1]
            j1steps := i
            j2steps := i
            
            for {
                j2steps = j1steps + steps2
                
                if j2steps < len(s){
                    result = result + s[j2steps:j2steps+1]
                }else{break}
                
                j1steps = j1steps + steps
                if  j1steps < len(s)  {
                    result = result + s[j1steps:j1steps+1]
                }else{break}
            }
        }
        
    }
    
    
    return result
    
    
    
}