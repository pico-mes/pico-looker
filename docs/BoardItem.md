# BoardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentCreatedBy** | **string** | Name of user who created the content this item is based on | [optional] [readonly] 
**ContentFavoriteId** | **int64** | Content favorite id associated with the item this content is based on | [optional] [readonly] 
**ContentMetadataId** | **int64** | Content metadata id associated with the item this content is based on | [optional] [readonly] 
**ContentUpdatedAt** | **string** | Last time the content that this item is based on was updated | [optional] [readonly] 
**DashboardId** | **int64** | Dashboard to base this item on | [optional] 
**Description** | **string** | The actual description for display | [optional] [readonly] 
**FavoriteCount** | **int64** | Number of times content has been favorited, if present | [optional] [readonly] 
**BoardSectionId** | **int64** | Associated Board Section | [optional] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Location** | **string** | The container folder name of the content | [optional] [readonly] 
**LookId** | **int64** | Look to base this item on | [optional] 
**LookmlDashboardId** | **string** | LookML Dashboard to base this item on | [optional] 
**Order** | **int64** | An arbitrary integer representing the sort order within the section | [optional] 
**Title** | **string** | The actual title for display | [optional] [readonly] 
**Url** | **string** | Relative url for the associated content | [optional] [readonly] 
**ViewCount** | **int64** | Number of times content has been viewed, if present | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


