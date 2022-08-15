//123

//-123

//120

//10

//1534236469

func reverse(x int) int {

   

    nev := false

   

    if x < 0 {

        nev = true

        x = 0 - x

    }

   

    //times := 0

    arr := []int{}

    for x >=10 {

        arr = append(arr,x%10)

        x = x/10

        //times++

    }

    arr = append(arr,x%10)

    //fmt.Println(times)

    //fmt.Println(arr)

    result := 0

    times := len(arr)-1

    for i := times ; i>=0 ; i--{

        //fmt.Println( int(math.Pow(10.0,float64(i))) )

        result = result + arr[times-i] *  int(math.Pow(10.0,float64(i)))

    }

   

    if nev{

        result = 0-result

    }

   

    if result > math.MaxInt32 -1 || result < math.MinInt32{

        return 0

    }

    return result

}