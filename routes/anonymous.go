package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type anonMail struct {
	EMail string `json:"email"`
}

func createAnonymous() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO check request header

		// TODO add hcaptcha

		// get the send email
		var received anonMail
		c.BindJSON(received)

		// TODO check email

		// TODO create api token

		// TODO store token into redis

		// TODO send email

		c.Status(http.StatusAccepted)
	}
}
