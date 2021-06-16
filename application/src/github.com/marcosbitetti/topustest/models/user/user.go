package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/marcosbitetti/topustest/services/database"
)

// "github.com/marcosbitetti/topustest/services/database"

type User struct {
	Id     string `json:"_id" bson:"_id"`
	Nome   string
	Sexo   byte
	Peso   float32
	Altura float32
	IMC    float32

	CreatedAt time.Time
	UpdatedAt time.Time
	Enabled   bool // used to soft delete
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

/*func (u User) Nome() string {
	return u.Nome
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
*/
func (u User) Calcular() float32 {
	u.IMC = u.Peso / (u.Altura * u.Altura)
	return u.IMC
}

func (u User) Update(nome string) string {
	return u.Nome
}

func (u User) Commit() bool {
	return true
}

func New(nome string, sexo byte, altura float32, peso float32, imc float32) User {
	t := time.Now()
	var u User = User{"", nome, sexo, altura, peso, imc, t, t, true}
	return u
}

/*
 * find by name
 */
func Find(nome string) []User {
	list := database.Find(bson.M{"nome": nome}, &User{}, database.UserCollection())
	var list2 []User = make([]User, len(list))
	for i, _u := range list {
		u, ok := _u.(*User)
		if ok {
			list2[i] = *u
		}
	}
	return list2
}

func FindAll() []User {
	list := database.FindAll(&User{}, database.UserCollection())
	var list2 []User = make([]User, len(list))
	for i, _u := range list {
		u, ok := _u.(*User)
		if ok {
			list2[i] = *u
		}
	}
	return list2
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
