package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"

	"github.com/udistrital/financiera_mongo_crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const TransactionCollection = "transacciones"

func RegistrarMovimiento(session *mgo.Session, options []interface{}) (err error) {
	fmt.Println("Registrar movimiento....")

	try.This(func() {
		var ops []txn.Op
		c := db.Cursor(session, TransactionCollection)
		runner := txn.NewRunner(c)

		for _, op := range options {
			ops = append(ops, op.(txn.Op))
		}

		id := bson.NewObjectId()
		if err = runner.Run(ops, id, nil); err != nil {
			fmt.Errorf("%s \n", err.Error())
			panic(err.Error())
		}
	}).Catch(func(e try.E) {
		beego.Error("Error en RegistrarMovimiento: ", e)
		panic(e)
	})

	return err
}
