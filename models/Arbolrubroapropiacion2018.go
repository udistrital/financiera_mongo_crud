package models

import (
	"fmt"

	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ArbolRubroApropiacion2018Collection = "arbolrubroapropiacion2018"
const ArbolRubroApropiacionCollection = "arbolrubroapropiacion"

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

type ArbolRubroApropiacion struct {
	Id                  string   `json:"_id" bson:"_id,omitempty"`
	Idpsql              string   `json:"idpsql"`
	Nombre              string   `json:"nombre"`
	Descripcion         string   `json:"descripcion"`
	Unidad_ejecutora    string   `json:"unidad_ejecutora"`
	Padre               string   `json:"padre"`
	Hijos               []string `json:"hijos"`
	Apropiacion_inicial int      `json:"apropiacion_inicial"`
}

func UpdateArbolRubroApropiacion(session *mgo.Session, j ArbolRubroApropiacion, id, vigencia string) error {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+vigencia)
	defer session.Close()
	// Update
	fmt.Println("id update: ", id)
	err := c.Update(bson.M{"_id": id}, &j)
	if err != nil {
		fmt.Println("updatw error")
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

func InsertArbolRubroApropiacion(session *mgo.Session, j *ArbolRubroApropiacion, vigencia string) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+vigencia)
	defer session.Close()
	c.Insert(&j)

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

func GetArbolRubroApropiacionById(session *mgo.Session, id, vigencia string) (*ArbolRubroApropiacion, error) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+vigencia)
	defer session.Close()
	var arbolRubroApropiacion *ArbolRubroApropiacion
	err := c.Find(bson.M{"_id": id}).One(&arbolRubroApropiacion)
	fmt.Println(arbolRubroApropiacion, " | ", err)
	return arbolRubroApropiacion, err
}

func DeleteArbolRubroApropiacion2018ById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
