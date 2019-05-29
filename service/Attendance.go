package service

import (
	"github.com/mumushuiding/go-attendance/model"
)

// SaveAttendance 保存考勤纪录
func SaveAttendance(a *model.Attendance) error {
	return a.Save()
}
