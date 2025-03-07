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

package server

import (
	"flag"
	scheduledmessages_helper "github.com/teamgram/teamgram-server/app/bff/scheduledmessages"

	"github.com/teamgram/proto/mtproto"
	account_helper "github.com/teamgram/teamgram-server/app/bff/account"
	authorization_helper "github.com/teamgram/teamgram-server/app/bff/authorization"
	autodownload_helper "github.com/teamgram/teamgram-server/app/bff/autodownload"
	"github.com/teamgram/teamgram-server/app/bff/bff/internal/config"
	chats_helper "github.com/teamgram/teamgram-server/app/bff/chats"
	configuration_helper "github.com/teamgram/teamgram-server/app/bff/configuration"
	contacts_helper "github.com/teamgram/teamgram-server/app/bff/contacts"
	dialogs_helper "github.com/teamgram/teamgram-server/app/bff/dialogs"
	drafts_helper "github.com/teamgram/teamgram-server/app/bff/drafts"
	emoji_helper "github.com/teamgram/teamgram-server/app/bff/emoji"
	files_helper "github.com/teamgram/teamgram-server/app/bff/files"
	folders_helper "github.com/teamgram/teamgram-server/app/bff/folders"
	gifs_helper "github.com/teamgram/teamgram-server/app/bff/gifs"
	langpack_helper "github.com/teamgram/teamgram-server/app/bff/langpack"
	messages_helper "github.com/teamgram/teamgram-server/app/bff/messages"
	miscellaneous_helper "github.com/teamgram/teamgram-server/app/bff/miscellaneous"
	notification_helper "github.com/teamgram/teamgram-server/app/bff/notification"
	nsfw_helper "github.com/teamgram/teamgram-server/app/bff/nsfw"
	photos_helper "github.com/teamgram/teamgram-server/app/bff/photos"
	promodata_helper "github.com/teamgram/teamgram-server/app/bff/promodata"
	qrcode_helper "github.com/teamgram/teamgram-server/app/bff/qrcode"
	reactions_helper "github.com/teamgram/teamgram-server/app/bff/reactions"
	reports_helper "github.com/teamgram/teamgram-server/app/bff/reports"
	secretchats_helper "github.com/teamgram/teamgram-server/app/bff/secretchats"
	sponsoredmessages_helper "github.com/teamgram/teamgram-server/app/bff/sponsoredmessages"
	stickers_helper "github.com/teamgram/teamgram-server/app/bff/stickers"
	themes_helper "github.com/teamgram/teamgram-server/app/bff/themes"
	tos_helper "github.com/teamgram/teamgram-server/app/bff/tos"
	twofa_helper "github.com/teamgram/teamgram-server/app/bff/twofa"
	updates_helper "github.com/teamgram/teamgram-server/app/bff/updates"
	usernames_helper "github.com/teamgram/teamgram-server/app/bff/usernames"
	users_helper "github.com/teamgram/teamgram-server/app/bff/users"
	wallpapers_helper "github.com/teamgram/teamgram-server/app/bff/wallpapers"
	webpage_helper "github.com/teamgram/teamgram-server/app/bff/webpage"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// tos_helper
		mtproto.RegisterRPCTosServer(
			grpcServer,
			tos_helper.New(tos_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// reports_helper
		mtproto.RegisterRPCReportsServer(
			grpcServer,
			reports_helper.New(reports_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// configuration_helper
		mtproto.RegisterRPCConfigurationServer(
			grpcServer,
			configuration_helper.New(configuration_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// qrcode_helper
		mtproto.RegisterRPCQrCodeServer(
			grpcServer,
			qrcode_helper.New(qrcode_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				KV:                c.KV,
				UserClient:        c.BizServiceClient,
				AuthSessionClient: c.AuthSessionClient,
				SyncClient:        c.SyncClient,
			}))

		// miscellaneous_helper
		mtproto.RegisterRPCMiscellaneousServer(
			grpcServer,
			miscellaneous_helper.New(miscellaneous_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// authorization_helper
		mtproto.RegisterRPCAuthorizationServer(
			grpcServer,
			authorization_helper.New(authorization_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				KV:                c.KV,
				Code:              c.Code,
				UserClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
				ChatClient:        c.BizServiceClient,
				StatusClient:      c.StatusClient,
				SyncClient:        c.SyncClient,
				MsgClient:         c.MsgClient,
			}, nil))

		// gifs_helper
		mtproto.RegisterRPCGifsServer(
			grpcServer,
			gifs_helper.New(gifs_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// gifs_helper
		mtproto.RegisterRPCPromoDataServer(
			grpcServer,
			promodata_helper.New(promodata_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// twofa_helper
		mtproto.RegisterRPCTwoFaServer(
			grpcServer,
			twofa_helper.New(twofa_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// chats_helper
		mtproto.RegisterRPCChatsServer(
			grpcServer,
			chats_helper.New(chats_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				ChatClient:        c.BizServiceClient,
				MsgClient:         c.MsgClient,
				DialogClient:      c.BizServiceClient,
				SyncClient:        c.SyncClient,
				MediaClient:       c.MediaClient,
				AuthsessionClient: c.AuthSessionClient,
				IdgenClient:       c.IdgenClient,
				MessageClient:     c.BizServiceClient,
			}))

		// files_helper
		mtproto.RegisterRPCFilesServer(
			grpcServer,
			files_helper.New(files_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DfsClient:     c.DfsClient,
				UserClient:    c.BizServiceClient,
				MediaClient:   c.MediaClient,
			}, nil))

		// webpage_helper
		mtproto.RegisterRPCWebPageServer(
			grpcServer,
			webpage_helper.New(webpage_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// secretchats_helper
		mtproto.RegisterRPCSecretChatsServer(
			grpcServer,
			secretchats_helper.New(secretchats_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// updates_helper
		mtproto.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UpdatesClient: c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
			}))

		// themes_helper
		mtproto.RegisterRPCThemesServer(
			grpcServer,
			themes_helper.New(themes_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// contacts_helper
		mtproto.RegisterRPCContactsServer(
			grpcServer,
			contacts_helper.New(contacts_helper.Config{
				RpcServerConf:  c.RpcServerConf,
				UserClient:     c.BizServiceClient,
				ChatClient:     c.BizServiceClient,
				UsernameClient: c.BizServiceClient,
				SyncClient:     c.SyncClient,
			}, nil))

		// dialogs_helper
		mtproto.RegisterRPCDialogsServer(
			grpcServer,
			dialogs_helper.New(dialogs_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UpdatesClient: c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				DialogClient:  c.BizServiceClient,
				SyncClient:    c.SyncClient,
				MessageClient: c.BizServiceClient,
			}, nil))

		// drafts_helper
		mtproto.RegisterRPCDraftsServer(
			grpcServer,
			drafts_helper.New(drafts_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DialogClient:  c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
				ChatClient:    c.BizServiceClient,
			}, nil))

		// emoji_helper
		mtproto.RegisterRPCEmojiServer(
			grpcServer,
			emoji_helper.New(emoji_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// folders_helper
		mtproto.RegisterRPCFoldersServer(
			grpcServer,
			folders_helper.New(folders_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// langpack_helper
		mtproto.RegisterRPCLangpackServer(
			grpcServer,
			langpack_helper.New(langpack_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// autodownload_helper
		mtproto.RegisterRPCAutoDownloadServer(
			grpcServer,
			autodownload_helper.New(autodownload_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// reactions_helper
		mtproto.RegisterRPCReactionsServer(
			grpcServer,
			reactions_helper.New(reactions_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// messages_helper
		mtproto.RegisterRPCMessagesServer(
			grpcServer,
			messages_helper.New(messages_helper.Config{
				RpcServerConf:  c.RpcServerConf,
				UserClient:     c.BizServiceClient,
				ChatClient:     c.BizServiceClient,
				MsgClient:      c.MsgClient,
				DialogClient:   c.BizServiceClient,
				IdgenClient:    c.IdgenClient,
				MessageClient:  c.BizServiceClient,
				MediaClient:    c.MediaClient,
				UsernameClient: c.BizServiceClient,
				SyncClient:     c.SyncClient,
			}))

		// notification_helper
		mtproto.RegisterRPCNotificationServer(
			grpcServer,
			notification_helper.New(notification_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}, nil))

		// users_helper
		mtproto.RegisterRPCUsersServer(
			grpcServer,
			users_helper.New(users_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
			}))

		// scheduledmessages_helper
		mtproto.RegisterRPCScheduledMessagesServer(
			grpcServer,
			scheduledmessages_helper.New(scheduledmessages_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// nsfw_helper
		mtproto.RegisterRPCNsfwServer(
			grpcServer,
			nsfw_helper.New(nsfw_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
			}))

		// sponsoredmessages_helper
		mtproto.RegisterRPCSponsoredMessagesServer(
			grpcServer,
			sponsoredmessages_helper.New(sponsoredmessages_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// stickers_helper
		mtproto.RegisterRPCStickersServer(
			grpcServer,
			stickers_helper.New(stickers_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// account_helper
		mtproto.RegisterRPCAccountServer(
			grpcServer,
			account_helper.New(account_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
				ChatClient:        c.BizServiceClient,
				SyncClient:        c.SyncClient,
			}))

		// photos_helper
		mtproto.RegisterRPCPhotosServer(
			grpcServer,
			photos_helper.New(photos_helper.Config{
				RpcServerConf: c.RpcServerConf,
				MediaClient:   c.MediaClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}))

		// usernames_helper
		mtproto.RegisterRPCUsernamesServer(
			grpcServer,
			usernames_helper.New(usernames_helper.Config{
				RpcServerConf:  c.RpcServerConf,
				UserClient:     c.BizServiceClient,
				UsernameClient: c.BizServiceClient,
				ChatClient:     c.BizServiceClient,
				SyncClient:     c.SyncClient,
			}, nil))

		// usernames_helper
		mtproto.RegisterRPCWallpapersServer(
			grpcServer,
			wallpapers_helper.New(wallpapers_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))
	})

	// logx.Must(err)

	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
