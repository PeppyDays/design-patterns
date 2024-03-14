package factory_method

import "fmt"

func getGun(gunType string) (Gunner, error) {
	switch gunType {
	case "AK47":
		return newAk47(), nil
	case "Musket":
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
