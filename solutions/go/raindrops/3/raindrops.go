package raindrops

import "fmt"

func Convert(number int) string {
    var a string
   if number%3==0{
    	a+="Pling"
    }
	if number%5==0{
    	a+="Plang"
    }
	if number%7==0{
    	a+= "Plong"
    }
	if a==""{
    	a=fmt.Sprint(number)
    }
return a
}
