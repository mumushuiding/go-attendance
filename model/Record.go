package model

// Record 原始打卡纪录
type Record struct {
	Model
	// 用户姓名
	UserName string `json:"userName"`
	// 公司
	Company string
	// 部门
	DeptName string `json:"deptName"`
	// 考勤日期: 格式: 2019/05/30 星期四
	WorkDate string `json:"workDate"`
	// 考勤时间,格式： 2019/05/30 16:30
	PlanCheckTime string `json:"planCheckTime"`
	// 打卡时间
	UserCheckTime string `json:"userCheckTime"`
	// 打卡结果
	UserCheckResult string `json:"userCheckResult"`
	// 打卡地址
	UserLocation string `json:"userLocation"`
	// 打卡备注
	UserRemark string `json:"userRemark"`
	// 打卡异常原因
	CheckExceptionReason string `json:"checkExceptionReason"`
	// 打卡设备
	UserDevice string `json:"userDevice"`
}
