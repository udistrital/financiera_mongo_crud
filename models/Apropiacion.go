package models

import (
	"github.com/udistrital/financiera_mongo_crud/db"
	"fmt"
"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ApropiacionCollection = "apropiacion"

type Apropiacion struct {
	Id            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Vigencia      int           `json:"vigencia"`
	Valor_inicial int           `json:"valor_inicial"`
	Rubro         Rubro       `json:"rubro"`
}

func UpdateApropiacion(session *mgo.Session, j Apropiacion, id string) error {
	c := db.Cursor(session, ApropiacionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertApropiacion(session *mgo.Session, j Apropiacion) bson.ObjectId {
	c := db.Cursor(session, ApropiacionCollection)
	defer session.Close()
	j.Id = bson.NewObjectId()
	c.Insert(j)
  return j.Id
}

func GetAllApropiacions(session *mgo.Session, query map[string]interface{}) []Apropiacion {
	c := db.Cursor(session, ApropiacionCollection)
	defer session.Close()
	fmt.Println("Getting all apropiacions")
	fmt.Println(query)
	var apropiacion []Apropiacion
	err := c.Find(query).All(&apropiacion)
	if err != nil {
		fmt.Println(err)
	}
	return apropiacion
}

func GetApropiacionById(session *mgo.Session, id string) ([]Apropiacion, error) {
	c := db.Cursor(session, ApropiacionCollection)
	defer session.Close()
	var apropiacions []Apropiacion
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&apropiacions)
	return apropiacions, err
}

func DeleteApropiacionById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ApropiacionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
