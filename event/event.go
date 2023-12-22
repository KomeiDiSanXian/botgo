package event

import (
	"encoding/json"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tidwall/gjson" // 由于回包的 d 类型不确定，gjson 用于从回包json中提取 d 并进行针对性的解析
)

var eventParseFuncMap = map[dto.OPCode]map[dto.EventType]eventParseFunc{
	dto.WSDispatchEvent: {
		dto.EventGuildCreate: guildHandler,
		dto.EventGuildUpdate: guildHandler,
		dto.EventGuildDelete: guildHandler,

		dto.EventChannelCreate: channelHandler,
		dto.EventChannelUpdate: channelHandler,
		dto.EventChannelDelete: channelHandler,

		dto.EventGuildMemberAdd:    guildMemberHandler,
		dto.EventGuildMemberUpdate: guildMemberHandler,
		dto.EventGuildMemberRemove: guildMemberHandler,

		dto.EventMessageCreate: messageHandler,
		dto.EventMessageDelete: messageDeleteHandler,

		dto.EventMessageReactionAdd:    messageReactionHandler,
		dto.EventMessageReactionRemove: messageReactionHandler,

		dto.EventAtMessageCreate:     atMessageHandler,
		dto.EventPublicMessageDelete: publicMessageDeleteHandler,

		dto.EventDirectMessageCreate: directMessageHandler,
		dto.EventDirectMessageDelete: directMessageDeleteHandler,

		dto.EventAudioStart:  audioHandler,
		dto.EventAudioFinish: audioHandler,
		dto.EventAudioOnMic:  audioHandler,
		dto.EventAudioOffMic: audioHandler,

		dto.EventMessageAuditPass:   messageAuditHandler,
		dto.EventMessageAuditReject: messageAuditHandler,

		dto.EventForumThreadCreate: threadHandler,
		dto.EventForumThreadUpdate: threadHandler,
		dto.EventForumThreadDelete: threadHandler,
		dto.EventForumPostCreate:   postHandler,
		dto.EventForumPostDelete:   postHandler,
		dto.EventForumReplyCreate:  replyHandler,
		dto.EventForumReplyDelete:  replyHandler,
		dto.EventForumAuditResult:  forumAuditHandler,

		dto.EventInteractionCreate: interactionHandler,

		dto.EventUserMessageCreate:  chatFromUserHandler,
		dto.EventGroupMessageCreate: chatFromGroupHandler,
		dto.EventGroupAddBot:        groupAddBotHandler,
		dto.EventGroupDelBot:        groupDelBotHandler,
		dto.EventGroupRejectMsg:     groupRejectMessageHandler,
		dto.EventGroupReciveMsg:     groupReciveMessageHandler,
		dto.EventUserAddBot:         userAddBotHandler,
		dto.EventUserDelBot:         userDelBotHandler,
		dto.EventUserReciveMsg:      userReciveMessageHandler,
		dto.EventUserRejectMsg:      userRejectMessageHandler,
	},
}

type eventParseFunc func(event *dto.WSPayload, message []byte) error

// ParseAndHandle 处理回调事件
func ParseAndHandle(payload *dto.WSPayload) error {
	// 指定类型的 handler
	if h, ok := eventParseFuncMap[payload.OPCode][payload.Type]; ok {
		return h(payload, payload.RawMessage)
	}
	// 透传handler，如果未注册具体类型的 handler，会统一投递到这个 handler
	if DefaultHandlers.Plain != nil {
		return DefaultHandlers.Plain(payload, payload.RawMessage)
	}
	return nil
}

// ParseData 解析数据
func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}

func guildHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGuildData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Guild != nil {
		return DefaultHandlers.Guild(payload, data)
	}
	return nil
}

func channelHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSChannelData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Channel != nil {
		return DefaultHandlers.Channel(payload, data)
	}
	return nil
}

func guildMemberHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGuildMemberData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GuildMember != nil {
		return DefaultHandlers.GuildMember(payload, data)
	}
	return nil
}

func messageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Message != nil {
		return DefaultHandlers.Message(payload, data)
	}
	return nil
}

func messageDeleteHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageDelete != nil {
		return DefaultHandlers.MessageDelete(payload, data)
	}
	return nil
}

func messageReactionHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSMessageReactionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageReaction != nil {
		return DefaultHandlers.MessageReaction(payload, data)
	}
	return nil
}

func atMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSATMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(payload, data)
	}
	return nil
}

func publicMessageDeleteHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSPublicMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.PublicMessageDelete != nil {
		return DefaultHandlers.PublicMessageDelete(payload, data)
	}
	return nil
}

func directMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSDirectMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.DirectMessage != nil {
		return DefaultHandlers.DirectMessage(payload, data)
	}
	return nil
}

func directMessageDeleteHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSDirectMessageDeleteData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.DirectMessageDelete != nil {
		return DefaultHandlers.DirectMessageDelete(payload, data)
	}
	return nil
}

func audioHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSAudioData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Audio != nil {
		return DefaultHandlers.Audio(payload, data)
	}
	return nil
}

func threadHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSThreadData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Thread != nil {
		return DefaultHandlers.Thread(payload, data)
	}
	return nil
}

func postHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSPostData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Post != nil {
		return DefaultHandlers.Post(payload, data)
	}
	return nil
}

func replyHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSReplyData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Reply != nil {
		return DefaultHandlers.Reply(payload, data)
	}
	return nil
}

func forumAuditHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSForumAuditData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.ForumAudit != nil {
		return DefaultHandlers.ForumAudit(payload, data)
	}
	return nil
}

func messageAuditHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSMessageAuditData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageAudit != nil {
		return DefaultHandlers.MessageAudit(payload, data)
	}
	return nil
}

func interactionHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSInteractionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Interaction != nil {
		return DefaultHandlers.Interaction(payload, data)
	}
	return nil
}

func chatFromUserHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSUserQuery{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.UserQuery != nil {
		return DefaultHandlers.UserQuery(payload, data)
	}
	return nil
}

func chatFromGroupHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupAtMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupAtMessage != nil {
		return DefaultHandlers.GroupAtMessage(payload, data)
	}
	return nil
}

func userAddBotHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSUserAddBot{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.UserAddBot != nil {
		return DefaultHandlers.UserAddBot(payload, data)
	}
	return nil
}

func userDelBotHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSUserDelBot{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.UserDelBot != nil {
		return DefaultHandlers.UserDelBot(payload, data)
	}
	return nil
}

func userReciveMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSUserReciveMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.UserReciveMessage != nil {
		return DefaultHandlers.UserReciveMessage(payload, data)
	}
	return nil
}

func userRejectMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSUserRejectMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.UserRejectMessage != nil {
		return DefaultHandlers.UserRejectMessage(payload, data)
	}
	return nil
}

func groupReciveMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupReciveMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupReciveMessage != nil {
		return DefaultHandlers.GroupReciveMessage(payload, data)
	}
	return nil
}

func groupRejectMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupRejectMessage{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupRejectMessage != nil {
		return DefaultHandlers.GroupRejectMessage(payload, data)
	}
	return nil
}

func groupAddBotHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSAddGroup{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.AddGroup != nil {
		return DefaultHandlers.AddGroup(payload, data)
	}
	return nil
}

func groupDelBotHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSQuitGroup{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.QuitGroup != nil {
		return DefaultHandlers.QuitGroup(payload, data)
	}
	return nil
}
