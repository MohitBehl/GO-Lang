package main

import "fmt"

type name struct {
	firstname, lastname, middlename string
}

func (obj *name) setdefault(){
	obj.firstname = "default"
	obj.middlename = "default"
	obj.lastname = "default"
}


func main() {
	// v := name{ firstname: "abc" , middlename: "ghi"}

	var v name
	v.setdefault()
	fmt.Println(v)

}
