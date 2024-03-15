package abstractfactory

type AdidasShoe struct {
	logo string
	size int
}

func (s *AdidasShoe) setLogo(logo string) {
	s.logo = logo
}

func (s *AdidasShoe) setSize(size int) {
	s.size = size
}

func (s *AdidasShoe) getLogo() string {
	return s.logo
}

func (s *AdidasShoe) getSize() int {
	return s.size
}

type AdidasShirt struct {
	logo string
	size int
}

func (s *AdidasShirt) setLogo(logo string) {
	s.logo = logo
}

func (s *AdidasShirt) setSize(size int) {
	s.size = size
}

func (s *AdidasShirt) getLogo() string {
	return s.logo
}

func (s *AdidasShirt) getSize() int {
	return s.size
}

type Adidas struct{}

func (a *Adidas) makeShoe() Shoe {
	return &AdidasShoe{
		logo: "adidas",
		size: 14,
	}
}

func (a *Adidas) makeShirt() Shirt {
	return &AdidasShirt{
		logo: "adidas",
		size: 14,
	}
}
