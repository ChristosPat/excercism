package phonenumber

import ("errors"
	"fmt")

func Number(phoneNumber string) (string, error) {
//strip the number
    var phoneA,phone string
    for i := range phoneNumber{
        if phoneNumber[i]>='0' && phoneNumber[i]<='9'{
            phoneA += string(phoneNumber[i])
        }
    }
//check country code
    if  len(phoneA) ==11 && phoneA[0]=='1'{
        phone=phoneA[1:]
    }else if len(phoneA)==10{
        phone=phoneA
    }else{
        return phoneNumber,errors.New("Invalid phone number")
    }
//check area code
    if phone[0] <'2'|| phone[0] >'9' || phone[3]<'2' || phone[3]>'9'{
        return phoneNumber,errors.New("Invalid are code")
    }
    return phone,nil
    }

    	

func AreaCode(phoneNumber string) (string, error) {
	phone,err:=Number(phoneNumber)
    return phone[:3],err
}
func Format(phoneNumber string) (string, error) {
	phone,err:=Number(phoneNumber)
    return fmt.Sprintf("(%s) %s-%s",phone[:3],phone[3:6],phone[6:]),err
}
