package v1

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
)

// PostUserMessage 单独发动消息给用户 (c2c)
func (o *openAPI) PostUserMessage(ctx context.Context, openID string, msg *dto.MessageToCreate) (*dto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("openid", openID).
		SetBody(msg).
		Post(o.getURL(usersMessageURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.Message), nil
}

// PostRichMedia 发送富文本消息到用户
func (o *openAPI) PostRichMediaToUser(ctx context.Context, openID string, media *dto.Media) (*dto.MediaReturnParam, error) {
	resp, err := o.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("openid", openID).
		SetBody(media).
		Post(o.getURL(usersFileURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.MediaReturnParam), nil
}
