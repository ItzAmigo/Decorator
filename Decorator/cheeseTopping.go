package Decorator

type CheeseTopping struct {
	burger      IBurger
	description string
}

func (c *CheeseTopping) GetPrice() int {
	return c.burger.GetPrice() + 3
}

func (c *CheeseTopping) GetDescription() string {
	return c.burger.GetDescription() + ", Cheese"
}

func NewCheese(burger IBurger) *CheeseTopping {
	cheese := new(CheeseTopping)
	cheese.burger = burger
	return cheese
}
