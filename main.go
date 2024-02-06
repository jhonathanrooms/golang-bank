package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "user=username dbname=dbname sslmode=disable"
)

// Definir estructuras de datos para usuarios, cuentas y transacciones

type Usuario struct {
	ID         int    `db:"id"`
	Nombre     string `db:"nombre"`
	Email      string `db:"email"`
	Contraseña string `db:"contraseña"`
}

type Cuenta struct {
	ID        int     `db:"id"`
	UsuarioID int     `db:"usuario_id"`
	Saldo     float64 `db:"saldo"`
}

type Transaccion struct {
	ID              int     `db:"id"`
	CuentaOrigenID  int     `db:"cuenta_origen_id"`
	CuentaDestinoID int     `db:"cuenta_destino_id"`
	Monto           float64 `db:"monto"`
	Fecha           string  `db:"fecha"`
}

func main() {
	// Conectar a la base de datos
	db, err := sqlx.Connect(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ejemplo de consulta a la base de datos
	var usuarios []Usuario
	err = db.Select(&usuarios, "SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir resultados
	fmt.Println("Usuarios:")
	for _, usuario := range usuarios {
		fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Nombre, usuario.Email)
	}
}
