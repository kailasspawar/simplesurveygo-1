package main

import (
   "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "sync"
    "time"
  )

 type Movies struct {
   Title string `json:"title,omitempty"`
   Year int  `json:"year,omitempty"`
   Director string `json:"director,omitempty"`
   Cast string `json:"cast,omitempty"`
   Genre string `json:"genre,omitempty"`
   Notes string `json:"notes,omitempty"`
}

func main(){
  
         session, err := mgo.Dial("127.0.0.1")
                 if err != nil {
                        panic(err)
                }
        defer session.Close()
     
    var waitGroup sync.WaitGroup

  waitGroup.Add(4)
	for query := 0; query < 4; query++ {
		go RunQuery(query, &waitGroup, session)	
   }

	waitGroup.Wait()
	log.Println("All Queries Completed")

}

func RunQuery(query int, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {

 	defer waitGroup.Done()
      sessionCopy := session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("prem").C("movies")

	log.Printf("RunQuery : %d : Executing\n", query)

	var movie []Movies
	 if err := c.Insert(&movie); err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	log.Printf("RunQuery : %d :", query)
}

