package isbnverifier

func IsValidISBN(isbn string) bool {
	    sum:=0
    dig:=0
    for i:=0;i<len(isbn);i++{
        if isbn[i]=='-'{
            continue
    	}else if isbn[i]>='0'&&isbn[i]<='9'{
            sum+=int(isbn[i]-'0')*(10-dig)
            dig ++
        }else if isbn[i]=='X'&& i==len(isbn)-1{
            sum+=10
            dig ++
        }else{
            return false
        }
    }
	if dig!=10{
        return false
    }
    return sum%11==0
    }