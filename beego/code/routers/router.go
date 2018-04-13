// @APIVersion 1.0.0
// @Title API
// @Description API Aplicacion Voto - Entidades Core
// @Contact ssierraf@correo.udistrital.edu.co
// @TermsOfServiceUrl http://oas.udistrital.edu.co/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
  beego.NSNamespace("/movimiento",
	    beego.NSInclude(
	      &controllers.MovimientoController{},
	    ),
   ),
  beego.NSNamespace("/apropiaciones",
	    beego.NSInclude(
	      &controllers.ApropiacionesController{},
	    ),
   ),
  beego.NSNamespace("/rubro",
	    beego.NSInclude(
	      &controllers.RubroController{},
	    ),
   ),
	)
	beego.AddNamespace(ns)
}