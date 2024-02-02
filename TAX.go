package main

import (
    "fmt"
    "math"
    )

func main(){
    var count_of_coins float64
    fmt.Scanf("%f",&count_of_coins)

    
    if count_of_coins <= 10000{
        fmt.Print(math.Round(float64((count_of_coins/100)*5)))
    }else if count_of_coins > 10000 && count_of_coins <=50000{
        count_of_coins = count_of_coins - 10000
        fmt.Print((math.Round(float64((count_of_coins/100)*10)+500)))
    }else if count_of_coins > 50000 && count_of_coins <=100000{
        count_of_coins = count_of_coins - 50000
        fmt.Print((math.Round(float64((count_of_coins/100)*15)+4500)))
    }else{
        count_of_coins = count_of_coins - 100000
        fmt.Print((math.Round(float64((count_of_coins/100)*20)+12000)))
    }
}
