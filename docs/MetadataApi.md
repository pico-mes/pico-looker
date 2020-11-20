# \MetadataApi

All URIs are relative to *https://picomes.cloud.looker.com:443/api/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionColumns**](MetadataApi.md#ConnectionColumns) | **Get** /connections/{connection_name}/columns | Get columns for a connection
[**ConnectionCostEstimate**](MetadataApi.md#ConnectionCostEstimate) | **Post** /connections/{connection_name}/cost_estimate | Estimate costs for a connection
[**ConnectionDatabases**](MetadataApi.md#ConnectionDatabases) | **Get** /connections/{connection_name}/databases | List accessible databases to this connection
[**ConnectionFeatures**](MetadataApi.md#ConnectionFeatures) | **Get** /connections/{connection_name}/features | Metadata features supported by this connection
[**ConnectionSchemas**](MetadataApi.md#ConnectionSchemas) | **Get** /connections/{connection_name}/schemas | Get schemas for a connection
[**ConnectionSearchColumns**](MetadataApi.md#ConnectionSearchColumns) | **Get** /connections/{connection_name}/search_columns | Search a connection for columns
[**ConnectionTables**](MetadataApi.md#ConnectionTables) | **Get** /connections/{connection_name}/tables | Get tables for a connection
[**ModelFieldnameSuggestions**](MetadataApi.md#ModelFieldnameSuggestions) | **Get** /models/{model_name}/views/{view_name}/fields/{field_name}/suggestions | Model field name suggestions



## ConnectionColumns

> []SchemaColumns ConnectionColumns(ctx, connectionName, optional)

Get columns for a connection

### Get the columns (and therefore also the tables) in a specific schema  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **optional.String**| For dialects that support multiple databases, optionally identify which to use | 
 **schemaName** | **optional.String**| Name of schema to use. | 
 **cache** | **optional.Bool**| True to fetch from cache, false to load fresh | 
 **tableLimit** | **optional.Int64**| limits the tables per schema returned | 
 **tableNames** | **optional.String**| only fetch columns for a given (comma-separated) list of tables | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]SchemaColumns**](SchemaColumns.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionCostEstimate

> CostEstimate ConnectionCostEstimate(ctx, connectionName, body, optional)

Estimate costs for a connection

### Connection cost estimating  Assign a `sql` statement to the body of the request. e.g., for Ruby, `{sql: 'select * from users'}`  **Note**: If the connection's dialect has no support for cost estimates, an error will be returned 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
**body** | [**CreateCostEstimate**](CreateCostEstimate.md)| SQL statement to estimate | 
 **optional** | ***ConnectionCostEstimateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionCostEstimateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CostEstimate**](CostEstimate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionDatabases

> []string ConnectionDatabases(ctx, connectionName)

List accessible databases to this connection

### List databases available to this connection  Certain dialects can support multiple databases per single connection. If this connection supports multiple databases, the database names will be returned in an array.  Connections using dialects that do not support multiple databases will return an empty array.  **Note**: [Connection Features](#!/Metadata/connection_features) can be used to determine if a connection supports multiple databases. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionFeatures

> ConnectionFeatures ConnectionFeatures(ctx, connectionName, optional)

Metadata features supported by this connection

### Retrieve metadata features for this connection  Returns a list of feature names with `true` (available) or `false` (not available)  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionFeaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionFeaturesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ConnectionFeatures**](ConnectionFeatures.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionSchemas

> []Schema ConnectionSchemas(ctx, connectionName, optional)

Get schemas for a connection

### Get the list of schemas and tables for a connection  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionSchemasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionSchemasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **optional.String**| For dialects that support multiple databases, optionally identify which to use | 
 **cache** | **optional.Bool**| True to use fetch from cache, false to load fresh | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionSearchColumns

> []ColumnSearch ConnectionSearchColumns(ctx, connectionName, optional)

Search a connection for columns

### Search a connection for columns matching the specified name  **Note**: `column_name` must be a valid column name. It is not a search pattern. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionSearchColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionSearchColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnName** | **optional.String**| Column name to find | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ColumnSearch**](ColumnSearch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionTables

> []SchemaTables ConnectionTables(ctx, connectionName, optional)

Get tables for a connection

### Get the list of tables for a schema  For dialects that support multiple databases, optionally identify which to use. If not provided, the default database for the connection will be used.  For dialects that do **not** support multiple databases, **do not use** the database parameter 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionTablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **optional.String**| Optional. Name of database to use for the query, only if applicable | 
 **schemaName** | **optional.String**| Optional. Return only tables for this schema | 
 **cache** | **optional.Bool**| True to fetch from cache, false to load fresh | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]SchemaTables**](SchemaTables.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelFieldnameSuggestions

> ModelFieldSuggestions ModelFieldnameSuggestions(ctx, modelName, viewName, fieldName, optional)

Model field name suggestions

### Field name suggestions for a model and view  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelName** | **string**| Name of model | 
**viewName** | **string**| Name of view | 
**fieldName** | **string**| Name of field to use for suggestions | 
 **optional** | ***ModelFieldnameSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModelFieldnameSuggestionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **term** | **optional.String**| Search term | 
 **filters** | **optional.String**| Suggestion filters | 

### Return type

[**ModelFieldSuggestions**](ModelFieldSuggestions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

