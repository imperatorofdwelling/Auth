package users

import (
	"github.com/imperatorofdwelling/Website-Auth/internal/models/entity"
	"time"

	"github.com/imperatorofdwelling/Website-Auth/internal/models/enum"
)

type Model struct {
	ID        string       `db:"id"`
	Email     string       `db:"email"`
	Name      string       `db:"name"`
	Surname   string       `db:"surname"`
	Password  string       `db:"password"` // Hashed
	Username  string       `db:"username"`
	Phone     string       `db:"phone"`
	Role      enum.Role    `db:"role"`
	City      enum.City    `db:"city"`
	Country   enum.Country `db:"country"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	Rating    float64      `db:"rating"`
}

func (m *Model) ToEntity() *entity.User {
	return &entity.User{
		ID:       entity.UserIDFromString(m.ID),
		Email:    m.Email,
		Name:     m.Name,
		Surname:  m.Surname,
		Username: m.Username,
		Phone:    m.Phone,
		Role:     m.Role,
		City:     m.City,
		Country:  m.Country,
		Rating:   m.Rating,
	}
}
