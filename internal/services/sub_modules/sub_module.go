package services
import ( 
	"setUp/internal/domain"
	"setUp/internal/services/permissions"
)
type SubModules struct {
	domain.BaseModel
	ModuleId    string       `gorm:"column:module_id;not null"`
	Name        string       `gorm:"column:name;null"`
	Const       string       `gorm:"column:const;not null"`
	Description string       `gorm:"column:description;null"`
	Permissions []services.Permissions `gorm:"foreignKey:ModuleId;column:permissions"`
}

func (SubModules) TableName() string {
    return "sub_modules"
}
