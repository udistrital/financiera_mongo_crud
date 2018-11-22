package models

import (
	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

// FuenteFinaciamientoPadre ...
type FuenteFinaciamientoPadre struct {
	ID              string      `orm:"size(128)"`
	UnidadEjecutora int         `json:"unidad_ejecutora"`
	Descripcion     string      `json:"descripcion"`
	IDPsql          int         `json:"idpsql"`
	Nombre          string      `json:"nombre"`
	TipoFuente      interface{} `json:"tipo_fuente"`
	ValorOriginal   float64     `json:"valor_original"`
}

// FuenteMovimiento ...
type FuenteMovimiento struct {
	ID                string
	Rubro             string
	DependenciaIDPsql string
	Saldo             float64
}

// ArbolRubroApropiacion2018Collection constante para la colección
const fuenteFinanciamientoPadre = "fuente_financiamiento_padre"
const fuenteMovimiento = "fuente_movimiento"

// InsertFuenteMovimiento función para registrar un documento de tipo fuente_movimiento
func InsertFuenteMovimiento(session *mgo.Session, j *FuenteMovimiento) {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	c.Insert(&j)
}

// InsertFuentFinanciamientoPadre función para registrar un documento de tipo fuente_financiamiento_padre
func InsertFuentFinanciamientoPadre(session *mgo.Session, j *FuenteFinaciamientoPadre) {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	c.Insert(&j)
}

// GetFuenteFinanciamientoPadreByID Obtener un documento por el id
func GetFuenteFinanciamientoPadreByID(session *mgo.Session, id string) *FuenteFinaciamientoPadre {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	var fuenteFinaciamientoPadre *FuenteFinaciamientoPadre
	err := c.Find(bson.M{"_id": id}).One(&fuenteFinaciamientoPadre)
	if err != nil {
		return nil
	}
	return fuenteFinaciamientoPadre
}

func EstructaRegistroFuentePadreTransaccion(session *mgo.Session, estructura *FuenteFinaciamientoPadre) (op txn.Op, err error) {
	estructura.ID = bson.NewObjectId().Hex()
	op = txn.Op{
		C:      fuenteFinanciamientoPadre,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}
