package utils

import (
	"github.com/gin-qt-business/app/utils/aes"
	"github.com/gin-qt-business/app/utils/check"
	"github.com/gin-qt-business/app/utils/jwt"
)

type Utils struct {
	Aes   aes.AES
	Jwt   jwt.JWT
	Check check.Check
}
