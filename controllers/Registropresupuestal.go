package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
  "strconv"
  "strings"
  "errors"
)

// Operaciones Crud RegistroPresupuestal
type RegistroPresupuestalController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 RegistroPresupuestal models.RegistroPresupuestal
// @Failure 403 :objectId is empty
// @router / [get]
func (j *RegistroPresupuestalController) GetAll() {
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

	obs := models.GetAllRegistroPresupuestals(session, query)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get RegistroPresupuestal by nombre
// @Param	nombre		path 	string	true		"El nombre de la RegistroPresupuestal a consultar"
// @Success 200 {object} models.RegistroPresupuestal
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *RegistroPresupuestalController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		registropresupuestal, err := models.GetRegistroPresupuestalById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = registropresupuestal
		}
	}
	j.ServeJSON()
}

// @Title Borrar RegistroPresupuestal
// @Description Borrar RegistroPresupuestal
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *RegistroPresupuestalController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteRegistroPresupuestalById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear RegistroPresupuestal
// @Description Crear RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal	true		"Body para la creacion de RegistroPresupuestal"
// @Success 200 {int} RegistroPresupuestal.Id
// @Failure 403 body is empty
// @router / [post]
func (j *RegistroPresupuestalController) Post() {
	var registropresupuestal models.RegistroPresupuestal
	json.Unmarshal(j.Ctx.Input.RequestBody, &registropresupuestal)
	fmt.Println(registropresupuestal)
	session,_ := db.GetSession()
	id := models.InsertRegistroPresupuestal(session,registropresupuestal)
	j.Data["json"] = id
	j.ServeJSON()
}

// @Title Update
// @Description update the RegistroPresupuestal
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *RegistroPresupuestalController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var registropresupuestal models.RegistroPresupuestal
	json.Unmarshal(j.Ctx.Input.RequestBody, &registropresupuestal)
	session,_ := db.GetSession()

	err := models.UpdateRegistroPresupuestal(session, registropresupuestal,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal	true		"Body para la creacion de RegistroPresupuestal"
// @Success 200 {int} RegistroPresupuestal.Id
// @Failure 403 body is empty
// @router / [options]
func (j *RegistroPresupuestalController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal true		"Body para la creacion de RegistroPresupuestal"
// @Success 200 {int} RegistroPresupuestal.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *RegistroPresupuestalController) RegistroPresupuestalDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}
