package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Movimiento
type MovimientoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Movimiento models.Movimiento
// @Failure 403 :objectId is empty
// @router / [get]
func (j *MovimientoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllMovimientos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Movimiento by nombre
// @Param	nombre		path 	string	true		"El nombre de la Movimiento a consultar"
// @Success 200 {object} models.Movimiento
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *MovimientoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		movimiento, err := models.GetMovimientoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = movimiento
		}
	}
	j.ServeJSON()
}

// @Title Borrar Movimiento
// @Description Borrar Movimiento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *MovimientoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteMovimientoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Movimiento
// @Description Crear Movimiento
// @Param	body		body 	models.Movimiento	true		"Body para la creacion de Movimiento"
// @Success 200 {int} Movimiento.Id
// @Failure 403 body is empty
// @router / [post]
func (j *MovimientoController) Post() {
	var movimiento models.Movimiento
	json.Unmarshal(j.Ctx.Input.RequestBody, &movimiento)
	fmt.Println(movimiento)
	session,_ := db.GetSession()
	models.InsertMovimiento(session,movimiento)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Movimiento
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *MovimientoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var movimiento models.Movimiento
	json.Unmarshal(j.Ctx.Input.RequestBody, &movimiento)
	session,_ := db.GetSession()

	err := models.UpdateMovimiento(session, movimiento,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Movimiento
// @Param	body		body 	models.Movimiento	true		"Body para la creacion de Movimiento"
// @Success 200 {int} Movimiento.Id
// @Failure 403 body is empty
// @router / [options]
func (j *MovimientoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Movimiento
// @Param	body		body 	models.Movimiento true		"Body para la creacion de Movimiento"
// @Success 200 {int} Movimiento.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *MovimientoController) MovimientoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}