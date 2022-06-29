package creational

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const (
	doors   = "doors"
	wheels  = "wheels"
	engine  = "engine"
	mirrors = "mirrors"
)

func TestCarBuilder(t *testing.T) {
	car := NewCarBuilder().SetDoors(doors).SetEngine(engine).SetMirrors(mirrors).SetWheels(wheels).Build()
	comparedCar := &Car{Wheels: wheels, Engine: engine, Mirrors: mirrors, Doors: doors}
	assert.DeepEqual(t, car, comparedCar)
	assert.Equal(t, car.GetInfo(), fmt.Sprintf("Vehicle info : Wheels %s - Doors %s, Mirrors : %s, Engine : %s", wheels, doors, mirrors, engine))

	moto := NewMotoBuilder().SetEngine(engine).SetWheels(wheels).SetDoors(doors).SetMirrors(wheels).Build()
	comparedMoto := &Moto{Wheels: wheels, Engine: engine}
	assert.DeepEqual(t, moto, comparedMoto)
	assert.Equal(t, moto.GetInfo(), fmt.Sprintf("Vehicle info : Wheels %s - Engine : %s", wheels, engine))
}
