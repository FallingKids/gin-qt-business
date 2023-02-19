package logic

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/constant"
	"github.com/gin-qt-business/app/dao"
	"github.com/gin-qt-business/app/type/data"
	"github.com/gin-qt-business/app/utils"
	pinyin "github.com/mozillazg/go-pinyin"
	"github.com/spf13/viper"
)

type UserLogic struct {
	LoginByPasswordParam data.LoginByPasswordReq
	RegisterReqParams    data.RegisterReq
	Ctx                  *gin.Context
}

func (userLogic *UserLogic) LoginByPassword() (*data.LoginByPasswordRes, error) {
	key := viper.Get("app.aes_key").(string)
	iv := viper.Get("app.aes_iv").(string)
	input := userLogic.LoginByPasswordParam
	password, err := utils.Decrypt(input.Password, []byte(key), []byte(iv))
	if err != nil {
		return nil, err
	}

	userDao := &dao.User{
		Uid:       input.Uid,
		Password:  string(password),
		DeletedAt: uint(constant.IS_NOT_DELETED),
	}

	count, err := userDao.CountUser()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf(string(constant.USER_IS_NOT_EXIST_ERROR))
	}

	jwt := &utils.JWT{
		Username: input.Uid,
		Password: string(password),
	}

	token, err := jwt.CreateToken()
	if err != nil {
		return nil, err
	}

	res := &data.LoginByPasswordRes{
		Token: token,
	}
	return res, nil
}

func (userLogic *UserLogic) Register() (data.RegisterRes, error) {
	key := viper.Get("app.aes_key").(string)
	iv := viper.Get("app.aes_iv").(string)

	input := userLogic.RegisterReqParams

	// uid 的命名规则为中文名称的，例如 胡奕舟 -> huyz
	pinyinArgs := pinyin.NewArgs()
	pinyinArgs.Style = pinyin.Normal
	pinyinStr := pinyin.Pinyin(input.Username, pinyinArgs)
	var uid string
	for _, s := range pinyinStr {
		if len(uid) == 0 {
			uid += string(s[0])
		} else {
			uid += string(s[0][0])
		}
	}

	password, err := utils.Decrypt(input.Password, []byte(key), []byte(iv))
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}

	user := dao.User{
		Uid: uid,
	}

	err = checkUid(&user)
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}

	user.Username = input.Username
	user.Password = string(password)
	user.Phone = input.Phone

	err = user.AddUser()
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}
	return data.RegisterRes{IsOk: true}, nil
}

func checkUid(user *dao.User) error {
	preUid := user.Uid
	count, err := user.CountUser()
	if err != nil {
		return err
	}
	if count > 0 {
		user.Uid = preUid + fmt.Sprint(time.Now().Month()) + fmt.Sprint(time.Now().Day()) + fmt.Sprint(time.Now().Second())
		time.Sleep(1 * time.Second)
		checkUid(user)
	}
	return nil
}
