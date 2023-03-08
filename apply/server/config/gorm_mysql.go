package config

//type Mysql struct {
//	Host     string `mapstructure:"host" json:"host"`
//	Port     int    `mapstructure:"port" json:"port"`
//	Username string `mapstructure:"db" json:"db"`
//	User     string `mapstructure:"user" json:"user"`
//	Password string `mapstructure:"password" json:"password"`
//	Db       string `mapstructure:"db" json:"db"`
//}
type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
