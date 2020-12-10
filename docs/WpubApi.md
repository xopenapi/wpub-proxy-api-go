# \WpubApi

All URIs are relative to *http://localhost:58200/app/wechatmp/proxy/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUser**](WpubApi.md#ListUser) | **Post** /listUser | 获取所有关注用户
[**SendMessage**](WpubApi.md#SendMessage) | **Post** /sendMessage | 发送客服消息
[**UpdateUserRemark**](WpubApi.md#UpdateUserRemark) | **Post** /updateUserRemark | 更新用户备注
[**UpdateUserTag**](WpubApi.md#UpdateUserTag) | **Post** /updateUserTag | 更新用户标签
[**UserDetail**](WpubApi.md#UserDetail) | **Post** /userDetail | 获取单个关注用户信息



## ListUser

> InlineResponse200 ListUser(ctx, body)

获取所有关注用户

获取所有关注用户

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ListUserReq**](ListUserReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> ApiResponse SendMessage(ctx, body)

发送客服消息

发送客服消息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SendMessageReq**](SendMessageReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserRemark

> ApiResponse UpdateUserRemark(ctx, body)

更新用户备注

更新用户备注

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateUserRemarkReq**](UpdateUserRemarkReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserTag

> ApiResponse UpdateUserTag(ctx, body)

更新用户标签

更新用户标签

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateUserTagReq**](UpdateUserTagReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDetail

> InlineResponse2001 UserDetail(ctx, body)

获取单个关注用户信息

获取单个关注用户信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserDetailReq**](UserDetailReq.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

