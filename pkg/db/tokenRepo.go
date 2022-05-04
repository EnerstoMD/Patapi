package db

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) SetRefreshToken(c *gin.Context, userID, tokenID string, expiresIn time.Duration) error {
	return repo.redisClient.Set(c, tokenID, userID, expiresIn).Err()
}

func (repo *DbSources) ValidateToken(c *gin.Context, userID, previoustokenID string) error {
	result, err := repo.redisClient.Get(c, previoustokenID).Result()
	if err != nil {
		return err
	}
	if result != userID {
		return errors.New("token invalid")
	}
	return nil
}

func (repo *DbSources) DeleteToken(c *gin.Context, token string) error {
	return repo.redisClient.Del(c, token).Err()
}
