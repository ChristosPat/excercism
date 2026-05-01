package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("only integers > 0")
	}
    	return collatz(n), nil
}

func collatz(n int)int{
    counter :=0
for n !=1 {
        if n%2==0{
        n= n/2            
        }else{
       n=3*n+1
            }
    counter ++
         }
return counter
    }

