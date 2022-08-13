//"PAYPALISHIRING"
//3
//"PAYPALISHIRING"
//4
//"A"
//1
//"A"
//2
//"ABC"
//2
func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    
    stringLen := len(s)
    
    midNum := numRows - 2
    partOfZNum := numRows + midNum
    
    
    eachPartColumnSize := 1+midNum
    baseColumn := (stringLen/partOfZNum) * eachPartColumnSize
    //fmt.Println(baseColumn)
    
    modOfZCount := 0
    modOfZCount = stringLen%partOfZNum
    
    
    //fmt.Println(modOfZCount)
    
    
    if modOfZCount == 0{
        baseColumn = baseColumn
    }else if modOfZCount <= numRows{
        baseColumn = baseColumn + 1
    }else{
        baseColumn = baseColumn + 1 + (modOfZCount - numRows ) 
    }
    
    //fmt.Println(baseColumn)

    //const arryColumn = 1000
    //const arryRow = 1000
    //var emptyArray [arryColumn][arryRow]string
    
    emptyArray := make([][]string, baseColumn)
    for i := range emptyArray{
        emptyArray[i] = make([]string , numRows)
    }
    
    index := 0 
    for i:=0 ; i<baseColumn ; i++{
        if i==0 || i%eachPartColumnSize ==0 {
            for j:=0 ; j < numRows; j++ {
                emptyArray[i][j] = s[index:index+1]
                index++
                if index >= stringLen{
                    break
                }
            }
        }else{
            pos := numRows -1 - i%eachPartColumnSize
            emptyArray[i][pos] = s[index:index+1]
            index++
        }
        if index >= stringLen{
            break
        }
    }
    //fmt.Println(emptyArray)
    
    result := ""
    for j:=0 ; j< numRows; j++{
        for i:=0; i < baseColumn; i++{
            if emptyArray[i][j] != ""{
                result=result+emptyArray[i][j]
            }
        }
    }
    
    
    return result
    
    
    
}