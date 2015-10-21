package main

import "fmt"
import "time"

var allCars map[int]*Car

type Car struct {
  currentRoad            *Road
  nextJunction           *Junction
  distanceToNextJunction float32
  speed                  float32
  destinationJunction    *Junction
}

func updateCar(car *Car, id int) {
  if(car.distanceToNextJunction <= 0) {
    if(car.nextJunction == nil && car.currentRoad == nil) {
      //this car is not in the simulation. TODO: remove from allCars
      return
    } else if(car.nextJunction == car.destinationJunction) {
      fmt.Print("Car ",id," has reached its destination")
      //Blank car's fields to show it cannot move anymore
      car.currentRoad = nil
      car.nextJunction = nil
      car.speed = 0
    } else {
      //Car picks route here:
      car.currentRoad = car.nextJunction.roadsOut[0]
      
      //Update car's fields with selected exit road from junction
      car.distanceToNextJunction = car.currentRoad.length
      car.nextJunction = car.currentRoad.nextJunction
      car.speed = car.currentRoad.speed
    }
  } else {
    //move along the road
    car.distanceToNextJunction -= car.speed
  }
}

func printCar(car *Car, id int) {
  if(car.nextJunction == nil && car.currentRoad == nil) {
    //this car no longer exists in the simulation. TODO: remove from allCars
    return
  } else {
    fmt.Print("Car ",id,": on road ")
    if car.currentRoad == nil { fmt.Print("none"); 
    } else { fmt.Print("id",car.currentRoad.neoId); } //else must have a } before it on the same line
    
    fmt.Print(", nextJunct:junct ",car.nextJunction.junctionNum,", distance to nextJunct: ",
      car.distanceToNextJunction," metres, speed: ", car.speed, " m/s")
    fmt.Println() //Println puts spaces between params, Print does not
  }
}

type Junction struct {
  neoId        int
  junctionNum  int
  roadsOut     map[int]*Road
}

type Road struct {
  neoId   int 
  length  float32
  speed   float32
  nextJunction *Junction
}

func update() {
  
  for key, value := range allCars {
    updateCar(value, key)
  }
  
  for id, car := range allCars {
    printCar(car, id)
  }

  fmt.Println("")
}

func main() {
  fmt.Println("Hello!")
  //declare instances
  allCars = make(map[int]*Car)
  
  //for objects(structs) always use pointer types for consistent syntax - so only use & for constructing and nowhere else
  junct1 := &Junction{27, 1, make(map[int]*Road) }
  junct2 := &Junction{28, 2, make(map[int]*Road) }
  myroad := &Road{23, 100,20, nil }
  
  //connect roads and junctions
  // j1 --myroad-> j2
  junct1.roadsOut[0] = myroad
  myroad.nextJunction = junct2
  
  allCars[0] = &Car{nil, junct1, 0, 0, junct2}
  
  //in Go all loops are made with a for statement, while(x) {...} is written for x {...}
  //brackets are not allowed around the for statement itself, strangely
  fmt.Println("Car init: "); printCar(allCars[0], 0)
  fmt.Println("")
  time.Sleep(1 * time.Second)
  for i := 0; i<10; i++ { 
    update()
    //sleep, just for a sense of "movement"
    time.Sleep(1 * time.Second)
  }
}
