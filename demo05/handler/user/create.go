package user

import (
	"net/http"

	"github.com/tongchao199/apiserver_demos/demo05/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", user.Username, user.Password)

	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrUserNameEmpty})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrPasswordEmpty})
		return
	}
}
