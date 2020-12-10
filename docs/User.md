# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribe** | **int32** | 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。 | [optional] 
**Openid** | **string** | 用户的标识，对当前公众号唯一 | [optional] 
**Nickname** | **string** | 用户的昵称 | [optional] 
**Sex** | **int32** | 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知 | [optional] 
**Language** | **string** | 用户的语言，简体中文为zh_CN | [optional] 
**City** | **string** | 用户所在城市 | [optional] 
**Province** | **string** | 用户所在省份 | [optional] 
**Country** | **string** | 用户所在国家 | [optional] 
**Headimgurl** | **string** | 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效 | [optional] 
**SubscribeTime** | **int64** | 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间 | [optional] 
**Unionid** | **string** | 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。 | [optional] 
**Remark** | **string** | 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注 | [optional] 
**Groupid** | **int64** | 用户所在的分组ID（兼容旧的用户分组接口） | [optional] 
**TagIdList** | **[]int64** |  | [optional] 
**TagList** | [**[]UserTag**](UserTag.md) | 用户被打上的标签ID列表 | [optional] 
**SubscribeScene** | **string** | ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_OTHERS 其他 | [optional] 
**QrScene** | **int64** | 二维码扫码场景（开发者自定义） | [optional] 
**QrSceneStr** | **string** | 二维码扫码场景描述（开发者自定义） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


