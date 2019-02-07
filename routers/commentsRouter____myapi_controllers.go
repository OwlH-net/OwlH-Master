package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myapi/controllers:MasterController"] = append(beego.GlobalControllerRouter["myapi/controllers:MasterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:MasterController"] = append(beego.GlobalControllerRouter["myapi/controllers:MasterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:MasterController"] = append(beego.GlobalControllerRouter["myapi/controllers:MasterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "CreateNode",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "UpdateNode",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "GetAllNodes",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "GetNode",
			Router: `/:nid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "DeleteNode",
			Router: `/:nid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "DeployNode",
			Router: `/:nid/deploy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "GetPong",
			Router: `/:nid/ping`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "DeployNode",
			Router: `/deploy/:nid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:NodeController"] = append(beego.GlobalControllerRouter["myapi/controllers:NodeController"],
		beego.ControllerComments{
			Method: "GetPong",
			Router: `/ping/:nid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
