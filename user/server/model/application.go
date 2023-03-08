package model

type Application struct {
	Id        int64  `gorm:"type:int(11) UNSIGNED AUTO_INCREMENT;primary_key" json:"id"`
	UserId    int64  `gorm:"type:int(11);index:user_id;" json:"user_id,omitempty"`
	Name      string `gorm:"type:varchar(100);index:name;" json:"name,omitempty"`
	AppId     string `gorm:"type:varchar(50);index:app_id;" json:"app_id,omitempty"`
	AppSecret string `gorm:"type:varchar(100);" json:"appSecret,omitempty"`
	Created   int64  `gorm:"type:int(11);not null DEFAULT '' COMMENT '时间戳'" json:"created,omitempty"`
}
