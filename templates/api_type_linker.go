package templatetype

import (
	"github.com/kataras/iris"
)

// LinkWithUserType ...
func LinkWithUserType(api iris.Party) {
	api.Get("/user/{id:string}", func(c iris.Context) {
		userID := c.Params().Get("id")

		c.JSON(iris.Map{
			"data":  userID,
			"error": nil,
		})
	})
}
