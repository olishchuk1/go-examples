package main
import "fmt"

func main(){
    // const pi = 3.14
    var age int= 18
    var float float64 = 23.4543
    var str = "Hello world"
    fmt.Println(age, float, "\n" + str)

    var web string = "In code we trust"
    fmt.Println (len(web))

    var num float32 = 4.222527454584272
    fmt.Printf("%T \n", num) //Показує тип даних
    fmt.Printf("%f \n", num) //За допомогою цієї функції можна виводити сталу
                             //кількість цифр після крапики
    var isDone bool = true
    fmt.Printf("%t \n", isDone)
}
