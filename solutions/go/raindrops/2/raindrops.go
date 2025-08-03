package raindrops

import "fmt"

func Convert(number int) string {
    var a string
    if number%3==0 && number%5==0 && number%7==0{
    	a="PlingPlangPlong"
        }else if number%5==0 && number%7==0{
    	a="PlangPlong"
        }else if number%3==0 && number%7==0{
    	a="PlingPlong"
        }else if number%3==0 && number%5==0{
    	a="PlingPlang"
       	}else if number%3==0{
    	a="Pling"
    }else if number%5==0{
    	a="Plang"
    }else if number%7==0{
    	a= "Plong"
    } else {
    	a=fmt.Sprint(number)
    }
return a
}
