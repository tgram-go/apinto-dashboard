package store

import (
	"context"
	"github.com/eolinker/apinto-dashboard/entry/monitor-entry"
	"github.com/eolinker/apinto-dashboard/entry/page-entry"
)

type IWarnStrategyIStore interface {
	IBaseStore[monitor_entry.WarnStrategy]
	GetByUuid(ctx context.Context, uuid string) (*monitor_entry.WarnStrategy, error)
	GetPage(ctx context.Context, namespaceId, partitionId int, name string, dimension []string, status, pageNum, pageSize int) ([]*monitor_entry.WarnStrategy, int64, error)
	GetAll(ctx context.Context, namespaceId, status int) ([]*monitor_entry.WarnStrategy, error)
	UpdateIsEnable(ctx context.Context, id int, isEnable bool) error
	GetByTitle(ctx context.Context, namespaceId, partitionId int, title string) (*monitor_entry.WarnStrategy, error)
	GetByPartitionId(ctx context.Context, namespaceId, partitionId int) ([]*monitor_entry.WarnStrategy, error)
}

type warnStrategyStore struct {
	*BaseStore[monitor_entry.WarnStrategy]
}

func newWarnStrategyIStore(db IDB) IWarnStrategyIStore {
	return &warnStrategyStore{BaseStore: CreateStore[monitor_entry.WarnStrategy](db)}
}

func (w *warnStrategyStore) GetByUuid(ctx context.Context, uuid string) (*monitor_entry.WarnStrategy, error) {
	return w.FirstQuery(ctx, "`uuid` = ?", []interface{}{uuid}, "")
}

func (w *warnStrategyStore) GetByTitle(ctx context.Context, namespaceId, partitionId int, title string) (*monitor_entry.WarnStrategy, error) {
	return w.FirstQuery(ctx, "`namespace` = ? and `partition_id` = ? and `title` = ?", []interface{}{namespaceId, partitionId, title}, "")
}

func (w *warnStrategyStore) UpdateIsEnable(ctx context.Context, id int, isEnable bool) error {
	_, err := w.UpdateWhere(ctx, &monitor_entry.WarnStrategy{Id: id}, map[string]interface{}{"is_enable": isEnable})
	if err != nil {
		return err
	}
	return nil
}

func (w *warnStrategyStore) GetAll(ctx context.Context, namespaceId, status int) ([]*monitor_entry.WarnStrategy, error) {
	db := w.DB(ctx).Where("`namespace` = ?", namespaceId)
	if status > -1 {
		db = db.Where("`is_enable` = ?", status)
	}
	list := make([]*monitor_entry.WarnStrategy, 0)
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (w *warnStrategyStore) GetByPartitionId(ctx context.Context, namespaceId, partitionId int) ([]*monitor_entry.WarnStrategy, error) {
	db := w.DB(ctx).Where("`namespace` = ?", namespaceId)
	db = db.Where("`partition_id` = ?", partitionId)
	list := make([]*monitor_entry.WarnStrategy, 0)
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (w *warnStrategyStore) GetPage(ctx context.Context, namespaceId, partitionId int, name string, dimension []string, status, pageNum, pageSize int) ([]*monitor_entry.WarnStrategy, int64, error) {
	db := w.DB(ctx).Where("`namespace` = ? and `partition_id` = ?", namespaceId, partitionId)
	if name != "" {
		db = db.Where("`title` like ?", "%"+name+"%")
	}
	if len(dimension) > 0 {
		db = db.Where("`dimension` in (?)", dimension)
	}
	if status > -1 {
		db = db.Where("`is_enable` = ?", status)
	}

	list := make([]*monitor_entry.WarnStrategy, 0)

	count := int64(0)

	if err := db.Model(list).Count(&count).Limit(pageSize).Offset(page_entry.PageIndex(pageNum, pageSize)).Order("`update_time` desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}
