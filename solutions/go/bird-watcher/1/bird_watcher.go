package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum :=0
    for i:=0; i<len(birdsPerDay);i++{
        sum=sum+birdsPerDay[i]
    }
return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	wsum := 0
    wind := week*7
    for i:= wind-7; i<wind; i++{
        wsum=wsum+birdsPerDay[i]
    }
return wsum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0; i<len(birdsPerDay); i=i+2 {
        birdsPerDay[i]=birdsPerDay[i]+1
    }
return birdsPerDay
}
