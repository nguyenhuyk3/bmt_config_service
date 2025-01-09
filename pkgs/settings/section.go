package settings

type Config struct {
	Server  serverSetting  `mapstructure:"server" json:"server"`
	Kafka   kafkaSetting   `mapstructure:"kafka" json:"kafka"`
	BMTUser postgreSetting `mapstructure:"bmt_user" json:"bmt_user"`
	Logger  loggerSetting  `mapstructure:"logger" json:"logger"`
}

type serverSetting struct {
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

type kafkaSetting struct {
	Address string `mapstructure:"address" json:"address"`
}

type postgreSetting struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	Dbname          string `mapstructure:"dbname" json:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime"`
}

type loggerSetting struct {
	LogLevel    string `mapstructure:"log_level" json:"log_level"`
	FileLogName string `mapstructure:"file_log_name" json:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size"`
	Compress    bool   `mapstructure:"compress" json:"compress"`
}
