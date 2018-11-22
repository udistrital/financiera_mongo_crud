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
		movimientosFuente              []map[string]interface{}
		options                        []interface{}
	)

	try.This(func() {
		json.Unmarshal(c.Ctx.Input.RequestBody, &fuente)

		err := formatdata.FillStruct(fuente["FuenteFinanciamiento"], &infoFuente)
		if err != nil {
			panic(err)
		}

		session, err := db.GetSession()
		if err != nil {
			beego.Error("error en la sesión")
			panic(err)
		}

		fuentePadre := models.GetFuenteFinanciamientoPadreByID(session, infoFuente["Codigo"].(string))

		if fuentePadre == nil { // en caso de que el padre sea nulo, se registra un nuevo padre
			err := formatdata.FillStruct(fuente["TipoFuenteFinanciamiento"], &tipoFuente)
			if err != nil { // error convirtiendo a tipo fuente
				panic(err)
			}

			fuentePadre = &models.FuenteFinaciamientoPadre{
				ID:              infoFuente["Codigo"].(string),
				UnidadEjecutora: int(infoFuente["UnidadEjecutora"].(float64)),
				Descripcion:     infoFuente["Descripcion"].(string),
				IDPsql:          int(infoFuente["Id"].(float64)),
				Nombre:          infoFuente["Nombre"].(string),
				TipoFuente:      tipoFuente["Nombre"],
				ValorOriginal:   calcularValorOriginal(fuente["AfectacionFuente"].([]interface{})),
			}
			op, err := models.EstructaRegistroFuentePadreTransaccion(session, fuentePadre)
			if err != nil {
				beego.Error("Error al creae estructura de fuente padre")
				panic(err)
			}
			options = append(options, op)
		}

		err = formatdata.FillStruct(fuente["AfectacionFuente"], &movimientosFuente)
		if err != nil {
			panic(err)
		}
		for _, v := range movimientosFuente {
			rubroAfecta := map[string]interface{}{
				"Rubro":      v["Rubro"].(string),
				"Dedepencia": int(v["Dependencia"].(float64)),
			}

			RubrosAfecta := []map[string]interface{}{rubroAfecta}

			movimiento := models.MovimientoCdp{
				IDPsql:          "3",
				RubrosAfecta:    RubrosAfecta,
				ValorOriginal:   v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Valor"].(float64),
				Tipo:            "fuente_financiamiento_" + v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["TipoMovimiento"].(map[string]interface{})["Nombre"].(string),
				Vigencia:        "2018",
				DocumentoPadre:  fuentePadre.ID,
				FechaRegistro:   v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Fecha"].(string),
				UnidadEjecutora: v["UnidadEjecutora"].(string),
			}

			op, err := models.EstrctTransaccionMov(session, &movimiento)
			if err != nil {
				beego.Error("Error en estructura de movimiento para fuente de financiamiento")
				panic(err)
			}
			options = append(options, op)
		}

		err = models.TrRegistroFuente(session, options)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		c.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("error en Post() ", e)
		c.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	c.ServeJSON()
}

func calcularValorOriginal(afectaciones []interface{}) (totalFuente float64) {
	for _, v := range afectaciones {
		for _, movimientos := range v.(map[string]interface{})["MovimientoFuenteFinanciamientoApropiacion"].([]interface{}) {
			totalFuente += movimientos.(map[string]interface{})["Valor"].(float64)
		}
	}
	return
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
