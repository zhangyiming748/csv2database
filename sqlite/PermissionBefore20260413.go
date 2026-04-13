package sqlite

import (
	"time"

	"gorm.io/gorm"
)

type PermissionBefore20260413 struct {
	Id             int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Username       string         `gorm:"size:50;not null;index;comment:用户名"`
	FullName       string         `gorm:"size:100;comment:全名"`
	Role           string         `gorm:"size:100;not null;index;comment:角色"`
	Type           string         `gorm:"size:50;comment:类型"`
	AssignmentType string         `gorm:"size:50;comment:分配类型"`
	Assignment     string         `gorm:"size:100;comment:分配"`
	StartDate      string         `gorm:"size:20;comment:开始日期"`
	EndDate        string         `gorm:"size:20;comment:结束日期"`
	ShortRoleDesc  string         `gorm:"size:500;comment:简短角色描述"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

/*
插入单条权限记录
*/
func (p *PermissionBefore20260413) InsertOne() error {
	db := GetSqlite()
	result := db.Create(p)
	return result.Error
}

/*
批量插入权限记录
*/
func InsertPermissions(permissions []PermissionBefore20260413) error {
	db := GetSqlite()
	result := db.Create(&permissions)
	return result.Error
}

/*
根据ID查询权限
*/
func FindPermissionByID(id int64) (*PermissionBefore20260413, error) {
	db := GetSqlite()
	var permission PermissionBefore20260413
	result := db.First(&permission, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &permission, nil
}

/*
根据用户名查询
*/
func FindPermissionsByUsername(username string) ([]PermissionBefore20260413, error) {
	db := GetSqlite()
	var permissions []PermissionBefore20260413
	result := db.Where("username = ?", username).Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return permissions, nil
}

/*
根据角色代码查询
*/
func FindPermissionsByRole(role string) ([]PermissionBefore20260413, error) {
	db := GetSqlite()
	var permissions []PermissionBefore20260413
	result := db.Where("role = ?", role).Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return permissions, nil
}

/*
查询所有权限记录
*/
func FindAllPermissions() ([]PermissionBefore20260413, error) {
	db := GetSqlite()
	var permissions []PermissionBefore20260413
	result := db.Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return permissions, nil
}

/*
更新权限记录
*/
func (p *PermissionBefore20260413) Update() error {
	db := GetSqlite()
	result := db.Save(p)
	return result.Error
}

/*
根据ID删除权限记录（软删除）
*/
func DeletePermissionByID(id int64) error {
	db := GetSqlite()
	result := db.Delete(&PermissionBefore20260413{}, id)
	return result.Error
}

/*
根据用户名删除（软删除）
*/
func DeletePermissionsByUsername(username string) error {
	db := GetSqlite()
	result := db.Where("username = ?", username).Delete(&PermissionBefore20260413{})
	return result.Error
}

/*
统计权限总数
*/
func CountPermissions() (int64, error) {
	db := GetSqlite()
	var count int64
	result := db.Model(&PermissionBefore20260413{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

/*
检查权限是否存在
*/
func (p *PermissionBefore20260413) Exists() (bool, error) {
	db := GetSqlite()
	var count int64
	result := db.Model(&PermissionBefore20260413{}).Where("username = ? AND role = ?", p.Username, p.Role).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
