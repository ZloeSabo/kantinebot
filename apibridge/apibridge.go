package apibridge

type BridgeProduct struct {
	name        string
	description string
	price       string
	kcal        string
}

func (p *BridgeProduct) Name() string {
	return p.name
}
