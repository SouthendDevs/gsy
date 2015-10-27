package main

import "container/list"

type Car struct {
  position float32
  velocity float32
}

func moveCar(car *Car, duration float32) {
  car.position += car.velocity*duration
}

////////////////////////////////

type Road struct {
  cars *list.List
}

func newRoad() *Road {
  road := &Road{ list.New() }
  return road
}

func driveCarOntoRoad(car *Car, road *Road) {
  car.position = 0
  road.cars.PushBack(car)
}

func positionOfCar(car *Car) float32 {
  return car.position
}
