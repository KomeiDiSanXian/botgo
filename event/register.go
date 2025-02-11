package event

import (
	"github.com/tencent-connect/botgo/dto"
)

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild       GuildEventHandler
	GuildMember GuildMemberEventHandler
	Channel     ChannelEventHandler

	Message             MessageEventHandler
	MessageReaction     MessageReactionEventHandler
	ATMessage           ATMessageEventHandler
	DirectMessage       DirectMessageEventHandler
	MessageAudit        MessageAuditEventHandler
	MessageDelete       MessageDeleteEventHandler
	PublicMessageDelete PublicMessageDeleteEventHandler
	DirectMessageDelete DirectMessageDeleteEventHandler

	Audio AudioEventHandler

	Thread     ThreadEventHandler
	Post       PostEventHandler
	Reply      ReplyEventHandler
	ForumAudit ForumAuditEventHandler

	Interaction InteractionEventHandler

	UserQuery          UserQueryEventHandler
	GroupAtMessage     GroupAtMessageEventHandler
	UserAddBot         UserAddBotEventHandler
	UserDelBot         UserDelBotEventHandler
	UserRejectMessage  UserRejectMessageEventHandler
	UserReciveMessage  UserReciveMessageEventHandler
	AddGroup           AddGroupEventHandler
	QuitGroup          QuitGroupEventHandler
	GroupRejectMessage GroupRejectMessageEventHandler
	GroupReciveMessage GroupReciveMessageEventHandler
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *dto.WSPayload, data *dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *dto.WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *dto.WSPayload, data *dto.WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *dto.WSPayload, data *dto.WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *dto.WSPayload, data *dto.WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *dto.WSPayload, data *dto.WSMessageData) error

// MessageDeleteEventHandler 消息事件 handler
type MessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSMessageDeleteData) error

// PublicMessageDeleteEventHandler 消息事件 handler
type PublicMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSPublicMessageDeleteData) error

// DirectMessageDeleteEventHandler 消息事件 handler
type DirectMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageDeleteData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *dto.WSPayload, data *dto.WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *dto.WSPayload, data *dto.WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *dto.WSPayload, data *dto.WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *dto.WSPayload, data *dto.WSMessageAuditData) error

// ThreadEventHandler 论坛主题事件 handler
type ThreadEventHandler func(event *dto.WSPayload, data *dto.WSThreadData) error

// PostEventHandler 论坛回帖事件 handler
type PostEventHandler func(event *dto.WSPayload, data *dto.WSPostData) error

// ReplyEventHandler 论坛帖子回复事件 handler
type ReplyEventHandler func(event *dto.WSPayload, data *dto.WSReplyData) error

// ForumAuditEventHandler 论坛帖子审核事件 handler
type ForumAuditEventHandler func(event *dto.WSPayload, data *dto.WSForumAuditData) error

// InteractionEventHandler 互动事件 handler
type InteractionEventHandler func(event *dto.WSPayload, data *dto.WSInteractionData) error

// UserQueryEventHandler 用户单聊事件 handler
type UserQueryEventHandler func(event *dto.WSPayload, data *dto.WSUserQuery) error

// GroupAtMessageEventHandler 群聊at机器人事件 handler
type GroupAtMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupAtMessage) error

// UserAddBotEventHandler 用户添加机器人事件 handler
type UserAddBotEventHandler func(event *dto.WSPayload, data *dto.WSUserAddBot) error

// UserDelBotEventHandler 用户删除机器人事件 handler
type UserDelBotEventHandler func(event *dto.WSPayload, data *dto.WSUserDelBot) error

// UserRejectMessageEventHandler 用户拒绝主动消息事件 handler
type UserRejectMessageEventHandler func(event *dto.WSPayload, data *dto.WSUserRejectMessage) error

// UserReciveMessageEventHandler 用户接受主动消息事件 handler
type UserReciveMessageEventHandler func(event *dto.WSPayload, data *dto.WSUserReciveMessage) error

// AddGroupEventHandler 群聊添加机器人事件 handler
type AddGroupEventHandler func(event *dto.WSPayload, data *dto.WSAddGroup) error

// QuitGroupEventHandler 群聊删除机器人事件 handler
type QuitGroupEventHandler func(event *dto.WSPayload, data *dto.WSQuitGroup) error

// GroupRejectMessageEventHandler 群聊拒绝主动消息事件 handler
type GroupRejectMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupRejectMessage) error

// GroupReciveMessageEventHandler 群聊接受主动消息事件 handler
type GroupReciveMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupReciveMessage) error

// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) dto.Intent {
	var i dto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | dto.EventToIntent(
				dto.EventAudioStart, dto.EventAudioFinish,
				dto.EventAudioOnMic, dto.EventAudioOffMic,
			)
		case InteractionEventHandler:
			DefaultHandlers.Interaction = handle
			i = i | dto.EventToIntent(dto.EventInteractionCreate)
		default:
		}
	}
	i = i | registerRelationHandlers(i, handlers...)
	i = i | registerMessageHandlers(i, handlers...)
	i = i | registerForumHandlers(i, handlers...)

	return i
}

// registerForumHandlers 注册论坛关系链相关handlers
func registerForumHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case ThreadEventHandler:
			DefaultHandlers.Thread = handle
			i = i | dto.EventToIntent(
				dto.EventForumThreadCreate, dto.EventForumThreadUpdate, dto.EventForumThreadDelete,
			)
		case PostEventHandler:
			DefaultHandlers.Post = handle
			i = i | dto.EventToIntent(dto.EventForumPostCreate, dto.EventForumPostDelete)
		case ReplyEventHandler:
			DefaultHandlers.Reply = handle
			i = i | dto.EventToIntent(dto.EventForumReplyCreate, dto.EventForumReplyDelete)
		case ForumAuditEventHandler:
			DefaultHandlers.ForumAudit = handle
			i = i | dto.EventToIntent(dto.EventForumAuditResult)
		default:
		}
	}
	return i
}

// registerRelationHandlers 注册频道关系链相关handlers
func registerRelationHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case GuildEventHandler:
			DefaultHandlers.Guild = handle
			i = i | dto.EventToIntent(dto.EventGuildCreate, dto.EventGuildDelete, dto.EventGuildUpdate)
		case GuildMemberEventHandler:
			DefaultHandlers.GuildMember = handle
			i = i | dto.EventToIntent(dto.EventGuildMemberAdd, dto.EventGuildMemberRemove, dto.EventGuildMemberUpdate)
		case ChannelEventHandler:
			DefaultHandlers.Channel = handle
			i = i | dto.EventToIntent(dto.EventChannelCreate, dto.EventChannelDelete, dto.EventChannelUpdate)
		default:
		}
	}
	return i
}

// registerMessageHandlers 注册消息相关的 handler
func registerMessageHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | dto.EventToIntent(dto.EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | dto.EventToIntent(dto.EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageCreate)
		case MessageDeleteEventHandler:
			DefaultHandlers.MessageDelete = handle
			i = i | dto.EventToIntent(dto.EventMessageDelete)
		case PublicMessageDeleteEventHandler:
			DefaultHandlers.PublicMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventPublicMessageDelete)
		case DirectMessageDeleteEventHandler:
			DefaultHandlers.DirectMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageDelete)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | dto.EventToIntent(dto.EventMessageReactionAdd, dto.EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | dto.EventToIntent(dto.EventMessageAuditPass, dto.EventMessageAuditReject)
		case UserQueryEventHandler:
			DefaultHandlers.UserQuery = handle
			i = i | dto.EventToIntent(dto.EventUserMessageCreate)
		case GroupAtMessageEventHandler:
			DefaultHandlers.GroupAtMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupMessageCreate)
		case AddGroupEventHandler:
			DefaultHandlers.AddGroup = handle
			i = i | dto.EventToIntent(dto.EventGroupAddBot)
		case QuitGroupEventHandler:
			DefaultHandlers.QuitGroup = handle
			i = i | dto.EventToIntent(dto.EventGroupDelBot)
		case GroupRejectMessageEventHandler:
			DefaultHandlers.GroupRejectMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupRejectMsg)
		case GroupReciveMessageEventHandler:
			DefaultHandlers.GroupReciveMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupReciveMsg)
		case UserAddBotEventHandler:
			DefaultHandlers.UserAddBot = handle
			i = i | dto.EventToIntent(dto.EventUserAddBot)
		case UserDelBotEventHandler:
			DefaultHandlers.UserDelBot = handle
			i = i | dto.EventToIntent(dto.EventUserDelBot)
		case UserReciveMessageEventHandler:
			DefaultHandlers.UserReciveMessage = handle
			i = i | dto.EventToIntent(dto.EventUserReciveMsg)
		case UserRejectMessageEventHandler:
			DefaultHandlers.UserRejectMessage = handle
			i = i | dto.EventToIntent(dto.EventUserRejectMsg) 
		default:
		}
	}
	return i
}
