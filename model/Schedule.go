package model

// Schedule 员工排班表
type Schedule struct {
	Model
	// 用户姓名
	UserName string `json:"userName"`
	// 公司
	Company string
	// 排班日期
	WorkDate string `json:"workDate"`
	// 班次id
	ClassID string `json:"classId"`
	// 是否休息
	IsRest bool `json:"isRest"`
}

// Save Save
func (c *Schedule) Save() error {
	return db.Create(c).Error
}

// DelScheduleByID 根据ID删除
func DelScheduleByID(id int) error {
	return db.Where("id=?", id).Delete(&Schedule{}).Error
}
