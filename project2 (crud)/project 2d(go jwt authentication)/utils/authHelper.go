package utils

import(
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error){
	userType :=c.GetString("user_type")
	err = nil
	if userType != role {
		err=errors.New("Unautherized to acess this resource")
		return err
	}
	return err
}

// ADMIN ONLY ACESS 
func MathchUserTypeToUid(c *gin.Context,userId string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err=nil
	if userType =="USER" &&uid != userId{
		err=errors.New("Unautherized to acess this resource")
		return err
	}
	CheckUserType(c, userType)
	return err
}