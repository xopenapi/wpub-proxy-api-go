/*
 * Uim mp api
 *
 * uim mp api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wpub

// LinkMessageEvent 链接消息事件
type LinkMessageEvent struct {
	ToUserName   string `json:"toUserName,omitempty"`
	FromUserName string `json:"fromUserName,omitempty"`
	CreateTime   int64  `json:"createTime,omitempty"`
	MsgType      string `json:"msgType,omitempty"`
	MsgId        int64  `json:"msgId,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Url          string `json:"url,omitempty"`
}
