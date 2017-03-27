package main

import (
    "fmt"
    )
//
// Abstract Stuff
//

// Component Interfaces

type Wheels interface {
    Friction(speed float64, num_wheels int) float64
}

type Motor interface {
    FuelConsumption(rpm float64, torque float64) float64
}

type Body interface {
    Drag(speed float64) float64
}

// Class with components

type Vehicle struct {
    wheels Wheels
    num_wheels int
    motor Motor
    body Body
}

func (v Vehicle) FuelConsumption(speed float64, rpm float64) float64 {
    var torque float64
    torque = v.wheels.Friction(speed, v.num_wheels) + v.body.Drag(speed) // ...some calculation...
    return v.motor.FuelConsumption(rpm, torque)
}

// Object composition often results in annoying forwarding methods

func (v Vehicle) WheelFriction(speed float64) float64 {
    return v.wheels.Friction(speed, v.num_wheels)
}

func (v Vehicle) BodyDrag(speed float64) float64 {
    return v.body.Drag(speed)
}

//
// Concrete Components
//
type LightDutyWheels struct {}
func (LightDutyWheels) Friction(speed float64, num_wheels int) float64 {
    // ...some calc...
    return float64(num_wheels) * speed
}

type HeavyDutyWheels struct {}
func (HeavyDutyWheels) Friction(speed float64, num_wheels int) float64 {
    // ...some calc...
    return float64(num_wheels) * speed * 2
}

type DieselMotor struct {}
func (DieselMotor) FuelConsumption(rpm float64, torque float64) float64 {
    // ... some calc ...
    return torque * rpm / 3.2e6
}

type PetrolMotor struct {}
func (PetrolMotor) FuelConsumption(rpm float64, torque float64) float64 {
    // ... some calc ...
    return torque * rpm / 2.3e6
}

type CarBody struct {}
func (CarBody) Drag(speed float64) float64 {
    // ... some calc ...
    return 340
}

type TruckBody struct {}
func (TruckBody) Drag(speed float64) float64 {
    // ... some calc ...
    return 650
}

type BusBody struct {}
func (BusBody) Drag(speed float64) float64 {
    // ... some calc ...
    return 800
}

type UteBody struct {}
func (UteBody) Drag(speed float64) float64 {
    // ... some calc ...
    return 470
}

//
// Instantiation
//

func main() {
    car := Vehicle { LightDutyWheels {}, 4, PetrolMotor {}, CarBody {} }

    truck := Vehicle { HeavyDutyWheels {}, 8, DieselMotor {}, TruckBody {} }

    ute := Vehicle { LightDutyWheels {}, 4, DieselMotor {}, UteBody {} }

    bus := Vehicle { HeavyDutyWheels {}, 6, DieselMotor {}, BusBody {} }

    fmt.Println("Car fuel consumption is", car.FuelConsumption(60, 5000))
    fmt.Println("Truck fuel consumption is", truck.FuelConsumption(60, 5000))
    fmt.Println("Ute fuel consumption is", ute.FuelConsumption(60, 5000))
    fmt.Println("Bus fuel consumption is", bus.FuelConsumption(60, 5000))
}
