package gsy.main

import org.scalatra._
import org.json4s.{DefaultFormats, Formats}
import org.scalatra.json._

case class GsyParams(simName: String, numberOfCars: Int)

class HelloWorld extends ScalatraFilter with JacksonJsonSupport {
  protected implicit val jsonFormats: Formats = DefaultFormats

  before() {
    contentType = formats("json")
  }

  get("/") {
    "Hello world!"
  }

  get("/config") {
    "Your config page goes here!"
  }

  get("/config/params") {
    GsyParams("Test Sim", 1)
  }

}