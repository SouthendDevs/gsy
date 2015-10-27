package main

type Car struct {
  position float32
  velocity float32
}

func moveCar(car *Car, duration float32) {
  car.position += car.velocity*duration
}
