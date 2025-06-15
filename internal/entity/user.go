package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ContextKey string

const (
	ContextUserID ContextKey = "user"
)

type User struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Password  password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type password struct {
	Text string
	Hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Text = text
	p.Hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.Hash, []byte(text))
}
