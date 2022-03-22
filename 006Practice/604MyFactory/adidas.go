package main

//Fabrica concreta
type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			brand: "adidas",
			size:  14,
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
