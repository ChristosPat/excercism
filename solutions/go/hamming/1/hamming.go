package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a)!=len(b){
        return 0,errors.New("Not even strands")
    }
    dist := 0
    for i,v:= range a{
        if v != rune(b[i]){
            dist +=1
        }
    }
    return dist,nil
}
