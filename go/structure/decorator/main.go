package decorator

import "fmt"

func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Println(pizzaWithCheeseAndTomato.getPrice())
}
