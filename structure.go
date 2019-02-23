 package main
 import "fmt"

 func main(){
   bob := Cats{"Bob", 4, 0.89}
   fmt.Println(bob.name, "is", bob.age, "years old")
   fmt.Println(bob.name, "fuction is", bob.test())

 }

 type Cats struct {
   name string
   age int
   happiness float64
 }


func (cat *Cats) test() float64 {
  return float64(cat.age) * cat.happiness
}
