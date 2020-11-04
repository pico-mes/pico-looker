# WhitelabelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**LogoFile** | **string** | Customer logo image. Expected base64 encoded data (write-only) | [optional] 
**LogoUrl** | **string** | Logo image url (read-only) | [optional] [readonly] 
**FaviconFile** | **string** | Custom favicon image. Expected base64 encoded data (write-only) | [optional] 
**FaviconUrl** | **string** | Favicon image url (read-only) | [optional] [readonly] 
**DefaultTitle** | **string** | Default page title | [optional] 
**ShowHelpMenu** | **bool** | Boolean to toggle showing help menus | [optional] 
**ShowDocs** | **bool** | Boolean to toggle showing docs | [optional] 
**ShowEmailSubOptions** | **bool** | Boolean to toggle showing email subscription options. | [optional] 
**AllowLookerMentions** | **bool** | Boolean to toggle mentions of Looker in emails | [optional] 
**AllowLookerLinks** | **bool** | Boolean to toggle links to Looker in emails | [optional] 
**CustomWelcomeEmailAdvanced** | **bool** | Allow subject line and email heading customization in customized emails” | [optional] 
**SetupMentions** | **bool** | Remove the word Looker from appearing in the account setup page | [optional] 
**AlertsLogo** | **bool** | Remove Looker logo from Alerts | [optional] 
**AlertsLinks** | **bool** | Remove Looker links from Alerts | [optional] 
**FoldersMentions** | **bool** | Remove Looker mentions in home folder page when you don’t have any items saved | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


