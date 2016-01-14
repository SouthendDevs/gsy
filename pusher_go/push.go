package main

import "github.com/pusher/pusher-http-go"
import "fmt"
import "github.com/subosito/gotenv"
import "os"
import "math/rand"
//import "math"       //"imported and not used" errors - why Google, why?
import "time"
import "strconv"

//import "reflect" //TEMP

var r *rand.Rand //pointer to a Rand object
var pushClient pusher.Client

//rNum := int( math.Floor(r.Float64() * 250) ) //0 to 250?

func init() {
  gotenv.Load()
  r = rand.New( rand.NewSource(time.Now().UnixNano()) )
    
  if(os.Getenv("PUSHERAPPID") == "") { panic(".env file or Pusher AppID missing! Ensure .env file from the #gosimulateyourself channel has been placed in this directory") }
  
  pushClient = pusher.Client{
    AppId: os.Getenv("PUSHERAPPID"),
    Key: os.Getenv("PUSHERKEY"),
    Secret: os.Getenv("PUSHERSECRET"),
  }
  
  fmt.Println("Loaded Pusher credentials from .env: AppId, Key, Secret: ", pushClient)
} 

func toString(num int) string {
  return strconv.Itoa(num) //itoa and atoi are the functions in C
}

func main(){
  fmt.Println("This is only for push and rendering testing, there is no simulation happening yet")
  startTime := time.Now().UnixNano()
  var numPushes = 50 

  var x0 = 100;
  var x1 = 450;

  for i := 0; i < numPushes; i++ {
    
    
    //fmt.Println("Pushing a random position: ", rNum)  
    
    data := map[string]string {
      "paused": "0",
      "veh0": toString(x0) + ",225,e,b00", //x,y,direction(n/e/s/w), colour, 12bit in hex
      "veh1": toString(x1) + ",200,w,0b0",
    }
  
    //fmt.Println("Pushing data: ", data)
    pushClient.Trigger("test_channel", "my_event", data)
    
    x0 += 5
    x1 -= 5
    //time.Sleep(200 * time.Millisecond)
  }

  endTime := time.Now().UnixNano()
  var nsTotal int64 = (endTime - startTime)
  var nsPerPush = nsTotal/int64(numPushes)  

  fmt.Println( nsTotal/int64(1000000), " msec total (not including initialisation time)" )
  fmt.Println( nsPerPush/int64(1000000), " msec per push" )
}
