package main 

import ("fmt")

func main(){
	
	var n bool= true //equivalency generates the boolean type
	m:= 1==2
	j:= 4==4
	fmt.Printf("%v,%T\n",n,n)
	fmt.Printf("%v, %T \n", m,m)
	fmt.Printf("%v, %T\n",j,j)
	}
