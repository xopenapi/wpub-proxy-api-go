# SendMessageReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromUser** | **string** | 发送人(运营号Id) | [optional] 
**ToUser** | **string** | 接收用户openid | [optional] 
**MsgType** | **string** | 消息类型:text 文本,image 图片,video 视频,voice 音频,music 音乐,articles 新闻 | [optional] 
**Text** | [**TextMessage**](TextMessage.md) |  | [optional] 
**Image** | [**ImageMessage**](ImageMessage.md) |  | [optional] 
**Voice** | [**VoiceMessage**](VoiceMessage.md) |  | [optional] 
**Video** | [**VideoMessage**](VideoMessage.md) |  | [optional] 
**Music** | [**MusicMessage**](MusicMessage.md) |  | [optional] 
**Articles** | [**ArticlesMessage**](ArticlesMessage.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


