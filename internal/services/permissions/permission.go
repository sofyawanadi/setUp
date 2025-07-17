package services

import (
	"setUp/internal/domain"
)

type Permissions struct {
	domain.BaseModel
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

func (Permissions) TableName() string {
	return "permissions"
}
