# \HomepageApi

All URIs are relative to *https://picomes.cloud.looker.com:443/api/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllPrimaryHomepageSections**](HomepageApi.md#AllPrimaryHomepageSections) | **Get** /primary_homepage_sections | Get All Primary homepage sections



## AllPrimaryHomepageSections

> []HomepageSection AllPrimaryHomepageSections(ctx, optional)

Get All Primary homepage sections

### Get information about the primary homepage's sections. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllPrimaryHomepageSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllPrimaryHomepageSectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]HomepageSection**](HomepageSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

