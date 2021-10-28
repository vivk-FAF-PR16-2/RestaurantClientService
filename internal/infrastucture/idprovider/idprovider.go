package idprovider

type IProvider interface {
	Get() int
}

type idProvider struct {
	id int
}

func NewProvider() IProvider {
	return &idProvider{
		id: 0,
	}
}

func (p *idProvider) Get() int {
	defer func() {
		p.id++
	}()
	return p.id
}
