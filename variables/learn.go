package variables

import (
	"fmt"
	"strconv"
)

var fullname string//global
var email string ="mark@gmail.com"
var c,python bool = true,false
const vat int = 7

func Learn() {
	fullname = "Mark Sakaberg"
	fmt.Printf("Hello %s Email %s ",fullname,email)
	fmt.Println(c,python)

	companyName := "Metaverse";
	isShow :=true
	age:=10
	fmt.Println(companyName,isShow,age)
	fmt.Printf("%T \n",isShow)

	f:=float64(age)
	fmt.Printf("%T \n",f)

	s:=strconv.Itoa(vat)
	fmt.Println("vat is " + s)
}