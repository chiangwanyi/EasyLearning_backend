package model

import (
	"easy_learning/db"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	Id        bson.ObjectId `bson:"_id"`
	UserName  string        `bson:"username"`
	Password  string        `bson:"password"`
	Email     string        `bson:"email"`
	Type      string        `bson:"type"`
	ClassId   string        `bson:"classid"`
	CreatedAt time.Time     `bson:"createdat"`
	UpdatedAt time.Time     `bson:"updatedat"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// CreateUser 创建用户
func (user *User) CreateUser() error {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	user.Id = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.ClassId = "[]"
	return client.Insert(user)
}

// FindUserById 通过 ID 查找用户
func FindUserById(id string) (user User, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user); err == nil {
		return user, nil
	} else {
		return User{}, err
	}
}

// FindUserByEmail 通过 Email 查找用户
func FindUserByEmail(email string) (user User, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"email": email}).One(&user); err == nil {
		return user, nil
	} else {
		return User{}, err
	}
}

// FindUserByUserName 通过 UserName 查找用户
func FindUserByUserName(username string) (user User, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"username": username}).One(&user); err == nil {
		return user, nil
	} else {
		return User{}, err
	}
}

