package controllers

import (
	"golang/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))
func Index(w http.ResponseWriter,r *http.Request){
	todososprodutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w,"index",todososprodutos)
}
 func New(w http.ResponseWriter,r *http.Request) {
	temp.ExecuteTemplate(w,"New",nil)
 }
 func Insert(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		
		precoConvertido,err :=  strconv.ParseFloat(preco, 64)
		 if err != nil {
			log.Println("erro na conversao do preco: ", err)
		 }
		 quantidadeConvertida,err := strconv.Atoi(quantidade)
		 if err != nil {
			log.Println("erro na conversao do quantidade: ", err)
		 }
		 models.CriarNovoProduto(nome, descricao, precoConvertido,quantidadeConvertida )

		 http.Redirect(w,r,"/",http.StatusMovedPermanently)
	}
 }
 func Delete(w http.ResponseWriter,r *http.Request){
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)

	http.Redirect(w,r,"/",301)
 }
 func Edit(w http.ResponseWriter,r *http.Request){
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w,"Edit",produto)
 }

 func Update(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade:= r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID:",err)
		}
		precoConvertidoParaFloat,err := strconv.ParseFloat(preco, 64)
		if err!= nil {
			log.Println("Erro na conversão do Preco:",err)
		}
		quantidadeConvertidaParaInt,err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da Quantidade:",err)
		}
		models.AtualizaProduto(idConvertidaParaInt,nome,descricao,precoConvertidoParaFloat,quantidadeConvertidaParaInt)

	}
	http.Redirect(w,r,"/",301)
 }


