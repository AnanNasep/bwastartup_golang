package user

import "time"

//Pemanggilan tabel user..
//pastikan semua kolom dipanggil beserta type nya
type User struct {
	ID             	int
	Name           	string
	Occupation    	string
	Email          	string
	PasswordHash   	string
	AvatarFileName 	string
	Role           	string
	CreatedAt      	time.Time
	UpdatedAt		time.Time
}