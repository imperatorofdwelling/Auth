package entity

import (
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/Website-Auth/internal/models/enum"
)

type UserID uuid.UUID

func (id UserID) String() string        { return uuid.UUID(id).String() }
func UserIDFromString(id string) UserID { return UserID(uuid.MustParse(id)) }

type User struct {
	ID       UserID       `json:"id"`
	Email    string       `json:"email"`
	Name     string       `json:"name"`
	Surname  string       `json:"surname"`
	Username string       `json:"username"`
	Phone    string       `json:"phone"`
	Role     enum.Role    `json:"role"`
	City     enum.City    `json:"city"`
	Country  enum.Country `json:"country"`
	Rating   float64      `json:"rating"`
}
