package main 

 import "fmt" 

func main(){
var i int
i = 0
i = (i + 1)
var j int
j = 0
j = (j + 1)
if (i <= 3) {
 fmt.Printf("%v\n", i)
i = (i + 1)

 }
for (j <= 3) {
 fmt.Printf("%v\n", j)
if (j == 2) {
 fmt.Printf("%v\n", 44)
 }

j = (j + 1)

 }

}