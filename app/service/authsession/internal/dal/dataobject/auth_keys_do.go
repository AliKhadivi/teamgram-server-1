/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type AuthKeysDO struct {
	Id        int64  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	Body      string `db:"body"`
	Deleted   bool   `db:"deleted"`
}
