package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

func connectDB() (*sql.DB, error) {
	connStr := "user=postgres password=Mut4* dbname=postgres sslmode=disable host=localhost port=5432"
	return sql.Open("postgres", connStr)
}

func insertRecord(db *sql.DB, ownerId int, ownerType, material string, quantity int, collection string) error {
	_, err := db.Exec("INSERT INTO datos_reco (ownerId, ownerType, material, quantity, collection) VALUES ($1, $2, $3, $4, $5)",
		ownerId, ownerType, material, quantity, collection)
	return err
}

func readCSV(filePath string, db *sql.DB, done chan<- bool) {
	defer func() { done <- true }()

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error al abrir el archivo CSV: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error al leer el archivo CSV: %v", err)
		return
	}

	if len(records) > 0 && records[0][0] == "ownerId" {
		records = records[1:]
	}

	for _, record := range records {
		if len(record) >= 5 {
			ownerId, err := strconv.Atoi(record[0])
			if err != nil {
				log.Printf("Error al convertir ownerId: %v", err)
				continue
			}
			quantity, err := strconv.Atoi(record[3])
			if err != nil {
				log.Printf("Error al convertir quantity: %v", err)
				continue
			}
			err = insertRecord(db, ownerId, record[1], record[2], quantity, record[4])
			if err != nil {
				log.Printf("Error al insertar en la base de datos: %v", err)
			}
		}
	}
}

func readExcel(filePath string, db *sql.DB, done chan<- bool) {
	defer func() { done <- true }()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Printf("Error al abrir el archivo Excel: %v", err)
		return
	}

	sheetName := file.GetSheetName(1)
	if sheetName == "" {
		log.Printf("No se pudo obtener el nombre de la hoja")
		return
	}

	rows, err := file.GetRows(sheetName)
	if err != nil {
		log.Printf("Error al leer las filas: %v", err)
		return
	}

	if len(rows) > 0 && rows[0][0] == "ownerId" {
		rows = rows[1:]
	}

	fmt.Println("Contenido del archivo Excel:")
	for _, row := range rows {
		if len(row) >= 5 {
			ownerId, err := strconv.Atoi(row[0])
			if err != nil {
				log.Printf("Error al convertir ownerId: %v", err)
				continue
			}
			quantity, err := strconv.Atoi(row[3])
			if err != nil {
				log.Printf("Error al convertir quantity: %v", err)
				continue
			}
			err = insertRecord(db, ownerId, row[1], row[2], quantity, row[4])
			if err != nil {
				log.Printf("Error al insertar en la base de datos: %v", err)
			}
		}
	}
}

func lambdaHandler(ctx context.Context) error {
	filePath := "datos-reco.csv"

	ext := strings.ToLower(filepath.Ext(filePath))

	done := make(chan bool)

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	switch ext {
	case ".csv":
		go readCSV(filePath, db, done)
	case ".xls", ".xlsx":
		go readExcel(filePath, db, done)
	default:
		log.Printf("Formato de archivo no soportado: %s", ext)
		return nil
	}

	<-done
	log.Printf("Datos insertados correctamente, formato de archivo: %s", ext)
	return nil
}

func main() {
	err := lambdaHandler(context.Background())
	if err != nil {
		log.Fatalf("Error en lambdaHandler: %v", err)
	}
}

