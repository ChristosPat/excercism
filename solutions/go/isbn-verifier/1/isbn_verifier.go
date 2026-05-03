package isbnverifier

func IsValidISBN(isbn string) bool {
    var nums []int
	sum :=0
    //remove dashes and check for invalid
    for i:=0;i<len(isbn)-1;i++{
        if isbn[i]=='-'{
            continue
        }else if isbn[i]>='0'&& isbn[i]<='9'{
            nums = append(nums,int(isbn[i]-'0'))
        }else{
            return false
        }
    }
    if len (nums)<9{
        return false
    }

    //check the last digit if it is X or number or invalid
    if isbn[len(isbn)-1]=='X'{
        nums=append(nums,10)
    }else if isbn[len(isbn)-1]>='0'&&isbn[len(isbn)-1]<='9'{
        nums=append(nums,int(isbn[len(isbn)-1]-'0'))
    }else{
        return false
    }

    //chek for valid length and the sum
    if len(nums)!=10{
        return false
    }
    j:=10
    
    for i:=0;i<10;i++{
        sum+=nums[i]*j
        j --
    }
    return sum%11==0
}
