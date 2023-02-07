package controller

import (
	"github.com/TheSaltiestFish/EasyDouyinApp/dal/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	res, err := db.QueryUserLogin(c, &username, &password)
	log.Println("QueryUser res=", res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
		})
		return
	}
	if res.ID != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	} else {
		//创建登录表user_login
		NewUserLogin := &db.UserLogin{
			UserName: username,
			Password: password,
			Token:    token,
		}
		userid, err := db.CreateUserLogin(c, NewUserLogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
			})
			return
		}

		log.Println("userid", userid)
		//创建用户信息表user
		NewUser := &db.User{
			UserID:   userid,
			UserName: username,
		}
		err = db.CreateUser(c, NewUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
			})
			return
		}

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    token,
		})
		return
	}

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	res, err := db.QueryUserLogin(c, &username, &password)
	log.Println(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
		})
		return
	}

	if res.ID != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(res.ID),
			Token:    token,
		})
		log.Println("+++++++登陆成功++++++++++++++++++++++++++++")
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
}

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	log.Println(userid)
	log.Println(token)

	//token鉴权
	tokenResp, err := db.QueryUserToken(c, &userid, &token)
	log.Println(tokenResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
		})
		return
	}
	if tokenResp.ID == 0 {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Token error! please login again!"},
		})
		return
	}

	//查用户信息
	res, err := db.QueryUser(c, &userid)
	log.Println(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Server error!"},
		})
		return
	}

	if res.ID != 0 {
		user := User{
			Id:            res.UserID,
			Name:          res.UserName,
			FollowCount:   res.FollowCount,
			FollowerCount: res.FollowerCount,
			IsFollow:      true,
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
		return
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
}
