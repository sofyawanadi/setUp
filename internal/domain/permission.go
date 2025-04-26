// user.go

package domain

type Permission struct {
	BaseModel
	RoleId string `gorm:"not null"`
	ModuleId string `gorm:"not null"`
	IsView bool `gorm:"default:false"`
	IsCreate bool `gorm:"default:false"`
	IsUpdate bool `gorm:"default:false"`
	IsDelete bool `gorm:"default:false"`
	// IsPrint bool `gorm:"default:false"`
	// IsExport bool `gorm:"default:false"`
	// IsImport bool `gorm:"default:false"`
	// IsApprove bool `gorm:"default:false"`
	// IsReject bool `gorm:"default:false"`
	// IsVerify bool `gorm:"default:false"`
}

type PermissionRepository interface {
	GetPermission(username string) (*Permission, error)
}
