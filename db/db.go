package db
import (
	"database/sql"
	_ "github.com/lib/pq"
)
func Conectacombancodedados() *sql.DB{
	conexao := "user=postgres dbname=Projeto_Livraria_FP password=postgres port = 5433 host=localhost sslmode=disable"
	db, err :=  sql.Open("postgres",conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}