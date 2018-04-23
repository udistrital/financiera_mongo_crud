package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const RegistroPresupuestalCollection = "registropresupuestal"

type RegistroPresupuestal struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Vigencia int `json:"vigencia"`
  FechaRegistro string `json:"fecha_registro"`
  Estado string `json:"estado"`
  Numero_Registro_Presupuestal int `json:"numero_registro_presupuestal"`
  Solicitud int`json:"solicitud"`
  Disponibilidad_apropiacion DisponibilidadApropiacion `json:"disponibilidad_apropiacion"`
}


func UpdateRegistroPresupuestal(session *mgo.Session, j RegistroPresupuestal, id string) error{
	c := db.Cursor(session,RegistroPresupuestalCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertRegistroPresupuestal(session *mgo.Session, j RegistroPresupuestal) bson.ObjectId {
	c := db.Cursor(session,RegistroPresupuestalCollection)
	defer session.Close()
  j.Id = bson.NewObjectId()
	c.Insert(j)
  return j.Id

}

func GetAllRegistroPresupuestals(session *mgo.Session, query map[string]interface{}) []RegistroPresupuestal {
	c := db.Cursor(session,RegistroPresupuestalCollection)
	defer session.Close()
    fmt.Println("Getting all registropresupuestals")
	var registropresupuestals []RegistroPresupuestal
	err := c.Find(query).All(&registropresupuestals)
	if err != nil {
		fmt.Println(err)
	}
	return registropresupuestals
}

func GetRegistroPresupuestalById(session *mgo.Session,id string) ([]RegistroPresupuestal,error) {
	c := db.Cursor(session, RegistroPresupuestalCollection)
	defer session.Close()
	var registropresupuestals []RegistroPresupuestal
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&registropresupuestals)
	return registropresupuestals,err
}

func DeleteRegistroPresupuestalById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, RegistroPresupuestalCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
