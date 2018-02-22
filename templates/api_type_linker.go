package templatetype

import (
	"github.com/kataras/iris"
)

// LinkWithUserType ...
func LinkWithUserType(api iris.Party) {
	api.Get("/user", func(c iris.Context) {
		c.JSON(iris.Map{
			"data":  "adas",
			"error": "nil",
		})
	})
}
