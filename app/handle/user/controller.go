package user

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
	errEnum "github.com/gin-qt-business/app/errors"
	"github.com/gin-qt-business/app/type/data"
	"github.com/gin-qt-business/app/utils/check"
)

type UserController struct{}

func (u *UserController) LoginByPassword(ctx *gin.Context) *base.Response {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return base.NewResponseError(errEnum.READALL_ERROR)
	}

	var input data.LoginByPasswordReq
	if err = json.Unmarshal(body, &input); err != nil {
		return base.NewResponseError(errEnum.JSON_UNMARSHAL_ERROR)
	}

	logicUser := UserLogic{
		LoginByPasswordParam: input,
		Ctx:                  ctx,
	}
	res, err := logicUser.LoginByPassword()
	if err != nil {
		return base.NewResponseError(err.Error())
	}

	return base.NewResponseSuccess(res)
}

func (u *UserController) Register(ctx *gin.Context) *base.Response {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return base.NewResponseError(errEnum.READALL_ERROR)
	}

	var input data.RegisterReq
	if err = json.Unmarshal(body, &input); err != nil {
		return base.NewResponseError(errEnum.JSON_UNMARSHAL_ERROR)
	}

	check := &check.Check{}

	if !check.CheckMobile(input.Phone) {
		return base.NewResponseError(errEnum.PARAMS_VALIDATE_ERROR)
	}

	logicUser := UserLogic{
		RegisterReqParams: input,
		Ctx:               ctx,
	}
	res, err := logicUser.Register()
	if err != nil {
		return base.NewResponseError(err.Error())
	}

	return base.NewResponseSuccess(res)
}
