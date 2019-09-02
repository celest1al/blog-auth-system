package services

import (
	dbp "github.com/celest1al/blog-auth-system/db"
	"github.com/celest1al/blog-auth-system/models"
	"github.com/celest1al/blog-auth-system/utils"
)

// ByRemember looks up a user by remember token,
// it will handle hashing for us
func ByRemember(token string) (*models.User, error) {
	hashedToken := utils.Hash(token)
	var user models.User
	db := dbp.New()
	err := db.Where("remember_token_hash=?", hashedToken).First(&user).Error
	if err != nil {
		return nil, err
	}
	db.Close()
	return &user, nil
}
