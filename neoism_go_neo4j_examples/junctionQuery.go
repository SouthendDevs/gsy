package main

import "fmt"
import "github.com/jmcvetta/neoism"
import "strconv"

const PASSWORD = "password" //Put the password in here, Mike posted it in the slack

var db *neoism.Database //The * means this is a pointer to a Database object

func initdb() {
  //neoism.Connect("http://username:password@server:port")
  var err error
  db, err = neoism.Connect("http://car_simulator:" + PASSWORD + "@carsimulator.sb05.stations.graphenedb.com:24789")
  //If error connecting to database, halt execution and display error message
  if(err != nil) { panic(err) }
  if(db == nil) { panic("Database object is null") }
}

////////////////////////////////


func queryJunction(junctNum int) {
  stmt := `
   MATCH (jone {id:"` + strconv.Itoa(junctNum) + `"})-[:ROAD]->(junction)
   RETURN junction
  `

  // query results
  res := []struct {
    Junction neoism.Node
  }{}

  // construct query
  cq := neoism.CypherQuery{
    Statement:  stmt,
    Result:     &res,
  }
  // execute query
  err := db.Cypher(&cq)
  if(err != nil){ panic(err) } //I would prefer exceptions but Go lacks these

  fmt.Printf("query result:\n")
  for i, _ := range res {
    n := res[i].Junction // Only one row of data returned
    fmt.Printf("  Node %d: %+v\n", i, n.Data )
  }
  
}

/////////////////////////////////////////

func main() {
  fmt.Println("Hello, connecting to gsy database...")
  initdb()
  fmt.Print("Which junction number do you want to query (1-9): ")
  junctNum := cliInputNum()
  fmt.Println("Querying junctions connected to junction", junctNum, "with ROAD relations: ") //using commas with Print/Println like this puts spaces in between
  queryJunction(junctNum)
}

func cliInputNum() int {
  var numStr string //auto-initialised to "", annoyingly
  fmt.Scanf("%s", &numStr) //& gets the pointer to - address of - numStr so Scanf can change its value
  // := will declare these variables to be the correct types returned by Atoi - int and error
  num, err := strconv.Atoi(numStr) //try to convert the string to int
  //if error converting the string, halt execution(panic) and display error message
  if(err != nil) {fmt.Println("Cannot convert str to int"); panic(err)} 
  return num
}
