package main
//UTF-8
import ("fmt")

func main(){
	s:="THis is a string"
	fmt.Printf("%v, %T\n",s,s)
	fmt.Printf("%v, %T\n",s[0], s[0])
	fmt.Printf("%v, %T\n",string(s[0]), s[0])
	//sstring concatentation
	m:="Niaje bro rada?"
	k:=m+s
	fmt.Println(k)
	//slicing of bits
	b:=[]byte(s)
	fmt.Printf("%v, %T\n",b,b)
}
