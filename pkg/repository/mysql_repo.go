package repository

import (
	"PUNK-Network/pkg/models"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// Загружаем переменные из .env файла
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Формируем строку подключения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	rootCertPath := os.Getenv("ROOT_CERT_PATH")

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(rootCertPath)

	if err != nil {
		log.Fatalf("Failed to read root certificate: %v", err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatalf("Failed to append PEM.")
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer conn.Close()

	q, err := conn.Query("SELECT version()")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer q.Close()

	var result string
	for q.Next() {
		if err := q.Scan(&result); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}
		fmt.Println(result)
	}
}

func GetDB() *sql.DB {
	return db
}

func AddUser(user models.User) error {
	_, err := db.Exec("INSERT INTO users (firstname, lastname, username, password, phone_id) VALUES (?, ?, ?, ?, ?)",
		user.Firstname, user.Lastname, user.Username, user.Password, user.PhoneID)
	return err
}

func GetProductByID(id int64) (models.Product, error) {
	var product models.Product
	row := db.QueryRow("SELECT id, title, description, price, seller FROM products WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Seller)
	return product, err
}
