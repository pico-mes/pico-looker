# OauthClientApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ClientGuid** | **string** | The globally unique id of this application | [optional] [readonly] 
**RedirectUri** | **string** | The uri with which this application will receive an auth code by browser redirect. | [optional] 
**DisplayName** | **string** | The application&#39;s display name | [optional] 
**Description** | **string** | A description of the application that will be displayed to users | [optional] 
**Enabled** | **bool** | When enabled is true, OAuth2 and API requests will be accepted from this app. When false, all requests from this app will be refused. | [optional] 
**GroupId** | **int64** | If set, only Looker users who are members of this group can use this web app with Looker. If group_id is not set, any Looker user may use this app to access this Looker instance | [optional] 
**TokensInvalidBefore** | [**time.Time**](time.Time.md) | All auth codes, access tokens, and refresh tokens issued for this application prior to this date-time for ALL USERS will be invalid. | [optional] [readonly] 
**ActivatedUsers** | [**[]UserPublic**](UserPublic.md) | All users who have been activated to use this app | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


