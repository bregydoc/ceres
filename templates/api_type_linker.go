package templatetype

import (
	"errors"
	"log"

	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
)

// Response ...
type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// LinkWithUserType ...
func LinkWithUserType(api iris.Party) {
	api.Get("/user/{id:string}", func(c iris.Context) {
		userID := c.Params().Get("id")
		ID, err := uuid.FromString(userID)
		if err != nil {
			log.Println("uuid.FromString(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		user, err := GetUserByID(ID)
		if err != nil {
			log.Println("GetUserByID(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  user,
			Error: err,
		})

	})

	api.Post("/user/create", func(c iris.Context) {
		user := new(User)
		err := c.ReadForm(user)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		// If not have Id
		if user.ID.String() == "" {
			newID, err := uuid.NewV4()
			if err != nil {
				log.Println("uuid.NewV4(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err,
				})
				return
			}
			user.ID = newID
		}

		returnedUser, err := CreateNewUser(user)
		if err != nil {
			log.Println("CreateNewUser(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  returnedUser,
			Error: err,
		})
	})

	api.Post("/user/update", func(c iris.Context) {
		user := new(User)
		err := c.ReadForm(user)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		// If not have Id
		if user.ID.String() == "" {
			log.Println("Update method needs Id")
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: errors.New("Update method needs Id"),
			})
			return
		}

		updatedUser, err := UpdateUser(user)
		if err != nil {
			log.Println("UpdateUser(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  updatedUser,
			Error: err,
		})

	})

	api.Post("/user/delete", func(c iris.Context) {
		type IDForm struct {
			ID uuid.UUID `json:"id"`
		}

		idForm := new(IDForm)
		err := c.ReadForm(idForm)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err,
			})
			return
		}

		returnedUser, err := DeleteUserByID(idForm.ID)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
		} else {
			c.StatusCode(iris.StatusOK)
		}

		c.JSON(Response{
			Data:  returnedUser,
			Error: err,
		})

	})
}
