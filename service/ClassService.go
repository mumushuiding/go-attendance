package service

import (
	"errors"
	"sync"

	"github.com/mumushuiding/util"

	"github.com/mumushuiding/go-attendance/model"
)

// ClassReceiver ClassReceiver
type ClassReceiver struct {
	Sections []*model.Section `json:"sections"`
	// 休息时间
	RestTime *model.Section `json:"restTime"`
	// 为true,则下班不打卡也不会生成异常
	IsOffDutyFreeCheck bool `json:"isOffDutyFreeCheck"`
	// 班次名称，如："早班"
	Serial string `json:"serial"`
	// 旷工迟到分钟数
	AbsentTime int `json:"absentTime"`
	// 允许迟到分钟数
	PermitLateMinutes int `json:"permitLateMinutes"`
	// 严重迟到分钟数
	SeriousLateTime int `json:"seriousLateTime"`
	// 休息时间是否打开
	IsRestTimeOpen bool `json:"isRestTimeOpen"`
	// 工作时间，单位是分钟
	WorkTime int `json:"workTime"`
	// 公司
	Company string `json:"company"`
	// 创建人
	UserName string `json:"userName"`
}

var saveClassLock sync.Mutex

// SaveClass 保存班次
func SaveClass(s *ClassReceiver) error {
	saveClassLock.Lock()
	defer saveClassLock.Unlock()
	// 判断班次名称是否已经存在
	exist, err := existsClassByComanyAndSerial(s.Company, s.Serial)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("已经存在班次【" + s.Serial + "】")
	}
	// 赋值
	var class = model.Class{}
	err = setClass(&class, s)
	// str, _ := util.ToJSONStr(&class)
	// fmt.Println(str)
	if err != nil {
		return err
	}
	return class.Save()
	// return nil
}
func existsClassByComanyAndSerial(company, serial string) (bool, error) {
	return model.ExistsClassByComanyAndSerial(company, serial)
}
func setClass(class *model.Class, receiver *ClassReceiver) error {
	var restTime, sections string
	var err error
	if receiver.RestTime != nil {
		restTime, err = util.ToJSONStr(receiver.RestTime)
		if err != nil {
			return err
		}
	}
	sections, err = util.ToJSONStr(receiver.Sections)
	if err != nil {
		return err
	}

	class.IsOffDutyFreeCheck = receiver.IsOffDutyFreeCheck
	class.Serial = receiver.Serial
	class.AbsentTime = receiver.AbsentTime
	class.PermitLateMinutes = receiver.PermitLateMinutes
	class.SeriousLateTime = receiver.SeriousLateTime
	class.IsRestTimeOpen = receiver.IsRestTimeOpen
	class.WorkTime = receiver.WorkTime
	class.Company = receiver.Company
	class.UserName = receiver.UserName
	class.RestTime = restTime
	class.Sections = sections

	return nil
}

// DelClassByID 删除班次
func DelClassByID(id int) error {
	return model.DelClassByID(id)
}

// UpdateClass 更新
func UpdateClass(s *model.Class) error {
	return s.Update()
}
