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

// Operaciones Crud Rubro
type RubroController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Rubro models.Rubro
// @Failure 403 :objectId is empty
// @router / [get]
func (j *RubroController) GetAll() {
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

	obs := models.GetAllRubros(session, query)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Rubro by nombre
// @Param	nombre		path 	string	true		"El nombre de la Rubro a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *RubroController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		rubro, err := models.GetRubroById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = rubro
		}
	}
	j.ServeJSON()
}

// @Title Borrar Rubro
// @Description Borrar Rubro
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *RubroController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteRubroById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Rubro
// @Description Crear Rubro
// @Param	body		body 	models.Rubro	true		"Body para la creacion de Rubro"
// @Success 200 {int} Rubro.Id
// @Failure 403 body is empty
// @router / [post]
func (j *RubroController) Post() {
	var rubro models.Rubro
	json.Unmarshal(j.Ctx.Input.RequestBody, &rubro)
	fmt.Println(rubro)
	session,_ := db.GetSession()
	id := models.InsertRubro(session,rubro)
	j.Data["json"] = id
	j.ServeJSON()
}

// @Title Update
// @Description update the Rubro
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *RubroController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var rubro models.Rubro
	json.Unmarshal(j.Ctx.Input.RequestBody, &rubro)
	session,_ := db.GetSession()

	err := models.UpdateRubro(session, rubro,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Rubro
// @Param	body		body 	models.Rubro	true		"Body para la creacion de Rubro"
// @Success 200 {int} Rubro.Id
// @Failure 403 body is empty
// @router / [options]
func (j *RubroController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Rubro
// @Param	body		body 	models.Rubro true		"Body para la creacion de Rubro"
// @Success 200 {int} Rubro.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *RubroController) RubroDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}
