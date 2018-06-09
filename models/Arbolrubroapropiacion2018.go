package models

import (
	"fmt"

	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ArbolRubroApropiacion2018Collection = "arbolrubroapropiacion2018"

type ArbolRubroApropiacion2018 struct {
	Id                  string   `json:"_id" bson:"_id,omitempty"`
	Idpsql              string   `json:"idpsql"`
	Nombre              string   `json:"nombre"`
	Descripcion         string   `json:"descripcion"`
	Unidad_ejecutora    string   `json:"unidad_ejecutora"`
	Padre               string   `json:"padre"`
	Hijos               []string `json:"hijos"`
	Apropiacion_inicial int      `json:"apropiacion_inicial"`
}

func UpdateArbolRubroApropiacion2018(session *mgo.Session, j ArbolRubroApropiacion2018, id string) error {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func RegistrarApropiacion(session *mgo.Session, j ArbolRubroApropiacion2018, vigencia string) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	if err := c.Insert(j); err != nil {
		panic(err)
	}
}

func InsertArbolRubroApropiacion2018(session *mgo.Session, j ArbolRubroApropiacion2018) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	c.Insert(j)

}

func GetAllArbolRubroApropiacion2018s(session *mgo.Session) []ArbolRubroApropiacion2018 {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	fmt.Println("Getting all arbolrubroapropiacion2018s")
	var arbolrubroapropiacion2018s []ArbolRubroApropiacion2018
	err := c.Find(bson.M{}).All(&arbolrubroapropiacion2018s)
	if err != nil {
		fmt.Println(err)
	}
	return arbolrubroapropiacion2018s
}

func GetArbolRubroApropiacion2018ById(session *mgo.Session, id string) ([]ArbolRubroApropiacion2018, error) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	var arbolrubroapropiacion2018s []ArbolRubroApropiacion2018
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&arbolrubroapropiacion2018s)
	return arbolrubroapropiacion2018s, err
}

func DeleteArbolRubroApropiacion2018ById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
