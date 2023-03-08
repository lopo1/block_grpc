package model

type Contracts struct {
	Id      int64  `gorm:"type:int(11) UNSIGNED AUTO_INCREMENT;primary_key" json:"id"`
	AppId   string `gorm:"type:varchar(50);index:app_id;" json:"app_id,omitempty"`
	Tag     string `gorm:"type:varchar(50);index:tag;not null DEFAULT '';COMMENT:'合约类型'" json:"tag,omitempty"`
	Created int64  `gorm:"type:int(11);not null DEFAULT '' COMMENT '时间戳'" json:"created,omitempty"`
}
