package decorator

type IPizza interface {
	getPrice() int
}
type VegPizza struct {
	cost int
}

func (pizza *VegPizza) getPrice() int {
	return pizza.cost
}

type TomatoToppings struct {
	pizza      IPizza
	tomatoCost int
}

func (t *TomatoToppings) getPrice() int {
	return t.tomatoCost + t.pizza.getPrice()
}

type CheeseToppings struct {
	pizza      IPizza
	cheeseCost int
}

func (c *CheeseToppings) getPrice() int {
	return c.cheeseCost + c.pizza.getPrice()
}
