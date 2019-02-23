package main
import "fmt"

func main(){
  var i = 0
  for i <= 10{
    fmt.Println(i)
    i++
  }

  for i := 0; i <=5; i++{
    fmt.Println(i)
  }
}

func PositiveSum(numbers []int) (sum int) {
  for _, num := range numbers {
    if num > 0 {
      sum = sum + num
    }
  }
  return sum
}
