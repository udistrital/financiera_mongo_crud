package models

import (
	"fmt"

	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const MovimientosCollection = "movimientos"

type MovimientoCdp struct {
	ID            string                   `json:"_id" bson:"_id,omitempty"`
	IDPsql        string                   `json:"idpsql"`
	RubrosAfecta  []map[string]interface{} `json:"rubros_afecta"`
	ValorOriginal float64                  `json:"valor_original"`
	// TotalAnulado      float64                  `json:"total_anulado"`
	// TotalComprometido float64                  `json:"total_comprometido"`
	Tipo            string `json:"tipo"`
	Vigencia        string `json:"vigencia"`
	DocumentoPadre  string `json:"documento_padre"`
	FechaRegistro   string `json:"fecha_registro"`
	UnidadEjecutora string `json:"unidad_ejecutora"`
}

// GetMovimientoByPsqlId Obtener un documento por el idpsql
func GetMovimientoByPsqlId(session *mgo.Session, id, tipo string) (*MovimientoCdp, error) {
	c := db.Cursor(session, MovimientosCollection)
	defer session.Close()
	var movimientoCdp *MovimientoCdp
	err := c.Find(bson.M{"idpsql": id, "tipo": tipo}).One(&movimientoCdp)
	return movimientoCdp, err
}

func EstrctTransaccionMov(session *mgo.Session, estructura *MovimientoCdp) (ops txn.Op, err error) {
	// id :=
	estructura.ID = bson.NewObjectId().Hex()
	op := txn.Op{
		C:      MovimientosCollection,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}

func EstrctUpdateTransaccionMov(session *mgo.Session, estructura *MovimientoCdp) (ops txn.Op, err error) {
	fmt.Println("Id:    ", estructura.ID)
	// estructura.ID = bson.NewObjectId().Hex()
	op := txn.Op{
		C:      MovimientosCollection,
		Id:     estructura.ID,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"rubrosafecta", estructura.RubrosAfecta}}}},
	}
	return op, err
}
