package main

import "fmt"

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
    if(car.nextJunction == car.destinationJunction) {
      fmt.Print("Car ",id," has reached its destination")
    } else {
      car.currentRoad = car.nextJunction.roadsOut[0]
      car.distanceToNextJunction = car.currentRoad.length
      car.nextJunction = car.currentRoad.nextJunction
      car.speed = car.currentRoad.speed
    }
  } else {
    car.distanceToNextJunction -= car.speed
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
  fmt.Println(" ")
  fmt.Println("print all cars:")
  for _, value := range allCars {
    fmt.Println(value)
  }
  
  fmt.Println("update all cars:")
  for key, value := range allCars {
    updateCar(value, key)
  }
  
  fmt.Println(" ")
}

func main() {
  fmt.Println("Hello!")
  //declare instances
  allCars = make(map[int]*Car)
   
  junct1 := Junction{27, 1, make(map[int]*Road) }
  junct2 := Junction{28, 2, make(map[int]*Road) }
  myroad := Road{23, 200,50, nil }
  
  //connect roads and junctions
  // j1 --myroad-> j2
  junct1.roadsOut[0] = &myroad
  myroad.nextJunction = &junct2
  
  fmt.Println("Go says I have to print this or no compilation for you: ")
  fmt.Println("junct1:", junct1)
  fmt.Println("myroad:", &myroad, " which is ", myroad)
  fmt.Println("junct2:", junct2)
  
  allCars[0] = &Car{nil, &junct1, 0, 0, &junct2}
  
  //fmt.Println("Car: ", allCars[0])

  update()
  update()
  update()
  update()
  update()
  update()
  update()
  update()
  update()
  
}
