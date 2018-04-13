package models

import (
	"api/db"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const MovimientoCollection = "movimiento"

type Movimiento struct {
	Id                  bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Numero              string        `json:"numero"`
	Estado_movimiento   string        `json:"estado_movimiento"`
	Fecha_movimiento    time.Time     `json:"fecha_movimiento"`
	Numero_oficio       int           `json:"numero_oficio"`
	Fecha_oficio        time.Time     `json:"fecha_oficio"`
	Descripcion         string        `json:"descripcion"`
	Unidad_ejecutora    int           `json:"unidad_ejecutora"`
	Apropiacion_destino int           `json:"apropiacion_destino"`
	Apropiacion_origen  int           `json:"apropiacion_origen"`
	Valor               float64       `json:"valor"`
	Tipo_movimiento     string        `json:"tipo_movimiento"`
}

func UpdateMovimiento(session *mgo.Session, j Movimiento, id string) error {
	c := db.Cursor(session, MovimientoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertMovimiento(session *mgo.Session, j Movimiento) {
	c := db.Cursor(session, MovimientoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllMovimientos(session *mgo.Session) []Movimiento {
	c := db.Cursor(session, MovimientoCollection)
	defer session.Close()
	fmt.Println("Getting all movimientos")
	var movimientos []Movimiento
	err := c.Find(bson.M{}).All(&movimientos)
	if err != nil {
		fmt.Println(err)
	}
	return movimientos
}

func GetMovimientoById(session *mgo.Session, id string) ([]Movimiento, error) {
	c := db.Cursor(session, MovimientoCollection)
	defer session.Close()
	var movimientos []Movimiento
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&movimientos)
	return movimientos, err
}

func DeleteMovimientoById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, MovimientoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
