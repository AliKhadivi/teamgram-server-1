/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"bytes"
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/teamgram/marmota/pkg/bytes2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/imaging"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/model"
)

// DfsUploadMp4DocumentMedia
// dfs.uploadMp4DocumentMedia creator:long media:InputMedia = Document;
func (c *DfsCore) DfsUploadMp4DocumentMedia(in *dfs.TLDfsUploadMp4DocumentMedia) (*mtproto.Document, error) {
	var (
		documentId = c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
		path       string
		err        error

		file      = in.GetMedia().GetFile()
		creatorId = in.GetCreator()
		media     = in.GetMedia()

		ext        = model.GetFileExtName(in.GetMedia().GetFile().GetName())
		extType    = model.GetStorageFileTypeConstructor(ext)
		accessHash = int64(extType)<<32 | int64(rand.Uint32())
	)

	var (
		thumbData []byte
		thumb     image.Image
		// photoId   = idgen.GetUUID()
		// ext2      = ".jpg"
		// extType2  = model.GetStorageFileTypeConstructor(ext2)
		// secretId  = int64(extType2)<<32 | int64(rand.Uint32())
	)

	// getFirstFrame
	tmpFileName := fmt.Sprintf("http://127.0.0.1:11701/dfs/file/%d_%d.mp4", creatorId, file.GetId())
	thumbData, err = c.svcCtx.FFmpegUtil.GetFirstFrame(tmpFileName)
	if err != nil {
		c.Logger.Errorf("getFirstFrameByPipe - error: %v", err)
		return nil, err
	}

	// 1. getFirstFrame
	/*
	   thumbs: [ vector<0x0>
	     { photoStrippedSize
	       type: "i" [STRING],
	       bytes: 01 28 18 CA A2 8A 2A 84 14 51 45 00 3C 63 68 F9... [91 BYTES],
	     },
	     { photoSize
	       type: "m" [STRING],
	       location: { fileLocationToBeDeprecated
	         volume_id: 500049000621 [LONG],
	         local_id: 17502 [INT],
	       },
	       w: 190 [INT],
	       h: 320 [INT],
	       size: 11452 [INT],
	     },
	   ],
	*/
	// build photoStrippedSize
	thumb, err = imaging.Decode(bytes.NewReader(thumbData))
	if err != nil {
		return nil, err

	}
	stripped := bytes2.NewBuffer(make([]byte, 0, 4096))
	if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
		err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 40, 0), 30)
	} else {
		err = imaging.EncodeStripped(stripped, imaging.Resize(thumb, 0, 40), 30)
	}
	if err != nil {
		return nil, err
	}

	// upload thumb
	var (
		mThumbData = bytes2.NewBuffer(make([]byte, 0, len(thumbData)))
		mThumb     image.Image
	)
	if thumb.Bounds().Dx() >= thumb.Bounds().Dy() {
		mThumb = imaging.Resize(thumb, 320, 0)
		// err = imaging.Encode(mThumbData, mThumb, 80)
	} else {
		mThumb = imaging.Resize(thumb, 0, 320)
		// err = imaging.Encode(mThumbData, imaging.Resize(thumb, 0, 320), 80)
	}

	err = imaging.EncodeJpeg(mThumbData, mThumb)
	if err != nil {
		return nil, err
	}

	// upload thumb
	path = fmt.Sprintf("%s/%d.dat", model.PhotoSZMediumType, documentId)
	// upload
	c.svcCtx.Dao.PutPhotoFile(c.ctx, path, mThumbData.Bytes())

	szList := []*mtproto.PhotoSize{
		mtproto.MakeTLPhotoStrippedSize(&mtproto.PhotoSize{
			Type:  model.PhotoSZStrippedType,
			Bytes: stripped.Bytes(),
		}).To_PhotoSize(),
		mtproto.MakeTLPhotoSize(&mtproto.PhotoSize{
			Type:  model.PhotoSZMediumType,
			W:     int32(mThumb.Bounds().Dx()),
			H:     int32(mThumb.Bounds().Dy()),
			Size2: int32(len(mThumbData.Bytes())),
		}).To_PhotoSize(),
	}

	// upload mp4 file
	fileInfo, err := c.svcCtx.Dao.GetFileInfo(c.ctx, creatorId, file.Id)
	if err != nil {
		c.Logger.Errorf("dfs.uploadDocumentFile - error: %v", err)
		return nil, err
	}
	c.svcCtx.Dao.SetCacheFileInfo(c.ctx, documentId, fileInfo)
	path = fmt.Sprintf("%d.dat", documentId)

	go func() {
		_, err2 := c.svcCtx.Dao.PutDocumentFile(c.ctx, path, c.svcCtx.Dao.NewSSDBReader(fileInfo))
		if err2 != nil {
			c.Logger.Errorf("dfs.PutDocumentFile - error: %v", err)
		}
	}()

	// gen attributes
	/*
	   attributes: [ vector<0x0>
	     { documentAttributeVideo
	       flags: 2 [INT],
	       round_message: [ SKIPPED BY BIT 0 IN FIELD flags ],
	       supports_streaming: YES [ BY BIT 1 IN FIELD flags ],
	       duration: 5 [INT],
	       w: 1280 [INT],
	       h: 720 [INT],
	     },
	     { documentAttributeFilename
	       file_name: "sample.mp4" [STRING],
	     },
	   ],
	*/

	attributes := make([]*mtproto.DocumentAttribute, 0, 2)
	attrVideo := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeVideo)
	if attrVideo != nil {
		attrVideo.SupportsStreaming = true
		attributes = append(attributes, attrVideo)
	}

	attrFileName := mtproto.GetDocumentAttribute(media.GetAttributes(), mtproto.Predicate_documentAttributeFilename)
	if attrFileName != nil {
		attributes = append(attributes, attrFileName)
	}

	// build document
	document := mtproto.MakeTLDocument(&mtproto.Document{
		Id:            documentId,
		AccessHash:    accessHash,
		FileReference: []byte{}, // TODO(@benqi): gen file_reference
		Date:          int32(time.Now().Unix()),
		MimeType:      "video/mp4",
		Size2:         int32(fileInfo.GetFileSize()),
		Thumbs:        szList,
		VideoThumbs:   nil,
		DcId:          1,
		Attributes:    attributes,
	}).To_Document()

	return document, nil
}
