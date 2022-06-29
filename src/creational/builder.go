package creational

import "fmt"

type Vehicle interface {
	GetInfo() string
}

type Car struct {
	Wheels  string
	Engine  string
	Mirrors string
	Doors   string
}

func (car *Car) GetInfo() string {
	return fmt.Sprintf("Vehicle info : Wheels %s - Doors %s, Mirrors : %s, Engine : %s", car.Wheels, car.Doors, car.Mirrors, car.Engine)
}

type Moto struct {
	Wheels string
	Engine string
}

func (moto *Moto) GetInfo() string {
	return fmt.Sprintf("Vehicle info : Wheels %s - Engine : %s", moto.Wheels, moto.Engine)
}

type VehicleBuilder interface {
	SetWheels(string) VehicleBuilder
	SetEngine(string) VehicleBuilder
	SetMirrors(string) VehicleBuilder
	SetDoors(string) VehicleBuilder
	Build() Vehicle
}

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() VehicleBuilder {
	return &CarBuilder{car: &Car{}}
}

func (carBuilder *CarBuilder) SetWheels(wheels string) VehicleBuilder {
	carBuilder.car.Wheels = wheels
	return carBuilder
}
func (carBuilder *CarBuilder) SetEngine(engine string) VehicleBuilder {
	carBuilder.car.Engine = engine
	return carBuilder
}
func (carBuilder *CarBuilder) SetMirrors(mirrors string) VehicleBuilder {
	carBuilder.car.Mirrors = mirrors
	return carBuilder
}
func (carBuilder *CarBuilder) SetDoors(doors string) VehicleBuilder {
	carBuilder.car.Doors = doors
	return carBuilder
}
func (carBuilder *CarBuilder) Build() Vehicle {
	return carBuilder.car
}

type MotoBuilder struct {
	moto *Moto
}

func NewMotoBuilder() VehicleBuilder {
	return &MotoBuilder{moto: &Moto{}}
}

func (motoBuilder *MotoBuilder) SetWheels(wheels string) VehicleBuilder {
	motoBuilder.moto.Wheels = wheels
	return motoBuilder
}
func (motoBuilder *MotoBuilder) SetEngine(engine string) VehicleBuilder {
	motoBuilder.moto.Engine = engine
	return motoBuilder
}
func (motoBuilder *MotoBuilder) SetMirrors(mirrors string) VehicleBuilder {
	println("No Mirrors for a moto")
	return motoBuilder
}
func (motoBuilder *MotoBuilder) SetDoors(doors string) VehicleBuilder {
	println("No Doors for a moto")
	return motoBuilder
}
func (motoBuilder *MotoBuilder) Build() Vehicle {
	return motoBuilder.moto
}
