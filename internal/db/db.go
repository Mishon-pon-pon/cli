package db

type DB struct {
	Connection struct {
		ConnectionString string `json:"connectionString"`
	}
}
