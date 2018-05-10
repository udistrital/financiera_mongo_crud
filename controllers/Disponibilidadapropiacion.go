package controllers

import (
  "github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
  "strconv"
  "strings"
  "errors"
)

// Operaciones Crud DisponibilidadApropiacion
type DisponibilidadApropiacionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 DisponibilidadApropiacion models.DisponibilidadApropiacion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *DisponibilidadApropiacionController) GetAll() {
	session,_ := db.GetSession()

  var query = make(map[string]interface{})

	if v := j.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				j.Data["json"] = errors.New("Consulta invalida")
				j.ServeJSON()
				return
			}

			if i, err := strconv.Atoi(kv[1]); err == nil {
				fmt.Println("puede ser convertido")
				k, v := kv[0], i
				query[k] = v
			} else {
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}
	}

	obs := models.GetAllDisponibilidadApropiacions(session, query)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get DisponibilidadApropiacion by nombre
// @Param	nombre		path 	string	true		"El nombre de la DisponibilidadApropiacion a consultar"
// @Success 200 {object} models.DisponibilidadApropiacion
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *DisponibilidadApropiacionController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		disponibilidadapropiacion, err := models.GetDisponibilidadApropiacionById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = disponibilidadapropiacion
		}
	}
	j.ServeJSON()
}

// @Title Borrar DisponibilidadApropiacion
// @Description Borrar DisponibilidadApropiacion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *DisponibilidadApropiacionController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteDisponibilidadApropiacionById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear DisponibilidadApropiacion
// @Description Crear DisponibilidadApropiacion
// @Param	body		body 	models.DisponibilidadApropiacion	true		"Body para la creacion de DisponibilidadApropiacion"
// @Success 200 {int} DisponibilidadApropiacion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *DisponibilidadApropiacionController) Post() {
	var disponibilidadapropiacion models.DisponibilidadApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &disponibilidadapropiacion)
	fmt.Println(disponibilidadapropiacion)
	session,_ := db.GetSession()
	id := models.InsertDisponibilidadApropiacion(session,disponibilidadapropiacion)
	j.Data["json"] = id
	j.ServeJSON()
}

// @Title Update
// @Description update the DisponibilidadApropiacion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *DisponibilidadApropiacionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var disponibilidadapropiacion models.DisponibilidadApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &disponibilidadapropiacion)
	session,_ := db.GetSession()

	err := models.UpdateDisponibilidadApropiacion(session, disponibilidadapropiacion,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear DisponibilidadApropiacion
// @Param	body		body 	models.DisponibilidadApropiacion	true		"Body para la creacion de DisponibilidadApropiacion"
// @Success 200 {int} DisponibilidadApropiacion.Id
// @Failure 403 body is empty
// @router / [options]
func (j *DisponibilidadApropiacionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear DisponibilidadApropiacion
// @Param	body		body 	models.DisponibilidadApropiacion true		"Body para la creacion de DisponibilidadApropiacion"
// @Success 200 {int} DisponibilidadApropiacion.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *DisponibilidadApropiacionController) DisponibilidadApropiacionDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}
