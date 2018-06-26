package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
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
	Id                  string                        `json:"_id" bson:"_id,omitempty"`
	Idpsql              string                        `json:"idpsql"`
	Nombre              string                        `json:"nombre"`
	Descripcion         string                        `json:"descripcion"`
	Unidad_ejecutora    string                        `json:"unidad_ejecutora"`
	Padre               string                        `json:"padre"`
	Hijos               []string                      `json:"hijos"`
	Apropiacion_inicial int                           `json:"apropiacion_inicial"`
	Movimientos         map[string]map[string]float64 `json:"movimientos"`
}

// type Movimiento struct {
// 	Mes_modificacion        float64 `json:"mes_modificacion"`
// 	Total_modificacion      float64 `json:"total_modificacion"`
// 	Total_adicion           float64 `json:"total_adicion"`
// 	Total_reduccion         float64 `json:"total_reduccion"`
// 	Traslados_total_credito float64 `json:"traslados_total_credito"`
// 	Traslados_contracredito float64 `json:"traslados_contracredito"`
// 	Total_anulado           float64 `json:"total_anulado"`
// 	Mes_cdp                 float64 `json:"mes_cdp"`
// 	Total_cdp               float64 `json:"total_cdp"`
// 	Mes_rp                  float64 `json:"mes_rp"`
// 	Total_rp                float64 `json:"total_rp"`
// 	Mes_op                  float64 `json:"mes_op"`
// 	Total_giro              float64 `json:"total_giro"`
// 	Mes_giro                float64 `json:"mes_giro"`
// }

func UpdateArbolRubroApropiacion(session *mgo.Session, j ArbolRubroApropiacion, id, ue, vigencia string) error {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
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

func RegistrarApropiacion(session *mgo.Session, j ArbolRubroApropiacion2018, ue, vigencia string) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	if err := c.Insert(j); err != nil {
		panic(err)
	}
}

func InsertArbolRubroApropiacion(session *mgo.Session, j *ArbolRubroApropiacion, ue, vigencia string) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
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

func GetArbolRubroApropiacionById(session *mgo.Session, id, ue, vigencia string) (*ArbolRubroApropiacion, error) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
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

func GetNodoApropiacion(session *mgo.Session, id, ue, vigencia string) (ArbolRubroApropiacion, error) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	var nodo ArbolRubroApropiacion
	err := c.Find(bson.M{"_id": id}).One(&nodo)
	return nodo, err
}

func GetRaicesApropiacion(session *mgo.Session, ue, vigencia string) ([]ArbolRubroApropiacion, error) {
	var roots []ArbolRubroApropiacion
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	err := c.Find(bson.M{
		"$or": []bson.M{bson.M{"padre": nil},
			bson.M{"padre": ""}},
		"idpsql":           bson.M{"$ne": nil},
		"unidad_ejecutora": bson.M{"$in": []string{"0", ue}},
	}).All(&roots)
	beego.Info("roots: ", roots)
	return roots, err
}

func CrearEstrctTransaccion(session *mgo.Session, estructuras []*ArbolRubroApropiacion, ue, vigencia string) error {
	var ops []txn.Op
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	runner := txn.NewRunner(c)
	for _, estructura := range estructuras {
		op := txn.Op{
			C:      ArbolRubroApropiacionCollection + "_" + vigencia + "_" + ue,
			Id:     estructura.Id,
			Assert: "d-",
			Update: estructura,
		}
		fmt.Println("estructura: ", estructura)
		ops = append(ops, op)
	}
	fmt.Println("ops: ", ops)
	id := bson.NewObjectId()
	err := runner.Run(ops, id, nil)
	return err
	// c := db.Cursor(session, ArbolRubrosCollection)
	// runner := txn.NewRunner(c)
	// ops := []txn.Op{{
	// 	C:      ArbolRubrosCollection,
	// 	Id:     rubroHijo.Id,
	// 	Assert: "d-",
	// 	Insert: rubroHijo,
	// }, {
	// 	C:      ArbolRubrosCollection,
	// 	Id:     rubroPadre.Id,
	// 	Assert: "d+",
	// 	Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	// }}
	// id := bson.NewObjectId() // Optional
	// err := runner.Run(ops, id, nil)
	// return err

	// return nil
}
