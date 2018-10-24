package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
	"github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
}

// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (c *FuenteFinanciamientoController) Post() {
	var (
		fuente, infoFuente, tipoFuente map[string]interface{}
		// infoFuente map[string]interface{}
		// tipoFuente
	)

	try.This(func() {
		json.Unmarshal(c.Ctx.Input.RequestBody, &fuente)
		beego.Info("fuente:", fuente["FuenteFinanciamiento"])

		err := formatdata.FillStruct(fuente["FuenteFinanciamiento"], &infoFuente)

		if err != nil {
			panic(err)
		}
		beego.Info("infoFuente: ", infoFuente["Codigo"])

		session, err := db.GetSession()
		if err != nil {
			panic(err)
		}

		fuentePadre := models.GetFuenteFinanciamientoPadreByID(session, infoFuente["Codigo"].(string))

		if fuentePadre == nil { // en caso de que el padre sea nulo, se registra un nuevo padre
			beego.Info("padre es nulo")
			err := formatdata.FillStruct(fuente["TipoFuenteFinanciamiento"], &tipoFuente)
			if err != nil { // error convirtiendo a tipo fuente
				panic(err)
			}

			fuentePadre = &models.FuenteFinaciamientoPadre{
				ID:              infoFuente["Codigo"].(string),
				UnidadEjecutora: 0,
				Descripcion:     infoFuente["Descripcion"].(string),
				IDPsql:          int(infoFuente["Id"].(float64)),
				Nombre:          infoFuente["Nombre"].(string),
				TipoFuente:      tipoFuente["Nombre"],
			}

			models.InsertFuentFinanciamientoPadre(session, fuentePadre)
		}

		beego.Info("fuentePadre: ", fuentePadre)

		defer session.Close()
		c.Data["json"] = "ok"
	}).Catch(func(e try.E) {
		beego.Info("error en Post() ", e)
		c.Data["json"] = e
	})

	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get FuenteFinanciamiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FuenteFinanciamiento
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FuenteFinanciamientoController) GetOne() {

}
