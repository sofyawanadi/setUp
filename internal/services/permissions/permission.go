// user.go

package domain

type Permission struct {
	BaseModel
	RoleId   string `gorm:"not null;column:role_id"`
	ModuleId string `gorm:"not null;column:module_id"`
	IsView   bool   `gorm:"default:false;column:is_view"`
	IsCreate bool   `gorm:"default:false;column:is_create"`
	IsUpdate bool   `gorm:"default:false;column:is_update"`
	IsDelete bool   `gorm:"default:false;column:is_delete"`
	// IsPrint bool `gorm:"default:false;column:is_print"`
	// IsExport bool `gorm:"default:false;column:is_export"`
	// IsImport bool `gorm:"default:false;column:is_import"`
	// IsApprove bool `gorm:"default:false;column:is_approve"`
	// IsReject bool `gorm:"default:false;column:is_reject"`
	// IsVerify bool `gorm:"default:false;column:is_verify"`
}

func (Permission) TableName() string {
    return "permissions"
}

type PermissionRepository interface {
	GetAll(username string) (*Permission, error)
}
