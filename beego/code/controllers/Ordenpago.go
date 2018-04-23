package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud OrdenPago
type OrdenPagoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 OrdenPago models.OrdenPago
// @Failure 403 :objectId is empty
// @router / [get]
func (j *OrdenPagoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllOrdenPagos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get OrdenPago by nombre
// @Param	nombre		path 	string	true		"El nombre de la OrdenPago a consultar"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *OrdenPagoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		ordenpago, err := models.GetOrdenPagoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = ordenpago
		}
	}
	j.ServeJSON()
}

// @Title Borrar OrdenPago
// @Description Borrar OrdenPago
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *OrdenPagoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteOrdenPagoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear OrdenPago
// @Description Crear OrdenPago
// @Param	body		body 	models.OrdenPago	true		"Body para la creacion de OrdenPago"
// @Success 200 {int} OrdenPago.Id
// @Failure 403 body is empty
// @router / [post]
func (j *OrdenPagoController) Post() {
	var ordenpago models.OrdenPago
	json.Unmarshal(j.Ctx.Input.RequestBody, &ordenpago)
	fmt.Println(ordenpago)
	session,_ := db.GetSession()
	models.InsertOrdenPago(session,ordenpago)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the OrdenPago
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *OrdenPagoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var ordenpago models.OrdenPago
	json.Unmarshal(j.Ctx.Input.RequestBody, &ordenpago)
	session,_ := db.GetSession()

	err := models.UpdateOrdenPago(session, ordenpago,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear OrdenPago
// @Param	body		body 	models.OrdenPago	true		"Body para la creacion de OrdenPago"
// @Success 200 {int} OrdenPago.Id
// @Failure 403 body is empty
// @router / [options]
func (j *OrdenPagoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear OrdenPago
// @Param	body		body 	models.OrdenPago true		"Body para la creacion de OrdenPago"
// @Success 200 {int} OrdenPago.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *OrdenPagoController) OrdenPagoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}