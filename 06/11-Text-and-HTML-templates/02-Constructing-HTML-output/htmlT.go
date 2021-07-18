package main

import (
	"database/sql"
	"fmt"

	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var (
	DATA  []Entry
	tFile string
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))

	// The template.ExecuteTemplate does all the work for us. Its first
	// parameter is the variable that holds the connection with the HTTP
	// client, its second parameter is the template file that will be
	// used for formatting, and its third parameter is the slice of
	// structures with the data.
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {

	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need Database File and Template File!")
		return
	}

	database := arguments[1]
	tFile = arguments[2]

	// sql.Open() opens the connection with the desired database
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Emptying database table")

	// With db.Exec() we can execute database commands without any feedback from them
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Populating", database)

	// The db.Prepare() allows us to execute a database command multiple times by changing
	// only its parameters and calling Exec() afterwards.
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?,?,?,?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var n, d, s int
	for rows.Next() {
		_ = rows.Scan(&n, &d, &s)
		temp := Entry{n, d, s}
		DATA = append(DATA, temp)
	}

	http.HandleFunc("/", myHandler)
	if err = http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
		return
	}

}
