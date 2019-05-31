package model

// Detail 考勤明细
type Detail struct {
	Model
	// 用户姓名
	UserName string `json:"userName"`
	// 公司
	Company string `json:"company"`
	// 部门
	DeptName string `json:"deptName"`
	// 工作日
	WorkDate string `json:"workDate"`
	// 班次信息,如："白班 08:00-17:30 "
	ClassInfo string `json:"classInfo"`
	// 出勤天数
	AttendDays int `json:"attendDays"`
	// 休息天数
	RestDays int `json:"restDays"`
	// 工作时常，单位分钟
	WorkTimeDuration int `json:"workTimeDuration"`
	// 迟到次数
	LateTimes int `json:"lateTimes"`
	// 迟到时长(分钟)
	LateDuration int `json:"lateDuration"`
	// 严重迟到次数
	SeriousLateTimes int `json:"seriousLateTimes"`
	// 严重迟到时长(分钟)
	SeriousLateDuration int `json:"seriousLateDuration"`
	// 旷工迟到天数
	AbsenteeismLateDays int `json:"absenteeismLateDays"`
	// 早退次数
	EarlyTimes int `json:"earlyTimes"`
	// 早退时长(分钟)
	EarlyDuration int `json:"earlyDuration"`
	// 上班缺卡次数
	OnDutyLackTimes int `json:"onDutyLackTimes"`
	// 下班缺卡次数
	OffDutyLackTimes int `json:"offDutyLackTimes"`
	// 旷工天数
	AllLackDays int `json:"allLackDays"`
	// 工作日加班(小时)
	WorkOvertimeWorkDayHour int `json:"workOvertimeWorkDayHour"`
	// 休息日加班(小时)
	WorkOvertimeRestDayHour int `json:"workOvertimeRestDayHour"`
	// 节假日加班(小时)
	WorkOvertimeHolidayHour int `json:"workOvertimeHolidayHour"`

	// 加班总时长
	V2workOvertimeDuration int `json:"V2workOvertimeDuration"`
	// 工作日（转加班费）（小时）
	V2workOvertimeFeeWorkDayHour int `json:"V2workOvertimeFeeWorkDayHour"`
	// 休息日（转加班费）（小时）
	V2workOvertimeFeeRestDayHour int `json:"V2workOvertimeFeeRestDayHour"`
	// 节假日（转加班费）（小时）
	V2workOvertimeFeeHolidayHour int `json:"V2workOvertimeFeeHolidayHour"`

	// 工作日（转调休）（小时）
	V2workOvertimeRestWorkDayHour int `json:"V2workOvertimeRestWorkDayHour"`
	// 休息日（转调休）（小时）
	V2workOvertimeRestRestDayHour int `json:"V2workOvertimeRestRestDayHour"`
	// 节假日（转调休）（小时）
	V2workOvertimeRestHolidayHour int `json:"V2workOvertimeRestHolidayHour"`
}
