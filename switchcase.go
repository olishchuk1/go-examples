package main
import "fmt"

func main(){
  var age = 5
  switch age {
  case 5: fmt.Println("You are 5 yo")
  case 15: fmt.Println("You are 15 yo")
  case 25: fmt.Println("You are 25 yo")
  case 65: fmt.Println("You are 65 yo")
  default: fmt.Println("I don't know how old are you")
  }
}
