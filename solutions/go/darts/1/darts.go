package darts



func Score(x, y float64) int {
	hit:=x*x+y*y
    switch  {
        case hit<=1:
            return 10
        case hit<=25:
        	return 5
        case hit<=100:
        	return 1
    }
    return 0
    }
    
