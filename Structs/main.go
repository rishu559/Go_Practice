package main

import "fmt"

type person struct{
	firstname string
	lastname string
	contact contact
}

type contact struct {
	number string 
	pin int
}

func main(){
	alex := person{
		firstname: "Alex" ,
		lastname: "Anderson",
		contact : contact{
			number : "9936767031",
			pin : 8765,
		},
		
		}
	// var alex person
	// alex.firstname = "Alex"
	// alex.lastname = "Anderson"
	
	fmt.Printf("%+v",alex)

}