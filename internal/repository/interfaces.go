type UserStore interface {
    GetUser(userID int) (string, error)
    CreateUser(name string) (int, error)
    DeleteUser(userID int) (bool, error)
}