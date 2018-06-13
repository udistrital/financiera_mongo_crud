package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
	"github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud ArbolRubroApropiacion2018
type ArbolRubroApropiacion2018Controller struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 ArbolRubroApropiacion2018 models.ArbolRubroApropiacion2018
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ArbolRubroApropiacion2018Controller) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllArbolRubroApropiacion2018s(session)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get ArbolRubroApropiacion2018 by nombre
// @Param	nombre		path 	string	true		"El nombre de la ArbolRubroApropiacion2018 a consultar"
// @Success 200 {object} models.ArbolRubroApropiacion2018
// @Failure 403 :uid is empty
// @router /:id/:vigencia [get]
func (j *ArbolRubroApropiacion2018Controller) Get() {
	id := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	if id != "" {
		arbolrubroapropiacion, err := models.GetArbolRubroApropiacionById(session, id, vigencia)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = arbolrubroapropiacion
		}
	}
	j.ServeJSON()
}

// @Title Borrar ArbolRubroApropiacion2018
// @Description Borrar ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ArbolRubroApropiacion2018Controller) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteArbolRubroApropiacion2018ById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear ArbolRubroApropiacion2018
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018	true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:vigencia [post]
func (j *ArbolRubroApropiacion2018Controller) Post() {
	vigencia := j.GetString(":vigencia")
	if vigencia != "" {
		var arbolrubroapropiacion *models.ArbolRubroApropiacion
		json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
		fmt.Println(arbolrubroapropiacion)
		session, _ := db.GetSession()
		models.InsertArbolRubroApropiacion(session, arbolrubroapropiacion, vigencia)
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = "vigencia null"
	}

	j.ServeJSON()
}

// @Title Update
// @Description update the ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId/:vigencia [put]
func (j *ArbolRubroApropiacion2018Controller) Put() {
	objectId := j.Ctx.Input.Param(":objectId")
	vigencia := j.Ctx.Input.Param(":vigencia")
	var arbolrubroapropiacion models.ArbolRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
	session, _ := db.GetSession()

	err := models.UpdateArbolRubroApropiacion(session, arbolrubroapropiacion, objectId, vigencia)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018	true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ArbolRubroApropiacion2018Controller) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ArbolRubroApropiacion2018Controller) ArbolRubroApropiacion2018DeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title RegistrarApropiacionInicial
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ApropiacionInicial"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router RegistrarApropiacionInicial/:vigencia [post]
func (j *ArbolRubroApropiacion2018Controller) RegistrarApropiacionInicial() {
	var (
		dataApropiacion map[string]interface{}
		rubro           models.ArbolRubros
	)
	try.This(func() {
		vigencia := j.Ctx.Input.Param(":vigencia")
		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataApropiacion); err == nil {
			beego.Info(vigencia)
			beego.Info(dataApropiacion)
			session, _ := db.GetSession()

			codigoRubro := dataApropiacion["Codigo"].(string)

			if rubro, err = models.GetArbolRubrosById(session, codigoRubro); err != nil {
				panic(err.Error())
			}

			nuevaApropiacion := models.ArbolRubroApropiacion2018{
				Id:                  codigoRubro,
				Idpsql:              strconv.Itoa(int(dataApropiacion["Id"].(float64))),
				Nombre:              dataApropiacion["Nombre"].(string),
				Descripcion:         "",
				Unidad_ejecutora:    dataApropiacion["UnidadEjecutora"].(string),
				Padre:               rubro.Padre,
				Hijos:               rubro.Hijos,
				Apropiacion_inicial: int(dataApropiacion["ApropiacionInicial"].(float64)),
			}

			if nuevaApropiacion.Padre == "" { // Si el rubro actual es una raíz, se hace un registro sencillo
				session, _ = db.GetSession()
				beego.Info("Es raíz")
				models.RegistrarApropiacion(session, nuevaApropiacion, vigencia)
			} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
				if err = construirRama(nuevaApropiacion.Id, vigencia, nuevaApropiacion.Apropiacion_inicial); err != nil {
					beego.Error("error en construir rama: ", err.Error())
					panic(err.Error())
				}
			}

			j.Data["json"] = map[string]interface{}{"Type": "success"}
		} else {
			panic(err.Error())
			beego.Error("unmarshal error: ", err.Error())
		}
	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})

	j.ServeJSON()
}

func construirRama(codigoRubro, vigencia string, nuevaApropiacion int) error {
	var (
		/*padreRubro,*/ actualRubro         models.ArbolRubros
		padreApropiacion, actualApropiacion *models.ArbolRubroApropiacion
		err                                 error
	)

	try.This(func() {
		session, _ := db.GetSession()
		actualRubro, err = models.GetArbolRubrosById(session, codigoRubro)
		session, _ = db.GetSession()
		padreApropiacion, _ = models.GetArbolRubroApropiacionById(session, actualRubro.Padre, vigencia)

		if padreApropiacion == nil {
			beego.Info("No está registrado en las apropiaciones")
			session, _ = db.GetSession()
			actualApropiacion = crearNuevaApropiacion(actualRubro, nuevaApropiacion)
			models.InsertArbolRubroApropiacion(session, actualApropiacion, vigencia)
			if actualApropiacion.Padre != "" {
				beego.Info("Tiene padre")
				construirRama(actualRubro.Padre, vigencia, actualApropiacion.Apropiacion_inicial)
			}
		} else {
			beego.Info("Está registrado en las apropiaciones")
			session, _ = db.GetSession()
			beego.Info(codigoRubro)
			apropiacionActualizada, _ := models.GetArbolRubroApropiacionById(session, codigoRubro, vigencia)
			apropiacionAnterior := 0
			session, _ = db.GetSession()
			if apropiacionActualizada != nil {
				apropiacionAnterior = apropiacionActualizada.Apropiacion_inicial
				apropiacionActualizada.Apropiacion_inicial = nuevaApropiacion
				models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, vigencia)
			} else {
				actualApropiacion = crearNuevaApropiacion(actualRubro, nuevaApropiacion)
				models.InsertArbolRubroApropiacion(session, actualApropiacion, vigencia)
			}

			propagarCambio(padreApropiacion.Id, vigencia, nuevaApropiacion-apropiacionAnterior)

		}

	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
	})
	return err
}

func propagarCambio(codigoRubro, vigencia string, valorPropagado int) error {
	var (
		/*padreRubro,*/ //actualRubro         models.ArbolRubros
		//padreApropiacion, actualApropiacion *models.ArbolRubroApropiacion
		err error
	)
	try.This(func() { // try catch para recibir errores
		beego.Info("Propagado: ", valorPropagado)
		beego.Info("Vigencia: ", vigencia)
		session, _ := db.GetSession()
		beego.Info(codigoRubro)
		apropiacionActualizada, err := models.GetArbolRubroApropiacionById(session, codigoRubro, vigencia)
		apropiacionActualizada.Apropiacion_inicial += valorPropagado
		beego.Info(apropiacionActualizada)
		//beego.Info(err.Error())
		if err != nil {
			panic(err.Error())
		}
		session, _ = db.GetSession()
		beego.Info("antes de otra cosa")
		models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, vigencia)
		beego.Info("otra cosa")
		if apropiacionActualizada.Padre != "" {
			//valorApropiacion := nuevaApropiacion - apropiacionAnterior //actualApropiacion.Apropiacion_inicial
			propagarCambio(apropiacionActualizada.Padre, vigencia, valorPropagado)
		}
	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
		err = errors.New("unknow error")
	})
	return err
}

func crearNuevaApropiacion(actualRubro models.ArbolRubros, nuevaApropiacion int) *models.ArbolRubroApropiacion {
	actualApropiacion := &models.ArbolRubroApropiacion{
		Id:                  actualRubro.Id,
		Idpsql:              actualRubro.Idpsql,
		Nombre:              actualRubro.Nombre,
		Descripcion:         actualRubro.Descripcion,
		Unidad_ejecutora:    actualRubro.Unidad_Ejecutora,
		Padre:               actualRubro.Padre,
		Hijos:               actualRubro.Hijos,
		Apropiacion_inicial: nuevaApropiacion,
	}
	return actualApropiacion
}
