package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func showMenu () {
	fmt.Println("\n####### MENU #######");
	fmt.Println("Contact Us:" + randomdata.PhoneNumber())


	fmt.Println("(1) Profit Calculator");
	fmt.Println("(2) Investment Calculator");
	fmt.Println("(3) Exit");
	fmt.Println("(4) See last calc");
	
	fmt.Print("Enter your choice: ");
}