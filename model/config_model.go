package model

type Config struct {
	DB struct {
		Driver   string `validate:"required" mapstructure:"driver"`
		Host     string `validate:"required" mapstructure:"host"`
		Port     int    `validate:"required" mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DbName   string `validate:"required" mapstructure:"dbname"`
		PoolSize uint64 `mapstructure:"pool_size"`
	} `mapstructure:"database"`

	Logger struct {
		Path         string `validate:"required" mapstructure:"path"`
		MaxAge       int    `validate:"required,gte=1,lte=365" mapstructure:"max_age"`
		RotationTime int    `validate:"required,gte=1,lte=365" mapstructure:"rotation_time"`
	} `mapstructure:"logger"`

	Env struct {
		Code     string `validate:"required" mapstructure:"code"`
		Timezone string `validate:"required" mapstructure:"timezone"`
	}
}
