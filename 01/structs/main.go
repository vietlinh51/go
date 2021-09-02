package main

import "fmt"

type Address struct {
	Location string
	City     string
	Country  string
}
type Person struct {
	Id       string
	FullName string
	Email    string
	Addr     Address
}

func (p Person) String() string {
	return fmt.Sprintf("{%s : %s : %s} lives at {%s, %s, %s}",
		p.Id, p.FullName, p.Email, p.Addr.Location, p.Addr.City, p.Addr.Country)

}

func main() {
	type personRequest struct {
		FullName string
		Email    string
		Location string
		City     string
		Country  string
	}

	pRequest := personRequest{
		FullName: "Trinh Minh Cuong",
		Email:    "cuong@techmaster.vn",
		Location: "12A Viwaseen Tower, 48 Tố Hữu",
		City:     "Hà nội",
		Country:  "Việt nam",
	}

	person := Person{
		Id:       "ox-13",
		FullName: pRequest.FullName,
		Email:    pRequest.Email,
		Addr: Address{
			Location: pRequest.Location,
			City:     pRequest.City,
			Country:  pRequest.Country,
		},
	}
	fmt.Println(person)
}
