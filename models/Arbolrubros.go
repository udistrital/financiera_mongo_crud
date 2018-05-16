package models

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const ArbolRubrosCollection = "arbolrubros"

type ArbolRubros struct {
	Id string `json:"_id" bson:"_id,omitempty"`
  Idpsql int `json:"idpsql"`
  Nombre string `json:"nombre"`
  Descripcion string `json:"descripcion"`
  Hijos []string `json:"hijos"`
  Padre string `json:"padre"`
}


func UpdateArbolRubros(session *mgo.Session, j ArbolRubros, id string) error{
	c := db.Cursor(session,ArbolRubrosCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertArbolRubros(session *mgo.Session, j ArbolRubros) {
	c := db.Cursor(session,ArbolRubrosCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllArbolRubross(session *mgo.Session) []ArbolRubros {
	c := db.Cursor(session,ArbolRubrosCollection)
	defer session.Close()
    fmt.Println("Getting all arbolrubross")
	var arbolrubross []ArbolRubros
	err := c.Find(bson.M{}).All(&arbolrubross)
	if err != nil {
		fmt.Println(err)
	}
	return arbolrubross
}

func GetArbolRubrosById(session *mgo.Session,id string) ([]ArbolRubros,error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var arbolrubross []ArbolRubros
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&arbolrubross)
	return arbolrubross,err
}

func DeleteArbolRubrosById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
