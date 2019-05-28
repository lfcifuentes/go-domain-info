package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"testTruora/data"
)



func InsertSearch(data string, url string) (string, error) {
	// conectarse
	db, errDb := sql.Open("postgres", "user=maxroach dbname=truora_test sslmode=disable port=26257")
	if errDb != nil {
		panic(errDb)
	}
	// insertar la información
	_, err := db.Exec(`INSERT INTO tbl_queries (url,info) VALUES ('`+url+`','`+data+`');`);

	if(err != nil){
		fmt.Println(err)
	}
	return "", nil

}

func CheckData(url string) bool {
	// conectarse
	db, errDb := sql.Open("postgres", "user=maxroach dbname=truora_test sslmode=disable port=26257")
	if errDb != nil {
		panic(errDb)
	}
	id := 0
	// optener el id
	db.QueryRow(`SELECT "querie_id" FROM tbl_queries WHERE url = '`+url+`';`).Scan(&id)

	if(id > 0){
		return false
	}
	return true
}

func ChangesData(d data.Data,url string) (bool,string) {
	// conectarse
	db, errDb := sql.Open("postgres", "user=maxroach dbname=truora_test sslmode=disable port=26257")
	if errDb != nil {
		panic(errDb)
	}
	var compare = data.Data{};
	// optener el id
	db.QueryRow("SELECT info FROM tbl_queries WHERE url = $1", url).Scan(&compare);

	return d.Compare(compare)
}

func UpdateServer(d string, url string) (string, error) {
	// conectarse
	db, errDb := sql.Open("postgres", "user=maxroach dbname=truora_test sslmode=disable port=26257")
	if errDb != nil {
		return "", errDb
	}
	// insertar la información
	_, err := db.Exec(`UPDATE tbl_queries SET info = '`+d+`' WHERE url='`+url+`';`);

	if(err != nil){
		fmt.Println(err)
	}
	return "", nil
}

func GetData() data.App  {
	// conectarse
	db, errDb := sql.Open("postgres", "user=maxroach dbname=truora_test sslmode=disable port=26257")
	if errDb != nil {
		panic(errDb)
	}
	var info = make([]data.DataSearch,1);
	// optener el id
	rows,err := db.Query("SELECT url,info  FROM tbl_queries")
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next()	{
		var url string
		var server string
		if err := rows.Scan(&url,&server); err != nil {
			log.Fatal(err)
		}
		var data data.DataSearch;
		errorjson := json.Unmarshal([]byte(server), &data)
		if(errorjson != nil){
			fmt.Printf("Error :%v",errorjson)
		}
		data.Url = url
		info = append(info, data)
	}

	return data.App{info}
}