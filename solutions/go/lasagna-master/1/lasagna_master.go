package lasagna

func PreparationTime(layers []string, averTime int ) (prepTime int){
	if averTime==0 {
        prepTime = len(layers)*2
    }else{
		prepTime= len(layers)*averTime
 }
return
}

func Quantities(layers []string) (noodles int, sauce float64){
	
    for i:=0; i<len(layers);i++{
        if layers[i]=="noodles"{
            noodles+=50
            
        }else if layers[i]=="sauce"{
        	sauce+=0.2
          
        }
    }
return
}
 
func AddSecretIngredient(friendsList,myList []string){
    
    myList[(len(myList)-1)]=friendsList[(len(friendsList)-1)]
}



func ScaleRecipe(quantities []float64, portions int)(result []float64){
    	a:=(float64(portions))/2
    for i:=0; i<len(quantities);i++{
        result=append(result,(quantities[i]*a))
    }
return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
