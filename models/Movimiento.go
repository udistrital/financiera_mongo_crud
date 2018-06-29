package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const MovimientosCollection = "movimientos"

type MovimientoCdp struct {
	ID                string                   `json:"_id" bson:"_id,omitempty"`
	IDPsql            string                   `json:"idpsql"`
	RubrosAfecta      []map[string]interface{} `json:"rubros_afecta"`
	ValorOriginal     float64                  `json:"valor_original"`
	TotalAnulado      float64                  `json:"total_anulado"`
	TotalComprometido float64                  `json:"total_comprometido"`
	Vigencia          string                   `json:"vigencia"`
}

func EstrctTransaccionMov(session *mgo.Session, estructura *MovimientoCdp) (ops txn.Op, err error) {
	// id :=
	estructura.ID = bson.NewObjectId().Hex()
	op := txn.Op{
		C:      MovimientosCollection,
		Id:     estructura.ID,
		Assert: "d-",
		//Insert: bson.D{{"$set", bson.D{{"rubros_afecta", estructura.RubrosAfecta}}}},
		Insert: estructura,
	}
	return op, err
}
