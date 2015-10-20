import javax.servlet.ServletContext

import gsy.main.HelloWorld
import org.scalatra.LifeCycle

class ScalatraBootstrap extends LifeCycle {
  override def init(context: ServletContext) {
    context mount (new HelloWorld, "/config/*")
  }
}