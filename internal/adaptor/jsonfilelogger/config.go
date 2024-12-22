package jsonfilelogger

type Config struct {
	FilePath         string `mapstructure:"file_path"`
	MaxSize          int    `mapstructure:"max_size"`
	MaxBackups       int    `mapstructure:"max_backups"`
	MaxAge           int    `mapstructure:"max_age"`
	LocalTime        bool   `mapstructure:"local_time"`
	Compress         bool   `mapstructure:"compress"`
	Console          bool   `mapstructure:"console"`
	File             bool   `mapstructure:"file"`
	EnableCaller     bool   `mapstructure:"enable_caller"`
	EnableStacktrace bool   `mapstructure:"enable_stack_trace"`
	Level            string `mapstructure:"level"`
}

/* for example:
MaxSize:  10, // megabytes
MaxBackups: 10, // megabytes
MaxAge:    30, // days
LocalTime: false,
Compress:  false,
*/
