package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Connection interface {
	Connect()
}

type Postgres struct{}

func (c *Postgres) Connect() {
	fmt.Println("Postgres connect")
}

type Redis struct{}

func (l *Redis) Connect() {
	fmt.Println("Redis connect")
}

type DBConnect struct {
	DB Connection
}

func (db *DBConnect) Connect() {
	db.DB.Connect()
}

// пример
//sqlConnection := &Postgres{}
//con := &DBConnect{DB: sqlConnection}
//con.DB.Connect()
