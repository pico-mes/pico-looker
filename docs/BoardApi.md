# \BoardApi

All URIs are relative to *https://picomes.cloud.looker.com:443/api/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllBoardItems**](BoardApi.md#AllBoardItems) | **Get** /board_items | Get All Board Items
[**AllBoardSections**](BoardApi.md#AllBoardSections) | **Get** /board_sections | Get All Board sections
[**AllBoards**](BoardApi.md#AllBoards) | **Get** /boards | Get All Boards
[**Board**](BoardApi.md#Board) | **Get** /boards/{board_id} | Get Board
[**BoardItem**](BoardApi.md#BoardItem) | **Get** /board_items/{board_item_id} | Get Board Item
[**BoardSection**](BoardApi.md#BoardSection) | **Get** /board_sections/{board_section_id} | Get Board section
[**CreateBoard**](BoardApi.md#CreateBoard) | **Post** /boards | Create Board
[**CreateBoardItem**](BoardApi.md#CreateBoardItem) | **Post** /board_items | Create Board Item
[**CreateBoardSection**](BoardApi.md#CreateBoardSection) | **Post** /board_sections | Create Board section
[**DeleteBoard**](BoardApi.md#DeleteBoard) | **Delete** /boards/{board_id} | Delete Board
[**DeleteBoardItem**](BoardApi.md#DeleteBoardItem) | **Delete** /board_items/{board_item_id} | Delete Board Item
[**DeleteBoardSection**](BoardApi.md#DeleteBoardSection) | **Delete** /board_sections/{board_section_id} | Delete Board section
[**SearchBoards**](BoardApi.md#SearchBoards) | **Get** /boards/search | Search Boards
[**UpdateBoard**](BoardApi.md#UpdateBoard) | **Patch** /boards/{board_id} | Update Board
[**UpdateBoardItem**](BoardApi.md#UpdateBoardItem) | **Patch** /board_items/{board_item_id} | Update Board Item
[**UpdateBoardSection**](BoardApi.md#UpdateBoardSection) | **Patch** /board_sections/{board_section_id} | Update Board section



## AllBoardItems

> []BoardItem AllBoardItems(ctx, optional)

Get All Board Items

### Get information about all board items. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllBoardItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllBoardItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **boardSectionId** | **optional.String**| Filter to a specific board section | 

### Return type

[**[]BoardItem**](BoardItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllBoardSections

> []BoardSection AllBoardSections(ctx, optional)

Get All Board sections

### Get information about all board sections. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllBoardSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllBoardSectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 

### Return type

[**[]BoardSection**](BoardSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllBoards

> []Board AllBoards(ctx, optional)

Get All Boards

### Get information about all boards. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllBoardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllBoardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Board**](Board.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Board

> Board Board(ctx, boardId, optional)

Get Board

### Get information about a board. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64**| Id of board | 
 **optional** | ***BoardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BoardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Board**](Board.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardItem

> BoardItem BoardItem(ctx, boardItemId, optional)

Get Board Item

### Get information about a board item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardItemId** | **int64**| Id of board item | 
 **optional** | ***BoardItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BoardItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardItem**](BoardItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardSection

> BoardSection BoardSection(ctx, boardSectionId, optional)

Get Board section

### Get information about a board section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardSectionId** | **int64**| Id of board section | 
 **optional** | ***BoardSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BoardSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardSection**](BoardSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBoard

> Board CreateBoard(ctx, body, optional)

Create Board

### Create a new board. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Board**](Board.md)| Board | 
 **optional** | ***CreateBoardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBoardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Board**](Board.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBoardItem

> BoardItem CreateBoardItem(ctx, body, optional)

Create Board Item

### Create a new board item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BoardItem**](BoardItem.md)| Board Item | 
 **optional** | ***CreateBoardItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBoardItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardItem**](BoardItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBoardSection

> BoardSection CreateBoardSection(ctx, body, optional)

Create Board section

### Create a new board section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BoardSection**](BoardSection.md)| Board section | 
 **optional** | ***CreateBoardSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBoardSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardSection**](BoardSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBoard

> string DeleteBoard(ctx, boardId)

Delete Board

### Delete a board. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64**| Id of board | 

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


## DeleteBoardItem

> string DeleteBoardItem(ctx, boardItemId)

Delete Board Item

### Delete a board item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardItemId** | **int64**| Id of board_item | 

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


## DeleteBoardSection

> string DeleteBoardSection(ctx, boardSectionId)

Delete Board section

### Delete a board section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardSectionId** | **int64**| Id of board section | 

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


## SearchBoards

> []Board SearchBoards(ctx, optional)

Search Boards

### Search Boards  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchBoardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchBoardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**| Matches board title. | 
 **createdAt** | **optional.String**| Matches the timestamp for when the board was created. | 
 **firstName** | **optional.String**| The first name of the user who created this board. | 
 **lastName** | **optional.String**| The last name of the user who created this board. | 
 **fields** | **optional.String**| Requested fields. | 
 **favorited** | **optional.Bool**| Return favorited boards when true. | 
 **creatorId** | **optional.String**| Filter on boards created by a particular user. | 
 **sorts** | **optional.String**| The fields to sort the results by | 
 **page** | **optional.Int64**| The page to return. | 
 **perPage** | **optional.Int64**| The number of items in the returned page. | 
 **offset** | **optional.Int64**| The number of items to skip before returning any. (used with limit and takes priority over page and per_page) | 
 **limit** | **optional.Int64**| The maximum number of items to return. (used with offset and takes priority over page and per_page) | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Board**](Board.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBoard

> Board UpdateBoard(ctx, boardId, body, optional)

Update Board

### Update a board definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64**| Id of board | 
**body** | [**Board**](Board.md)| Board | 
 **optional** | ***UpdateBoardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBoardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Board**](Board.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBoardItem

> BoardItem UpdateBoardItem(ctx, boardItemId, body, optional)

Update Board Item

### Update a board item definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardItemId** | **int64**| Id of board item | 
**body** | [**BoardItem**](BoardItem.md)| Board Item | 
 **optional** | ***UpdateBoardItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBoardItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardItem**](BoardItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBoardSection

> BoardSection UpdateBoardSection(ctx, boardSectionId, body, optional)

Update Board section

### Update a board section definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardSectionId** | **int64**| Id of board section | 
**body** | [**BoardSection**](BoardSection.md)| Board section | 
 **optional** | ***UpdateBoardSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBoardSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**BoardSection**](BoardSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

