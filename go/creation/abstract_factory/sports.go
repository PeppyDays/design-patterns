package abstractfactory

import "fmt"

type Shoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type SportsFactory interface {
	makeShoe() Shoe
	makeShirt() Shirt
}

func GetSportsFactory(brand string) (SportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
