// @APIVersion 1.0.0
// @Title API
// @Description API Aplicacion Voto - Entidades Core
// @Contact ssierraf@correo.udistrital.edu.co
// @TermsOfServiceUrl http://oas.udistrital.edu.co/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/financiera_mongo_crud/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
  beego.NSNamespace("/movimiento",
	    beego.NSInclude(
	      &controllers.MovimientoController{},
	    ),
   ),
  beego.NSNamespace("/apropiacion",
	    beego.NSInclude(
	      &controllers.ApropiacionController{},
	    ),
   ),
  beego.NSNamespace("/rubro",
	    beego.NSInclude(
	      &controllers.RubroController{},
	    ),
   ),
  beego.NSNamespace("/disponibilidadapropiacion",
	    beego.NSInclude(
	      &controllers.DisponibilidadApropiacionController{},
	    ),
   ),
  beego.NSNamespace("/registropresupuestal",
	    beego.NSInclude(
	      &controllers.RegistroPresupuestalController{},
	    ),
   ),
  beego.NSNamespace("/ordenpago",
	    beego.NSInclude(
	      &controllers.OrdenPagoController{},
	    ),
   ),
	)
	beego.AddNamespace(ns)
}
