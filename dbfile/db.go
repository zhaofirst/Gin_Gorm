package dbfile

// 初始化 数据库
import (
	"demo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client defines the db client.
type Client struct {
	db *gorm.DB
	// other method
}

// NewClient returns the db client.
func NewClient(dsn string) (*Client, error) {

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.User{})
	return &Client{db: db}, nil
}
