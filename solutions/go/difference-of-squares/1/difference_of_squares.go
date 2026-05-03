package differenceofsquares


func SquareOfSum(n int) int {
    sum:=0
	for i:=0;i <= n;i++{
        sum+=i
    }
    return sum*sum
}
func SumOfSquares(n int) int {
    squar:=0
	for i:=0;i <= n;i++{
        squar+=i*i
    }
    return squar
}

func Difference(n int) int {
	return SquareOfSum(n)-SumOfSquares(n)
}
