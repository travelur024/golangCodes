package main

type iShoe interface {
	setSize(size int)
	setBrand(brand string)
	getSize() int
	getBrand() string
}

type shoe struct {
	brand string
	size  int
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) setBrand(brand string) {
	s.brand = brand
}
func (s *shoe) getBrand() string {
	return s.brand
}
func (s *shoe) getSize() int {
	return s.size
}
