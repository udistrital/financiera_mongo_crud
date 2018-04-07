package models

import (
	"api/db"
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "reflect"
)

const ApropiacionesCollection = "apropiaciones"

type Apropiaciones struct {
	Id            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Vigencia      int           `json:"vigencia"`
	Valor_inicial float64       `json:"valor_inicial"`
	Movimientos   []Movimiento  `json:"movimientos"`
}

func UpdateApropiaciones(session *mgo.Session, j Apropiaciones, id string) error {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertApropiaciones(session *mgo.Session, j Apropiaciones) {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllApropiacioness(session *mgo.Session, query map[string]interface{}) []Apropiaciones {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	var buffer bytes.Buffer
	buffer.WriteString("")
	fmt.Println("Getting all apropiacioness")
	var apropiacioness []Apropiaciones
	// if len(query) != 0 {
	// 	for key, value := range query {
	// 		buffer.WriteString(key)
	// 		buffer.WriteString(":")
	// 		buffer.WriteString(value)
	// 		buffer.WriteString(",")
	// 	}
	// }
	fmt.Println(query)
	err := c.Find(query).All(&apropiacioness)
	if err != nil {
		fmt.Println(err)
	}
	return apropiacioness
}

func GetApropiacionesById(session *mgo.Session, id string) ([]Apropiaciones, error) {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	var apropiacioness []Apropiaciones
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&apropiacioness)
	return apropiacioness, err
}

func DeleteApropiacionesById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ApropiacionesCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
