package models

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const AgendaCollection = "agenda"

type Agenda struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
  Apellido string `json:"apellido"`
  Telefono string `json:"telefono"`
}


func UpdateAgenda(session *mgo.Session, j Agenda, id string) error{
	c := db.Cursor(session,AgendaCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertAgenda(session *mgo.Session, j Agenda) {
	c := db.Cursor(session,AgendaCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllAgendas(session *mgo.Session) []Agenda {
	c := db.Cursor(session,AgendaCollection)
	defer session.Close()
    fmt.Println("Getting all agendas")
	var agendas []Agenda
	err := c.Find(bson.M{}).All(&agendas)
	if err != nil {
		fmt.Println(err)
	}
	return agendas
}

func GetAgendaById(session *mgo.Session,id string) ([]Agenda,error) {
	c := db.Cursor(session, AgendaCollection)
	defer session.Close()
	var agendas []Agenda
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&agendas)
	return agendas,err
}

func DeleteAgendaById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, AgendaCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
