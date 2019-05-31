package model

// Group 考核组
type Group struct {
	Model
	// 地址，如：福州日报社
	Address      string
	AddressCount int `json:"addressCount"`
	// 班次
	ClassesIds []int `json:"classesIds"`
	// 具体班次信息
	ClassesList []string `json:"classesList"`
	// 默认班次
	DefaultClassID       int  `json:"defaultClassId"`
	EnableEmpSelectClass bool `json:"enableEmpSelectClass"`
	// 管理员列表
	ManagerList []*User `json:"managerList"`
	// 考核组成员人数
	MemberCount int `json:"memberCount"`
	// 考核组名称
	Name    string
	Company string
	// 考勤时间
	SelectedClass []*Class `json:"selectedClass"`
	// 类型有：FIXED(固定班制),TURN(轮班制)
	Type string
	// wifi名称
	Wifi      string
	WifiCount int `json:"wifiCount"`
	// 工作日,如 [0,362510019,362510019,362510019,362510019,362510019,0]，0表示不上班，362510019为班次Id
	WorkDayList []int `json:"workDayList"`
	// 打卡地点
	Locations []*Location
	// 有效范围，单位米
	Offset int
}

// User User
type User struct {
	UserID int `json:"userId"`
	Name   string
}

// Wifi 根据WiFi考勤
type Wifi struct {
	Key int
	// MAC地址:如：00:15:5d:90:3e:0e
	MacAddress string `json:"macAddress"`
	// 名称
	Ssid string
}

// Location 打卡地点
type Location struct {
	Key int
	// 打卡地址
	Address string
	// 纬度
	Latitude float64
	// 经度
	Longitude float64
	// 名称
	Title    string
	Accuracy int
}
