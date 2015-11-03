//run me with "$ go run pointerIntro.go"

package main

import "fmt"

type User struct {
  name string
}

func demo1() {
  //make a variable x of type User - initialised to nil to start with
  var x User
  
  //construct a User
  x = User{"Andy"}
  
  //the = operator copies the object itself (:= is type inference - it declares y with the correct type automatically),
  y := x
  //so we now have 2 User objects
  
  y.name = "Dan"
  fmt.Println("demo1:",x,y)
}

func demo2() {
  //make a variable x of type *User - User pointer -  initialised to nil to start with
  //use pointer types for all objects - recommended
  var x *User
  
  //Construct a User then get a pointer to it (get the memory address-of x with &x)
  x = &User{"Andy"}
  
  //the = operator now copies the POINTER - x and y are now both pointers that point(refer) to the same object, 
  //like objects in Ruby/Java/Python etc by default
  y := x
  //the pointer itself doesn't have a name field - Go automatically dereferences it to the object we want
  // - we can always use .
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
  //type inference: x is auto-initialised to a User (not *User)
  x := User{"Andy"}
  fmt.Println("before demo4func(x):", x.name)
  //the User object is COPIED into the function
  demo4func(x)
  //we still have our original User which was not modified
  fmt.Println("after demo4func(x):", x.name)
}

////////////////

func demo5func(user *User) {
  //this is still pass by value - the value(memory address) of the pointer(reference) has been passed in
  //which still refers to the original object, not a copy of the object itself
  fmt.Println("inside demo5func(x):", user.name)
  //we edit the original object - Go automatically deferences it with .
  user.name = "John"
  fmt.Println("inside demo5func(x):", user.name)
}

func demo5() {
  fmt.Println("")
  //type inference: (:=) x is auto-initialised to a *User (not User)
  x := &User{"Andy"} //we construct it and then get a pointer to it (get the memory address-of x with &x)
  fmt.Println("before demo5func(x):", x.name)
  //the User POINTER is copied into the function - so the function can access the original object
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




