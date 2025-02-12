package settings

type basicSetting struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username,omitempty"`
	Password string `mapstructure:"password" json:"password,omitempty"`
	Database int    `mapstructure:"database" json:"database,omitempty"`
}

type postgreSetting struct {
	BasicSetting    basicSetting `mapstructure:",squash" json:"basic_settings"`
	DbName          string       `mapstructure:"db_name" json:"db_name"`
	MaxIdleConns    int          `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns    int          `mapstructure:"max_open_conns" json:"max_open_conns"`
	ConnMaxLifetime int          `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime"`
}

type mailSetting struct {
	BasicSetting basicSetting `mapstructure:",squash" json:"basic_settings"`
}

type redisSetting struct {
	BasicSetting basicSetting `mapstructure:",squash" json:"basic_settings"`
}

// type loggerSetting struct {
// 	LogLevel    string `mapstructure:"log_level" json:"log_level"`
// 	FileLogName string `mapstructure:"file_log_name" json:"file_log_name"`
// 	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups"`
// 	MaxAge      int    `mapstructure:"max_age" json:"max_age"`
// 	MaxSize     int    `mapstructure:"max_size" json:"max_size"`
// 	Compress    bool   `mapstructure:"compress" json:"compress"`
// }
