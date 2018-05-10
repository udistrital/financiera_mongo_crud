package models

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const RubroCollection = "rubro"

type Rubro struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Codigo string `json:"codigo"`
  Nombre string `json:"nombre"`
  Entidad string `json:"entidad"`
  Descripcion string `json:"descripcion"`
  Unidad_ejecutora int `json:"unidad_ejecutora"`
  Hijos []string `json:"hijos"`
}


func UpdateRubro(session *mgo.Session, j Rubro, id string) error{
	c := db.Cursor(session,RubroCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertRubro(session *mgo.Session, j Rubro) bson.ObjectId {
	c := db.Cursor(session,RubroCollection)
	defer session.Close()
  j.Id = bson.NewObjectId()
	c.Insert(j)
  return j.Id
}

func GetAllRubros(session *mgo.Session, query map[string]interface{}) []Rubro {
	c := db.Cursor(session,RubroCollection)
	defer session.Close()
    fmt.Println("Getting all rubros")
	var rubros []Rubro
	err := c.Find(query).All(&rubros)
	if err != nil {
		fmt.Println(err)
	}
	return rubros
}

func GetRubroById(session *mgo.Session,id string) ([]Rubro,error) {
	c := db.Cursor(session, RubroCollection)
	defer session.Close()
	var rubros []Rubro
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&rubros)
	return rubros,err
}

func DeleteRubroById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, RubroCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
