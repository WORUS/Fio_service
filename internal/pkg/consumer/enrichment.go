package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Nationality struct {
	Country struct {
		Country string `json:"country_id"`
	}
}

var client *http.Client

func (u *User) GetNationality() {
	url := "https://api.nationalize.io/?name=" + u.Name
	u.GetJSON(url)
}
func (u *User) GetAge() {
	url := "https://api.agify.io/?name=" + u.Name
	u.GetJSON(url)
}

func (u *User) GetGender() {
	url := "https://api.genderize.io/?name=" + u.Name
	u.GetJSON(url)
}

func (u *User) GetJSON(url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	log.Println(resp.Body)

	return json.NewDecoder(resp.Body).Decode(u)
}

func (u *User) Enrich() error {
	client = &http.Client{
		Timeout: time.Second * 5,
	}

	u.GetAge()
	u.GetGender()
	u.GetNationality()

	fmt.Printf("Name: %s,\nSurname: %s,\nAge: %d,\nGender: %s,\nNationality: %s",
		u.Name, u.Surname, u.Age, u.Gender, u.Country[0].CountryId)

	return nil
}
