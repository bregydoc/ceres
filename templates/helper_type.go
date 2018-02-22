package templatetype

import (
	"encoding/json"
	"log"

	"github.com/asdine/storm"
	"github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// GetUserByID ...
func GetUserByID(id uuid.UUID) (*User, error) {
	user := new(User)

	if err := DbUser.One("ID", id, user); err != nil {
		if err == storm.ErrNotFound {
			log.Println("User not found")
			return user, err
		}
		log.Println("DbUser.Find(), ", err.Error())
		return user, err
	}
	return user, nil
}

// GetUserByAnyField ...
func GetUserByAnyField(field string, value interface{}) (*User, error) {
	user := new(User)

	if err := DbUser.One(field, value, user); err != nil {
		if err == storm.ErrNotFound {
			log.Println("User not found")
			return user, err
		}
		log.Println("DbUser.Find(), ", err.Error())
		return user, err
	}
	return user, nil
}

// GetAllUsers ...
func GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := DbUser.All(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}

// CreateNewUserWithRawData ...
func CreateNewUserWithRawData(data []byte) (*User, error) {
	user := new(User)
	err := json.Unmarshal(data, user)
	if err != nil {
		log.Println("json.Unmarshal(), ", err.Error())
		return user, err
	}
	err = DbUser.Save(user)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("User already exist")
			return user, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return user, err
	}

	return user, nil
}

// CreateNewUser ...
func CreateNewUser(user *User) (*User, error) {

	err := DbUser.Save(user)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			log.Println("User already exist")
			return user, err
		}
		log.Println("json.Unmarshal(), ", err.Error())
		return user, err
	}

	return user, nil
}

// DeleteUserByID ...
func DeleteUserByID(id uuid.UUID) (*User, error) {

	user, err := GetUserByID(id)
	if err != nil {
		log.Println("GetUserByID(), ", err.Error())
		return user, err
	}

	err = DbUser.DeleteStruct(user)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return user, err
	}

	return user, nil

}

// DeleteUserByStruct ...
func DeleteUserByStruct(user *User) (*User, error) {
	err := DbUser.DeleteStruct(user)
	if err != nil {
		log.Println("DeleteStruct(), ", err.Error())
		return user, err
	}

	return user, nil
}

// UpdateUser ...
func UpdateUser(user *User) (*User, error) {
	err := DbUser.Update(user)
	if err != nil {
		log.Println("DbUser.Update(), ", err.Error())
		return user, err
	}
	return user, nil
}
