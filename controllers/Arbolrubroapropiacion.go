package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
	"github.com/udistrital/financiera_mongo_crud/db"
	"github.com/udistrital/financiera_mongo_crud/models"
)

// ArbolRubroApropiacionController struct del controlador, utiliza los atributos y funciones de un controlador de beego
type ArbolRubroApropiacionController struct {
	beego.Controller
}

// Get Método Get de HTTP
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
	session, errGetSession := db.GetSession()

	if errGetSession == nil {
		if id != "" {
			arbolrubroapropiacion, err := models.GetArbolRubroApropiacionById(session, id, unidadEjecutora, vigencia)
			if err != nil {
				j.Data["json"] = err.Error()
			} else {
				j.Data["json"] = arbolrubroapropiacion
			}
		}
	}else{
		j.Data["json"] = errGetSession.Error()
	}

	j.ServeJSON()
}

// Delete elimina
// @Title Delete ArbolRubroApropiacion2018
// @Description Borrar ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ArbolRubroApropiacionController) Delete() {
	session, errGetSession := db.GetSession()

	if errGetSession == nil {
		objectID := j.Ctx.Input.Param(":objectId")
		result, errDelete := models.DeleteArbolRubroApropiacion2018ById(session, objectID)
		if errDelete == nil{
			j.Data["json"] = result
		}else{
			j.Data["json"] = errDelete.Error()
		}

	}else{
		j.Data["json"] = errGetSession.Error()
	}

	j.ServeJSON()
}

// Post Método Post de HTTP
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
		if errorUnmarshal := json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion); errorUnmarshal == nil {
			fmt.Println(arbolrubroapropiacion)
			session, errGetSession := db.GetSession()

			if errGetSession == nil{
				models.InsertArbolRubroApropiacion(session, arbolrubroapropiacion, unidadEjecutora, vigencia)
				j.Data["json"] = "insert success!"
			}else{
				j.Data["json"] = errGetSession.Error();
			}


		}else{
			j.Data["json"] = "insert fail!"
		}

	} else {
		j.Data["json"] = "vigencia null"
	}

	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update the ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId/:vigencia/:unidadEjecutora [put]
func (j *ArbolRubroApropiacionController) Put() {
	objectID := j.Ctx.Input.Param(":objectId")
	vigencia := j.Ctx.Input.Param(":vigencia")
	unidadEjecutora := j.Ctx.Input.Param(":unidadEjecutora")
	var arbolrubroapropiacion models.ArbolRubroApropiacion
	if errorUnmarshal := json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion); errorUnmarshal == nil{
		session, errGetSession := db.GetSession()

		if errGetSession == nil{
			err := models.UpdateArbolRubroApropiacion(session, arbolrubroapropiacion, objectID, unidadEjecutora, vigencia)
			if err != nil {
				j.Data["json"] = err.Error()
			} else {
				j.Data["json"] = "update success!"
			}
		}else{
			j.Data["json"] = errGetSession.Error();
		}


	}else{
		j.Data["json"] = errorUnmarshal.Error()
	}

	j.ServeJSON()
}

// Options options
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

// ArbolRubroApropiacion2018DeleteOptions ArbolRubroApropiacion2018DeleteOptions
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

// ArbolApropiacion devuelve un árbol desde la raiz indicada
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
	session, errGetSession := db.GetSession()
	var errConv error
	if errGetSession == nil{
		var arbolApropacionessGrande []map[string]interface{}

		raiz, err := models.GetNodoApropiacion(session, nodoRaiz, ueStr, vigencia)

		if err == nil {
			arbolApropiaciones := make(map[string]interface{})
			arbolApropiaciones["Id"], errConv = strconv.Atoi(raiz.Idpsql)

			if errConv != nil {
				beego.Info("error conviertiendo tipos")
			}

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

	}else{
		j.Data["json"] = errGetSession.Error()
	}

	j.ServeJSON()
}

// RaicesArbolApropiacion ...
// @Title RaicesArbolApropiacion
// @Description RaicesArbolApropiacion
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /RaicesArbolApropiacion/:unidadEjecutora/:vigencia [get]
func (j *ArbolRubroApropiacionController) RaicesArbolApropiacion() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")
	session, errGetSession := db.GetSession()

	if errGetSession == nil {
		var roots []map[string]interface{}
		raices, err := models.GetRaicesApropiacion(session, ueStr, vigencia)
		for i := 0; i < len(raices); i++ {
			idPsql, errorID := strconv.Atoi(raices[i].Idpsql)
			 if errorID != nil {
				 idPsql = 0
				 beego.Info( "error en conversion:", errorID)
			 }

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
	}else{
			j.Data["json"] = errGetSession.Error()
	}


	j.ServeJSON()
}

// Obtiene y devuelve el nodo hijo de la apropiación, devolviendolo en un objeto tipo json (map[string]interface{})
// Se devuelve un objeto de este tipo y no de models con el fin de utilizar la estructura de json utilizada ya en el cliente
// y no tener que hacer grandes modificaciones en el
func getHijoApropiacion(id, ue, vigencia string) map[string]interface{} {
	var errorConv error
	session, errGetSession := db.GetSession()

	if errGetSession == nil {
		rubroHijo, errorGet := models.GetArbolRubroApropiacionById(session, id, ue, vigencia)
		hijo := make(map[string]interface{})
		if rubroHijo != nil && errorGet == nil {
			if rubroHijo.Id != "" {
				hijo["Id"], errorConv = strconv.Atoi(rubroHijo.Idpsql)
				if errorConv != nil {
					beego.Error("Error en conversion", errorConv)
				}

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
		}else{
			beego.Info("Error en get", errorGet)
		}

		return hijo
	}else{
		return nil
	}

}

// RegistrarApropiacionInicial ...
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
			session, errGetSession := db.GetSession()

			if errGetSession == nil {
				codigoRubro := dataApropiacion["Codigo"].(string)
				unidadEjecutora := dataApropiacion["UnidadEjecutora"].(string)
				if rubro, err = models.GetArbolRubrosById(session, codigoRubro); err != nil {
					panic(err.Error())
				}

				nuevaApropiacion := models.ArbolRubroApropiacion{
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
					session, errGetSession := db.GetSession()

					if errGetSession == nil {
						models.InsertArbolRubroApropiacion(session, &nuevaApropiacion, unidadEjecutora, vigencia)
					}else{
						beego.Error("error en GetSession", errGetSession)
					}


				} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
					if err = construirRama(nuevaApropiacion.Id, unidadEjecutora, vigencia, nuevaApropiacion.Idpsql, nuevaApropiacion.Apropiacion_inicial); err != nil {
						beego.Error("error en construir rama: ", err.Error())
						panic(err.Error())
					}
				}

				j.Data["json"] = map[string]interface{}{"Type": "success"}

			}else{
				panic(err.Error())
				beego.Error("session error: ", errGetSession.Error())
			}

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

// Construye la rama a partir de un registro de apropiación inicial
func construirRama(codigoRubro, ue, vigencia, idApr string, nuevaApropiacion int) error {
	var (
		actualRubro                         models.ArbolRubros
		padreApropiacion, actualApropiacion *models.ArbolRubroApropiacion
		err                                 error
	)

	try.This(func() {
		session, errGetSession := db.GetSession()

		if errGetSession == nil {
			actualRubro, err = models.GetArbolRubrosById(session, codigoRubro)
			actualRubro.Unidad_Ejecutora = ue
			session, errGetSession := db.GetSession()

			if errGetSession == nil {
				padreApropiacion, _ = models.GetArbolRubroApropiacionById(session, actualRubro.Padre, ue, vigencia)

				if padreApropiacion == nil {
					session, errGetSession := db.GetSession()

					if errGetSession == nil {
						actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
						models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
						if actualApropiacion.Padre != "" {
							if errorConstruir := construirRama(actualRubro.Padre, ue, vigencia, actualRubro.Idpsql, actualApropiacion.Apropiacion_inicial); errorConstruir != nil{
								beego.Error("error en construir rama: ", errorConstruir.Error())
								panic(errorConstruir.Error())
							}
						}
					}else{
						beego.Error("error en GetSession", errGetSession)
					}

				} else {
					session, errGetSession := db.GetSession()

					if errGetSession == nil {
						apropiacionActualizada, _ := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
						apropiacionAnterior := 0
						session, errGetSession := db.GetSession()

						if errGetSession == nil {
							if apropiacionActualizada != nil {
								apropiacionAnterior = apropiacionActualizada.Apropiacion_inicial
								apropiacionActualizada.Apropiacion_inicial = nuevaApropiacion
								if errorUpdate := models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia); errorUpdate != nil{
									beego.Error("Error al hacer update en arbol Rubro Apropiacion", errorUpdate.Error())
								}
							} else {
								actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
								models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
							}

							if errorPropagar :=propagarCambio(padreApropiacion.Id, ue, vigencia, nuevaApropiacion-apropiacionAnterior); errorPropagar != nil{
								beego.Error("Error en propagar cambio", errorPropagar.Error())
							}
						}else{
							beego.Error("error en GetSession", errGetSession)
						}


					}else{
						beego.Error("error en GetSession", errGetSession)
					}



				}
			}else{
				beego.Error("error en GetSession", errGetSession)
			}


		}else{
			beego.Error("error en GetSession", errGetSession)
		}



	}).Catch(func(e try.E) {
		beego.Error("catch error: ", e)
	})
	return err
}

// Propaga el cambio de la apropiación desde la hoja hasta la raiz, verificando recursivamente si el rubro que se está obteniendo tiene un padre o no
func propagarCambio(codigoRubro, ue, vigencia string, valorPropagado int) error {
	var err error

	try.This(func() { // try catch para recibir errores

		session, errGetSession := db.GetSession()

		if errGetSession == nil {
			apropiacionActualizada, err := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
			apropiacionActualizada.Apropiacion_inicial += valorPropagado

			if err != nil {
				panic(err.Error())
			}
			session, errGetSession := db.GetSession()

			if errGetSession == nil {
				if errorUpdate := models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia); errorUpdate != nil{
					beego.Error("Error Update Arbol Rubro Apropiacion: ", errorUpdate.Error())
				}

				if apropiacionActualizada.Padre != "" {
					if errorPropagar := propagarCambio(apropiacionActualizada.Padre, ue, vigencia, valorPropagado); errorPropagar != nil{
						beego.Error("Error en propagar Cambio: ", errorPropagar.Error())
					}
				}

			}else {
				beego.Error("error en GetSession", errGetSession.Error())
			}


		}else{
			beego.Error("error en GetSession", errGetSession.Error())
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

var tipoTotal string
var tipoMovimiento string
var tipoMovimientoPadre string

//@Title SaldoCDP
//

// @Title RegistrarMovimiento
// @Description Registra los movimientos (como cdp, rp, ver variable tipoMovimiento) y los propaga tanto en la colección
// arbolrubrosapropiacion_[vigencia]_[unidad_ejecutura], como en la colección movimientos. Utiliza la función registrarValores para registrar los valores,
// y se le envian como párametro el nombre de los movimientos que se van a guardar en el atributo movimiento de la colección arbolrubrosapropiacion,
// al igual que se envia la variable dataValor, que son los valores del movimiento enviados desde el api_mid_financiera

// @Param	body		body 	models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router RegistrarMovimiento/:tipoPago [post]
func (j *ArbolRubroApropiacionController) RegistrarMovimiento() {
	var dataValor map[string]interface{}

	try.This(func() {

		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataValor); err != nil {
			panic(err.Error())
		}

		switch tipoMovimiento = j.GetString(":tipoPago"); tipoMovimiento {
		//rp
		case "Cdp":
			tipoTotal = "TotalComprometidoCdp"
			tipoMovimientoPadre = "Apr"
			registrarValores(dataValor, "total_cdp", "mes_cdp")
		case "Rp":
			tipoTotal = "TotalComprometidoRp"
			tipoMovimientoPadre = "Cdp"
			registrarValores(dataValor, "total_rp", "mes_rp")
		case "AnulacionRp":
			tipoTotal = "TotalAnuladoRp"
			tipoMovimientoPadre = "Rp"
			registrarValores(dataValor, "total_anulado_rp", "mes_anulado_rp")
		case "AnulacionCdp":
			tipoTotal = "TotalAnuladoCdp"
			tipoMovimientoPadre = "Cdp"
			registrarValores(dataValor, "total_anulado_cdp", "mes_anulado_cdp")
		case "Adicion": //Adición a la apropiación inicial
			tipoTotal = "AdicionApr"
			tipoMovimientoPadre = ""
			registrarValores(dataValor, "total_adicion", "mes_modificacion")
		case "ModificacionApr": // traslado de apropiación
			beego.Info("Modificación de apropiación.....")
			registrarModifacionApr(dataValor)
		}

		j.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		beego.Error("catch error registrar movimiento: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

// De acuerdo a los valores que recibe, se hacen las modificaciones en el arbolrubroapropiacion y también en la colección de movimientos
// Parámetros: Recibe los valores correspondientes a la modificación, el mes correspondiente de la modificaicón
func registrarModifacionApr(dataValor map[string]interface{}) (err error) {
	var ops []interface{}

	try.This(func() {
		unidadEjecutora := strconv.Itoa(int(dataValor["UnidadEjecutora"].(float64)))
		fechaRegistro := dataValor["FechaMovimiento"].(string)
		vigencia := strconv.Itoa(int(dataValor["Vigencia"].(float64)))
		mes, errorFecha := time.Parse("2006-01-02", fechaRegistro)
		if (errorFecha != nil){
			beego.Info ("error en fecha", errorFecha.Error())
		}

		opsApr := registrarValoresModf(dataValor["Afectacion"].([]interface{}), strconv.Itoa(int(mes.Month())), vigencia, unidadEjecutora)

		for _, v := range dataValor["Afectacion"].([]interface{}) {

			value := v.(map[string]interface{}) // Convierte el elemento v en un map[string]inerface{}, para evitar una conversión constante del mismo

			tipoMovimiento := value["TipoMovimiento"].(string)

			modificacionApr := models.MovimientoCdp{
				IDPsql:          strconv.Itoa(int(dataValor["Id"].(float64))),
				Tipo:            tipoMovimiento,
				Vigencia:        vigencia,
				DocumentoPadre:  strconv.Itoa(int(value["Apropiacion"].(float64))),
				FechaRegistro:   fechaRegistro,
				UnidadEjecutora: unidadEjecutora,
			}
			modificacionApr.RubrosAfecta = append(modificacionApr.RubrosAfecta, value)

			session, errGetSession := db.GetSession()

			if errGetSession == nil {
				op, err := models.EstrctTransaccionMov(session, &modificacionApr)
				if err != nil {
					panic(err.Error())
				}
				ops = append(ops, op)
			}else{
				beego.Error("error en GetSession", errGetSession.Error())
			}



		}
		ops = append(ops, opsApr...)
		beego.Info("ops en registarModifcacionApr: ")
		for i := range ops {
			fmt.Println(ops[i], "\n......")
		}
		session, errGetSession := db.GetSession()
		if errGetSession == nil {
			err = models.RegistrarMovimiento(session, ops)
		}else{
			beego.Error("error en GetSession", errGetSession.Error())
		}

	}).Catch(func(e try.E) {
		beego.Error("catch error registrar modificación apropiación")
		panic(e)
	})
	return err
}

// Crea un CDP para las modificaciones de apropiación inicial que lo necesitan
func crearCdp(dataMovimiento map[string]interface{}, unidadEjecutora, fechaRegistro, vigencia string) (op interface{}) {
	var err error // error handle variable

	try.This(func() {
		rubrosAfecta := make(map[string]interface{})
		rubrosAfecta["Rubro"] = dataMovimiento["CuentaCredito"].(string)
		if dataMovimiento["TipoMovimiento"] == "Traslado" {
			rubrosAfecta["Rubro"] = dataMovimiento["CuentaContraCredito"].(string)
		}

		rubrosAfecta["Valor"] = dataMovimiento["Valor"].(float64)
		rubrosAfecta["Apropiacion"] = strconv.Itoa(int(dataMovimiento["Apropiacion"].(float64)))
		cdp := models.MovimientoCdp{
			IDPsql:          strconv.Itoa(int(dataMovimiento["Disponibilidad"].(float64))),
			Tipo:            "Cdp",
			Vigencia:        vigencia,
			DocumentoPadre:  "0",
			FechaRegistro:   fechaRegistro,
			UnidadEjecutora: unidadEjecutora,
		}
		cdp.RubrosAfecta = append(cdp.RubrosAfecta, rubrosAfecta)

		session, errGetSession := db.GetSession()

		if errGetSession == nil{
			op, err = models.EstrctTransaccionMov(session, &cdp)
			if err != nil {
				panic(err.Error())
			}
		}else{
			beego.Error("error en GetSession", errGetSession.Error())
		}

	}).Catch(func(e try.E) {
		beego.Error("catch error en crearCdp")
		panic(e)
	})
	return
}

// Registrar valores de modificación en arbolrubroapropiacion
func registrarValoresModf(dataModificacion []interface{}, mes, vigencia, ue string) (ops []interface{}) {
	// var err error
	nuevoValor := make(map[string]map[string]map[string]float64)

	try.This(func() {

		for _, d := range dataModificacion {
			data := d.(map[string]interface{})
			data["Mes"] = mes

			if nuevoValor[data["CuentaCredito"].(string)] == nil {
				nuevoValor[data["CuentaCredito"].(string)] = make(map[string]map[string]float64)
			}

			if nuevoValor[data["CuentaCredito"].(string)][mes] == nil {
				nuevoValor[data["CuentaCredito"].(string)][mes] = make(map[string]float64)
			}

			if data["TipoMovimiento"].(string) != "Traslado" {
				formatModifGeneral(data, nuevoValor)
			} else {

				formatModifTraslado(data, nuevoValor)
				beego.Info("nuevoValor: ", nuevoValor)
			}

		}
		beego.Info("nuevoValor: ", nuevoValor)
		for k, v := range nuevoValor {
			op, err := prograpacionValores(k, mes, vigencia, ue, v[mes])
			if err != nil {
				panic(err.Error())
			}
			ops = append(ops, op...)
		}
		beego.Info("ops: ", ops)
	}).Catch(func(e try.E) {
		beego.Error("catch error en registrarValoresModificaciones")
		panic(e)
	})
	return
}

// Formatea las modificaciones de tipo: reducción, adición, suspensión
func formatModifGeneral(data map[string]interface{}, res map[string]map[string]map[string]float64) {
	// if res[data["CuentaContraCredito"].(string)] == nil {
	// 	res[data["CuentaContraCredito"].(string)] = make(map[string]map[string]float64)

	// 	if res[data["CuentaContraCredito"].(string)][data["Mes"].(string)] == nil {
	// 		res[data["CuentaContraCredito"].(string)][data["Mes"].(string)] = make(map[string]float64)
	// 	}

	// 	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_credito"] += data["Valor"].(float64)
	// 	res[data["CuentaContraCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_contra_credito"] += data["Valor"].(float64)
	// } else {
	// 	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)] += data["Valor"].(float64)
	// }

	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)] += data["Valor"].(float64)

}

// Formatea las modificaciones de traslado
func formatModifTraslado(data map[string]interface{}, res map[string]map[string]map[string]float64) {
	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_credito"] += data["Valor"].(float64)
	res[data["CuentaContraCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_contra_credito"] += data["Valor"].(float64)

}

// Itera sobre cada uno de los objetos que estén en el atributo Afectacion enviado desde el api_mid_financiera, que tienen la información necesaria del movimiento.
// Mientras se itera en cada uno de los elementos, se crean las variable rubro, unidadEjecutora y vigencia, para que se pueda buscar el nodo correspondiente en
// la colección arbolrubrosapropiacion_[vigencia]_[unidadEjecutora], luego se comprueba si dicho nodo tiene movimientosAsociados a el. En caso de no tener ninguno
// se instancia un nuevo atributo para que tenga esos valores, luego se guardan los valores enviados desde el api_mid_finciera en la variable nuevoValor y se envian
// como parametro para la función propagarValores, la cuál propaga los valores en el arbolrubrosapropiaciones, devolviendo un arrreglo de interfaces op
// Para la transacción que se llevará acabo
func registrarValores(dataValor map[string]interface{}, total, mes string) (err error) {
	try.This(func() {

		var (
			op  []interface{} // operación para la transacción
			ops []interface{} // todas las operaciones de la transacción
		)

		for _, v := range dataValor["Afectacion"].([]interface{}) {
			rubro := v.(map[string]interface{})["Rubro"].(string)
			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
			vigencia := dataValor["Vigencia"].(string)

			session, errGetSession := db.GetSession()

			if errGetSession == nil {
				beego.Info(rubro, " |", unidadEjecutora, " |", vigencia)
				rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

				if err != nil {
					panic(err.Error())
				}

				nuevoValor := make(map[string]float64)

				if len(rubroApropiacion.Movimientos) == 0 {
					rubroApropiacion.Movimientos = make(map[string]map[string]float64)
					rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] = make(map[string]float64)
				}

				if rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] == nil {
					rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] = make(map[string]float64)
				}

				nuevoValor[mes] = v.(map[string]interface{})["Valor"].(float64)
				nuevoValor[total] = v.(map[string]interface{})["Valor"].(float64)

				rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][mes] = v.(map[string]interface{})["Valor"].(float64)
				rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][total] += v.(map[string]interface{})["Valor"].(float64)

				//models.UpdateArbolRubroApropiacion(session, *rubroApropiacion, rubroApropiacion.Id, rubroApropiacion.Unidad_ejecutora, vigencia)
				ops, err = prograpacionValores(rubroApropiacion.Id, dataValor["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
				if err != nil {
					panic(err.Error())
				}
			}else{
				beego.Error("error en GetSession", errGetSession.Error())
			}


		}

		op, err = registrarDocumentoMovimiento(dataValor, total, mes)
		ops = append(ops, op...)

		session, errGetSession := db.GetSession()

		if errGetSession == nil{
			err = models.RegistrarMovimiento(session, ops)
		}else{
			beego.Error("error en GetSession", errGetSession.Error())
		}

	}).Catch(func(e try.E) {
		beego.Error("catch error registrar valores: ", e)
		panic(e)
	})
	return err
}

func registrarDocumentoMovimiento(dataValor map[string]interface{}, total, mes string) (ops []interface{}, err error) {
	try.This(func() {
		var rubrosAfecta []map[string]interface{}

		documentoPadre, _ := dataValor["Disponibilidad"].(float64)

		for _, rubroAfecta := range dataValor["Afectacion"].([]interface{}) {
			//rubroAfecta.(map[string]interface{})[tipoTotal] = 0.0
			//rubroAfecta.(map[string]interface{})["TotalComprometido"] = 0.0

			rubrosAfecta = append(rubrosAfecta, rubroAfecta.(map[string]interface{}))
		}

		movimiento := models.MovimientoCdp{
			IDPsql:         strconv.Itoa(int(dataValor["Id"].(float64))),
			RubrosAfecta:   rubrosAfecta,
			Tipo:           tipoMovimiento,
			Vigencia:       dataValor["Vigencia"].(string),
			DocumentoPadre: strconv.Itoa(int(documentoPadre)), // si el documento padre esta vacio (no tiene) el valor guardado es 0 (?)
		}
		session, errGetSession := db.GetSession()

		if errGetSession == nil {
			op, err := models.EstrctTransaccionMov(session, &movimiento)
			if err != nil {
				panic(err.Error())
			}
			ops = append(ops, op)

			opp, err := propagarValorMovimientos(movimiento.DocumentoPadre, movimiento, tipoMovimiento) // opp son los movimientos a propagar en la tx de mongodb

			ops = append(ops, opp...)
			if err != nil {
				panic(err.Error())
			}
		}else{
			beego.Error("error en GetSession", errGetSession.Error())
		}

	}).Catch(func(e try.E) {
		beego.Error("error en registrar RP ", e)
		panic(e)
	})
	return ops, err
}

// H
func propagarValorMovimientos(documentoPadre string, Rp models.MovimientoCdp, tMovimiento string) (op []interface{}, err error) {
	session, errGetSession := db.GetSession()

	if errGetSession == nil {
		selectTipoMovimientoPadre(tMovimiento)
		padre, _ := models.GetMovimientoByPsqlId(session, documentoPadre, tipoMovimientoPadre)

		if padre != nil {
			afectacionWalk(&Rp, padre)
			session, errGetSession = db.GetSession()

			if errGetSession == nil {
				opM, err := models.EstrctUpdateTransaccionMov(session, padre) //opM es la tx del movimiento a actualizar
				if err != nil {
					panic(err.Error())
				}

				op = append(op, opM)

				opp, err := propagarValorMovimientos(padre.DocumentoPadre, Rp, tipoMovimientoPadre)
				if err != nil {
					panic(err.Error())
				}
				op = append(op, opp...)

			}else{
				beego.Error("error en GetSession", errGetSession.Error())
			}

		}
		// ???
		for _, imp := range op {
			beego.Info("ops........ controller ", imp, "\n")

		}
	}else{
		beego.Error("error en GetSession", errGetSession.Error())
	}

	return
}

// afectacionWalk itera en todos los elementos de RubrosAfecta del apuntador rp, y luego itera en todos los elementos de RubroAfecta del apuntador RP
// hasta encontrar los elementos que el movimiento de tipo RP afecta al CDP, en este punto se pueden dar dos acciones dependiendo de la variable tipoTotal:
// 1. Que el movimiento aún no tenga el atributo tipoTotal registrado
// 2. Que el movimiento ya tenga el atributo tipoTotal registrado
// Si el movimiento aún no tiene registrado el atributo tipoTotal, se crea y se le asigna el valor que viene del RP (el cual afectaria al CDP)
// Si el movimiento ya tiene registrado el atributo tipoTotal, se modifica su valor sumándole el que viene del RP (el cual afectaria a los correspondientes rubros del CDP)
// Finalmente los apuntadores son modificados y continuan su proceso en la función propagarValorMovimientos
func afectacionWalk(Rp, Cdp *models.MovimientoCdp) {
	for _, rubroRp := range Rp.RubrosAfecta {
		for i := 0; i < len(Cdp.RubrosAfecta); i++ {
			if Cdp.RubrosAfecta[i]["Rubro"].(string) == rubroRp["Rubro"].(string) {
				if Cdp.RubrosAfecta[i][tipoTotal] != nil {
					Cdp.RubrosAfecta[i][tipoTotal] = Cdp.RubrosAfecta[i][tipoTotal].(float64) + rubroRp["Valor"].(float64)
				} else {
					Cdp.RubrosAfecta[i][tipoTotal] = rubroRp["Valor"].(float64)
				}
			}
		}
	}

}

func prograpacionValores(rubro, mes, vigencia, ue string, valorPrograpado map[string]float64) (ops []interface{}, err error) {
	try.This(func() {

		session, _ := db.GetSession()

		apropiacionPadre, err := models.GetArbolRubroApropiacionById(session, rubro, ue, vigencia)

		var apropiacionesCdp []*models.ArbolRubroApropiacion
		if err != nil {
			panic(err.Error())
		}

		for apropiacionPadre != nil {
			beego.Info()
			if apropiacionPadre.Movimientos[mes] == nil {
				apropiacionPadre.Movimientos[mes] = make(map[string]float64)
			}

			if len(apropiacionPadre.Movimientos) == 0 {
				apropiacionPadre.Movimientos = make(map[string]map[string]float64)
				apropiacionPadre.Movimientos[mes] = valorPrograpado
			} else {
				for key, value := range valorPrograpado {

					if apropiacionPadre.Movimientos[mes][key] != 0 {
						if strings.Contains(key, "mes") {
							apropiacionPadre.Movimientos[mes][key] = value

						} else {
							apropiacionPadre.Movimientos[mes][key] += value
						}
					} else {
						apropiacionPadre.Movimientos[mes][key] = value
					}
				}
			}

			apropiacionesCdp = append(apropiacionesCdp, apropiacionPadre)

			if apropiacionPadre.Padre != "" {
				session, _ = db.GetSession()
				apropiacionPadre, err = models.GetArbolRubroApropiacionById(session, apropiacionPadre.Padre, ue, vigencia)
			} else {
				apropiacionPadre = nil
			}

			if err != nil {
				panic(err.Error())
			}

		}
		session, _ = db.GetSession()
		options, err := models.EstrctTransaccionArbolApropiacion(session, apropiacionesCdp, ue, vigencia)
		if err != nil {
			beego.Error("Error en transacción de arbolRbubroApropiacion")
			panic(err.Error())
		}
		for _, obj := range options {
			ops = append(ops, obj)
		}

	}).Catch(func(e try.E) {
		beego.Error("catch error prograpacionValores: ", e)
		panic(e)
	})

	return ops, err
}

func selectTipoMovimientoPadre(tipoHijo string) {
	switch tipoMovimiento = tipoHijo; tipoMovimiento {
	//rp
	case "Cdp":
		tipoMovimientoPadre = "Apr"
	case "Rp":
		tipoMovimientoPadre = "Cdp"
	case "AnulacionRp":
		tipoMovimientoPadre = "Rp"
	case "AnulacionCdp":
		tipoMovimientoPadre = "Cdp"
	default:
		tipoMovimientoPadre = ""
	}
}

// @Title SaldoApropiacion
// @Description Devuelve el saldo de una apropiación especifica
// @Param	body		body 	models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router /SaldoApropiacion/:rubro/:unidadEjecutora/:vigencia [get]
func (j *ArbolRubroApropiacionController) SaldoApropiacion() {
	try.This(func() {
		var (
			rubroParam    string
			unidadEParam  int
			vigenciaParam int
			err           error
		)
		response := make(map[string]float64)
		rubroParam = j.GetString(":rubro")
		if unidadEParam, err = j.GetInt(":unidadEjecutora"); err != nil {
			panic(err.Error())
		}

		if vigenciaParam, err = j.GetInt(":vigencia"); err != nil {
			panic(err.Error())
		}

		session, _ := db.GetSession()
		rubro, err := models.GetArbolRubroApropiacionById(session, rubroParam, strconv.Itoa(unidadEParam), strconv.Itoa(vigenciaParam))
		beego.Info("rubros: ", rubro)
		for _, value := range rubro.Movimientos {
			for key, data := range value {
				response[key] += data
			}
		}
		response["valor_inicial"] = float64(rubro.Apropiacion_inicial)
		if err != nil {
			panic(err.Error())
		}

		j.Data["json"] = response
	}).Catch(func(e try.E) {
		j.Data["json"] = e
	})

	j.ServeJSON()
}

// SaldoMovimiento Devuelve un objeto con el saldo del cdp
// @Title SaldoMovimiento
// @Description Devuelve el saldo de un CDP especifico
// @Param	idPsql		path 	int	true		"idPsql del documento"
// @Param	rubro		path 	string	true		"código del rubro"
// @Param	fuente		query	string false		"fuente de financiamiento"
// @Success 200 {string} success
// @Failure 403 error
// @router /SaldoMovimiento/:idPsql/:rubro/:tipoMovimiento [get]
func (j *ArbolRubroApropiacionController) SaldoMovimiento() {
	try.This(func() {
		var (
			cdpID    int
			err      error
			response map[string]interface{}
		)

		cdpID, err = j.GetInt(":idPsql") // id psql del cdp
		if err != nil {
			panic(err.Error())
		}
		rubro := j.GetString(":rubro")
		fuente := j.GetString("fuente")
		tipoMovimiento := j.GetString(":tipoMovimiento")
		session, _ := db.GetSession()
		cdp, err := models.GetMovimientoByPsqlId(session, strconv.Itoa(cdpID), tipoMovimiento)
		if err != nil {
			panic(err.Error())
		}

		for _, value := range cdp.RubrosAfecta {
			if value["FuenteCodigo"] == nil && value["Rubro"].(string) == rubro && fuente == "" {
				response = value
			} else if value["Rubro"].(string) == rubro && value["FuenteCodigo"].(string) == fuente {
				response = value
			}
		}
		delete(response, "FuenteNombre")
		delete(response, "Rubro")
		delete(response, "Apropiacion")
		delete(response, "FuenteCodigo")
		delete(response, "UnidadEjecutora")

		j.Data["json"] = response
	}).Catch(func(e try.E) {
		j.Data["json"] = e
	})

	j.ServeJSON()
}
