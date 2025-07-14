package services
import ( 
	"setUp/internal/domain"
)

type User struct {
	domain.BaseModel
	Username string `gorm:"column:username;unique;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Address  string `gorm:"column:address;null"`
	Token    string `gorm:"-"`
}
func (User) TableName() string {
    return "users"
}

type PostUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=100"`
	Username string `json:"username" form:"username" validate:"required"`
}