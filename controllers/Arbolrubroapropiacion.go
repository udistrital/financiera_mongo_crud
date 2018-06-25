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
type ArbolRubroApropiacionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 ArbolRubroApropiacion2018 models.ArbolRubroApropiacion2018
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ArbolRubroApropiacionController) GetAll() {
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
// @router /:id/:vigencia/:unidadEjecutora [get]
func (j *ArbolRubroApropiacionController) Get() {
	id := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	session, _ := db.GetSession()
	if id != "" {
		arbolrubroapropiacion, err := models.GetArbolRubroApropiacionById(session, id, unidadEjecutora, vigencia)
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
func (j *ArbolRubroApropiacionController) Delete() {
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
// @router /:vigencia/:unidadEjecutora [post]
func (j *ArbolRubroApropiacionController) Post() {
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	if vigencia != "" {
		var arbolrubroapropiacion *models.ArbolRubroApropiacion
		json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
		fmt.Println(arbolrubroapropiacion)
		session, _ := db.GetSession()
		models.InsertArbolRubroApropiacion(session, arbolrubroapropiacion, unidadEjecutora, vigencia)
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
// @router /:objectId/:vigencia/:unidadEjecutora [put]
func (j *ArbolRubroApropiacionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")
	vigencia := j.Ctx.Input.Param(":vigencia")
	unidadEjecutora := j.Ctx.Input.Param(":unidadEjecutora")
	var arbolrubroapropiacion models.ArbolRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
	session, _ := db.GetSession()

	err := models.UpdateArbolRubroApropiacion(session, arbolrubroapropiacion, objectId, unidadEjecutora, vigencia)
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
func (j *ArbolRubroApropiacionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ArbolRubroApropiacionController) ArbolRubroApropiacion2018DeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight ArbolApropiacion
// @Description Devuelve un nivel del árbol de apropiaciones
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {object} models.Object
// @Failure 403 body is empty
// @router /ArbolApropiacion/:raiz/:unidadEjecutora/:vigencia [get]
func (j *ArbolRubroApropiacionController) ArbolApropiacion() {
	nodoRaiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var arbolApropacionessGrande []map[string]interface{}

	raiz, err := models.GetNodoApropiacion(session, nodoRaiz, ueStr, vigencia)

	if err == nil {
		arbolApropiaciones := make(map[string]interface{})
		arbolApropiaciones["Id"], _ = strconv.Atoi(raiz.Idpsql)
		arbolApropiaciones["Codigo"] = raiz.Id
		arbolApropiaciones["Nombre"] = raiz.Nombre
		arbolApropiaciones["IsLeaf"] = true
		arbolApropiaciones["UnidadEjecutora"] = raiz.Unidad_ejecutora
		arbolApropiaciones["ApropiacionInicial"] = raiz.Apropiacion_inicial

		var hijos []interface{}
		for j := 0; j < len(raiz.Hijos); j++ {
			hijo := getHijoApropiacion(raiz.Hijos[j], ueStr, vigencia)
			if len(hijo) > 0 {
				arbolApropiaciones["IsLeaf"] = false
				hijos = append(hijos, hijo)
			}
		}
		arbolApropiaciones["Hijos"] = hijos
		arbolApropacionessGrande = append(arbolApropacionessGrande, arbolApropiaciones)

		j.Data["json"] = arbolApropacionessGrande
	} else {
		j.Data["json"] = err
	}

	j.ServeJSON()
}

// @Title RaicesArbolApropiacion
// @Description RaicesArbolApropiacion
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /RaicesArbolApropiacion/:unidadEjecutora/:vigencia [get]
func (j *ArbolRubroApropiacionController) RaicesArbolApropiacion() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var roots []map[string]interface{}
	raices, err := models.GetRaicesApropiacion(session, ueStr, vigencia)
	for i := 0; i < len(raices); i++ {
		idPsql, _ := strconv.Atoi(raices[i].Idpsql)
		root := map[string]interface{}{
			"Id":                 idPsql,
			"Codigo":             raices[i].Id,
			"Nombre":             raices[i].Nombre,
			"Hijos":              raices[i].Hijos,
			"IsLeaf":             true,
			"UnidadEjecutora":    raices[i].Unidad_ejecutora,
			"ApropiacionInicial": raices[i].Apropiacion_inicial,
		}
		if len(raices[i].Hijos) > 0 {
			var hijos []map[string]interface{}
			root["IsLeaf"] = false
			for j := 0; j < len(root["Hijos"].([]string)); j++ {
				hijo := getHijoApropiacion(root["Hijos"].([]string)[j], ueStr, vigencia)
				if len(hijo) > 0 {
					hijos = append(hijos, hijo)
				}
			}
			root["Hijos"] = hijos
		}
		roots = append(roots, root)
	}

	if err != nil {
		j.Data["json"] = err
	} else {
		j.Data["json"] = roots
	}

	j.ServeJSON()
}

func getHijoApropiacion(id, ue, vigencia string) map[string]interface{} {
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetArbolRubroApropiacionById(session, id, ue, vigencia)
	hijo := make(map[string]interface{})
	if rubroHijo != nil {
		if rubroHijo.Id != "" {
			hijo["Id"], _ = strconv.Atoi(rubroHijo.Idpsql)
			hijo["Codigo"] = rubroHijo.Id
			hijo["Nombre"] = rubroHijo.Nombre
			hijo["IsLeaf"] = false
			hijo["UnidadEjecutora"] = rubroHijo.Unidad_ejecutora
			hijo["ApropiacionInicial"] = rubroHijo.Apropiacion_inicial
			if len(rubroHijo.Hijos) == 0 {
				hijo["IsLeaf"] = true
				hijo["Hijos"] = nil
				return hijo
			}
		}
	}

	return hijo
}

// @Title RegistrarApropiacionInicial
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ApropiacionInicial"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router RegistrarApropiacionInicial/:vigencia [post]
func (j *ArbolRubroApropiacionController) RegistrarApropiacionInicial() {
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
			unidadEjecutora := dataApropiacion["UnidadEjecutora"].(string)
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
				models.RegistrarApropiacion(session, nuevaApropiacion, unidadEjecutora, vigencia)
			} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
				if err = construirRama(nuevaApropiacion.Id, unidadEjecutora, vigencia, nuevaApropiacion.Idpsql, nuevaApropiacion.Apropiacion_inicial); err != nil {
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

func construirRama(codigoRubro, ue, vigencia, idApr string, nuevaApropiacion int) error {
	var (
		/*padreRubro,*/ actualRubro         models.ArbolRubros
		padreApropiacion, actualApropiacion *models.ArbolRubroApropiacion
		err                                 error
	)

	try.This(func() {
		session, _ := db.GetSession()
		actualRubro, err = models.GetArbolRubrosById(session, codigoRubro)
		session, _ = db.GetSession()
		padreApropiacion, _ = models.GetArbolRubroApropiacionById(session, actualRubro.Padre, ue, vigencia)

		if padreApropiacion == nil {
			beego.Info("No está registrado en las apropiaciones")
			session, _ = db.GetSession()
			actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
			models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
			if actualApropiacion.Padre != "" {
				beego.Info("Tiene padre")
				construirRama(actualRubro.Padre, ue, vigencia, actualRubro.Idpsql, actualApropiacion.Apropiacion_inicial)
			}
		} else {
			beego.Info("Está registrado en las apropiaciones")
			session, _ = db.GetSession()
			beego.Info(codigoRubro)
			apropiacionActualizada, _ := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
			apropiacionAnterior := 0
			session, _ = db.GetSession()
			if apropiacionActualizada != nil {
				apropiacionAnterior = apropiacionActualizada.Apropiacion_inicial
				apropiacionActualizada.Apropiacion_inicial = nuevaApropiacion
				models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia)
			} else {
				actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
				models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
			}

			propagarCambio(padreApropiacion.Id, ue, vigencia, nuevaApropiacion-apropiacionAnterior)

		}

	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
	})
	return err
}

func propagarCambio(codigoRubro, ue, vigencia string, valorPropagado int) error {
	var err error

	try.This(func() { // try catch para recibir errores

		session, _ := db.GetSession()
		apropiacionActualizada, err := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
		apropiacionActualizada.Apropiacion_inicial += valorPropagado

		if err != nil {
			panic(err.Error())
		}
		session, _ = db.GetSession()
		models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia)

		if apropiacionActualizada.Padre != "" {
			propagarCambio(apropiacionActualizada.Padre, ue, vigencia, valorPropagado)
		}
	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
		err = errors.New("unknow error")
	})
	return err
}

func crearNuevaApropiacion(actualRubro models.ArbolRubros, aprId string, nuevaApropiacion int) *models.ArbolRubroApropiacion {
	actualApropiacion := &models.ArbolRubroApropiacion{
		Id:                  actualRubro.Id,
		Idpsql:              aprId,
		Nombre:              actualRubro.Nombre,
		Descripcion:         actualRubro.Descripcion,
		Unidad_ejecutora:    actualRubro.Unidad_Ejecutora,
		Padre:               actualRubro.Padre,
		Hijos:               actualRubro.Hijos,
		Apropiacion_inicial: nuevaApropiacion,
	}
	return actualApropiacion
}

// @Title RegistrarCdp
// @Description Crear RegistrarCdp, este servicio se utiliza sólo para registrar por primera vez el CDP
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de RegistrarCdp"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router RegistrarCdp/ [post]
func (j *ArbolRubroApropiacionController) RegistrarCdp() {
	try.This(func() {
		var cdpData map[string]interface{}
		err := json.Unmarshal(j.Ctx.Input.RequestBody, &cdpData)

		if err != nil {
			panic(err.Error())
		}

		for _, v := range cdpData["Afectacion"].([]interface{}) {
			rubro := v.(map[string]interface{})["Rubro"].(string)
			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
			vigencia := strconv.Itoa(int(cdpData["Vigencia"].(float64)))

			session, _ := db.GetSession()
			rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

			if err != nil {
				panic(err.Error())
			}

			nuevoValor := make(map[string]float64)

			if len(rubroApropiacion.Movimientos) == 0 {
				rubroApropiacion.Movimientos = make(map[string]map[string]float64)
			}

			nuevoValor["mes_cdp"] = v.(map[string]interface{})["Valor"].(float64)
			nuevoValor["total_cdp"] = v.(map[string]interface{})["Valor"].(float64)

			rubroApropiacion.Movimientos[cdpData["MesRegistro"].(string)] = nuevoValor
			session, _ = db.GetSession()
			models.UpdateArbolRubroApropiacion(session, *rubroApropiacion, rubroApropiacion.Id, rubroApropiacion.Unidad_ejecutora, vigencia)
			prograpacionValores(rubroApropiacion.Padre, cdpData["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
		}

		j.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar cdp: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

// @Title AnulacionCdp
// @Description Crear y propagar anulaciones de CDP
// @Param	body		body 	models.ArbolRubroApropiacion true "Body para la creación de anulaciones"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router AnulacionCdp/ [post]
func (j *ArbolRubroApropiacionController) RegistraAnulacionCdp() {
	try.This(func() {
		var anulacionData map[string]interface{}

		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &anulacionData); err != nil {
			panic(err.Error())
		}

		beego.Info(anulacionData)

		for _, v := range anulacionData["Afectacion"].([]interface{}) {
			rubro := v.(map[string]interface{})["Rubro"].(string)
			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
			vigencia := anulacionData["Vigencia"].(string)

			session, _ := db.GetSession()
			rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

			if err != nil {
				panic(err.Error())
			}

			nuevoValor := make(map[string]float64)

			nuevoValor["mes_anulado"] = v.(map[string]interface{})["Valor"].(float64)
			nuevoValor["total_anulado"] = v.(map[string]interface{})["Valor"].(float64)

			rubroApropiacion.Movimientos[anulacionData["MesRegistro"].(string)]["mes_anulado"] = v.(map[string]interface{})["Valor"].(float64)
			rubroApropiacion.Movimientos[anulacionData["MesRegistro"].(string)]["total_anulado"] += v.(map[string]interface{})["Valor"].(float64)
			session, _ = db.GetSession()
			models.UpdateArbolRubroApropiacion(session, *rubroApropiacion, rubroApropiacion.Id, rubroApropiacion.Unidad_ejecutora, vigencia)
			prograpacionValores(rubroApropiacion.Padre, anulacionData["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
		}
		j.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar anulacion cdp: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

// @Title RegistrarMovimiento
// @Description Crear y propagar Valores de movimientos en arbol apropiaciones
// @Param	body		body 	models.ArbolRubroApropiacion true "Body para la movimiento en arbol apropiaciones"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router RegistrarMovimiento/:tipoPago [post]
func (j *ArbolRubroApropiacionController) RegistrarMovimiento() {
	var dataValor map[string]interface{}

	try.This(func() {

		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataValor); err != nil {
			panic(err.Error())
		}

		switch tipoMovimiento := j.GetString(":tipoPago"); tipoMovimiento {
		//rp
		case "rp":
			registrarValores(dataValor, "total_rp", "mes_rp")
		case "anulacion":
			registrarValores(dataValor, "total_anulado", "mes_anulado")
		}

		j.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar movimiento: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

func registrarValores(dataValor map[string]interface{}, total, mes string) {
	try.This(func() {

		beego.Info(dataValor)

		for _, v := range dataValor["Afectacion"].([]interface{}) {
			rubro := v.(map[string]interface{})["Rubro"].(string)
			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
			vigencia := dataValor["Vigencia"].(string)

			session, _ := db.GetSession()
			rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

			if err != nil {
				panic(err.Error())
			}

			nuevoValor := make(map[string]float64)

			nuevoValor[mes] = v.(map[string]interface{})["Valor"].(float64)
			nuevoValor[total] = v.(map[string]interface{})["Valor"].(float64)

			rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][mes] = v.(map[string]interface{})["Valor"].(float64)
			rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][total] += v.(map[string]interface{})["Valor"].(float64)
			session, _ = db.GetSession()
			models.UpdateArbolRubroApropiacion(session, *rubroApropiacion, rubroApropiacion.Id, rubroApropiacion.Unidad_ejecutora, vigencia)
			prograpacionValores(rubroApropiacion.Padre, dataValor["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
		}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar valores: ", e)
	})
}

// @Title RegistrarRp
// @Description Crear y propagar Valores de RP
// @Param	body		body 	models.ArbolRubroApropiacion true "Body para la creación de RP"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router RegistrarRp/ [post]
func (j *ArbolRubroApropiacionController) RegistrarRp() {
	try.This(func() {
		var anulacionData map[string]interface{}

		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &anulacionData); err != nil {
			panic(err.Error())
		}

		beego.Info(anulacionData)

		for _, v := range anulacionData["Afectacion"].([]interface{}) {
			rubro := v.(map[string]interface{})["Rubro"].(string)
			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
			vigencia := anulacionData["Vigencia"].(string)

			session, _ := db.GetSession()
			rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

			if err != nil {
				panic(err.Error())
			}

			nuevoValor := make(map[string]float64)

			nuevoValor["mes_rp"] = v.(map[string]interface{})["Valor"].(float64)
			nuevoValor["total_rp"] = v.(map[string]interface{})["Valor"].(float64)

			rubroApropiacion.Movimientos[anulacionData["MesRegistro"].(string)]["mes_rp"] = v.(map[string]interface{})["Valor"].(float64)
			rubroApropiacion.Movimientos[anulacionData["MesRegistro"].(string)]["total_rp"] += v.(map[string]interface{})["Valor"].(float64)
			session, _ = db.GetSession()
			models.UpdateArbolRubroApropiacion(session, *rubroApropiacion, rubroApropiacion.Id, rubroApropiacion.Unidad_ejecutora, vigencia)
			prograpacionValores(rubroApropiacion.Padre, anulacionData["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
		}
		j.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar anulacion cdp: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

func prograpacionValores(padreRubro, mes, vigencia, ue string, valorPrograpado map[string]float64) (err error) {
	try.This(func() {
		// session, _ := db.GetSession()
		// apropiacionPadre, err := models.GetArbolRubroApropiacionById(session, padreRubro, ue, vigencia)

		// if err != nil {
		// 	panic(err.Error())
		// }

		// if len(apropiacionPadre.Movimientos) == 0 {
		// 	apropiacionPadre.Movimientos = make(map[string]map[string]float64)
		// 	apropiacionPadre.Movimientos[mes] = valorPrograpado
		// } else {
		// 	for key, value := range valorPrograpado {
		// 		if apropiacionPadre.Movimientos[mes][key] != 0 {
		// 			apropiacionPadre.Movimientos[mes][key] += value
		// 		} else {
		// 			apropiacionPadre.Movimientos[mes][key] = value
		// 		}
		// 	}
		// }

		// session, _ = db.GetSession()
		// models.UpdateArbolRubroApropiacion(session, *apropiacionPadre, apropiacionPadre.Id, apropiacionPadre.Unidad_ejecutora, vigencia)

		// if apropiacionPadre.Padre != "" {
		// 	prograpacionValores(apropiacionPadre.Padre, mes, vigencia, ue, valorPrograpado)
		// }

		//variables
		session, _ := db.GetSession()
		apropiacionPadre, err := models.GetArbolRubroApropiacionById(session, padreRubro, ue, vigencia)
		if err != nil {
			panic(err.Error())
		}

		//models.UpdateArbolRubroApropiacion(session, *apropiacionPadre, apropiacionPadre.Id, apropiacionPadre.Unidad_ejecutora, vigencia)

		for apropiacionPadre != nil {

			if len(apropiacionPadre.Movimientos) == 0 {
				apropiacionPadre.Movimientos = make(map[string]map[string]float64)
				apropiacionPadre.Movimientos[mes] = valorPrograpado
			} else {
				for key, value := range valorPrograpado {
					if apropiacionPadre.Movimientos[mes][key] != 0 {
						apropiacionPadre.Movimientos[mes][key] += value
					} else {
						apropiacionPadre.Movimientos[mes][key] = value
					}
				}
			}
			session, _ := db.GetSession()
			if apropiacionPadre.Padre != "" {
				apropiacionPadre, err = models.GetArbolRubroApropiacionById(session, apropiacionPadre.Padre, ue, vigencia)
			} else {
				apropiacionPadre = nil
			}

			if err != nil {
				panic(err.Error())
			}
			beego.Info("Apr ", apropiacionPadre)

		}

	}).Catch(func(e try.E) {
		beego.Error("catch error propagarCdp: ", e)
	})

	return err
}
