package factory_method

type musket struct {
	Gun
}

func newMusket() Gunner {
	return &musket{
		Gun: Gun{
			name:  "Musket Gun",
			power: 1,
		},
	}
}
