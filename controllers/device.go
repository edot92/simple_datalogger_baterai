package controllers

import (
	"encoding/json"

	"github.com/edot92/simple_datalogger_baterai/models"

	"github.com/astaxie/beego"
)

type DeviceController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Param    body        body    models.device.ParamChart  true "tes"
// @Success 200 {object} models.Object "sukses"
// @Failure 403 :objectId is empty
// @router /chart [post]
func (c *DeviceController) Chart() {
	var ob models.ParamChart
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	resModel, _ := models.GetRealtimeChart(ob.JenisChart)
	c.Data["json"] = map[string]interface{}{"data": resModel}
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Success 200 {object} models.Object "sukses"
// @Failure 403 :objectId is empty
// @router /realtimebox [get]
func (c *DeviceController) Realtimebox() {

	resModel, _ := models.GetRealtimeBox()
	c.Data["json"] = map[string]interface{}{"data": resModel}
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Success 200 {object} models.Object "sukses"
// @Failure 403 :objectId is empty
// @router /historylast [get]
func (c *DeviceController) Historylast() {

	resModel, _ := models.GetHistoryLast()
	c.Data["json"] = map[string]interface{}{"data": resModel}
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}

// @Title Get Historybydaterange
// @Description find object by objectid
// @Param    body        body    models.device.Paramhistorydate  true "tes"
// @Success 200 {object} models.Object "sukses"
// @Failure 403 :objectId is empty
// @router /historybydaterange [post]
func (c *DeviceController) Historybydaterange() {
	var ob models.Paramhistorydate
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// fmt.Println(ob)
	// os.Exit(1)
	resModel, _ := models.GetHistoryDateRange(ob)
	c.Data["json"] = map[string]interface{}{"data": resModel}

	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}
