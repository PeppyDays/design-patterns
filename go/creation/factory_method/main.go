package factory_method

import "fmt"

func main() {
	ak47, _ := getGun("AK47")
	musket, _ := getGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g Gunner) {
	fmt.Printf("Gun: %s", g.getName())
}
