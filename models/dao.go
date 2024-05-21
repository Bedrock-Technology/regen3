package models

import "gorm.io/gorm"

func GetCursor(orm *gorm.DB) (cursor *Cursor, err error) {
	err = orm.First(&cursor).Error
	return
}

func SaveCursor(orm *gorm.DB, cursor *Cursor) error {
	err := orm.Save(cursor).Error
	return err
}

func GetAllPods(orm *gorm.DB) (pods map[string]Pod, err error) {
	podsSlice := make([]Pod, 0)
	pods = make(map[string]Pod)
	err = orm.Find(&podsSlice).Error
	for _, pod := range podsSlice {
		pods[pod.Address] = pod
	}
	return pods, err
}
