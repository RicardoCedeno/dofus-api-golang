package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	connectionString := "user=postgres password=7dejulio dbname=golang sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Error al obtener la conexión de bajo nivel %w ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Error al verificar la conexión: %w ", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
