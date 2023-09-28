package models
import "golang/db"
type Produto struct {
	Id int;
	Nome string;
	Descricao string;
	Preco float64;
	Quantidade int;
}

func BuscaTodosOsProdutos() []Produto{
	db := db.Conectacombancodedados()

	selectdetodososprodutos, err := db.Query("select * from produtos")
	if err != nil{
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectdetodososprodutos.Next(){
		var id int
		var nome,descricao string
		var preco float64
		var quantidade int

		err = selectdetodososprodutos.Scan(&id,&nome,&descricao,&preco,&quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos,p)
		
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int){
		db := db.Conectacombancodedados()

		insereDadosNoBanco, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values ($1,$2,$3,$4)")
		if err != nil {
			panic(err.Error())
		}

		insereDadosNoBanco.Exec(nome,descricao,preco,quantidade)
		
		defer db.Close()
}

func DeletaProduto(id string){
	db := db.Conectacombancodedados()

	deletarOProduto,err := db.Prepare("delete from produtos where id=$1")
	if err != nil{
	panic(err.Error())
 }
 deletarOProduto.Exec(id)
 defer db.Close()
}