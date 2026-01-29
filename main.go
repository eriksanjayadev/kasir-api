package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"kasir-api/database"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

config := Config{
	Port: viper.GetString("PORT")
	DBConn: viper.GetString("DB_CONN")
}


db, err := database.InitDB(config.DBConn)
if err != nil {
	log.Fatal("Failed to initialize database: ", err)
}

defer db.close()

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"message": message,
	})
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getCategories(w http.ResponseWriter) {
	JSON(w, http.StatusContinue, categories)
}

func getProducts(w http.ResponseWriter) {
	JSON(w, http.StatusContinue, produk)
}

func main() {

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server failed to start")
	}
}
