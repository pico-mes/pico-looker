# SshTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelId** | **string** | Unique ID for the tunnel | [optional] [readonly] 
**SshServerId** | **string** | SSH Server ID | [optional] 
**SshServerName** | **string** | SSH Server name | [optional] [readonly] 
**SshServerHost** | **string** | SSH Server Hostname or IP Address | [optional] [readonly] 
**SshServerPort** | **int64** | SSH Server port | [optional] [readonly] 
**SshServerUser** | **string** | Username used to connect to the SSH Server | [optional] [readonly] 
**LastAttempt** | **string** | Time of last connect attempt | [optional] [readonly] 
**LocalHostPort** | **int64** | Localhost Port used by the Looker instance to connect to the remote DB | [optional] [readonly] 
**DatabaseHost** | **string** | Hostname or IP Address of the Database Server | [optional] 
**DatabasePort** | **int64** | Port that the Database Server is listening on | [optional] 
**Status** | **string** | Current connection status for this Tunnel | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


