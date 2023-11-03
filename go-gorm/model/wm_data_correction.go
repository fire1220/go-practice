package model

type WmDataCorrection struct {
	ID           int64 `json:"id" gorm:"column:id"`                                     // 主键id
	Mid          int64 `json:"mid" gorm:"column:mid"`                                   // 学员id
	OperateTypeA int64 `json:"operate_type_a" gorm:"column:operate_type;default:null"`  // 使用数据库的默认值
	OperateTypeB int64 `json:"operate_type_b" gorm:"column:operate_type;default:"`      // 使用数据库的默认值
	OperateTypeC int64 `json:"operate_type_c" gorm:"column:operate_type;AUTOINCREMENT"` // 使用数据库的默认值(原意是设置自动增长：创建的时候可以把空值，走MySQL的默认值，创建数据表的时候会设置为自动增长)
	OperateType  int64 `json:"operate_type" gorm:"column:operate_type;default:()"`      // 使用数据库的默认值
}

func (m *WmDataCorrection) TableName() string {
	return "wm_data_correction"
}
