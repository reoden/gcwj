package main 

 import "fmt" 

func main(){
var x int
x = 4
var y int
y = 5
var z int
z = 6
x = 4
y = 5

fmt.Printf("%v\n", ((x * y) + 2))

fmt.Printf("%v\n", ((2 * 3) + (((5 * 2) / 2) * 4)))

fmt.Printf("%v\n", (15 - 5))

fmt.Printf("%v\n", (z >= x))

fmt.Printf("%v\n", (x < 51))

fmt.Printf("%v\n", (x >= 51))

if (x >= 51) { fmt.Printf("%v\n", x)
 }else {
 fmt.Printf("%v\n", z)
 }
if (x > 4) { fmt.Printf("%v\n", y)
 }else {
 fmt.Printf("%v\n", x)
 }
}