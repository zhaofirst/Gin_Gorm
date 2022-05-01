package dbfile

import (
	"demo/models"
	"log"
)

// CreateUser creates user to the db.
func (c *Client) CreateUser(user models.User) error {
	res := c.db.Create(&user)
	if res.Error != nil {
		log.Printf("Create user %v, error: %v\n", user, res.Error)
		return res.Error
	}

	return nil
}

// ReadUser reads user from the db.
func (c *Client) ReadUser(id int) (models.User, error) {
	var user models.User
	res := c.db.First(&user, id)
	if res.Error != nil {
		log.Printf("Query user by id: %d, error: %v\n", id, res.Error)
		return user, res.Error
	}

	return user, nil
}

// UpdateUser updates user to the db.
func (c *Client) UpdateUser(user models.User) error {
	res := c.db.Save(&user)
	if res.Error != nil {
		log.Printf("Save user %v, error: %v\n", user, res.Error)
		return res.Error
	}

	return nil
}

// DeleteUser deletes user from the db.
func (c *Client) DeleteUser(id int) error {
	res := c.db.Delete(&models.User{ID :id})
	if res.Error != nil {
		log.Printf("Delete user by id: %d, error: %v\n", id, res.Error)
		return res.Error
	}
	return nil
}
