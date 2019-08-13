package db

// Mongo mongo
type Mongo struct {
	Address  string
	Username string
	Password string
}

var mongoOP *MongoOP

// GetInstance get instance
func (mongo *Mongo) GetInstance() (*MongoOP, error) {
	if mongoOP == nil {
		err := mongo.Connect()
		if err != nil {
			// do something
		}
	}
	return mongoOP, nil
}

// Connect connect to mongo
func (mongo *Mongo) Connect() error {
	return nil
}

// DisConnect disconnect from mongo
func (mongo *Mongo) DisConnect() error {
	return nil
}
