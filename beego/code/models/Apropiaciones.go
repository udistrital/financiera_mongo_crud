package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const ApropiacionesCollection = "apropiaciones"

type Apropiaciones struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Vigencia int `json:"vigencia"`
  Valor_inicial float64 `json:"valor_inicial"`
  Movimientos []Movimiento `json:"movimientoss"`
}


func UpdateApropiaciones(session *mgo.Session, j Apropiaciones, id string) error{
	c := db.Cursor(session,ApropiacionesCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertApropiaciones(session *mgo.Session, j Apropiaciones) {
	c := db.Cursor(session,ApropiacionesCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllApropiacioness(session *mgo.Session) []Apropiaciones {
	c := db.Cursor(session,ApropiacionesCollection)
	defer session.Close()
    fmt.Println("Getting all apropiacioness")
	var apropiacioness []Apropiaciones
	err := c.Find(bson.M{}).All(&apropiacioness)
	if err != nil {
		fmt.Println(err)
	}
	return apropiacioness
}

func GetApropiacionesById(session *mgo.Session,id string) ([]Apropiaciones,error) {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	var apropiacioness []Apropiaciones
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&apropiacioness)
	return apropiacioness,err
}

func DeleteApropiacionesById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}