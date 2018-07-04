package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "AgendaDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id/:vigencia/:unidadEjecutora`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "ArbolRubroApropiacion2018DeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId/:vigencia/:unidadEjecutora`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:vigencia/:unidadEjecutora`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "ArbolApropiacion",
			Router: `/ArbolApropiacion/:raiz/:unidadEjecutora/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "RaicesArbolApropiacion",
			Router: `/RaicesArbolApropiacion/:unidadEjecutora/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "RegistrarApropiacionInicial",
			Router: `RegistrarApropiacionInicial/:vigencia`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubroApropiacionController"],
		beego.ControllerComments{
			Method: "RegistrarMovimiento",
			Router: `RegistrarMovimiento/:tipoPago`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "ArbolRubrosDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "ArbolRubro",
			Router: `/ArbolRubro/:raiz/:unidadEjecutora`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "RaicesArbol",
			Router: `/RaicesArbol/:unidadEjecutora`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "EliminarRubro",
			Router: `/eliminarRubro/:idPsql`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:ArbolRubrosController"],
		beego.ControllerComments{
			Method: "RegistrarRubro",
			Router: `/registrarRubro`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/financiera_mongo_crud/controllers:MovimientosController"],
		beego.ControllerComments{
			Method: "AgendaDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

}
