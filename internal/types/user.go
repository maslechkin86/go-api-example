package types

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) IsValid() bool {
	return len(u.Name) > 0 && u.ID >= 0
}
