package abstractfactory

type NikeShoe struct {
	logo string
	size int
}

func (s *NikeShoe) setLogo(logo string) {
	s.logo = logo
}

func (s *NikeShoe) setSize(size int) {
	s.size = size
}

func (s *NikeShoe) getLogo() string {
	return s.logo
}

func (s *NikeShoe) getSize() int {
	return s.size
}

type NikeShirt struct {
	logo string
	size int
}

func (s *NikeShirt) setLogo(logo string) {
	s.logo = logo
}

func (s *NikeShirt) setSize(size int) {
	s.size = size
}

func (s *NikeShirt) getLogo() string {
	return s.logo
}

func (s *NikeShirt) getSize() int {
	return s.size
}

type Nike struct{}

func (a *Nike) makeShoe() Shoe {
	return &NikeShoe{
		logo: "adidas",
		size: 14,
	}
}

func (a *Nike) makeShirt() Shirt {
	return &NikeShirt{
		logo: "adidas",
		size: 14,
	}
}
