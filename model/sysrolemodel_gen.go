// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"
)

type (
	sysRoleModel interface {
		Insert(ctx context.Context, data *SysRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRole, error)
		Update(ctx context.Context, data *SysRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRole struct {
		Id          int64          `db:"id"`           // 标识
		Name        sql.NullString `db:"name"`         // 角色名称
		Code        sql.NullString `db:"code"`         // 角色编码
		Description sql.NullString `db:"description"`  // 描述
		CreatedBy   sql.NullInt64  `db:"created_by"`   // 创建人
		CreatedTime sql.NullTime   `db:"created_time"` // 创建时间
		UpdatedBy   sql.NullInt64  `db:"updated_by"`   // 更新人
		UpdatedTime sql.NullTime   `db:"updated_time"` // 更新时间
	}
)

func newSysRoleModel(conn sqlx.SqlConn) *defaultSysRoleModel {
	return &defaultSysRoleModel{
		conn:  conn,
		table: "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysRoleModel) FindOne(ctx context.Context, id int64) (*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	var resp SysRole
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleModel) Insert(ctx context.Context, data *SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Code, data.Description, data.CreatedBy, data.CreatedTime, data.UpdatedBy, data.UpdatedTime)
	return ret, err
}

func (m *defaultSysRoleModel) Update(ctx context.Context, data *SysRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Code, data.Description, data.CreatedBy, data.CreatedTime, data.UpdatedBy, data.UpdatedTime, data.Id)
	return err
}

func (m *defaultSysRoleModel) tableName() string {
	return m.table
}