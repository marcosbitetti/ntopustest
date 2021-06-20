package user

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/marcosbitetti/ntopustest/services/database"
)

// "github.com/marcosbitetti/topustest/services/database"

type User struct {
	Id     string `json:"_id" bson:"_id,omitempty"`
	Nome   string
	Sexo   byte
	Altura float32
	Peso   float32
	IMC    float32

	CreatedAt time.Time
	UpdatedAt time.Time
	Enabled   bool // used to soft delete, not implemented at this time
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

// Map string to iso to make JSON more readable
var MapSexo = map[string]byte{
	"?": NOT_KNOW,
	"M": MALE,
	"F": FEMALE,
	"_": NOT_APPLICABLE,
}

func (u User) Calcular() float32 {
	u.IMC = u.Peso / (u.Altura * u.Altura)
	return u.IMC
}

func (u User) Update(nome string, sexo byte, altura float32, peso float32, imc float32) (string, bool) {
	// Nome changes? verify rules
	if u.Nome != nome {
		if Exists(nome) {
			return "user-exists", false
		}
	}

	u.Nome = nome
	u.Sexo = sexo
	u.Altura = altura
	u.Peso = peso
	u.IMC = imc
	u.UpdatedAt = time.Now()
	ok := database.Update(bson.M{"nome": u.Nome}, u, database.UserCollection())

	return u.Nome, ok
}

func (u User) Delete() bool {
	return database.Delete(bson.M{"nome": u.Nome}, database.UserCollection())
}

func (u User) Commit() bool {
	if u.Id != "" {
		log.Println("update called")
		return false
	}
	return database.Insert(u, database.UserCollection())
}

func New(nome string, sexo byte, altura float32, peso float32, imc float32) (User, bool) {
	if Exists(nome) {
		return User{}, false
	}

	t := time.Now()
	var u User = User{"", nome, sexo, altura, peso, imc, t, t, true}
	return u, true
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

func Exists(nome string) bool {
	list := Find(nome)
	return len(list) > 0
}

func Delete(u User) bool {
	return false
}
