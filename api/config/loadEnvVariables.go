package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	// Define la ruta al archivo .env en la carpeta padre
	envPath := filepath.Join("..", ".env")
    
	// Carga el archivo .env
	err := godotenv.Load(envPath)
	if err != nil {
		err2 := godotenv.Load(".env")
		if err2 != nil {
			log.Fatalf("Error cargando el archivo .env: %v", err)
			return
		}
		fmt.Println("Carga de variables de entorno exitosa")
		return
	}
}