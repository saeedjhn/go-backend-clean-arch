package logger

type Config struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	LocalTime  bool   `mapstructure:"local_time"`
	Compress   bool   `mapstructure:"compress"`
}

/* for example:
Filename: "./logs/log.json",
MaxSize:  10, // megabytes
MaxBackups: 10, // megabytes
MaxAge:    30, // days
LocalTime: false,
Compress:  false,
*/
