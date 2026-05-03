package luhn

func Valid(id string) bool {
	
    var numId []int 

 	for i:=0;i<len(id);i++{
        if id[i]==' '{
            continue
        }
        a:=int(id[i]-'0')
        if id[i]<'0' || id[i]>'9'{
            return false
            }else{
            numId = append(numId,a)
        }
	}
    
    if len(numId)<2{
       return false
       }
    
    sum:=0
    d:=0
    for i:=len(numId)-1;i>=0;i--{
        if d==0{
           sum+=numId[i] 
            d=1
        }else{
            a:=numId[i]*2
            if a>9{
                a=a-9
                sum+=a
                }else{
                sum+=a
            }
            d=0
        }
    }
    
    if sum%10==0{
        return true
    }
    return false
}
