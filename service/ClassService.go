package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// SaveClass 保存班次
func SaveClass(s *model.Class) error {
	return s.Save()
}

// DelClassByID 删除班次
func DelClassByID(id int) error {
	return model.DelClassByID(id)
}

// UpdateClass 更新
func UpdateClass(s *model.Class) error {
	return s.Update()
}
