package model

type User struct {
	Id       int64  `gorm:"type:int(11) UNSIGNED AUTO_INCREMENT;primary_key" json:"id"`
	Uuid     string `gorm:"type:varchar(100);index:uuid;" json:"uuid,omitempty"`
	Username string `gorm:"type:varchar(100);index:username;" json:"username,omitempty"`
	Address  string `gorm:"type:varchar(100);index:address;" json:"address,omitempty"`
	Email    string `gorm:"type:varchar(50);index:email;" json:"email,omitempty"`
	Phone    string `gorm:"type:varchar(20);index:phone;" json:"phone,omitempty"`
	Password string `gorm:"type:varchar(100);not null DEFAULT '';COMMENT:'唯一id';" json:"password,omitempty"`
	Nonce    string `gorm:"type:varchar(25);not null DEFAULT '' COMMENT '随机nonce'" json:"nonce,omitempty"`
	Created  int64  `gorm:"type:int(11);not null DEFAULT '' COMMENT '时间戳'" json:"created,omitempty"`
	UpData   int64  `gorm:"type:int(11);not null DEFAULT '' COMMENT '时间戳'" json:"upData,omitempty"`
}
