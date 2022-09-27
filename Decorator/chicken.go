package Decorator

type Chicken struct {
	description string
}

func (c *Chicken) GetPrice() int {
	return 12
}

func (c *Chicken) GetDescription() string {
	return c.description
}

func NewChicken() *Chicken {
	chicken := new(Chicken)
	chicken.description = "Beef"
	return chicken
}
