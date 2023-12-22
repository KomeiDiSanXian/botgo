package v1

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
)

// PostGroupMessage 发动消息到群
func (o *openAPI) PostGroupMessage(ctx context.Context, groupOpenID string, msg *dto.MessageToCreate) (*dto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("group_openid", groupOpenID).
		SetBody(msg).
		Post(o.getURL(groupsMessageURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.Message), nil
}

// PostRichMediaToGroup 发送富文本消息到群
func (o *openAPI) PostRichMediaToGroup(ctx context.Context, groupOpenID string, media *dto.Media) (*dto.MediaReturnParam, error) {
	resp, err := o.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("group_openid", groupOpenID).
		SetBody(media).
		Post(o.getURL(groupsFileURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.MediaReturnParam), nil
}
