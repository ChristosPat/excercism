package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen" : 6,
        "dozen" : 12,
        "small_gross" : 120,
        "gross" : 144,
        "great_gross" : 1728,
    }
return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_,ok := units[(unit)]
    if ok==false{
        return false
    }else{
    	_,ok:=bill[item]
        if ok==false{
            bill[item]=units[unit]
        }else{
        	bill[item]+=units[unit]
        }
    return true
    }
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_,aok:=bill[item]
    _,bok:=units[unit]
    if !aok{
        return false
    }else if !bok{
        return false
    }else if bill[item]-units[unit]==0{
        delete(bill, (item))
        return true
    }else if bill[item]-units[unit]<0{
        return false
    }else{
    	bill[item]-=units[unit]
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_,ok:=bill[item]
    if !ok{
        return  0,false
    }else{
    	return bill[item],true
    }
}
