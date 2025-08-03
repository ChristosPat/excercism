package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	a:= (successRate/100)
    return float64(productionRate)*a
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	a := float64(productionRate)*(successRate/100)/60
    return int(a)
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	a := carsCount / 10
    b := carsCount % 10
    c := a*95000+b*10000
    return uint(c)
}
