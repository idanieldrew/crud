package databse

import "fmt"

type Conf struct {
	User     string
	Database string
	Password string
}

func ConnectToDb(c Conf) string {
	connection := fmt.Sprintf("%s:%s@/%s", c.User, c.Password, c.Database)
	return connection
}
