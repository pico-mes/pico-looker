# \CommandApi

All URIs are relative to *https://picomes.cloud.looker.com:443/api/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](CommandApi.md#CreateCommand) | **Post** /commands | Create a custom command
[**DeleteCommand**](CommandApi.md#DeleteCommand) | **Delete** /commands/{command_id} | Delete a custom command
[**GetAllCommands**](CommandApi.md#GetAllCommands) | **Get** /commands | Get All Commands
[**UpdateCommand**](CommandApi.md#UpdateCommand) | **Patch** /commands/{command_id} | Update a custom command



## CreateCommand

> Command CreateCommand(ctx, body)

Create a custom command

### Create a new command. # Required fields: [:name, :linked_content_id, :linked_content_type] # `linked_content_type` must be one of [\"dashboard\", \"lookml_dashboard\"] # 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Command**](Command.md)| Writable command parameters | 

### Return type

[**Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommand

> DeleteCommand(ctx, commandId)

Delete a custom command

### Delete an existing custom command. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **int64**| ID of a command | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCommands

> []Command GetAllCommands(ctx, optional)

Get All Commands

### Get All Commands. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllCommandsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllCommandsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentId** | **optional.String**| Id of the associated content. This must be accompanied with content_type. | 
 **contentType** | **optional.String**| Type of the associated content. This must be accompanied with content_id. | 
 **limit** | **optional.Int64**| Number of results to return. | 

### Return type

[**[]Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommand

> Command UpdateCommand(ctx, commandId, body)

Update a custom command

### Update an existing custom command. # Optional fields: ['name', 'description'] # 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **int64**| ID of a command | 
**body** | [**UpdateCommand**](UpdateCommand.md)| Re-writable command parameters | 

### Return type

[**Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

