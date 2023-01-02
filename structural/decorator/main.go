package decorator

import "fmt"

func Demo() {
	pizza := &VegPizza{
		cost: 100,
	}
	pizzaWithTomato := &TomatoToppings{
		pizza:      pizza,
		tomatoCost: 30,
	}

	pizzaWithTomatoAndCheese := &CheeseToppings{
		pizza:      pizzaWithTomato,
		cheeseCost: 40,
	}
	fmt.Printf("Price of veg Pizza with tomato and cheese topping is %d\n", pizzaWithTomatoAndCheese.getPrice())

}
