/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package chat_helper

import (
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dao/mysql_dao"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}

type (
	ChatParticipantsDAO = mysql_dao.ChatParticipantsDAO
	ChatParticipantsDO  = dataobject.ChatParticipantsDO
)

var (
	NewChatParticipantsDAO = mysql_dao.NewChatParticipantsDAO
)
