package token

import "time"

type Config struct {
	AccessTokenSecret      string        `mapstructure:"secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}
