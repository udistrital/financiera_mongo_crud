package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "ApropiacionDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "DisponibilidadApropiacionDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "MovimientoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "OrdenPagoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "RegistroPresupuestalDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RubroController"] = append(beego.GlobalControllerRouter["api/controllers:RubroController"],
		beego.ControllerComments{
			Method: "RubroDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

}
