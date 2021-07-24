package user

//sercive ini dependency dari handler

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	//login
	Login(input LoginInput)(User, error)
	//cari email 
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	//upload avatar
	SaveAvatar(ID int, fileLocation string)(User, error)
	//middleware update
	GetUserByID(ID int)(User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil	
}

//login
func (s *service) Login(input LoginInput)(User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err	
	}
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err	
	}
	return user, nil
}
//mapping struck input ke struck user
//simpan struck user melalui repository


//cari email ketika input 
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err	
	}	
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//Dapatkan user berdasarkan ID
	//upload attribute avatar file name
	//simpan perubahan avatar file

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err	
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err	
	}
	return updatedUser, nil
}

//MIDDLEWARE UPDATE
func (s *service) GetUserByID(ID int)(User, error){
	user, err := s.repository.FindByID(ID)
	if err != nil{
		return user, err 
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}
	return user, nil
}