package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const DisponibilidadApropiacionCollection = "disponibilidadapropiacion"

type DisponibilidadApropiacion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Valor int `json:"valor"`
  Fuente_financiamiento int `json:"fuente_financiamiento"`
  Apropiacion Apropiacion `json:"apropiacion"`
}


func UpdateDisponibilidadApropiacion(session *mgo.Session, j DisponibilidadApropiacion, id string) error{
	c := db.Cursor(session,DisponibilidadApropiacionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertDisponibilidadApropiacion(session *mgo.Session, j DisponibilidadApropiacion) bson.ObjectId {
	c := db.Cursor(session,DisponibilidadApropiacionCollection)
	defer session.Close()
  j.Id = bson.NewObjectId()
	c.Insert(j)
  return j.Id

}

func GetAllDisponibilidadApropiacions(session *mgo.Session, query map[string]interface{}) []DisponibilidadApropiacion {
	c := db.Cursor(session,DisponibilidadApropiacionCollection)
	defer session.Close()
    fmt.Println("Getting all disponibilidadapropiacions")
	var disponibilidadapropiacions []DisponibilidadApropiacion
	err := c.Find(query).All(&disponibilidadapropiacions)
	if err != nil {
		fmt.Println(err)
	}
	return disponibilidadapropiacions
}

func GetDisponibilidadApropiacionById(session *mgo.Session,id string) ([]DisponibilidadApropiacion,error) {
	c := db.Cursor(session, DisponibilidadApropiacionCollection)
	defer session.Close()
	var disponibilidadapropiacions []DisponibilidadApropiacion
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&disponibilidadapropiacions)
	return disponibilidadapropiacions,err
}

func DeleteDisponibilidadApropiacionById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, DisponibilidadApropiacionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
