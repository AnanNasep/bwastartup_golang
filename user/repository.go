package user

//repository ini dependency dari sercive

import "gorm.io/gorm"

// repository adalah penyimpanan...
type Repository interface{
	Save(user User)(User, error)
}


type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Save(user User)(User, error){
	err := r.db.Create(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}