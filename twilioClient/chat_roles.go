package twilioClient

import (
	"golang.org/x/net/context"
	"net/url"
	"strings"
)

/**
 * Doc: https://www.twilio.com/docs/chat/rest/roles
 */
const chatRoleApiPath string = "Roles"

type ChatRole struct {
	Sid          string           `json:"sid"`
	AccountSid   string           `json:"account_sid"`
	ServiceSid   string           `json:"service_sid"`
	DateCreated  string           `json:"date_created"`
	DateUpdated  string           `json:"date_updated"`
	FriendlyName string           `json:"friendly_name"`
	Type         ChatRoleType     `json:"type"`
	Permissions  []ChatPermission `json:"permissions"`
	Url          string           `json:"url"`
}

type ChatRoleType string

const Channel = ChatRoleType("channel")
const Deployment = ChatRoleType("deployment")

func (s ChatRoleType) ToString() string {
	switch s {
	case Channel:
		return "channel"
	case Deployment:
		return "deployment"
	default:
		return strings.Title(string(s))
	}
}

type ChatPermission string

// deployment only
const CreateChannel = ChatPermission("createChannel")
const JoinChannel = ChatPermission("joinChannel")

// channel only
const SendMessage = ChatPermission("sendMessage")
const SendMediaMessage = ChatPermission("sendMediaMessage")
const LeaveChannel = ChatPermission("leaveChannel")
const DeleteOwnMessage = ChatPermission("deleteOwnMessage")

// common
const DestroyChannel = ChatPermission("destroyChannel")
const InviteMember = ChatPermission("inviteMember")
const RemoveMember = ChatPermission("removeMember")
const EditChannelName = ChatPermission("editChannelName")
const EditChannelAttributes = ChatPermission("editChannelAttributes")
const AddMember = ChatPermission("addMember")
const EditOwnMessage = ChatPermission("editOwnMessage")
const EditAnyMessage = ChatPermission("editAnyMessage")
const EditOwnMessageAttributes = ChatPermission("editOwnMessageAttributes")
const EditAnyMessageAttributes = ChatPermission("editAnyMessageAttributes")
const DeleteAnyMessage = ChatPermission("deleteAnyMessage")
const EditOwnUserInfo = ChatPermission("editOwnUserInfo")
const EditAnyUserInfo = ChatPermission("editAnyUserInfo")

func (s ChatPermission) ToString() string {
	switch s {
	case CreateChannel:
		return "createChannel"
	case JoinChannel:
		return "joinChannel"
	case SendMessage:
		return "sendMessage"
	case SendMediaMessage:
		return "sendMediaMessage"
	case LeaveChannel:
		return "leaveChannel"
	case DeleteOwnMessage:
		return "deleteOwnMessage"
	case DestroyChannel:
		return "destroyChannel"
	case InviteMember:
		return "inviteMember"
	case RemoveMember:
		return "removeMember"
	case EditChannelName:
		return "editChannelName"
	case EditChannelAttributes:
		return "editChannelAttributes"
	case AddMember:
		return "addMember"
	case EditOwnMessage:
		return "editOwnMessage"
	case EditAnyMessage:
		return "editAnyMessage"
	case EditOwnMessageAttributes:
		return "editOwnMessageAttributes"
	case EditAnyMessageAttributes:
		return "editAnyMessageAttributes"
	case DeleteAnyMessage:
		return "deleteAnyMessage"
	case EditOwnUserInfo:
		return "editOwnUserInfo"
	case EditAnyUserInfo:
		return "editAnyUserInfo"
	default:
		return strings.Title(string(s))
	}
}

type ChatRolePage struct {
	Page
	ChatRoles []*ChatRole `json:"roles"`
}

type ChatRoleService struct {
	client *TwilioClient
}

func (service *ChatRoleService) Create(ctx context.Context, data url.Values) (*ChatRole, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatRoleApiPath
	app := new(ChatRole)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ChatRoleService) Update(ctx context.Context, appSid string, data url.Values) (*ChatRole, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatRoleApiPath + "/" + appSid

	app := new(ChatRole)
	err := service.client.post(targetUrl, data, app)
	return app, err
}

func (service *ChatRoleService) Delete(ctx context.Context, appSid string) error {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatRoleApiPath + "/" + appSid

	return service.client.delete(targetUrl)
}

func (service *ChatRoleService) Get(ctx context.Context, appSid string) (*ChatRole, error) {
	var targetUrl = service.client.ChatBaseUrl() + "/" + chatRoleApiPath + "/" + appSid

	app := new(ChatRole)
	err := service.client.get(targetUrl, url.Values{}, app)

	return app, err
}

func (service *ChatRoleService) GetPage(ctx context.Context, data url.Values) (*ChatRolePage, error) {
	iter := service.GetPageIterator(data)
	return iter.Next(ctx)
}

type ChatRolePageIterator struct {
	p *PageIterator
}

func (service *ChatRoleService) GetPageIterator(data url.Values) *ChatRolePageIterator {
	iter := NewPageIterator(service.client, data, service.client.ChatBaseUrl()+"/"+chatRoleApiPath)
	return &ChatRolePageIterator{
		p: iter,
	}
}

func (c *ChatRolePageIterator) Next(ctx context.Context) (*ChatRolePage, error) {
	cp := new(ChatRolePage)
	err := c.p.Next(ctx, cp)
	if err != nil {
		return nil, err
	}
	c.p.SetNextPageURI(cp.NextPageURI)
	return cp, nil
}
