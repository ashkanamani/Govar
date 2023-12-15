package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
)


var Db *sql.DB

func init() {
	var err error 
	Db, err = sql.Open("postgres", "user=postgres dbname=govar sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("cannot generate uuid", err)
		
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid
}

func Encrypt(plaintext string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}