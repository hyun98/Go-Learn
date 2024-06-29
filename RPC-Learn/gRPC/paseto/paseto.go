package paseto

import (
	"RPC-Learn/config"
	auth "RPC-Learn/gRPC/proto"
	"crypto/rand"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(config *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(config.Paseto.Key),
	}
}

func (pm *PasetoMaker) CreateNewToken(authData *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	if read, err := rand.Read(randomBytes); err != nil {
		return "", err
	} else {
		return pm.Pt.Encrypt(pm.Key, authData, read)
	}
}

func (pm *PasetoMaker) VerifyToken(token string) error {
	var authData *auth.AuthData
	return pm.Pt.Decrypt(token, pm.Key, authData, nil)
}
