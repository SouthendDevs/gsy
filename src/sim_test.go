package main

import "testing"

/*
func TestOneCarMovesFromPositionOneToPositionTwo(t *testing.T) {
	//construct a car with position state 1
	mycar := &Car{1, 0}
	//move the car's position from 1 to 2
	moveCar(mycar,1)
	
	if mycar.position != 2 { t.Fatal("car position not 2, some other value") }
}

func TestOneCarMovesFromPositionOneToPositionTen(t *testing.T) {
	//construct a car with position state 1
	mycar := &Car{1, 0}
	//move the car's position from 1 to 2
	moveCar(mycar,9)
	
	if mycar.position != 10 { t.Fatal("car position not 10 some other value") }
}
*/

func TestCarCanMove100MetresInTenSecondsAtTenMetresPerSecond(t *testing.T) {
  //make a car with position 0 and velocity 10
  mycar := &Car{0,10}
  //update the car to move for 10 seconds
  moveCar(mycar,10)
  
  if mycar.position != 100 { t.Fatal("car position not 100m after 10s at 10m/s") }
}
