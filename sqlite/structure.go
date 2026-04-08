package sqlite

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id              int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Role            string         `gorm:"size:100;not null;index;comment:角色"`
	RootRoleDesc    string         `gorm:"size:255;comment:根角色描述"`
	DerivedRole     string         `gorm:"size:100;not null;index;comment:PT-派生角色"`
	DerivedRoleDesc string         `gorm:"size:500;comment:PT-派生描述"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

/*
插入单条角色记录
*/
func (r *Role) InsertOne() error {
	db := GetSqlite()
	result := db.Create(r)
	return result.Error
}

/*
批量插入角色记录
*/
func InsertRoles(roles []Role) error {
	db := GetSqlite()
	result := db.Create(&roles)
	return result.Error
}

/*
根据ID查询角色
*/
func FindRoleByID(id int64) (*Role, error) {
	db := GetSqlite()
	var role Role
	result := db.First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

/*
根据角色代码查询
*/
func FindRoleByRoleCode(roleCode string) ([]Role, error) {
	db := GetSqlite()
	var roles []Role
	result := db.Where("role = ?", roleCode).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

/*
根据派生角色代码查询
*/
func FindRoleByDerivedRole(derivedRole string) ([]Role, error) {
	db := GetSqlite()
	var roles []Role
	result := db.Where("derived_role = ?", derivedRole).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

/*
查询所有角色记录
*/
func FindAllRoles() ([]Role, error) {
	db := GetSqlite()
	var roles []Role
	result := db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

/*
更新角色记录
*/
func (r *Role) Update() error {
	db := GetSqlite()
	result := db.Save(r)
	return result.Error
}

/*
根据ID删除角色记录（软删除）
*/
func DeleteRoleByID(id int64) error {
	db := GetSqlite()
	result := db.Delete(&Role{}, id)
	return result.Error
}

/*
根据角色代码删除（软删除）
*/
func DeleteRoleByRoleCode(roleCode string) error {
	db := GetSqlite()
	result := db.Where("role = ?", roleCode).Delete(&Role{})
	return result.Error
}

/*
统计角色总数
*/
func CountRoles() (int64, error) {
	db := GetSqlite()
	var count int64
	result := db.Model(&Role{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

/*
检查角色是否存在
*/
func (r *Role) Exists() (bool, error) {
	db := GetSqlite()
	var count int64
	result := db.Model(&Role{}).Where("role = ? AND derived_role = ?", r.Role, r.DerivedRole).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
