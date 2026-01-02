package test

import (
	"go-clean-architecture-pzn/entity"
	"strconv"

	"github.com/google/uuid"
)

func ClearAll() {

}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data: %+v", err)
	}
}

func ClearContact() {
	err := db.Where("id is not null").Delete(&entity.Contact{}).Error
	if err != nil {
		log.Fatalf("Failed clear contact data: %+v", err)
	}
}

func ClearAddresses() {
	err := db.Where("id is not null").Delete(&entity.Address{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data: %+v", err)
	}
}

func CreateContacts(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		contact := &entity.Contact{
			ID:        uuid.NewString(),
			FirstName: "Contact",
			LastName:  strconv.Itoa(i),
			Email:     "contact" + strconv.Itoa(i) + "@example.com",
			Phone:     "0999" + strconv.Itoa(i),
			UserId:    user.ID,
		}

		err := db.Create(contact).Error
		if err != nil {
			log.Fatalf("Failed create contact data: %+v", err)
		}
	}
}
