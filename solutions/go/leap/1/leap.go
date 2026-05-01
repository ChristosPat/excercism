//package leap is calculating if a given year is leap or not
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%100==0 && year%400==0{
        return true
    }else if year%100!=0 && year%4==0{
        return true
    }
    return false
    }
