package main

import (
	"bytes"
	"compress/gzip"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"regexp"
)

var db *sqlx.DB

type Person struct {
	Name        string
	Address     string
	PhoneNumber PhoneNumber
}

type PhoneNumber string

var phoneNumberCheck = regexp.MustCompile(`\(\d{3}\) \d{3}-\d{4}`)

func (p PhoneNumber) Value() (driver.Value, error) {
	matched := phoneNumberCheck.Match([]byte(p))
	if !matched {
		return driver.Value(""), fmt.Errorf("Number '%s' not a valid PhoneNumber format.", p)
	}
	return driver.Value(string(p)), nil
}

type Compressed struct {
	Content GzippedText
}

type GzippedText []byte

func (g GzippedText) Value() (driver.Value, error) {
	b := make([]byte, 0, len(g))
	buf := bytes.NewBuffer(b)
	w := gzip.NewWriter(buf)
	w.Write(g)
	w.Close()
	return buf.Bytes(), nil
}

func (g *GzippedText) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("Incompatible type for GzippedText")
	}
	reader, err := gzip.NewReader(bytes.NewReader(source))
	defer reader.Close()
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	*g = GzippedText(b)
	return nil
}

func PrintIfErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	iq := "INSERT INTO person (name, address, phonenumber) VALUES (:name, :address, :phonenumber)"
	p := Person{Name: "Bin", Address: "Smuth", PhoneNumber: PhoneNumber("(212) 343-1928")}
	_, err := db.NamedExec(iq, &p)
	if err != nil {
		log.Fatal(err)
	}
	invalid := Person{"Ma'a", "Nonu", PhoneNumber("123.456.1234")}
	_, err = db.NamedExec(iq, &invalid)
	PrintIfErr(err)
	_, err = db.Exec("INSERT INTO person (name, address, phonenumber) VALUES (?, ?, ?)", "Ma'a", "Nonu", PhoneNumber("123.456.7890"))
	PrintIfErr(err)

	_, err = db.Exec("INSERT INTO compressed (content) VALUES (?)", GzippedText("Hello, world."))
	PrintIfErr(err)
	row := db.QueryRow("SELECT * FROM compressed LIMIT 1;")
	var raw []byte
	var unz GzippedText
	row.Scan(&raw)
	fmt.Println("Raw:", raw)
	row = db.QueryRow("SELECT * FROM compressed LIMIT 1;")
	err = row.Scan(&unz)
	PrintIfErr(err)
	fmt.Println("Unzipped:", string(unz))

}

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "/dev/shm/typetest.db")
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS person ( name text, address text, phonenumber text);")
	db.Exec("CREATE TABLE IF NOT EXISTS compressed (content text);")
}
