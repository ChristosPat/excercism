package sumofmultiples

func exist (list []int,num int)bool{
    for _,i := range list{
        if num == i{
            return true
        }
    }
    return false
}

func SumMultiples(limit int, divisors ...int) int {
	multi := []int{}
	result :=0

    for _,div:=range divisors{
        if div==0{
            continue
        }
    	for multiple:=div;multiple<limit;multiple+=div{
            if !exist(multi,multiple){
                multi=append(multi,multiple)
            }
        }
    }
   
	for _,i:= range multi{
        result+=i
        
    }
  return result
}
