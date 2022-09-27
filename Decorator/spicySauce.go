package Decorator

type SpicySauce struct {
	burger      IBurger
	description string
}

func (s *SpicySauce) GetPrice() int {
	return s.burger.GetPrice() + 3
}

func (s *SpicySauce) GetDescription() string {
	return s.burger.GetDescription() + ", Spicy"
}

func NewSpicy(burger IBurger) *SpicySauce {
	spicy := new(SpicySauce)
	spicy.burger = burger
	return spicy
}
