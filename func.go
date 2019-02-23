package main
import "fmt"

func main(){
  var a = 5
  var b = 10
  var r int
  r = summ(a, b)
  fmt.Println(r)

  a, b = summ2(a, b)
  fmt.Println(a, b)

  var num = 3
  square := func() int {
    num *= num
    return num
    }
  fmt.Println(square())
}

func summ(num_1 int, num_2 int) int {
  var res int
  res = num_1 + num_2
  return res
}

func summ2(num_1 int, num_2 int) (int, int) {
  var res int
  res =  num_1 + num_2
  return res, num_1*num_2
}
