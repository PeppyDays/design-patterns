package factory_method

type AK47 struct {
	Gun
}

func newAk47() Gunner {
	return &AK47{
		Gun: Gun{
			name:  "AK47 Gun",
			power: 4,
		},
	}
}
