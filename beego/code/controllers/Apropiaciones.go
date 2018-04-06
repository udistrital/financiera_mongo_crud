package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Apropiaciones
type ApropiacionesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Apropiaciones models.Apropiaciones
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ApropiacionesController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllApropiacioness(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Apropiaciones by nombre
// @Param	nombre		path 	string	true		"El nombre de la Apropiaciones a consultar"
// @Success 200 {object} models.Apropiaciones
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ApropiacionesController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		apropiaciones, err := models.GetApropiacionesById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = apropiaciones
		}
	}
	j.ServeJSON()
}

// @Title Borrar Apropiaciones
// @Description Borrar Apropiaciones
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ApropiacionesController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteApropiacionesById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Apropiaciones
// @Description Crear Apropiaciones
// @Param	body		body 	models.Apropiaciones	true		"Body para la creacion de Apropiaciones"
// @Success 200 {int} Apropiaciones.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ApropiacionesController) Post() {
	var apropiaciones models.Apropiaciones
	json.Unmarshal(j.Ctx.Input.RequestBody, &apropiaciones)
	fmt.Println(apropiaciones)
	session,_ := db.GetSession()
	models.InsertApropiaciones(session,apropiaciones)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Apropiaciones
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ApropiacionesController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var apropiaciones models.Apropiaciones
	json.Unmarshal(j.Ctx.Input.RequestBody, &apropiaciones)
	session,_ := db.GetSession()

	err := models.UpdateApropiaciones(session, apropiaciones,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Apropiaciones
// @Param	body		body 	models.Apropiaciones	true		"Body para la creacion de Apropiaciones"
// @Success 200 {int} Apropiaciones.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ApropiacionesController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Apropiaciones
// @Param	body		body 	models.Apropiaciones true		"Body para la creacion de Apropiaciones"
// @Success 200 {int} Apropiaciones.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ApropiacionesController) ApropiacionesDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}