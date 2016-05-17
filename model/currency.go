package model

type Currency struct {
	name string
	converter float64
}

func NewCurrency(name string, converter float64) *Currency {
	return &Currency{
		name,
		converter,
	}
}

func (c Currency) SetName(name string) {
	c.name = name
}

func (c Currency) GetName() string {
	return c.name
}

func (c Currency) SetConverter(converter float64) {
	c.converter = converter
}

func (c Currency) GetConverter() float64 {
	return c.converter
}