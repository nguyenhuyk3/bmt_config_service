package settings

type Config struct {
	Server  serverSetting `mapstructure:"server" json:"server"`
	BMTUser bmtUser       `mapstructure:"bmt_user" json:"bmt_user"`
}

type serverSetting struct {
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

type bmtUser struct {
	Databse postgreSetting `mapstructure:"database" json:"database"`
	Mail    mailSetting    `mapstructure:"mail" json:"mail"`
}
