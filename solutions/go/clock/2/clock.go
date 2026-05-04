package clock

import "fmt"

// Define the Clock type here.
type Clock struct{
    hour int
    minutes int
}

func New(h, m int) Clock {
    //count everything in minutes and normalize 0<totalMin<1440 
    //1440= total minutes of a day
    totalMin := ((h*60+m)%1440+1440)%1440
    c:=Clock{}
    c.hour=totalMin/60
    c.minutes=totalMin%60
	return c
}

func (c Clock) Add(m int) Clock {
	return New(c.hour,c.minutes+m)
    }

func (c Clock) Subtract(m int) Clock {
	return New(c.hour,c.minutes-m)
}

func (c Clock) String() string {
	a:= fmt.Sprintf("%02d:%02d",c.hour,c.minutes)
    return a
}
