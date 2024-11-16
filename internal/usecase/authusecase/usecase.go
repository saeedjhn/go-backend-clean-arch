package authusecase

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
)

type Config struct {
	AccessTokenSecret      string        `mapstructure:"secret"`
	RefreshTokenSecret     string        `mapstructure:"refresh_secret"`
	AccessTokenSubject     string        `mapstructure:"access_subject"`
	RefreshTokenSubject    string        `mapstructure:"refresh_subject"`
	AccessTokenExpiryTime  time.Duration `mapstructure:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `mapstructure:"refresh_token_expire_duration"`
}

type Interactor struct {
	config Config
	//token  *token.Token
}

// var _ userservice.AuthGenerator = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(config Config) *Interactor {
	return &Interactor{config: config}
}

func (i Interactor) CreateAccessToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: req.Subject,
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.AccessTokenExpiryTime * time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(req.ExpireTime)),
		},
		UserID: req.User.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := accessToken.SignedString([]byte(a.config.AccessTokenSecret))
	tokenString, err := accessToken.SignedString([]byte(req.Secret))
	if err != nil {
		return userauthservicedto.CreateTokenResponse{}, err
	}

	return userauthservicedto.CreateTokenResponse{Token: tokenString}, err
}

func (i Interactor) CreateRefreshToken(
	req userauthservicedto.CreateTokenRequest,
) (userauthservicedto.CreateTokenResponse, error) {
	//rt, err := i.token.CreateRefreshToken(
	//	req.User.ID,
	//	i.config.RefreshTokenSecret,
	//	i.config.RefreshTokenSubject,
	//	i.config.RefreshTokenExpiryTime,
	//)
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: req.Subject,
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.config.AccessTokenExpiryTime * time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(req.ExpireTime)),
		},
		UserID: req.User.ID,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := accessToken.SignedString([]byte(a.config.AccessTokenSecret))
	tokenString, err := accessToken.SignedString([]byte(req.Secret))
	if err != nil {
		return userauthservicedto.CreateTokenResponse{}, err
	}

	return userauthservicedto.CreateTokenResponse{Token: tokenString}, err
}

func (i Interactor) IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (i Interactor) ParseToken(req userauthservicedto.ParseTokenRequest) (userauthservicedto.ParseTokenResponse[*Claims], error) {
	token, err := jwt.ParseWithClaims(
		req.Token,
		&Claims{},
		func(_ *jwt.Token) (interface{}, error) {
			return []byte(req.Secret), nil // secret is accessTokenSecret or refreshTokenSecret_
		},
	)
	if err != nil {
		return userauthservicedto.ParseTokenResponse[*Claims]{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return userauthservicedto.ParseTokenResponse[*Claims]{Claims: claims}, nil
	}

	return userauthservicedto.ParseTokenResponse[*Claims]{}, err
}

//func (i Interactor) ParseAccessToken(requestToken string) (*Claims, error) {
//	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType
//	// token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
//	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//	//		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//	//	}
//	//	return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
//	// })
//
//	token, err := jwt.ParseWithClaims(
//		requestToken,
//		&Claims{},
//		func(_ *jwt.Token) (interface{}, error) {
//			return []byte(i.config.AccessTokenSecret), nil // secret is accessTokenSecret or refreshTokenSecret_
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	claims, ok := token.Claims.(*Claims)
//	if ok && token.Valid {
//		return claims, nil
//	}
//	return nil, err
//}
//
//func (i Interactor) ParseRefreshToken(requestToken string) (*Claims, error) {
//	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType
//	// token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
//	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//	//		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//	//	}
//	//	return []byte(secret), nil // secret is accessTokenSecret or refreshTokenSecret_
//	// })
//
//	token, err := jwt.ParseWithClaims(
//		requestToken,
//		&Claims{},
//		func(_ *jwt.Token) (interface{}, error) {
//			return []byte(i.config.RefreshTokenSecret), nil // secret is accessTokenSecret or refreshTokenSecret_
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	claims, ok := token.Claims.(*Claims)
//	if ok && token.Valid {
//		return claims, nil
//	}
//	return nil, err
//}
