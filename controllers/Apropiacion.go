package controllers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud Apropiacion
type ApropiacionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Apropiacion models.Apropiacion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ApropiacionController) GetAll() {
	session, _ := db.GetSession()

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

	obs := models.GetAllApropiacions(session, query)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Apropiacion by nombre
// @Param	nombre		path 	string	true		"El nombre de la Apropiacion a consultar"
// @Success 200 {object} models.Apropiacion
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ApropiacionController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		apropiacion, err := models.GetApropiacionById(session, id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = apropiacion
		}
	}
	j.ServeJSON()
}

// @Title Borrar Apropiacion
// @Description Borrar Apropiacion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ApropiacionController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteApropiacionById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Apropiacion
// @Description Crear Apropiacion
// @Param	body		body 	models.Apropiacion	true		"Body para la creacion de Apropiacion"
// @Success 200 {int} Apropiacion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ApropiacionController) Post() {
	var apropiacion models.Apropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &apropiacion)
	fmt.Println(apropiacion)
	session, _ := db.GetSession()
	id := models.InsertApropiacion(session, apropiacion)
	j.Data["json"] = id
	j.ServeJSON()
}

// @Title Update
// @Description update the Apropiacion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ApropiacionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var apropiacion models.Apropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &apropiacion)
	session, _ := db.GetSession()

	err := models.UpdateApropiacion(session, apropiacion, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Apropiacion
// @Param	body		body 	models.Apropiacion	true		"Body para la creacion de Apropiacion"
// @Success 200 {int} Apropiacion.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ApropiacionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Apropiacion
// @Param	body		body 	models.Apropiacion true		"Body para la creacion de Apropiacion"
// @Success 200 {int} Apropiacion.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ApropiacionController) ApropiacionDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}
