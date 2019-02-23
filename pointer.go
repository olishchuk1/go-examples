package main
import "fmt"

func main(){
  var x = 0
  pointer(&x)
  fmt.Println(x)
}

func pointer (y *int) {
  *y = 2
}
