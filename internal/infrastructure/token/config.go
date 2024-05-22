package token

import "time"

type Config struct {
	AccessTokenSecret      string        `mapstructure:"jwt_secret"`
	RefreshTokenSecret     string        `mapstructure:"jwt_refresh_secret"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"jwt_access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"jwt_refresh_token_expire_duration"`
}
