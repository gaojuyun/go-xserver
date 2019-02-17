/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package db

import (
	"errors"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

// RoleList : 代表 1 个 redis 对象
type RoleList struct {
	Key   string
	slot1 uint64
	slot2 uint64
	slot3 uint64
	slot4 uint64
	slot5 uint64

	dirtyDataInRoleList               map[string]interface{}
	dirtyDataForStructFiledInRoleList map[string]interface{}
	isLoadInRoleList                  bool
	dbKeyInRoleList                   string
	ddbNameInRoleList                 string
	expireInRoleList                  uint
}

// NewRoleList : NewRoleList 的构造函数
func NewRoleList(dbName string, key string) *RoleList {
	return &RoleList{
		Key:                               key,
		ddbNameInRoleList:                 dbName,
		dbKeyInRoleList:                   "RoleList:" + key,
		dirtyDataInRoleList:               make(map[string]interface{}),
		dirtyDataForStructFiledInRoleList: make(map[string]interface{}),
	}
}

// HasKey : 是否存在 KEY
//          返回值，若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (objRoleList *RoleList) HasKey() (int, error) {
	db := go_redis_orm.GetDB(objRoleList.ddbNameInRoleList)
	val, err := redis.Int(db.Do("EXISTS", objRoleList.dbKeyInRoleList))
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Load : 从 redis 加载数据
func (objRoleList *RoleList) Load() error {
	if objRoleList.isLoadInRoleList == true {
		return errors.New("alreay load")
	}
	db := go_redis_orm.GetDB(objRoleList.ddbNameInRoleList)
	val, err := redis.Values(db.Do("HGETALL", objRoleList.dbKeyInRoleList))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	var data struct {
		Slot1 uint64 `redis:"slot1"`
		Slot2 uint64 `redis:"slot2"`
		Slot3 uint64 `redis:"slot3"`
		Slot4 uint64 `redis:"slot4"`
		Slot5 uint64 `redis:"slot5"`
	}
	if err := redis.ScanStruct(val, &data); err != nil {
		return err
	}
	objRoleList.slot1 = data.Slot1
	objRoleList.slot2 = data.Slot2
	objRoleList.slot3 = data.Slot3
	objRoleList.slot4 = data.Slot4
	objRoleList.slot5 = data.Slot5
	objRoleList.isLoadInRoleList = true
	return nil
}

// Save : 保存数据到 redis
func (objRoleList *RoleList) Save() error {
	if len(objRoleList.dirtyDataInRoleList) == 0 && len(objRoleList.dirtyDataForStructFiledInRoleList) == 0 {
		return nil
	}
	for k := range objRoleList.dirtyDataForStructFiledInRoleList {
		_ = k

	}
	db := go_redis_orm.GetDB(objRoleList.ddbNameInRoleList)
	if _, err := db.Do("HMSET", redis.Args{}.Add(objRoleList.dbKeyInRoleList).AddFlat(objRoleList.dirtyDataInRoleList)...); err != nil {
		return err
	}
	if objRoleList.expireInRoleList != 0 {
		if _, err := db.Do("EXPIRE", objRoleList.dbKeyInRoleList, objRoleList.expireInRoleList); err != nil {
			return err
		}
	}
	objRoleList.dirtyDataInRoleList = make(map[string]interface{})
	objRoleList.dirtyDataForStructFiledInRoleList = make(map[string]interface{})
	return nil
}

// Delete : 从 redis 删除数据
func (objRoleList *RoleList) Delete() error {
	db := go_redis_orm.GetDB(objRoleList.ddbNameInRoleList)
	_, err := db.Do("DEL", objRoleList.dbKeyInRoleList)
	if err == nil {
		objRoleList.isLoadInRoleList = false
		objRoleList.dirtyDataInRoleList = make(map[string]interface{})
		objRoleList.dirtyDataForStructFiledInRoleList = make(map[string]interface{})
	}
	return err
}

// IsLoad : 是否已经从 redis 导入数据
func (objRoleList *RoleList) IsLoad() bool {
	return objRoleList.isLoadInRoleList
}

// Expire : 向 redis 设置该对象过期时间
func (objRoleList *RoleList) Expire(v uint) {
	objRoleList.expireInRoleList = v
}

// DirtyData : 获取该对象目前已脏的数据
func (objRoleList *RoleList) DirtyData() (map[string]interface{}, error) {
	for k := range objRoleList.dirtyDataForStructFiledInRoleList {
		_ = k

	}
	data := make(map[string]interface{})
	for k, v := range objRoleList.dirtyDataInRoleList {
		data[k] = v
	}
	objRoleList.dirtyDataInRoleList = make(map[string]interface{})
	objRoleList.dirtyDataForStructFiledInRoleList = make(map[string]interface{})
	return data, nil
}

// Save2 : 保存数据到 redis 的第 2 种方法
func (objRoleList *RoleList) Save2(dirtyData map[string]interface{}) error {
	if len(dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(objRoleList.ddbNameInRoleList)
	if _, err := db.Do("HMSET", redis.Args{}.Add(objRoleList.dbKeyInRoleList).AddFlat(dirtyData)...); err != nil {
		return err
	}
	if objRoleList.expireInRoleList != 0 {
		if _, err := db.Do("EXPIRE", objRoleList.dbKeyInRoleList, objRoleList.expireInRoleList); err != nil {
			return err
		}
	}
	return nil
}

// GetSlot1 : 获取字段值
func (objRoleList *RoleList) GetSlot1() uint64 {
	return objRoleList.slot1
}

// GetSlot2 : 获取字段值
func (objRoleList *RoleList) GetSlot2() uint64 {
	return objRoleList.slot2
}

// GetSlot3 : 获取字段值
func (objRoleList *RoleList) GetSlot3() uint64 {
	return objRoleList.slot3
}

// GetSlot4 : 获取字段值
func (objRoleList *RoleList) GetSlot4() uint64 {
	return objRoleList.slot4
}

// GetSlot5 : 获取字段值
func (objRoleList *RoleList) GetSlot5() uint64 {
	return objRoleList.slot5
}

// SetSlot1 : 设置字段值
func (objRoleList *RoleList) SetSlot1(value uint64) {
	objRoleList.slot1 = value
	objRoleList.dirtyDataInRoleList["slot1"] = value
}

// SetSlot2 : 设置字段值
func (objRoleList *RoleList) SetSlot2(value uint64) {
	objRoleList.slot2 = value
	objRoleList.dirtyDataInRoleList["slot2"] = value
}

// SetSlot3 : 设置字段值
func (objRoleList *RoleList) SetSlot3(value uint64) {
	objRoleList.slot3 = value
	objRoleList.dirtyDataInRoleList["slot3"] = value
}

// SetSlot4 : 设置字段值
func (objRoleList *RoleList) SetSlot4(value uint64) {
	objRoleList.slot4 = value
	objRoleList.dirtyDataInRoleList["slot4"] = value
}

// SetSlot5 : 设置字段值
func (objRoleList *RoleList) SetSlot5(value uint64) {
	objRoleList.slot5 = value
	objRoleList.dirtyDataInRoleList["slot5"] = value
}
