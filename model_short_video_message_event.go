/*
 * Uim mp api
 *
 * uim mp api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wpub

// ShortVideoMessageEvent 小视频消息事件
type ShortVideoMessageEvent struct {
	ToUserName    string `json:"toUserName,omitempty"`
	FromUserName  string `json:"fromUserName,omitempty"`
	CreateTime    int64  `json:"createTime,omitempty"`
	MsgType       string `json:"msgType,omitempty"`
	MsgId         int64  `json:"msgId,omitempty"`
	MediaId       string `json:"mediaId,omitempty"`
	MediaURL      string `json:"mediaURL,omitempty"`
	ThumbMediaURL string `json:"thumbMediaURL,omitempty"`
	ThumbMediaId  string `json:"thumbMediaId,omitempty"`
}
