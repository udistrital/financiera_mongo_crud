package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const ArbolRubrosCollection = "arbol_rubro"

type ArbolRubros struct {
	Id          string   `json:"_id" bson:"_id,omitempty"`
	Idpsql      int      `json:"idpsql"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Hijos       []string `json:"hijos"`
	Padre       string   `json:"padre"`
}

func UpdateArbolRubros(session *mgo.Session, j ArbolRubros, id string) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertArbolRubros(session *mgo.Session, j ArbolRubros) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	err := c.Insert(j)
	return err
}

func GetAllArbolRubross(session *mgo.Session) []ArbolRubros {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	fmt.Println("Getting all arbolrubross")
	var arbolrubross []ArbolRubros
	err := c.Find(bson.M{}).All(&arbolrubross)
	if err != nil {
		fmt.Println(err)
	}
	return arbolrubross
}

func GetArbolRubrosById(session *mgo.Session, id string) (ArbolRubros, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var arbolrubross ArbolRubros
	err := c.Find(bson.M{"_id": id}).One(&arbolrubross)
	return arbolrubross, err
}

func DeleteArbolRubrosById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}

func GetNodo(session *mgo.Session, id string) (ArbolRubros, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var nodo ArbolRubros
	err := c.Find(bson.M{"_id": id}).One(&nodo)
	return nodo, err
}

func RubroTransacton(rubroPadre, rubroHijo ArbolRubros, session *mgo.Session) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	beego.Info(rubroPadre.Id)
	runner := txn.NewRunner(c)
	beego.Error("error 0")
	ops := []txn.Op{{
		C:      "arbol_rubro",
		Id:     rubroHijo.Id,
		Assert: "d-",
		Insert: rubroHijo,
	}, {
		C:      "arbol_rubro",
		Id:     rubroPadre.Id,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	}}
	beego.Error("error 1")
	id := bson.NewObjectId() // Optional
	beego.Error("error 2")
	err := runner.Run(ops, id, nil)
	beego.Error("error 3")
	return err
}
