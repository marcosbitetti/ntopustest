package user

import (
	db "github.com/marcosbitetti/topustest/services/database"
)

type User struct {
	nome   string
	sexo   byte
	peso   float32
	altura float32
	imc    float32
}

/*
 * ISO/IEC 5218
 * ref https://en.wikipedia.org/wiki/ISO/IEC_5218
 */
const (
	NOT_KNOW       byte = 0
	MALE           byte = 1
	FEMALE         byte = 2
	NOT_APPLICABLE byte = 9
)

func (u User) Nome() string {
	return u.nome
}

func (u User) Sexo() byte {
	return u.sexo
}

func (u User) Peso() float32 {
	return u.peso
}

func (u User) Altura() float32 {
	return u.altura
}

func (u User) IMC() float32 {
	return u.imc
}

func (u User) Calcular() float32 {
	u.imc = u.peso / (u.altura * u.altura)
	return u.imc
}

func (u User) Update(nome string) string {
	return u.nome
}

func (u User) Commit() bool {
	return true
}

func New(nome string, sexo byte, altura float32, peso float32, imc float32) User {
	var u User = User{nome, sexo, altura, peso, imc}
	return u
}

/*
 * find by name
 */
func Find(nome string) []User {
	var users []User
	users = append(users, New("miguel", MALE, 1.2, 60.0, 0.2))
	if db.DB() {
		users = append(users, New("ruel", MALE, 1.2, 60.0, 0.2))
	}
	return users
}

/*
 * verifica se o nome já existe
 */
func Collide(nome string) bool {
	return false
}

func Delete(u User) bool {
	return false
}
