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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type ChatsDAO struct {
	db *sqlx.DB
}

func NewChatsDAO(db *sqlx.DB) *ChatsDAO {
	return &ChatsDAO{db}
}

// Insert
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)
func (dao *ChatsDAO) Insert(ctx context.Context, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// InsertTx
// insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)
func (dao *ChatsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, default_banned_rights, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :default_banned_rights, :date)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

// Select
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = :id
func (dao *ChatsDAO) Select(ctx context.Context, id int64) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// SelectLastCreator
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = :creator_user_id order by `date` desc limit 1
func (dao *ChatsDAO) SelectLastCreator(ctx context.Context, creator_user_id int64) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, creator_user_id)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectLastCreator(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
	if rows.Next() {
		// TODO(@benqi): not use reflect
		err = rows.StructScan(do)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectLastCreator(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

// UpdateTitle
// update chats set title = :title, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateTitle(ctx context.Context, title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, title, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

// UpdateTitleTx
// update chats set title = :title, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateTitleTx(tx *sqlx.Tx, title string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, title, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

// UpdateAbout
// update chats set about = :about where id = :id
func (dao *ChatsDAO) UpdateAbout(ctx context.Context, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, about, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// UpdateAboutTx
// update chats set about = :about where id = :id
func (dao *ChatsDAO) UpdateAboutTx(tx *sqlx.Tx, about string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

// SelectByIdList
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (dao *ChatsDAO) SelectByIdList(ctx context.Context, idList []int32) (rList []dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChatsDO
	for rows.Next() {
		v := dataobject.ChatsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// SelectByIdListWithCB
// select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (:idList)
func (dao *ChatsDAO) SelectByIdListWithCB(ctx context.Context, idList []int32, cb func(i int, v *dataobject.ChatsDO)) (rList []dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo_id, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	if len(idList) == 0 {
		rList = []dataobject.ChatsDO{}
		return
	}

	query, a, err = sqlx.In(query, idList)
	if err != nil {
		// r sql.Result
		logx.WithContext(ctx).Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectByIdList(_), error: %v", err)
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

	var values []dataobject.ChatsDO
	for rows.Next() {
		v := dataobject.ChatsDO{}

		// TODO(@benqi): not use reflect
		err = rows.StructScan(&v)
		if err != nil {
			logx.WithContext(ctx).Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

// UpdateParticipantCount
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateParticipantCount(ctx context.Context, participant_count int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_count, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

// UpdateParticipantCountTx
// update chats set participant_count = :participant_count, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateParticipantCountTx(tx *sqlx.Tx, participant_count int32, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_count, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

// UpdatePhotoId
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdatePhotoId(ctx context.Context, photo_id int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, photo_id, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePhotoId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePhotoId(_), error: %v", err)
	}

	return
}

// UpdatePhotoIdTx
// update chats set photo_id = :photo_id, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdatePhotoIdTx(tx *sqlx.Tx, photo_id int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo_id = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photo_id, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePhotoId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePhotoId(_), error: %v", err)
	}

	return
}

// UpdateDefaultBannedRights
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDefaultBannedRights(ctx context.Context, default_banned_rights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, default_banned_rights, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

// UpdateDefaultBannedRightsTx
// update chats set default_banned_rights = :default_banned_rights, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDefaultBannedRightsTx(tx *sqlx.Tx, default_banned_rights int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, default_banned_rights, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

// UpdateVersion
// update chats set version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateVersion(ctx context.Context, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

// UpdateVersionTx
// update chats set version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateVersionTx(tx *sqlx.Tx, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

// UpdateDeactivated
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDeactivated(ctx context.Context, deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, deactivated, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

// UpdateDeactivatedTx
// update chats set deactivated = :deactivated, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateDeactivatedTx(tx *sqlx.Tx, deactivated bool, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, deactivated, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

// UpdateMigratedTo
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateMigratedTo(ctx context.Context, migrated_to_id int64, migrated_to_access_hash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, migrated_to_id, migrated_to_access_hash, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}

// UpdateMigratedToTx
// update chats set migrated_to_id = :migrated_to_id, migrated_to_access_hash = :migrated_to_access_hash, participant_count = 0, deactivated = 1, version = version + 1 where id = :id
func (dao *ChatsDAO) UpdateMigratedToTx(tx *sqlx.Tx, migrated_to_id int64, migrated_to_access_hash int64, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, migrated_to_id, migrated_to_access_hash, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}
