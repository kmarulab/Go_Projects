package main

import ("fmt"
//	"strconv"//converts int to strings 
)

func main(){
	var i float32=42.5
	fmt.Printf("%v , %T\n",i,i)
	var j int
	j= int(i)//int conversion operator
	fmt.Printf("%v, %T\n",j,j) 
//	var m string
//	m= strconv.Itoa(i)
//	fmt.Printf("%v,%T\n", m,m)
}
