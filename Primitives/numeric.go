//this program is for display only wont run at all
//Floating point follow IEEE 754 standard
package main

import("fmt")

func main(){
	var a int8 //-128 to 127
	var b int16// -32768 to 32767
	var c int32//-2147483648 to 2147483647
	var d int64// -9223372036854775808 to 9223372036854775807
	var e uint8// 0-255
	var f uint16//0-65535
	var g uint32//0-4294967295
	//operations are the same as for other languages
	//operations occur on variables of the same type without your consent
	//consent comes when you implicitly do type conversions
	//bit operators are also present
	//^ represents exlusive OR
	//&^(AND NOT(not understood))
	//bit shifting also occurs both left and right
	var h float32 3.14
	var i float64 2.1E14
	var j float32 12.e72
	var k complex64= 1+2i
	var l complex128=12+4i
	//to get the real part or imaginary part
	fmt.Printf("%v, %T\n", real(k),real(k))
	fmt.Printf("%v, %T\n", imag(k),imag(k))
	// values will be float32s
	var m complex128=complex(5,12)
	//declaring a complex number using the inbuilt funxtion
}
