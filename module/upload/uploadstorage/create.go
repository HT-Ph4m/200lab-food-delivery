package uploadstorage

import (
	"200lab-project-1/common"
	"context"
)

func (store sqlStore) CreateImage(context context.Context, data *common.Image) error {
	//db := store.db
	//
	//if err := db.Table(common.Image{}.TableName()).
	//	Where("id in (?)", ids).
	//	Delete(nil).
	//	Error; err != nil {
	//	return err
	//}

	return nil
}
