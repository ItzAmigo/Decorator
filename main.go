package main

import (
	"Decorator/Decorator"
	"fmt"
)

func main() {
	var burger1 Decorator.IBurger = Decorator.NewBeef()
	burger1 = Decorator.NewCheese(burger1)
	burger1 = Decorator.NewCheese(burger1)
	burger1 = Decorator.NewCheese(burger1)
	fmt.Println("Triple CheeseBeef: "+burger1.GetDescription(), burger1.GetPrice())

	var burger2 Decorator.IBurger = Decorator.NewChicken()
	burger2 = Decorator.NewSpicy(burger2)
	burger2 = Decorator.NewCheese(burger2)
	fmt.Println("Spicy ChickenCheese: "+burger1.GetDescription(), burger2.GetPrice())
}
