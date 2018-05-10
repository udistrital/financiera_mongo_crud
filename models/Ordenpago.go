package models

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const OrdenPagoCollection = "ordenpago"

type OrdenPago struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Vigencia int `json:"vigencia"`
  Valor_base int `json:"valor_base"`
  Unidad_ejecutora int `json:"unidad_ejecutora"`
  Forma_pago int `json:"forma_pago"`
  Registro_presupuestal RegistroPresupuestal `json:"registro_presupuestals"`
}


func UpdateOrdenPago(session *mgo.Session, j OrdenPago, id string) error{
	c := db.Cursor(session,OrdenPagoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertOrdenPago(session *mgo.Session, j OrdenPago) {
	c := db.Cursor(session,OrdenPagoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllOrdenPagos(session *mgo.Session) []OrdenPago {
	c := db.Cursor(session,OrdenPagoCollection)
	defer session.Close()
    fmt.Println("Getting all ordenpagos")
	var ordenpagos []OrdenPago
	err := c.Find(bson.M{}).All(&ordenpagos)
	if err != nil {
		fmt.Println(err)
	}
	return ordenpagos
}

func GetOrdenPagoById(session *mgo.Session,id string) ([]OrdenPago,error) {
	c := db.Cursor(session, OrdenPagoCollection)
	defer session.Close()
	var ordenpagos []OrdenPago
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&ordenpagos)
	return ordenpagos,err
}

func DeleteOrdenPagoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, OrdenPagoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}
