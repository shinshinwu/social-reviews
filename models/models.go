package models

import (
  "fmt"
  "time"
  "social-reviews/db"
  "golang.org/x/crypto/bcrypt"
  "github.com/astaxie/beego"
)

type User struct {
  ID                int     `gorm:"primary_key"`
  UserName          string  `gorm:"size:255;not null;unique_index"` // Default size for string is 255, reset it with this tag
  Email             string  `gorm:"size:255;unique_index"`
  EncryptedPassword string  `gorm:"size:255;not null"`
  CreatedAt         time.Time
  UpdatedAt         time.Time

  Reviews         []Review
}

type Review struct {
	ID              int    `gorm:"primary_key"`
  User            User
  UserID          int     `gorm:"index"`
	Title           string  `gorm:"not null"`
	Comment         string  `gorm:"type:text;not null"`
  IsApproved      bool    `gorm:"default:1"`
  CreatedAt       time.Time
  UpdatedAt       time.Time
}

func init() {
  db.RegisterModel(&User{})
  db.RegisterModel(&Review{})
}

func AllUsers() (users []*User, err error) {
  res := db.Conn.Debug().Order("created_at desc").Find(&users)
  err = res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding users: %v", res.Error.Error())
  }
  return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) AsMap() map[string]string {
  user := make(map[string]string)
  user["ID"] = fmt.Sprintf("%v",u.ID)
  user["UserName"] = u.UserName
  user["Email"] = u.Email
  return user
}

func FindUser(id int) (*User, error) {
  user := &User{ID: id}
  res := db.Conn.Debug().First(&user, "id=?", id)
  err := res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
  }
  return user, err
}

func CreateUser(userName string, email string, password string) (*User, error) {
  pwHash, _ := HashPassword(password)
  user := &User{UserName: userName, Email: email, EncryptedPassword: pwHash}
  res := db.Conn.Debug().Create(user)
  err := res.Error
  return user, err
}

func UserLogin(email string, password string) (user *User, err error) {
  user = &User{}
  pwHash, _ := HashPassword(password)
  db := db.Conn.Where("email = ? AND encrypted_password = ?", email, pwHash).First(user).Scan(user)
  err = db.Error
  if err != nil {
    return nil, err
  }
  return
}

func AllReviews() (reviews []*Review, err error) {
  res := db.Conn.Debug().Preload("User").Order("created_at desc").Find(&reviews)
  err = res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding reviews: %v", res.Error.Error())
  }
  return
}

func CreateReview(userId int, title string, comment string) (*Review, error) {
  review := &Review{UserID: userId, Title: title, Comment: comment}
  res := db.Conn.Debug().Create(review)
  err := res.Error
  return review, err
}
