# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/common.proto](#api_v1_common-proto)
    - [RowStatus](#share-api-v1-RowStatus)
  
- [api/v1/user_service.proto](#api_v1_user_service-proto)
    - [CreateUserAccessTokenRequest](#share-api-v1-CreateUserAccessTokenRequest)
    - [CreateUserAccessTokenResponse](#share-api-v1-CreateUserAccessTokenResponse)
    - [CreateUserRequest](#share-api-v1-CreateUserRequest)
    - [CreateUserResponse](#share-api-v1-CreateUserResponse)
    - [DeleteUserAccessTokenRequest](#share-api-v1-DeleteUserAccessTokenRequest)
    - [DeleteUserAccessTokenResponse](#share-api-v1-DeleteUserAccessTokenResponse)
    - [DeleteUserRequest](#share-api-v1-DeleteUserRequest)
    - [DeleteUserResponse](#share-api-v1-DeleteUserResponse)
    - [GetUserRequest](#share-api-v1-GetUserRequest)
    - [GetUserResponse](#share-api-v1-GetUserResponse)
    - [ListUserAccessTokensRequest](#share-api-v1-ListUserAccessTokensRequest)
    - [ListUserAccessTokensResponse](#share-api-v1-ListUserAccessTokensResponse)
    - [ListUsersRequest](#share-api-v1-ListUsersRequest)
    - [ListUsersResponse](#share-api-v1-ListUsersResponse)
    - [UpdateUserRequest](#share-api-v1-UpdateUserRequest)
    - [UpdateUserResponse](#share-api-v1-UpdateUserResponse)
    - [User](#share-api-v1-User)
    - [UserAccessToken](#share-api-v1-UserAccessToken)
  
    - [UserService](#share-api-v1-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/common.proto


 


<a name="share-api-v1-RowStatus"></a>

### RowStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROW_STATUS_UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| ARCHIVED | 2 |  |


 

 

 



<a name="api_v1_user_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/user_service.proto



<a name="share-api-v1-CreateUserAccessTokenRequest"></a>

### CreateUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |
| description | [string](#string) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |






<a name="share-api-v1-CreateUserAccessTokenResponse"></a>

### CreateUserAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [UserAccessToken](#share-api-v1-UserAccessToken) |  |  |






<a name="share-api-v1-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#share-api-v1-User) |  |  |






<a name="share-api-v1-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#share-api-v1-User) |  |  |






<a name="share-api-v1-DeleteUserAccessTokenRequest"></a>

### DeleteUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |
| access_token | [string](#string) |  | access_token is the access token to delete. |






<a name="share-api-v1-DeleteUserAccessTokenResponse"></a>

### DeleteUserAccessTokenResponse







<a name="share-api-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |






<a name="share-api-v1-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="share-api-v1-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |






<a name="share-api-v1-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#share-api-v1-User) |  |  |






<a name="share-api-v1-ListUserAccessTokensRequest"></a>

### ListUserAccessTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |






<a name="share-api-v1-ListUserAccessTokensResponse"></a>

### ListUserAccessTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_tokens | [UserAccessToken](#share-api-v1-UserAccessToken) | repeated |  |






<a name="share-api-v1-ListUsersRequest"></a>

### ListUsersRequest







<a name="share-api-v1-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#share-api-v1-User) | repeated |  |






<a name="share-api-v1-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#share-api-v1-User) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="share-api-v1-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#share-api-v1-User) |  |  |






<a name="share-api-v1-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{username} |
| id | [string](#string) |  |  |
| username | [string](#string) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |
| password | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="share-api-v1-UserAccessToken"></a>

### UserAccessToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| description | [string](#string) |  |  |
| issued_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 


<a name="share-api-v1-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListUsers | [ListUsersRequest](#share-api-v1-ListUsersRequest) | [ListUsersResponse](#share-api-v1-ListUsersResponse) | ListUsers returns a list of users. |
| GetUser | [GetUserRequest](#share-api-v1-GetUserRequest) | [GetUserResponse](#share-api-v1-GetUserResponse) | GetUser gets a user by name. |
| CreateUser | [CreateUserRequest](#share-api-v1-CreateUserRequest) | [CreateUserResponse](#share-api-v1-CreateUserResponse) | CreateUser creates a new user. |
| UpdateUser | [UpdateUserRequest](#share-api-v1-UpdateUserRequest) | [UpdateUserResponse](#share-api-v1-UpdateUserResponse) | UpdateUser updates a user. |
| DeleteUser | [DeleteUserRequest](#share-api-v1-DeleteUserRequest) | [DeleteUserResponse](#share-api-v1-DeleteUserResponse) | DeleteUser deletes a user. |
| ListUserAccessTokens | [ListUserAccessTokensRequest](#share-api-v1-ListUserAccessTokensRequest) | [ListUserAccessTokensResponse](#share-api-v1-ListUserAccessTokensResponse) | ListUserAccessTokens returns a list of access tokens for a user. |
| CreateUserAccessToken | [CreateUserAccessTokenRequest](#share-api-v1-CreateUserAccessTokenRequest) | [CreateUserAccessTokenResponse](#share-api-v1-CreateUserAccessTokenResponse) | CreateUserAccessToken creates a new access token for a user. |
| DeleteUserAccessToken | [DeleteUserAccessTokenRequest](#share-api-v1-DeleteUserAccessTokenRequest) | [DeleteUserAccessTokenResponse](#share-api-v1-DeleteUserAccessTokenResponse) | DeleteUserAccessToken deletes an access token for a user. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

