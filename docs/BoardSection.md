# BoardSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Time at which this section was created. | [optional] [readonly] 
**DeletedAt** | [**time.Time**](time.Time.md) | Time at which this section was deleted. | [optional] 
**Description** | **string** | Description of the content found in this section. | [optional] 
**BoardId** | **int64** | Id reference to parent board | [optional] 
**BoardItems** | [**[]BoardItem**](BoardItem.md) | Items in the board section | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**ItemOrder** | **[]int64** | ids of the board items in the order they should be displayed | [optional] 
**Title** | **string** | Name of row | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Time at which this section was last updated. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


