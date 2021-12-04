package model

import (
	"demo/dao"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Phone     string `gorm:"not null;unique"`
	Nickname  string `gorm:"not null"`
	Realname  string `gorm:"not null"`
	AvatarUrl string
}

func (u *User) Create() (err error) {
	// 加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	u.Password = string(hash) // 保存在数据库的密码
	if err := dao.DB.Create(&u).Error; err != nil {
		log.Printf("create err:%v\n", err)
		return err
	}
	return
}
func (u *User) Get() (getUser User, err error) {
	if err = dao.DB.Where(&u).Find(&getUser).Error; err != nil {
		log.Printf("err:%v\n", err)
	}
	return
}

// 已废弃
func (u *User) UpdateOld(newPassword string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = dao.DB.Model(&u).Update("Password", string(hash)).Error; err != nil {
		log.Printf("update err:%v\n", err)
	}
	return
}
func (u *User) Update(newUser User) (err error) {
	if newUser.Password != "" { // 说明是来改密码的
		newPassword := newUser.Password
		hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return err
		}
		if err = dao.DB.Model(&u).Update("Password", string(hash)).Error; err != nil {
			log.Printf("update err:%v\n", err)
			return err
		}
	}
	if err = dao.DB.Model(&u).Updates(newUser).Error; err != nil {
		log.Printf("update err:%v\n", err)
	}
	return
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return
}
