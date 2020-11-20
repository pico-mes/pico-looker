# Board

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentMetadataId** | **int64** | Id of associated content_metadata record | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Date of board creation | [optional] [readonly] 
**DeletedAt** | [**time.Time**](time.Time.md) | Date of board deletion | [optional] 
**Description** | **string** | Description of the board | [optional] 
**BoardSections** | [**[]BoardSection**](BoardSection.md) | Sections of the board | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**SectionOrder** | **[]int64** | ids of the board sections in the order they should be displayed | [optional] 
**Title** | **string** | Title of the board | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Date of last board update | [optional] [readonly] 
**UserId** | **int64** | User id of board creator | [optional] [readonly] 
**PrimaryHomepage** | **bool** | Whether the board is the primary homepage or not | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


