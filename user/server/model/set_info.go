package model

type SetInfo struct {
	Id                   int64  `gorm:"type:int(11) UNSIGNED AUTO_INCREMENT;primary_key" json:"id"`
	Name                 string `gorm:"type:varchar(50);index:name;" json:"name,omitempty"`
	MaxApplicationNumber int64  `gorm:"type:int(11);not null DEFAULT 5;COMMENT:'最大应用'" json:"maxApplicationNumber,omitempty"`
	MaxAddressNumber     int64  `gorm:"type:int(11);not null DEFAULT 50;COMMENT:'最大地址创建'" json:"maxAddressNumber,omitempty"`
	Created              int64  `gorm:"type:int(11);not null DEFAULT '' COMMENT '时间戳'" json:"created,omitempty"`
}
