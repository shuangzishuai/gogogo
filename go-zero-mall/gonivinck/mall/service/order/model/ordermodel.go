package model

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		TxInsert(tx *sql.Tx, data *Order) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Order) error
		FindByAllByUid(uid int64) ([]*Order, error)
		FindOneByUid(uid int64) (*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

func (c customOrderModel) FindByAllByUid(uid int64) ([]*Order, error) {
	//TODO implement me
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid`  = ?", orderRows, c.table)
	err := c.QueryRowNoCache(&resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) FindAllByUid(uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid`  = ?", orderRows, m.table)
	err := m.QueryRowNoCache(&resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func (m *defaultOrderModel) TxInsert(tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := tx.Exec(query, data.Uid, data.Pid, data.Amount, data.Status)

	return ret, err
}

func (m *defaultOrderModel) TxUpdate(tx *sql.Tx, data *Order) error {
	productIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return tx.Exec(query, data.Uid, data.Pid, data.Amount, data.Status, data.Id)
	}, productIdKey)
	return err
}

func (m *defaultOrderModel) FindOneByUid(uid int64) (*Order, error) {
	var resp Order
	query := fmt.Sprintf("select %s from %s where `uid` = ? order by create_time desc limit 1", orderRows, m.table)
	err := m.QueryRowNoCache(&resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
