package main

import(
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)


struct Command{
	PreCommand string
	ExeCommand string
	PostComannd string
 }

struct OutPut{
	output string
}


func (com Command) ExecutePre(db *sql.DB){
	result,err:=db.Query(com.PreCommand)
	if err!=nil{
		panic(err)
	}
	
}

func (com Command)ExecuteExe(){
	
}

func (com Command)ExecutePost(){
	
}
