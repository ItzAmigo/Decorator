package Decorator

type Beef struct {
	description string
}

func (b *Beef) GetPrice() int {
	return 15
}

func (b *Beef) GetDescription() string {
	return b.description
}

func NewBeef() *Beef {
	beef := new(Beef)
	beef.description = "Beef"
	return beef
}
