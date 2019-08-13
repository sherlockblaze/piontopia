package db

// Mysql mysql connection info
type Mysql struct {
	Address  string
	Username string
	Password string
}

var mysqlOP *MysqlOP

// GetInstance return singleton instance of mysql Conn
func (mysql *Mysql) GetInstance() (*MysqlOP, error) {
	if mysqlOP == nil {
		err := mysql.Connect()
		if err != nil {
			// do something
		}
	}
	return mysqlOP, nil
}

// Connect connect to mysql
func (mysql *Mysql) Connect() error {
	return nil
}

// DisConnect disconnect from mysql
func (mysql *Mysql) DisConnect() error {
	return nil
}
