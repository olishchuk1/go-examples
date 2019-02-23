package main
import "fmt"

func main(){
  var arr[3] int
  arr[0] = 45
  arr[1] = 7
  arr[2] = 745
  fmt.Println(arr[1])

  nums := [4]float64 {3.14, 2.13, 1.0, 7.77}
  for i, value := range nums{
    fmt.Println(value, "-", i)
  }
}
