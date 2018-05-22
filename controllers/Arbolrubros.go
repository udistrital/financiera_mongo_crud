package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud ArbolRubros
type ArbolRubrosController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 ArbolRubros models.ArbolRubros
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ArbolRubrosController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllArbolRubross(session)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get ArbolRubros by nombre
// @Param	nombre		path 	string	true		"El nombre de la ArbolRubros a consultar"
// @Success 200 {object} models.ArbolRubros
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ArbolRubrosController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		arbolrubros, err := models.GetArbolRubrosById(session, id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = arbolrubros
		}
	}
	j.ServeJSON()
}

// @Title Borrar ArbolRubros
// @Description Borrar ArbolRubros
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ArbolRubrosController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteArbolRubrosById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear ArbolRubros
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros	true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ArbolRubrosController) Post() {
	var arbolrubros models.ArbolRubros
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubros)
	fmt.Println(arbolrubros)
	session, _ := db.GetSession()
	models.InsertArbolRubros(session, arbolrubros)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the ArbolRubros
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ArbolRubrosController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var arbolrubros models.ArbolRubros
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubros)
	session, _ := db.GetSession()

	err := models.UpdateArbolRubros(session, arbolrubros, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros	true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ArbolRubrosController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ArbolRubrosController) ArbolRubrosDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Registra rubro
// @Description Convierte la estructura del api crud para registrarla en un documento mongo
// @Param	body		interface{}	true		"Body para la creacion de un nuevo rubro"
// @Success 201 {string} recibido
// @Failure 403 body is empty
// @router /registrarRubro [post]
func (j *ArbolRubrosController) RegistrarRubro() {
	var rubroData interface{}
	var rubroPadre string = ""
	session, _ := db.GetSession()
	json.Unmarshal(j.Ctx.Input.RequestBody, &rubroData)
	beego.Info("rubroData: ", rubroData)

	rubroDataHijo := rubroData.(map[string]interface{})["RubroHijo"].(map[string]interface{})

	if rubroDataPadre := rubroData.(map[string]interface{})["RubroPadre"].(map[string]interface{}); rubroDataPadre != nil {
		rubroPadre = rubroDataPadre["Codigo"].(string)
	}

	nuevoRubro := models.ArbolRubros{
		Id:          rubroDataHijo["Codigo"].(string),
		Idpsql:      -1,
		Nombre:      rubroDataHijo["Nombre"].(string),
		Descripcion: rubroDataHijo["Descripcion"].(string),
		Hijos:       nil,
		Padre:       rubroPadre}

	models.InsertArbolRubros(session, nuevoRubro)
	beego.Info(nuevoRubro)
	beego.Info(reflect.TypeOf(rubroData))
	j.Data["json"] = "recibido"
	j.ServeJSON()
}
