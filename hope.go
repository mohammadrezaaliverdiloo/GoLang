package main 

import (
    "fmt"    
    )
    
func main (){
    
    var HopeEnd int
    var HopeStep int
    Hope := [] string {}
    fmt.Scanf("%d",&HopeStep)
    fmt.Scanf("%d",&HopeEnd)
    for i:=1;i<=HopeEnd;i++{
         if i%HopeStep==0{
             Hope =append(Hope,"Hope")
             fmt.Print("\n")
             for j:=1;j <= len(Hope);j++{
                fmt.Print("Hope ") 
             }
         }else{
             fmt.Print("\n",i)
        }
    }
}    
