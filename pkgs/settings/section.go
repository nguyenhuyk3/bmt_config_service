package settings

type Config struct {
	Server  ServerSetting  `mapstructure:"server" json:"server"`
	BMTUser PostgreSetting `mapstructure:"bmt_user" json:"bmt_user"`
	Logger  LoggerSetting  `mapstructure:"logger" json:"logger"`
}

type ServerSetting struct {
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

type PostgreSetting struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	Dbname          string `mapstructure:"dbname" json:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level" json:"log_level"`
	FileLogName string `mapstructure:"file_log_name" json:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size"`
	Compress    bool   `mapstructure:"compress" json:"compress"`
}
