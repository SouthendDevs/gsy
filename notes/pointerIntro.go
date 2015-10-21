//run me with "$ go run pointerIntro.go"

package main

import "fmt"

type User struct {
  name string
}

func demo1() {
  var x User
  x = User{"Andy"}
  
  //the = operator COPIES the object (:= is type inference - it declares y with the correct type automatically),
  y := x
  //so we now have 2 User objects
  
  y.name = "Dan"
  fmt.Println("demo1:",x,y)
}

func demo2() {
  var x *User //use pointer types everywhere - recommended
  x = &User{"Andy"} //we declare it and then get a pointer to it (get the memory address-of x with &x)
  //the = operator now copies the POINTER - x and y now both point(refer) to the same object, as in Ruby/Java/Python etc
  y := x
  y.name = "Dan"
  fmt.Println("demo2:",x,y) //x and y both point to Dan
}

func demo3() {
  var x User //no pointer type
  x = User{"Andy"}
  //& gets the address-of (pointer or reference to) - x and y again both point to the same object
  y := &x
  y.name = "Dan"
  fmt.Println("demo3:",x,y) //x and y both point to Dan
}

////////////////

func demo4func(user User) {
  //pass by value - a COPY of Andy has been passed in
  fmt.Println("inside demo4func(x):", user.name)
  //we edit the copy we have recieved
  user.name = "John"
  fmt.Println("inside demo4func(x):", user.name)
  //the copy is now destroyed by the GC, the original object is unmodified
}

func demo4() {
  fmt.Println("")
  var x User
  x = User{"Andy"}
  fmt.Println("before demo4func(x):", x.name)
  //the User object is COPIED into the function
  demo4func(x)
  //we still have our original User which was not modified
  fmt.Println("after demo4func(x):", x.name)
}

////////////////

func demo5func(user *User) {
  //this is still pass by value - the value(memory address) of the pointer(reference) has been passed 
  //which still refers to the original object, not a copy
  fmt.Println("inside demo5func(x):", user.name)
  //we edit the original object
  user.name = "John"
  fmt.Println("inside demo5func(x):", user.name)
}

func demo5() {
  fmt.Println("")
  var x *User //x now points to a User, it is not a User itself
  x = &User{"Andy"} //we declare it and then get a pointer to it (get the memory address-of x with &x)
  fmt.Println("before demo5func(x):", x.name)
  //the User pointer is COPIED into the function - so the function can access the original object
  demo5func(x)
  //the original User is now modified, as you would expect
  fmt.Println("after demo5func(x):", x.name)
}


func main() {
  demo1()
  demo2()
  demo3()
  demo4()
  demo5()
}
