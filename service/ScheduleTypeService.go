package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// SaveScheduleType 保存班次
func SaveScheduleType(s *model.ScheduleType) error {
	return s.Save()
}

// DelScheduleTypeByID 删除班次
func DelScheduleTypeByID(id int) error {
	return model.DelScheduleTypeByID(id)
}

// UpdateScheduleType 更新
func UpdateScheduleType(s *model.ScheduleType) error {
	return s.Update()
}
