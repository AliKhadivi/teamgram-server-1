/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type MessagesDAO struct {
	db *sqlx.DB
}

func NewMessagesDAO(db *sqlx.DB) *MessagesDAO {
	return &MessagesDAO{db}
}

// InsertOrReturnId
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :date2) on duplicate key update id = last_insert_id(id)
func (dao *MessagesDAO) InsertOrReturnId(ctx context.Context, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :date2) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

// InsertOrReturnIdTx
// insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :date2) on duplicate key update id = last_insert_id(id)
func (dao *MessagesDAO) InsertOrReturnIdTx(tx *sqlx.Tx, do *dataobject.MessagesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into messages(user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2) values (:user_id, :user_message_box_id, :dialog_id1, :dialog_id2, :dialog_message_id, :sender_user_id, :peer_type, :peer_id, :random_id, :message_filter_type, :message_data, :message, :mentioned, :media_unread, :pinned, :date2) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrReturnId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrReturnId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrReturnId(%v)_error: %v", do, err)
	}

	return
}

// SelectByRandomId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where sender_user_id = :sender_user_id and random_id = :random_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByRandomId(ctx context.Context, sender_user_id int64, random_id int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where sender_user_id = ? and random_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sender_user_id, random_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByRandomId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByRandomId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageIdList(ctx context.Context, user_id int64, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and deleted = 0 and user_message_box_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and deleted = 0 and user_message_box_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageIdListWithCB(ctx context.Context, user_id int64, idList []int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and deleted = 0 and user_message_box_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByMessageId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByMessageId(ctx context.Context, user_id int64, user_message_box_id int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectByMessageDataIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where deleted = 0 and dialog_message_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageDataIdList(ctx context.Context, idList []int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where deleted = 0 and dialog_message_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByMessageDataIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageDataIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByMessageDataIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where deleted = 0 and dialog_message_id in (:idList) order by user_message_box_id desc
func (dao *MessagesDAO) SelectByMessageDataIdListWithCB(ctx context.Context, idList []int64, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where deleted = 0 and dialog_message_id in (?) order by user_message_box_id desc"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByMessageDataIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageDataIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByMessageDataId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and dialog_message_id = :dialog_message_id and deleted = 0 limit 1
func (dao *MessagesDAO) SelectByMessageDataId(ctx context.Context, user_id int64, dialog_message_id int64) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and dialog_message_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_message_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMessageDataId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMessageDataId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectBackwardByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetIdLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectBackwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetIdLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetIdLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetIdLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectBackwardByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetDateLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardByOffsetDateLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectBackwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 < :date2 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardByOffsetDateLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, date2 int64, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardByOffsetDateLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardByOffsetDateLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetDateLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, date2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardByOffsetDateLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardByOffsetDateLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and date2 >= :date2 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardByOffsetDateLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, date2 int64, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and date2 >= ? and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, date2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardByOffsetDateLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardByOffsetDateLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerUserMessageId
// select user_message_box_id, message_box_type from messages where user_id = :peerId and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1)
func (dao *MessagesDAO) SelectPeerUserMessageId(ctx context.Context, peerId int32, user_id int64, user_message_box_id int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_message_box_id, message_box_type from messages where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peerId, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerUserMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerUserMessageId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectPeerUserMessage
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :peerId and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id and deleted = 0 limit 1)
func (dao *MessagesDAO) SelectPeerUserMessage(ctx context.Context, peerId int64, user_id int64, user_message_box_id int32) (rValue *dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and deleted = 0 and dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ? and deleted = 0 limit 1)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peerId, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerUserMessage(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.MessagesDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerUserMessage(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectPeerDialogMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != :user_id and dialog_message_id in (select dialog_message_id from messages where user_id = :user_id and user_message_box_id in (:idList)) and deleted = 0
func (dao *MessagesDAO) SelectPeerDialogMessageIdList(ctx context.Context, user_id int64, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != ? and dialog_message_id in (select dialog_message_id from messages where user_id = ? and user_message_box_id in (?)) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerDialogMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != :user_id and dialog_message_id in (select dialog_message_id from messages where user_id = :user_id and user_message_box_id in (:idList)) and deleted = 0
func (dao *MessagesDAO) SelectPeerDialogMessageIdListWithCB(ctx context.Context, user_id int64, idList []int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != ? and dialog_message_id in (select dialog_message_id from messages where user_id = ? and user_message_box_id in (?)) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerDialogMessageIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogMessageListByMessageId
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id) and deleted = 0
func (dao *MessagesDAO) SelectDialogMessageListByMessageId(ctx context.Context, user_id int64, user_message_box_id int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ?) and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageListByMessageId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogMessageListByMessageId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogMessageListByMessageIdWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where dialog_message_id = (select dialog_message_id from messages where user_id = :user_id and user_message_box_id = :user_message_box_id) and deleted = 0
func (dao *MessagesDAO) SelectDialogMessageListByMessageIdWithCB(ctx context.Context, user_id int64, user_message_box_id int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where dialog_message_id = (select dialog_message_id from messages where user_id = ? and user_message_box_id = ?) and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageListByMessageId(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogMessageListByMessageId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogLastMessageId
// select user_message_box_id from messages where user_id = :user_id and dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2 and deleted = 0 order by user_message_box_id desc limit 1
func (dao *MessagesDAO) SelectDialogLastMessageId(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64) (rValue int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and dialog_id1 = ? and dialog_id2 = ? and deleted = 0 order by user_message_box_id desc limit 1"
	err = dao.db.Get(ctx, &rValue, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("get in SelectDialogLastMessageId(_), error: %v", err)
	}

	return
}

// SelectDialogsByMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectDialogsByMessageIdList(ctx context.Context, user_id int64, idList []int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogsByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogsByMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) SelectDialogsByMessageIdListWithCB(ctx context.Context, user_id int64, idList []int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)

	if len(idList) == 0 {
		rList = []dataobject.MessagesDO{}
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogsByMessageIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogsByMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogLastMessageList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectDialogLastMessageList(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogLastMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogLastMessageListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectDialogLastMessageListWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogLastMessageList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogLastMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// DeleteMessagesByMessageIdList
// update messages set deleted = 1 where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) DeleteMessagesByMessageIdList(ctx context.Context, user_id int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

// DeleteMessagesByMessageIdListTx
// update messages set deleted = 1 where user_id = :user_id and user_message_box_id in (:idList) and deleted = 0
func (dao *MessagesDAO) DeleteMessagesByMessageIdListTx(tx *sqlx.Tx, user_id int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set deleted = 1 where user_id = ? and user_message_box_id in (?) and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteMessagesByMessageIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteMessagesByMessageIdList(_), error: %v", err)
	}

	return
}

// SelectDialogMessageIdList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectDialogMessageIdList(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectDialogMessageIdListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectDialogMessageIdListWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectDialogMessageIdList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectDialogMessageIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerMessageList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != :user_id and dialog_message_id = :dialog_message_id and deleted = 0
func (dao *MessagesDAO) SelectPeerMessageList(ctx context.Context, user_id int64, dialog_message_id int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != ? and dialog_message_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_message_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerMessageList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPeerMessageListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != :user_id and dialog_message_id = :dialog_message_id and deleted = 0
func (dao *MessagesDAO) SelectPeerMessageListWithCB(ctx context.Context, user_id int64, dialog_message_id int64, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id != ? and dialog_message_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_message_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPeerMessageList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPeerMessageList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// UpdateMediaUnread
// update messages set media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMediaUnread(ctx context.Context, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMediaUnreadTx
// update messages set media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMediaUnreadTx(tx *sqlx.Tx, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMentionedAndMediaUnread
// update messages set mentioned = 0, media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMentionedAndMediaUnread(ctx context.Context, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

// UpdateMentionedAndMediaUnreadTx
// update messages set mentioned = 0, media_unread = 0 where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateMentionedAndMediaUnreadTx(tx *sqlx.Tx, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set mentioned = 0, media_unread = 0 where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMentionedAndMediaUnread(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMentionedAndMediaUnread(_), error: %v", err)
	}

	return
}

// SelectByMediaType
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectByMediaType(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, message_filter_type int32, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, message_filter_type, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMediaType(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMediaType(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByMediaTypeWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectByMediaTypeWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, message_filter_type int32, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, message_filter_type, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByMediaType(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByMediaType(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPhoneCallList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectPhoneCallList(ctx context.Context, user_id int64, message_filter_type int32, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_filter_type, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPhoneCallList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPhoneCallList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPhoneCallListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and message_filter_type = :message_filter_type and user_message_box_id < :user_message_box_id and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectPhoneCallListWithCB(ctx context.Context, user_id int64, message_filter_type int32, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and message_filter_type = ? and user_message_box_id < ? and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, message_filter_type, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPhoneCallList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPhoneCallList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// Search
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 and message != '' and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) Search(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, q2 string, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Search(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Search(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SearchWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and deleted = 0 and message != '' and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, q2 string, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Search(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Search(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SearchGlobal
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and user_message_box_id < :user_message_box_id and deleted = 0 and message != '' and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchGlobal(ctx context.Context, user_id int64, user_message_box_id int32, q2 string, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchGlobal(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SearchGlobal(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SearchGlobalWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and user_message_box_id < :user_message_box_id and deleted = 0 and message != '' and message like :q2 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SearchGlobalWithCB(ctx context.Context, user_id int64, user_message_box_id int32, q2 string, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and user_message_box_id < ? and deleted = 0 and message != '' and message like ? order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, user_message_box_id, q2, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SearchGlobal(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SearchGlobal(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectBackwardUnreadMentionsByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardUnreadMentionsByOffsetIdLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectBackwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id < :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit :limit
func (dao *MessagesDAO) SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id < ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectBackwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardUnreadMentionsByOffsetIdLimit
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardUnreadMentionsByOffsetIdLimit(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectForwardUnreadMentionsByOffsetIdLimitWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and user_message_box_id >= :user_message_box_id and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit :limit
func (dao *MessagesDAO) SelectForwardUnreadMentionsByOffsetIdLimitWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, user_message_box_id int32, limit int32, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and user_message_box_id >= ? and mentioned = 1 and media_unread = 1 and deleted = 0 order by user_message_box_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2, user_message_box_id, limit)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectForwardUnreadMentionsByOffsetIdLimit(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPinnedList
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedList(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPinnedList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectPinnedListWithCB
// select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedListWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, cb func(i int, v *dataobject.MessagesDO)) (rList []dataobject.MessagesDO, err error) {
	var (
		query = "select user_id, user_message_box_id, dialog_id1, dialog_id2, dialog_message_id, sender_user_id, peer_type, peer_id, random_id, message_filter_type, message_data, message, mentioned, media_unread, pinned, date2 from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectPinnedList(_), error: %v", err)
		return
	}

	defer func() {
		rows.Close()
		if err == nil && cb != nil {
			for i := 0; i < len(rList); i++ {
				cb(i, &rList[i])
			}
		}
	}()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectPinnedList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectLastTwoPinnedList
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (dao *MessagesDAO) SelectLastTwoPinnedList(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = dao.db.Select(ctx, &rList, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectLastTwoPinnedList(_), error: %v", err)
	}

	return
}

// SelectLastTwoPinnedListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2
func (dao *MessagesDAO) SelectLastTwoPinnedListWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, cb func(i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc limit 2"
	err = dao.db.Select(ctx, &rList, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectLastTwoPinnedList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// UpdatePinned
// update messages set pinned = :pinned where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdatePinned(ctx context.Context, pinned bool, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, pinned, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

// UpdatePinnedTx
// update messages set pinned = :pinned where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdatePinnedTx(tx *sqlx.Tx, pinned bool, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, pinned, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePinned(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePinned(_), error: %v", err)
	}

	return
}

// SelectPinnedMessageIdList
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedMessageIdList(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = dao.db.Select(ctx, &rList, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPinnedMessageIdList(_), error: %v", err)
	}

	return
}

// SelectPinnedMessageIdListWithCB
// select user_message_box_id from messages where user_id = :user_id and (dialog_id1 = :dialog_id1 and dialog_id2 = :dialog_id2) and pinned = 1 and deleted = 0 order by user_message_box_id desc
func (dao *MessagesDAO) SelectPinnedMessageIdListWithCB(ctx context.Context, user_id int64, dialog_id1 int64, dialog_id2 int64, cb func(i int, v int32)) (rList []int32, err error) {
	var query = "select user_message_box_id from messages where user_id = ? and (dialog_id1 = ? and dialog_id2 = ?) and pinned = 1 and deleted = 0 order by user_message_box_id desc"
	err = dao.db.Select(ctx, &rList, query, user_id, dialog_id1, dialog_id2)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectPinnedMessageIdList(_), error: %v", err)
	}

	if cb != nil {
		for i := 0; i < len(rList); i++ {
			cb(i, rList[i])
		}
	}

	return
}

// UpdateUnPinnedByIdList
// update messages set pinned = 0 where user_id = :user_id and user_message_box_id in (:idList)
func (dao *MessagesDAO) UpdateUnPinnedByIdList(ctx context.Context, user_id int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = 0 where user_id = ? and user_message_box_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateUnPinnedByIdList(_), error: %v", err)
	}

	return
}

// UpdateUnPinnedByIdListTx
// update messages set pinned = 0 where user_id = :user_id and user_message_box_id in (:idList)
func (dao *MessagesDAO) UpdateUnPinnedByIdListTx(tx *sqlx.Tx, user_id int64, idList []int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set pinned = 0 where user_id = ? and user_message_box_id in (?)"
		a       []interface{}
		rResult sql.Result
	)

	if len(idList) == 0 {
		return
	}

	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(tx.Context()).Errorf("sqlx.In in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateUnPinnedByIdList(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateUnPinnedByIdList(_), error: %v", err)
	}

	return
}

// UpdateEditMessage
// update messages set message_data = :message_data, message = :message where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateEditMessage(ctx context.Context, message_data string, message string, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, message_data, message, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}

// UpdateEditMessageTx
// update messages set message_data = :message_data, message = :message where user_id = :user_id and user_message_box_id = :user_message_box_id
func (dao *MessagesDAO) UpdateEditMessageTx(tx *sqlx.Tx, message_data string, message string, user_id int64, user_message_box_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update messages set message_data = ?, message = ? where user_id = ? and user_message_box_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, message_data, message, user_id, user_message_box_id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateEditMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateEditMessage(_), error: %v", err)
	}

	return
}
