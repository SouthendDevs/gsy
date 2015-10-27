package main

import "testing"

func TestCarCanMove100MetresIn10SecondsAt10MetresPerSecond(t *testing.T) {
  //make a car with position 0 and velocity 10
  mycar := &Car{0,10}
  //update the car to move for 10 seconds
  moveCar(mycar,10)
  
  if mycar.position != 100 { t.Fatal("car position not 100m after 10s at 10m/s") }
}

func TestCarCanDriveOntoARoad (t *testing.T) {
  //make a car with position 0 and velocity 10
  mycar := &Car{0,10}
  //make a road
  myroad := newRoad()
  //put car onto road
  driveCarOntoRoad(mycar, myroad)
  
  if positionOfCar(mycar) != 0 { t.Fatal("car not at start of road") }
}
