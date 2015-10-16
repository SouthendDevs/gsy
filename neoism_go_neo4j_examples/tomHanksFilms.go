package main

import "fmt"
import "github.com/jmcvetta/neoism"

var db *neoism.Database //The * means this is a pointer to a Database object

func initdb() {
  var err error
  
  //neoism.Connect("http://username:password@server:port")
  panic("This script works on the default example films and actors database, you will need that running on localhost. Put your password in the Connect statement below then comment out this panic line. The default user is neo4j")
  
  db, err = neoism.Connect("http://neo4j:yourpassword@localhost:7474/")
  if(err != nil) { panic(err) }
  if(db == nil) { panic("Database object is null") }
}

////////////////////////////////


func queryTomHanks() {
  //it only works with the variable name "movie" for some reason - that is the type of the node
  stmt := `
   MATCH (tom {name:"Tom Hanks"})-[:ACTED_IN]->(movie)
   RETURN movie
  `

  // query results
  res := []struct {
    Movie neoism.Node
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
    n := res[i].Movie // Only one row of data returned
    fmt.Printf("  Node %d: %+v\n", i, n.Data )
  }
  
}

/////////////////////////////////////////

func main() {
  fmt.Println("Hello")
  initdb()
  fmt.Println("Querying films Tom Hanks acted in:")
  queryTomHanks()
}
