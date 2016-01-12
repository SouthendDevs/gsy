package main

import "github.com/pusher/pusher-http-go"
import "fmt"
import "github.com/subosito/gotenv"
import "os"
import "math/rand"
import "math"
import "time"
import "strconv"

func init() {
    gotenv.Load()
} 
 

 
func main(){
  r := rand.New( rand.NewSource(time.Now().UnixNano()) )
  
  client := pusher.Client{
    AppId: os.Getenv("PUSHERAPPID"),
    Key: os.Getenv("PUSHERKEY"),
    Secret: os.Getenv("PUSHERSECRET"),
  }
  
  fmt.Println("Loaded credentials from env: Pusher AppId, Key, Secret: ", client)
 
  startTime := time.Now().UnixNano()
  var numPushes int64 = 50 

  for i := 0; i < int(numPushes); i++ {
    rNum := int( math.Floor(r.Float64() * 250) ) //0 to 250?
    
    fmt.Println("Pushing a random position: ", rNum)  
    
    data := map[string]string {
      "direction": "2", //1: move up
      "position": strconv.Itoa(rNum),  //amount to move - Itoa converts from integer to string
    }
  
    fmt.Println("Pushing data: ", data)
    
    client.Trigger("test_channel", "my_event", data)
    fmt.Println("pushed random position")
  }

  endTime := time.Now().UnixNano()
  var nsTotal int64 = (endTime - startTime)
  var nsPerPush = nsTotal/numPushes  

  fmt.Println( nsTotal/int64(1000000), " msec total (not including initialisation time)" )
  fmt.Println( nsPerPush/int64(1000000), " msec per push" )

}
