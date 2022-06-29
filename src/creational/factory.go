package creational

type CarType int

const (
	AudiPrice       = 10.45
	VolkswagenPrice = 5.67

	Audi CarType = iota
	Volkswagen
)

type iCar interface {
	GetPrice() float64
}

type cAudi struct {
}

func (audi *cAudi) GetPrice() float64 {
	return AudiPrice
}

type cVolkswagen struct {
}

func (volkswagen *cVolkswagen) GetPrice() float64 {
	return VolkswagenPrice
}

type CarFactory struct {
}

func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

func (carFactory *CarFactory) GetCar(carType CarType) iCar {
	switch carType {
	case Audi:
		return &cAudi{}
	case Volkswagen:
		return &cVolkswagen{}
	default:
		return nil
	}
}
