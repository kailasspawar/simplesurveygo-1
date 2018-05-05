package dao

import (
         "gopkg.in/mgo.v2"
         "gopkg.in/mgo.v2/bson"
    )

 
func deactivate() interface{} {
session := MgoSession.Clone()
        defer session.Close()
   var response []interface{}

   c := session.DB("simplesurveys").C("survey")
  
   query := c.RemoveAll(bson.M{"status":false})
   err := query.All(&response)
        if err!= nil{
		return nil
	}else{
      return   response
   }
}

func main(){

   go deactivate()
}

