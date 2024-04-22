package user

import "fmt"

// In-memory user store (replace with database interaction)
var users []User

// CreateUser creates a new user
func CreateUser(name string) (*User, error) {
	newUser := User{
		ID:   getNextID(users),
		Name: name,
	}
	users = append(users, newUser)
	return &newUser, nil
}

// GetUser retrieves a user by ID
func GetUser(id uint) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// UpdateUser updates an existing user
func UpdateUser(id uint, name string) error {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

// DeleteUser removes a user by ID
func DeleteUser(id uint) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}
