package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			size:  19,
			brand: "Nike bb",
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
