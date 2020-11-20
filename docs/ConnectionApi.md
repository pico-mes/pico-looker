# \ConnectionApi

All URIs are relative to *https://picomes.cloud.looker.com:443/api/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllConnections**](ConnectionApi.md#AllConnections) | **Get** /connections | Get All Connections
[**AllDialectInfos**](ConnectionApi.md#AllDialectInfos) | **Get** /dialect_info | Get All Dialect Infos
[**AllSshServers**](ConnectionApi.md#AllSshServers) | **Get** /ssh_servers | Get All SSH Servers
[**AllSshTunnels**](ConnectionApi.md#AllSshTunnels) | **Get** /ssh_tunnels | Get All SSH Tunnels
[**Connection**](ConnectionApi.md#Connection) | **Get** /connections/{connection_name} | Get Connection
[**CreateConnection**](ConnectionApi.md#CreateConnection) | **Post** /connections | Create Connection
[**CreateSshServer**](ConnectionApi.md#CreateSshServer) | **Post** /ssh_servers | Create SSH Server
[**CreateSshTunnel**](ConnectionApi.md#CreateSshTunnel) | **Post** /ssh_tunnels | Create SSH Tunnel
[**DeleteConnection**](ConnectionApi.md#DeleteConnection) | **Delete** /connections/{connection_name} | Delete Connection
[**DeleteConnectionOverride**](ConnectionApi.md#DeleteConnectionOverride) | **Delete** /connections/{connection_name}/connection_override/{override_context} | Delete Connection Override
[**DeleteSshServer**](ConnectionApi.md#DeleteSshServer) | **Delete** /ssh_server/{ssh_server_id} | Delete SSH Server
[**DeleteSshTunnel**](ConnectionApi.md#DeleteSshTunnel) | **Delete** /ssh_tunnel/{ssh_tunnel_id} | Delete SSH Tunnel
[**SshPublicKey**](ConnectionApi.md#SshPublicKey) | **Get** /ssh_public_key | Get SSH Public Key
[**SshServer**](ConnectionApi.md#SshServer) | **Get** /ssh_server/{ssh_server_id} | Get SSH Server
[**SshTunnel**](ConnectionApi.md#SshTunnel) | **Get** /ssh_tunnel/{ssh_tunnel_id} | Get SSH Tunnel
[**TestConnection**](ConnectionApi.md#TestConnection) | **Put** /connections/{connection_name}/test | Test Connection
[**TestConnectionConfig**](ConnectionApi.md#TestConnectionConfig) | **Put** /connections/test | Test Connection Configuration
[**TestSshServer**](ConnectionApi.md#TestSshServer) | **Get** /ssh_server/{ssh_server_id}/test | Test SSH Server
[**TestSshTunnel**](ConnectionApi.md#TestSshTunnel) | **Get** /ssh_tunnel/{ssh_tunnel_id}/test | Test SSH Tunnel
[**UpdateConnection**](ConnectionApi.md#UpdateConnection) | **Patch** /connections/{connection_name} | Update Connection
[**UpdateSshServer**](ConnectionApi.md#UpdateSshServer) | **Patch** /ssh_server/{ssh_server_id} | Update SSH Server
[**UpdateSshTunnel**](ConnectionApi.md#UpdateSshTunnel) | **Patch** /ssh_tunnel/{ssh_tunnel_id} | Update SSH Tunnel



## AllConnections

> []DbConnection AllConnections(ctx, optional)

Get All Connections

### Get information about all connections. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllConnectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllDialectInfos

> []DialectInfo AllDialectInfos(ctx, optional)

Get All Dialect Infos

### Get information about all dialects. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllDialectInfosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllDialectInfosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DialectInfo**](DialectInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllSshServers

> []SshServer AllSshServers(ctx, optional)

Get All SSH Servers

### Get information about all SSH Servers. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllSshServersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllSshServersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]SshServer**](SshServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllSshTunnels

> []SshTunnel AllSshTunnels(ctx, optional)

Get All SSH Tunnels

### Get information about all SSH Tunnels. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllSshTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllSshTunnelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]SshTunnel**](SshTunnel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Connection

> DbConnection Connection(ctx, connectionName, optional)

Get Connection

### Get information about a connection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnection

> DbConnection CreateConnection(ctx, body)

Create Connection

### Create a connection using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DbConnection**](DbConnection.md)| Connection | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSshServer

> SshServer CreateSshServer(ctx, body)

Create SSH Server

### Create an SSH Server. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SshServer**](SshServer.md)| SSH Server | 

### Return type

[**SshServer**](SshServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSshTunnel

> SshTunnel CreateSshTunnel(ctx, body)

Create SSH Tunnel

### Create an SSH Tunnel 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SshTunnel**](SshTunnel.md)| SSH Tunnel | 

### Return type

[**SshTunnel**](SshTunnel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> string DeleteConnection(ctx, connectionName)

Delete Connection

### Delete a connection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionOverride

> string DeleteConnectionOverride(ctx, connectionName, overrideContext)

Delete Connection Override

### Delete a connection override. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
**overrideContext** | **string**| Context of connection override | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshServer

> string DeleteSshServer(ctx, sshServerId)

Delete SSH Server

### Delete an SSH Server. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshServerId** | **string**| Id of SSH Server | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshTunnel

> string DeleteSshTunnel(ctx, sshTunnelId)

Delete SSH Tunnel

### Delete an SSH Tunnel 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshTunnelId** | **string**| Id of SSH Tunnel | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshPublicKey

> SshPublicKey SshPublicKey(ctx, )

Get SSH Public Key

### Get the SSH public key  Get the public key created for this instance to identify itself to a remote SSH server. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SshPublicKey**](SshPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshServer

> SshServer SshServer(ctx, sshServerId)

Get SSH Server

### Get information about an SSH Server. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshServerId** | **string**| Id of SSH Server | 

### Return type

[**SshServer**](SshServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshTunnel

> SshTunnel SshTunnel(ctx, sshTunnelId)

Get SSH Tunnel

### Get information about an SSH Tunnel. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshTunnelId** | **string**| Id of SSH Tunnel | 

### Return type

[**SshTunnel**](SshTunnel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnection

> []DbConnectionTestResult TestConnection(ctx, connectionName, optional)

Test Connection

### Test an existing connection.  Note that a connection's 'dialect' property has a 'connection_tests' property that lists the specific types of tests that the connection supports.  This API is rate limited.  Unsupported tests in the request will be ignored. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***TestConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tests** | [**optional.Interface of []string**](string.md)| Array of names of tests to run | 

### Return type

[**[]DbConnectionTestResult**](DBConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionConfig

> []DbConnectionTestResult TestConnectionConfig(ctx, body, optional)

Test Connection Configuration

### Test a connection configuration.  Note that a connection's 'dialect' property has a 'connection_tests' property that lists the specific types of tests that the connection supports.  This API is rate limited.  Unsupported tests in the request will be ignored. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DbConnection**](DbConnection.md)| Connection | 
 **optional** | ***TestConnectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tests** | [**optional.Interface of []string**](string.md)| Array of names of tests to run | 

### Return type

[**[]DbConnectionTestResult**](DBConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSshServer

> SshServer TestSshServer(ctx, sshServerId)

Test SSH Server

### Test the SSH Server 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshServerId** | **string**| Id of SSH Server | 

### Return type

[**SshServer**](SshServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSshTunnel

> SshTunnel TestSshTunnel(ctx, sshTunnelId)

Test SSH Tunnel

### Test the SSH Tunnel 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshTunnelId** | **string**| Id of SSH Tunnel | 

### Return type

[**SshTunnel**](SshTunnel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> DbConnection UpdateConnection(ctx, connectionName, body)

Update Connection

### Update a connection using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
**body** | [**DbConnection**](DbConnection.md)| Connection | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshServer

> SshServer UpdateSshServer(ctx, sshServerId, body)

Update SSH Server

### Update an SSH Server. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshServerId** | **string**| Id of SSH Server | 
**body** | [**SshServer**](SshServer.md)| SSH Server | 

### Return type

[**SshServer**](SshServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshTunnel

> SshTunnel UpdateSshTunnel(ctx, sshTunnelId, body)

Update SSH Tunnel

### Update an SSH Tunnel 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshTunnelId** | **string**| Id of SSH Tunnel | 
**body** | [**SshTunnel**](SshTunnel.md)| SSH Tunnel | 

### Return type

[**SshTunnel**](SshTunnel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

