package controller

import (
	"net/http"

	"github.com/mumushuiding/go-attendance/service"

	"github.com/mumushuiding/util"
)

// SaveClass 保存班次
func SaveClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.ResponseErr(w, "只支持POST方法！！")
		return
	}
	var class service.ClassReceiver
	err := util.Body2Struct(r, &class)
	if err != nil {
		util.ResponseErr(w, err)
		return
	}
	if len(class.Company) == 0 {
		util.ResponseErr(w, "公司 company 不能为空！！")
		return
	}
	if len(class.Sections) == 0 {
		util.ResponseErr(w, "班次的时间段section不能为空")
		return
	}
	if len(class.Serial) == 0 {
		util.ResponseErr(w, "班次名称 serial 不能为空 ")
		return
	}
	err = service.SaveClass(&class)
	if err != nil {
		util.ResponseErr(w, err)
		return
	}
	util.ResponseOk(w)
}
