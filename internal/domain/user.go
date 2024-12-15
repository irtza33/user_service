type User struct {
    UserID int    `json:"user_id"`
    Name   string `json:"name"`
}

func (u *User) GetUserID() int {
    return u.UserID
}

func (u *User) GetName() string {
    return u.Name
}

func NewUser(userID int, name string) *User {
    return &User{
        UserID: userID,
        Name:   name,
    }
}