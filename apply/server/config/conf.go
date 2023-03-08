package config

type ServerConfig struct {
	Port         int64        `mapstructure:"port" json:"port"`
	JWT          JWT          `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql        Mysql        `mapstructure:"mysql" json:"mysql"`
	Pgsql        Pgsql        `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	System       System       `mapstructure:"system" json:"system" yaml:"system"`
	PushUrl      string       `mapstructure:"push_url" json:"push_url"`
	WsUrl        []string     `mapstructure:"ws_url" json:"ws_url"`
	HttpUrl      []string     `mapstructure:"http_url" json:"http_url"`
	Contract     ContractList `mapstructure:"contract" json:"contract"`
	Circle       Circle       `mapstructure:"circle" json:"circle"`
	BlockPrivate BlockPrivate `mapstructure:"block_private" json:"block_private"`
}

type ContractList struct {
	Estate     string `mapstructure:"estate" json:"estate"`
	External   string `mapstructure:"external" json:"external"`
	Membership string `mapstructure:"membership" json:"membership"`
	Owner      string `mapstructure:"owner" json:"owner"`
	Usdc       string `mapstructure:"usdc" json:"usdc"`
}
type BlockPrivate struct {
	Key    string `mapstructure:"key" json:"key"`
	Cipher string `mapstructure:"cipher" json:"cipher"`
	Wallet string `mapstructure:"wallet" json:"wallet"`
}

type Circle struct {
	HttpUrl  string `mapstructure:"http_url" json:"http_url"`
	Apikey   string `mapstructure:"apikey" json:"apikey"`
	WalletId string `mapstructure:"wallet_id" json:"wallet_id"`
}
