package models

import (
	"api/db"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const RubroCollection = "rubro"

type Rubro struct {
	Id               bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	Codigo           string          `json:"codigo"`
	Nombre           string          `json:"nombre"`
	Entidad          string          `json:"entidad"`
	Descripcion      string          `json:"descripcion"`
	Unidad_ejecutora int             `json:"unidad_ejecutora"`
	Apropiaciones    []Apropiaciones `json:"apropiaciones"`
}

func UpdateRubro(session *mgo.Session, j Rubro, id string) error {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertRubro(session *mgo.Session, j Rubro) {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllRubros(session *mgo.Session) []Rubro {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	fmt.Println("Getting all rubros")
	var rubros []Rubro
	err := c.Find(bson.M{}).All(&rubros)
	if err != nil {
		fmt.Println(err)
	}
	return rubros
}

func GetRubroById(session *mgo.Session, id string) ([]Rubro, error) {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	var rubros []Rubro
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&rubros)
	return rubros, err
}

func DeleteRubroById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
