package user

// ini cuman format aja
// maksudnya untuk meng custom si nama-nama dari json nya..
// contohnya seperti huruf nya besar jadi kecil oleh `json:"id"`
// atau hanya menampilkan beberapa field dari email

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageURL: 	user.AvatarFileName,
	}
	return formatter
}