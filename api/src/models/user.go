package models

import (
	//"github.com/makersacademy/go-react-acebook-template/api/src/controllers"
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // gorm.Model creates the following common fields automatically; ID (unit / gorm:"primaryKey"), CreatedAt (time.Time), UpdatedAt(time.Time), DeletedAt (gorm.DeletedAt / gorm:"index")
	Email      string `json:"email"`
	Password   string `json:"password"`
	Username   string `json:"username"`
	PhotoURL   string `json:"image"`
}

// This function creates a new record in the database
func (user *User) Save() (*User, error) {
	// existingUser, err := FindUserByEmail(user.Email)
	// if existingUser != nil {
	// 	return &User{}, err
	// }

	err := Database.Create(user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// FindUser(id) finds and returns the first record in the database where the id matches the id given
func FindUser(id string) (*User, error) {
	var user User
	err := Database.Where("id = ?", id).First(&user).Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// FindUserByEmail(email) finds and returns the first record in the database where the email matches the email given
func FindUserByEmail(email string) (*User, error) {
	var user User
	err := Database.Where("email = ?", email).First(&user).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SeedUserIfNotExist() {
	var Usercount int64
	var Postcount int64
	Database.Model(&User{}).Count(&Usercount)
	Database.Model(&Post{}).Count(&Postcount)
	if Usercount == 0 || Postcount == 0 {
		user := User{
			Email:    "capy@bara.com",
			Password: "capybara!",
			Username: "Mr Capybara",
			PhotoURL: "https://i.ibb.co/bsn7QMT/user-image.png",
		}
		user.Save()

		UserIDString := strconv.Itoa(int(user.ID))

		post := Post{
			Message: "Welcome to Acebook! What's on your mind? Write your first post here.",
			UserID:  UserIDString,
		}
		post.Save()
	}

}
