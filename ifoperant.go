package main
import "fmt"

func main (){
  var age = 8
  if age < 5 {
    fmt.Println("Less then five!")
  } else if age == 5{
    fmt.Println("Equals five!")
  } else if (age > 5) && (age < 18){
    var grade = age - 5
    fmt.Println("You should go to", grade, "grade")
  } else {
    fmt.Println("You should go to the university")
  }
}
