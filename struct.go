package main

import (
	"fmt"
	"strconv"
)

//Insert struct declaration here
type customer struct {
	fName            string
	lName            string
	age              int
	subscriber       bool
	homeAddress      string
	phone            int
	creditAvailable  float32
	currentCartCost  float32
	currentOrderCost float32
}

func main() {
	//Insert code here
	customer1 := customer{
		fName:            "Maya",
		lName:            "Fey",
		age:              18,
		subscriber:       false,
		homeAddress:      "Kurain",
		phone:            29941224,
		creditAvailable:  2000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}
	customer2 := customer{
		fName:            "Mia",
		lName:            "Fey",
		age:              25,
		subscriber:       true,
		homeAddress:      "Kurain",
		phone:            29201224,
		creditAvailable:  9000.00,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	customerBase := []customer{customer1, customer2}

	for _, person := range customerBase {
		printCustomerData(person)
	}

}
func printCustomerData(person customer) {
	fmt.Println("Person's name is : " + person.fName + " " + person.lName + ".")
	fmt.Println("Age is :" + strconv.Itoa(person.age) + ".")
	fmt.Println("Subscribed? : " + strconv.FormatBool(person.subscriber) + ".")
	fmt.Println("Home address is : " + person.homeAddress + ".")
	fmt.Print("Phone Number:")
	fmt.Println(person.phone)
	fmt.Print("Credit Available :")
	fmt.Println(person.creditAvailable)
	fmt.Print("Current Cart Cost :")
	fmt.Println(person.currentCartCost)
	fmt.Print("Current Order Cost :")
	fmt.Println(person.currentOrderCost)

}
