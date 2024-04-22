package user

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// Add more fields as needed
}

func getNextID(users []User) uint {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}
