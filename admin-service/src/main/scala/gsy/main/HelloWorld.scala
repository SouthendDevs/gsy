package gsy.main

import org.scalatra._

class HelloWorld extends ScalatraFilter {
  get("/") {"Hello world!"}

  get("/config") {"Your config page goes here!"}
}