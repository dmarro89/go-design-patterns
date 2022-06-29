package creational

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewFactory(t *testing.T) {
	assert.DeepEqual(t, NewCarFactory(), &CarFactory{})
}

func TestGetCar(t *testing.T) {
	factory := NewCarFactory()
	audi := factory.GetCar(Audi)
	volkswagen := factory.GetCar(Volkswagen)
	assert.DeepEqual(t, audi, &cAudi{})
	assert.Equal(t, audi.GetPrice(), AudiPrice)
	assert.DeepEqual(t, volkswagen, &cVolkswagen{})
	assert.Equal(t, volkswagen.GetPrice(), VolkswagenPrice)
	assert.Equal(t, factory.GetCar(10), nil)
}
