package main
//naming controls visibility
//lowercase variables are scoped to the package
//uppercase variables scoped to the outside
//blockscope within function
//length of variable should reflect the lifespan of the variable eg counters
//longer names for variables with longer lifespan
//acronymns put in uppercase for readablity
import (
	"fmt"
)
var( b float64 = 43
	m float32 =36
)
func main(){
	var i int=30
	k := 99
	var j="Konyagi"
//shadowing- innermost variable takes precedence
	fmt.Println(m)
	var m float32=45
	fmt.Println(m)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(i)
	fmt.Printf("%v, %T", j,j)
	fmt.Println("Empty")
	fmt.Println(b)
	}
