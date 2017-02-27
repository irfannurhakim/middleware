package middleware

import (
	"encoding/json"

	"github.com/irfannurhakim/goreq"
	"github.com/irfannurhakim/models"
	"github.com/labstack/echo"
)

// ResponseMessage is
type ResponseMessage struct {
	Data models.User `json:"data"`
}

// Profile is
func Profile(identityHost string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// get request object
			req := c.Request()

			// get Tenant ID
			tenantID := c.Param("tenant_id")

			// get auth header token
			authHeader := req.Header.Get("Authorization")

			request := goreq.Request{
				Method: "GET",
				Uri:    identityHost + "/" + tenantID + "/api/v1/me",
			}

			request.AddHeader("Authorization", authHeader)

			res, err := request.Do()
			if err != nil {
				return err
			}

			responseMessage := new(ResponseMessage)
			bodyStr, _ := res.Body.ToString()
			json.Unmarshal([]byte(bodyStr), &responseMessage)

			c.Set("user_profile", responseMessage.Data)

			return next(c)
		}
	}
}
