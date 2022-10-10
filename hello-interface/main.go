// based on https://www.alexedwards.net/blog/interfaces-explained
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

// Create a Customer type
type Customer struct {
	Name string
	Age  int
}

func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

func main() {
	// Initialize a Count object and pass it to WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog().
	count := Count(3)
	WriteLog(count)

	// customer
	c := &Customer{Name: "Alice", Age: 21}

	// write customer in buffer
	var buf bytes.Buffer
	var err = c.WriteJSON(&buf)
	if err != nil {
		log.Fatalln(err)
	}

	// write customer to file
	f, err := os.Create("tmp_customer")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	err = c.WriteJSON(f)
	if err != nil {
		log.Fatalln(err)
	}

	//writeToDb()

	// wildcard interface{}
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	fmt.Printf("%+v", person)

	// type assert value `age` back to an `int` before using it. As in map still has type of `interface{}`
	age, ok := person["age"].(int)
	if !ok {
		log.Fatal("could not assert value to int")
		return
	}

	person["age"] = age + 1
	fmt.Printf("%+v", person)
}

func writeToDb() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}
