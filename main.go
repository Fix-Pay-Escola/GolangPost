package main

import (
	"net/http"
	_ "github.com/lib/pq"
	"golang/routes"
)

type Produto struct {
	id int;
	Nome string;
	Descricao string;
	Preco float64;
	Quantidade int;
}



func main(){
	routes.CarregaRotas()
	http.ListenAndServe(":5500",nil)
	
}


