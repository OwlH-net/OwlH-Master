// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"owlhmaster/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/master",
			beego.NSInclude(
				&controllers.MasterController{},
			),
		),
		beego.NSNamespace("/node",
			beego.NSInclude(
				&controllers.NodeController{},
			),
		),
		beego.NSNamespace("/ruleset",
			beego.NSInclude(
				&controllers.RulesetController{},
			),
		),
		beego.NSNamespace("/stap",
			beego.NSInclude(
				&controllers.StapController{},
			),
		),
		beego.NSNamespace("/rulesetSource",
			beego.NSInclude(
				&controllers.RulesetSourceController{},
			),
		),
		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),	
	)
	beego.AddNamespace(ns)
}
