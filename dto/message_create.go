package dto

import "github.com/tencent-connect/botgo/dto/keyboard"

// MessageToCreate 发送消息结构体定义，同时复用给v2接口
type MessageToCreate struct {
	Content string `json:"content,omitempty"`
	Embed   *Embed `json:"embed,omitempty"`
	Ark     *Ark   `json:"ark,omitempty"`
	Image   string `json:"image,omitempty"`
	// 要回复的消息id，为空是主动消息，公域机器人会异步审核，不为空是被动消息，公域机器人会校验语料
	MsgID            string                    `json:"msg_id,omitempty"`
	MessageReference *MessageReference         `json:"message_reference,omitempty"`
	Markdown         *Markdown                 `json:"markdown,omitempty"`
	Keyboard         *keyboard.MessageKeyboard `json:"keyboard,omitempty"` // 消息按钮组件
	EventID          string                    `json:"event_id,omitempty"` // 要回复的事件id, 逻辑同MsgID

	// v2特有字段

	Type       int    `json:"msg_type,omitempty"` // 消息类型：0 是文本，2 是 markdown， 3 ark，4 embed，7 media 富媒体
	Media      *Media `json:"med,omitempty"`      // 富文本消息
	MessageSeq int    `json:"msg_seq,omitempty"` //  回复消息的序号，与 msg_id 联合使用，避免相同消息id回复重复发送，不填默认是1，相同的 msg_id + msg_seq 重复发送会失败
}

// MessageReference 引用消息
type MessageReference struct {
	MessageID             string `json:"message_id"`               // 消息 id
	IgnoreGetMessageError bool   `json:"ignore_get_message_error"` // 是否忽律获取消息失败错误
}

// Markdown markdown 消息
type Markdown struct {
	TemplateID int               `json:"template_id"` // 模版 id
	Params     []*MarkdownParams `json:"params"`      // 模版参数
	Content    string            `json:"content"`     // 原生 markdown
}

// MarkdownParams markdown 模版参数 键值对
type MarkdownParams struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// SettingGuideToCreate 发送引导消息的结构体
type SettingGuideToCreate struct {
	Content      string        `json:"content,omitempty"`       // 频道内发引导消息可以带@
	SettingGuide *SettingGuide `json:"setting_guide,omitempty"` // 设置引导
}

// SettingGuide 设置引导
type SettingGuide struct {
	// 频道ID, 当通过私信发送设置引导消息时，需要指定guild_id
	GuildID string `json:"guild_id"`
}

// Media 富文本消息
//
// https://bot.q.qq.com/wiki/develop/api-v2/server-inter/message/send-receive/rich-media.html
type Media struct {
	FileType    int    `json:"file_type"`    //媒体类型：1 图片，2 视频，3 语音，4 文件（暂不开放）资源格式要求 图片：png/jpg，视频：mp4，语音：silk
	URL         string `json:"url"`          // 需要发送媒体资源的url
	SendMessage bool   `json:"srv_send_msg"` // 设置 true 会直接发送消息到目标端，且会占用主动消息频次
}
