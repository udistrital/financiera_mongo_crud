package controllers

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Agenda
type AgendaController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Agenda models.Agenda
// @Failure 403 :objectId is empty
// @router / [get]
func (j *AgendaController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllAgendas(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Agenda by nombre
// @Param	nombre		path 	string	true		"El nombre de la Agenda a consultar"
// @Success 200 {object} models.Agenda
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *AgendaController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		agenda, err := models.GetAgendaById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = agenda
		}
	}
	j.ServeJSON()
}

// @Title Borrar Agenda
// @Description Borrar Agenda
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *AgendaController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteAgendaById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Agenda
// @Description Crear Agenda
// @Param	body		body 	models.Agenda	true		"Body para la creacion de Agenda"
// @Success 200 {int} Agenda.Id
// @Failure 403 body is empty
// @router / [post]
func (j *AgendaController) Post() {
	var agenda models.Agenda
	json.Unmarshal(j.Ctx.Input.RequestBody, &agenda)
	fmt.Println(agenda)
	session,_ := db.GetSession()
	models.InsertAgenda(session,agenda)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Agenda
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *AgendaController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var agenda models.Agenda
	json.Unmarshal(j.Ctx.Input.RequestBody, &agenda)
	session,_ := db.GetSession()

	err := models.UpdateAgenda(session, agenda,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Agenda
// @Param	body		body 	models.Agenda	true		"Body para la creacion de Agenda"
// @Success 200 {int} Agenda.Id
// @Failure 403 body is empty
// @router / [options]
func (j *AgendaController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Agenda
// @Param	body		body 	models.Agenda true		"Body para la creacion de Agenda"
// @Success 200 {int} Agenda.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *AgendaController) AgendaDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}
