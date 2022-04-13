package auth

import (
    "github.com/dgrijalva/jwt-go"
    "os"
)

type Payload struct {
    Phone string
}

func (p *Payload) Valid() error {
    return nil
}

type JWTManager struct {
    secret string
}

var GlobalJWTManager *JWTManager

func NewJWTManager(secret string) *JWTManager {
    return &JWTManager{
        secret: secret,
    }
}

func GetGlobalJWTManager() *JWTManager {
    if GlobalJWTManager != nil {
        return GlobalJWTManager
    }
    GlobalJWTManager = NewJWTManager(os.Getenv("JWT_SECRET"))
    return GlobalJWTManager
}

func (j *JWTManager) GenerateToken(phone string) string {
    payload := &Payload{
        Phone: phone,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    signedToken, _ := token.SignedString([]byte(j.secret))
    return signedToken
}

func (j *JWTManager) VerifyToken(token string) (*Payload, error) {
    keyFunc := func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, ErrInvalidToken
        }
        return []byte(j.secret), nil
    }

    parsedToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
    if err != nil {
        return nil, ErrInvalidToken
    }

    payload, ok := parsedToken.Claims.(*Payload)
    if !ok {
        return nil, ErrInvalidToken
    }
    return payload, nil
}
