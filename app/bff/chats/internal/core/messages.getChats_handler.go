// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (c *ChatsCore) MessagesGetChats(in *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	chats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: make([]*mtproto.Chat, 0, len(in.Id)),
	}).To_Messages_Chats()

	for _, id := range in.Id {
		chat, _ := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
			ChatId: id,
		})
		if chat != nil {
			chats.Chats = append(chats.Chats, chat.ToUnsafeChat(c.MD.UserId))
		}
	}

	return chats, nil
}
