# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-inspire.proto](#npool/cloud-hashing-inspire.proto)
    - [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting)
    - [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting)
    - [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated)
    - [CouponPool](#cloud.hashing.inspire.v1.CouponPool)
    - [CreateAgencySettingRequest](#cloud.hashing.inspire.v1.CreateAgencySettingRequest)
    - [CreateAgencySettingResponse](#cloud.hashing.inspire.v1.CreateAgencySettingResponse)
    - [CreateAppCouponSettingRequest](#cloud.hashing.inspire.v1.CreateAppCouponSettingRequest)
    - [CreateAppCouponSettingResponse](#cloud.hashing.inspire.v1.CreateAppCouponSettingResponse)
    - [CreateCouponAllocatedRequest](#cloud.hashing.inspire.v1.CreateCouponAllocatedRequest)
    - [CreateCouponAllocatedResponse](#cloud.hashing.inspire.v1.CreateCouponAllocatedResponse)
    - [CreateCouponPoolRequest](#cloud.hashing.inspire.v1.CreateCouponPoolRequest)
    - [CreateCouponPoolResponse](#cloud.hashing.inspire.v1.CreateCouponPoolResponse)
    - [CreateNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.CreateNewUserRewardSettingRequest)
    - [CreateNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.CreateNewUserRewardSettingResponse)
    - [CreatePurchaseInvitationRequest](#cloud.hashing.inspire.v1.CreatePurchaseInvitationRequest)
    - [CreatePurchaseInvitationResponse](#cloud.hashing.inspire.v1.CreatePurchaseInvitationResponse)
    - [CreateRegistrationInvitationRequest](#cloud.hashing.inspire.v1.CreateRegistrationInvitationRequest)
    - [CreateRegistrationInvitationResponse](#cloud.hashing.inspire.v1.CreateRegistrationInvitationResponse)
    - [CreateUserInvitationCodeRequest](#cloud.hashing.inspire.v1.CreateUserInvitationCodeRequest)
    - [CreateUserInvitationCodeResponse](#cloud.hashing.inspire.v1.CreateUserInvitationCodeResponse)
    - [GetAgencySettingByAppRequest](#cloud.hashing.inspire.v1.GetAgencySettingByAppRequest)
    - [GetAgencySettingByAppResponse](#cloud.hashing.inspire.v1.GetAgencySettingByAppResponse)
    - [GetAgencySettingRequest](#cloud.hashing.inspire.v1.GetAgencySettingRequest)
    - [GetAgencySettingResponse](#cloud.hashing.inspire.v1.GetAgencySettingResponse)
    - [GetAppCouponSettingByAppRequest](#cloud.hashing.inspire.v1.GetAppCouponSettingByAppRequest)
    - [GetAppCouponSettingByAppResponse](#cloud.hashing.inspire.v1.GetAppCouponSettingByAppResponse)
    - [GetAppCouponSettingRequest](#cloud.hashing.inspire.v1.GetAppCouponSettingRequest)
    - [GetAppCouponSettingResponse](#cloud.hashing.inspire.v1.GetAppCouponSettingResponse)
    - [GetCouponAllocatedRequest](#cloud.hashing.inspire.v1.GetCouponAllocatedRequest)
    - [GetCouponAllocatedResponse](#cloud.hashing.inspire.v1.GetCouponAllocatedResponse)
    - [GetCouponPoolRequest](#cloud.hashing.inspire.v1.GetCouponPoolRequest)
    - [GetCouponPoolResponse](#cloud.hashing.inspire.v1.GetCouponPoolResponse)
    - [GetCouponPoolsByAppReleaserRequest](#cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserRequest)
    - [GetCouponPoolsByAppReleaserResponse](#cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserResponse)
    - [GetCouponPoolsByAppRequest](#cloud.hashing.inspire.v1.GetCouponPoolsByAppRequest)
    - [GetCouponPoolsByAppResponse](#cloud.hashing.inspire.v1.GetCouponPoolsByAppResponse)
    - [GetCouponsAllocatedByAppRequest](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppRequest)
    - [GetCouponsAllocatedByAppResponse](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppResponse)
    - [GetCouponsAllocatedByAppUserRequest](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserRequest)
    - [GetCouponsAllocatedByAppUserResponse](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserResponse)
    - [GetNewUserRewardSettingByAppRequest](#cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppRequest)
    - [GetNewUserRewardSettingByAppResponse](#cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppResponse)
    - [GetNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.GetNewUserRewardSettingRequest)
    - [GetNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.GetNewUserRewardSettingResponse)
    - [GetPurchaseInvitationByAppOrderRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderRequest)
    - [GetPurchaseInvitationByAppOrderResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderResponse)
    - [GetPurchaseInvitationRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationRequest)
    - [GetPurchaseInvitationResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationResponse)
    - [GetPurchaseInvitationsByAppRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppRequest)
    - [GetPurchaseInvitationsByAppResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppResponse)
    - [GetRegistrationInvitationRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationRequest)
    - [GetRegistrationInvitationResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationResponse)
    - [GetRegistrationInvitationsByAppInviterRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterRequest)
    - [GetRegistrationInvitationsByAppInviterResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterResponse)
    - [GetRegistrationInvitationsByAppRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppRequest)
    - [GetRegistrationInvitationsByAppResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppResponse)
    - [GetUserInvitationCodeByAppUserRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserRequest)
    - [GetUserInvitationCodeByAppUserResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserResponse)
    - [GetUserInvitationCodeByCodeRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeRequest)
    - [GetUserInvitationCodeByCodeResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeResponse)
    - [GetUserInvitationCodeRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeRequest)
    - [GetUserInvitationCodeResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeResponse)
    - [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting)
    - [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation)
    - [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation)
    - [UpdateAgencySettingRequest](#cloud.hashing.inspire.v1.UpdateAgencySettingRequest)
    - [UpdateAgencySettingResponse](#cloud.hashing.inspire.v1.UpdateAgencySettingResponse)
    - [UpdateAppCouponSettingRequest](#cloud.hashing.inspire.v1.UpdateAppCouponSettingRequest)
    - [UpdateAppCouponSettingResponse](#cloud.hashing.inspire.v1.UpdateAppCouponSettingResponse)
    - [UpdateCouponAllocatedRequest](#cloud.hashing.inspire.v1.UpdateCouponAllocatedRequest)
    - [UpdateCouponAllocatedResponse](#cloud.hashing.inspire.v1.UpdateCouponAllocatedResponse)
    - [UpdateNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.UpdateNewUserRewardSettingRequest)
    - [UpdateNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.UpdateNewUserRewardSettingResponse)
    - [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode)
    - [VersionResponse](#cloud.hashing.inspire.v1.VersionResponse)
  
    - [CloudHashingInspire](#cloud.hashing.inspire.v1.CloudHashingInspire)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-inspire.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-inspire.proto



<a name="cloud.hashing.inspire.v1.AgencySetting"></a>

### AgencySetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RegistrationRewardThreshold | [int32](#int32) |  |  |
| RegistrationCouponID | [string](#string) |  |  |
| KycRewardThreshold | [int32](#int32) |  |  |
| KycCouponID | [string](#string) |  |  |
| PurchaseRewardPercent | [int32](#int32) |  |  |
| PurchaseRewardChainLevels | [int32](#int32) |  |  |






<a name="cloud.hashing.inspire.v1.AppCouponSetting"></a>

### AppCouponSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| DominationLimit | [double](#double) |  |  |
| TotalLimit | [int32](#int32) |  |  |






<a name="cloud.hashing.inspire.v1.CouponAllocated"></a>

### CouponAllocated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Used | [bool](#bool) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.CouponPool"></a>

### CouponPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Denomination | [double](#double) |  |  |
| Circulation | [int32](#int32) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Messge | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.CreateAgencySettingRequest"></a>

### CreateAgencySettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreateAgencySettingResponse"></a>

### CreateAgencySettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreateAppCouponSettingRequest"></a>

### CreateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreateAppCouponSettingResponse"></a>

### CreateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreateCouponAllocatedRequest"></a>

### CreateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) |  |  |






<a name="cloud.hashing.inspire.v1.CreateCouponAllocatedResponse"></a>

### CreateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) |  |  |






<a name="cloud.hashing.inspire.v1.CreateCouponPoolRequest"></a>

### CreateCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud.hashing.inspire.v1.CouponPool) |  |  |






<a name="cloud.hashing.inspire.v1.CreateCouponPoolResponse"></a>

### CreateCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud.hashing.inspire.v1.CouponPool) |  |  |






<a name="cloud.hashing.inspire.v1.CreateNewUserRewardSettingRequest"></a>

### CreateNewUserRewardSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreateNewUserRewardSettingResponse"></a>

### CreateNewUserRewardSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.CreatePurchaseInvitationRequest"></a>

### CreatePurchaseInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.CreatePurchaseInvitationResponse"></a>

### CreatePurchaseInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.CreateRegistrationInvitationRequest"></a>

### CreateRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.CreateRegistrationInvitationResponse"></a>

### CreateRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.CreateUserInvitationCodeRequest"></a>

### CreateUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode) |  |  |






<a name="cloud.hashing.inspire.v1.CreateUserInvitationCodeResponse"></a>

### CreateUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode) |  |  |






<a name="cloud.hashing.inspire.v1.GetAgencySettingByAppRequest"></a>

### GetAgencySettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetAgencySettingByAppResponse"></a>

### GetAgencySettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetAgencySettingRequest"></a>

### GetAgencySettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetAgencySettingResponse"></a>

### GetAgencySettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetAppCouponSettingByAppRequest"></a>

### GetAppCouponSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetAppCouponSettingByAppResponse"></a>

### GetAppCouponSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetAppCouponSettingRequest"></a>

### GetAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetAppCouponSettingResponse"></a>

### GetAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponAllocatedRequest"></a>

### GetCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponAllocatedResponse"></a>

### GetCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolRequest"></a>

### GetCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolResponse"></a>

### GetCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud.hashing.inspire.v1.CouponPool) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserRequest"></a>

### GetCouponPoolsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserResponse"></a>

### GetCouponPoolsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud.hashing.inspire.v1.CouponPool) | repeated |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolsByAppRequest"></a>

### GetCouponPoolsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponPoolsByAppResponse"></a>

### GetCouponPoolsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud.hashing.inspire.v1.CouponPool) | repeated |  |






<a name="cloud.hashing.inspire.v1.GetCouponsAllocatedByAppRequest"></a>

### GetCouponsAllocatedByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponsAllocatedByAppResponse"></a>

### GetCouponsAllocatedByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) | repeated |  |






<a name="cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserRequest"></a>

### GetCouponsAllocatedByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserResponse"></a>

### GetCouponsAllocatedByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) | repeated |  |






<a name="cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppRequest"></a>

### GetNewUserRewardSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppResponse"></a>

### GetNewUserRewardSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetNewUserRewardSettingRequest"></a>

### GetNewUserRewardSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetNewUserRewardSettingResponse"></a>

### GetNewUserRewardSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderRequest"></a>

### GetPurchaseInvitationByAppOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderResponse"></a>

### GetPurchaseInvitationByAppOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationRequest"></a>

### GetPurchaseInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationResponse"></a>

### GetPurchaseInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppRequest"></a>

### GetPurchaseInvitationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppResponse"></a>

### GetPurchaseInvitationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PurchaseInvitation](#cloud.hashing.inspire.v1.PurchaseInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationRequest"></a>

### GetRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationResponse"></a>

### GetRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterRequest"></a>

### GetRegistrationInvitationsByAppInviterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterResponse"></a>

### GetRegistrationInvitationsByAppInviterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppRequest"></a>

### GetRegistrationInvitationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppResponse"></a>

### GetRegistrationInvitationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud.hashing.inspire.v1.RegistrationInvitation) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserRequest"></a>

### GetUserInvitationCodeByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserResponse"></a>

### GetUserInvitationCodeByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeRequest"></a>

### GetUserInvitationCodeByCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeResponse"></a>

### GetUserInvitationCodeByCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeRequest"></a>

### GetUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.GetUserInvitationCodeResponse"></a>

### GetUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud.hashing.inspire.v1.UserInvitationCode) |  |  |






<a name="cloud.hashing.inspire.v1.NewUserRewardSetting"></a>

### NewUserRewardSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RegistrationCouponID | [string](#string) |  |  |
| KycCouponID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.PurchaseInvitation"></a>

### PurchaseInvitation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| InvitationCodeID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.RegistrationInvitation"></a>

### RegistrationInvitation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InviterID | [string](#string) |  |  |
| InviteeID | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateAgencySettingRequest"></a>

### UpdateAgencySettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateAgencySettingResponse"></a>

### UpdateAgencySettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AgencySetting](#cloud.hashing.inspire.v1.AgencySetting) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateAppCouponSettingRequest"></a>

### UpdateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateAppCouponSettingResponse"></a>

### UpdateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud.hashing.inspire.v1.AppCouponSetting) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateCouponAllocatedRequest"></a>

### UpdateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateCouponAllocatedResponse"></a>

### UpdateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud.hashing.inspire.v1.CouponAllocated) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateNewUserRewardSettingRequest"></a>

### UpdateNewUserRewardSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.UpdateNewUserRewardSettingResponse"></a>

### UpdateNewUserRewardSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [NewUserRewardSetting](#cloud.hashing.inspire.v1.NewUserRewardSetting) |  |  |






<a name="cloud.hashing.inspire.v1.UserInvitationCode"></a>

### UserInvitationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InvitationCode | [string](#string) |  |  |






<a name="cloud.hashing.inspire.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.inspire.v1.CloudHashingInspire"></a>

### CloudHashingInspire
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.inspire.v1.VersionResponse) | Method Version |
| CreateNewUserRewardSetting | [CreateNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.CreateNewUserRewardSettingRequest) | [CreateNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.CreateNewUserRewardSettingResponse) |  |
| GetNewUserRewardSetting | [GetNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.GetNewUserRewardSettingRequest) | [GetNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.GetNewUserRewardSettingResponse) |  |
| GetNewUserRewardSettingByApp | [GetNewUserRewardSettingByAppRequest](#cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppRequest) | [GetNewUserRewardSettingByAppResponse](#cloud.hashing.inspire.v1.GetNewUserRewardSettingByAppResponse) |  |
| UpdateNewUserRewardSetting | [UpdateNewUserRewardSettingRequest](#cloud.hashing.inspire.v1.UpdateNewUserRewardSettingRequest) | [UpdateNewUserRewardSettingResponse](#cloud.hashing.inspire.v1.UpdateNewUserRewardSettingResponse) |  |
| CreateAgencySetting | [CreateAgencySettingRequest](#cloud.hashing.inspire.v1.CreateAgencySettingRequest) | [CreateAgencySettingResponse](#cloud.hashing.inspire.v1.CreateAgencySettingResponse) |  |
| GetAgencySetting | [GetAgencySettingRequest](#cloud.hashing.inspire.v1.GetAgencySettingRequest) | [GetAgencySettingResponse](#cloud.hashing.inspire.v1.GetAgencySettingResponse) |  |
| GetAgencySettingByApp | [GetAgencySettingByAppRequest](#cloud.hashing.inspire.v1.GetAgencySettingByAppRequest) | [GetAgencySettingByAppResponse](#cloud.hashing.inspire.v1.GetAgencySettingByAppResponse) |  |
| UpdateAgencySetting | [UpdateAgencySettingRequest](#cloud.hashing.inspire.v1.UpdateAgencySettingRequest) | [UpdateAgencySettingResponse](#cloud.hashing.inspire.v1.UpdateAgencySettingResponse) |  |
| CreatePurchaseInvitation | [CreatePurchaseInvitationRequest](#cloud.hashing.inspire.v1.CreatePurchaseInvitationRequest) | [CreatePurchaseInvitationResponse](#cloud.hashing.inspire.v1.CreatePurchaseInvitationResponse) |  |
| GetPurchaseInvitation | [GetPurchaseInvitationRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationRequest) | [GetPurchaseInvitationResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationResponse) |  |
| GetPurchaseInvitationsByApp | [GetPurchaseInvitationsByAppRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppRequest) | [GetPurchaseInvitationsByAppResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationsByAppResponse) |  |
| GetPurchaseInvitationByAppOrder | [GetPurchaseInvitationByAppOrderRequest](#cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderRequest) | [GetPurchaseInvitationByAppOrderResponse](#cloud.hashing.inspire.v1.GetPurchaseInvitationByAppOrderResponse) |  |
| CreateRegistrationInvitation | [CreateRegistrationInvitationRequest](#cloud.hashing.inspire.v1.CreateRegistrationInvitationRequest) | [CreateRegistrationInvitationResponse](#cloud.hashing.inspire.v1.CreateRegistrationInvitationResponse) |  |
| GetRegistrationInvitation | [GetRegistrationInvitationRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationRequest) | [GetRegistrationInvitationResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationResponse) |  |
| GetRegistrationInvitationsByApp | [GetRegistrationInvitationsByAppRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppRequest) | [GetRegistrationInvitationsByAppResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppResponse) |  |
| GetRegistrationInvitationsByAppInviter | [GetRegistrationInvitationsByAppInviterRequest](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterRequest) | [GetRegistrationInvitationsByAppInviterResponse](#cloud.hashing.inspire.v1.GetRegistrationInvitationsByAppInviterResponse) |  |
| CreateUserInvitationCode | [CreateUserInvitationCodeRequest](#cloud.hashing.inspire.v1.CreateUserInvitationCodeRequest) | [CreateUserInvitationCodeResponse](#cloud.hashing.inspire.v1.CreateUserInvitationCodeResponse) |  |
| GetUserInvitationCode | [GetUserInvitationCodeRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeRequest) | [GetUserInvitationCodeResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeResponse) |  |
| GetUserInvitationCodeByAppUser | [GetUserInvitationCodeByAppUserRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserRequest) | [GetUserInvitationCodeByAppUserResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeByAppUserResponse) |  |
| GetUserInvitationCodeByCode | [GetUserInvitationCodeByCodeRequest](#cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeRequest) | [GetUserInvitationCodeByCodeResponse](#cloud.hashing.inspire.v1.GetUserInvitationCodeByCodeResponse) |  |
| CreateCouponAllocated | [CreateCouponAllocatedRequest](#cloud.hashing.inspire.v1.CreateCouponAllocatedRequest) | [CreateCouponAllocatedResponse](#cloud.hashing.inspire.v1.CreateCouponAllocatedResponse) |  |
| GetCouponAllocated | [GetCouponAllocatedRequest](#cloud.hashing.inspire.v1.GetCouponAllocatedRequest) | [GetCouponAllocatedResponse](#cloud.hashing.inspire.v1.GetCouponAllocatedResponse) |  |
| GetCouponsAllocatedByApp | [GetCouponsAllocatedByAppRequest](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppRequest) | [GetCouponsAllocatedByAppResponse](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppResponse) |  |
| GetCouponsAllocatedByAppUser | [GetCouponsAllocatedByAppUserRequest](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserRequest) | [GetCouponsAllocatedByAppUserResponse](#cloud.hashing.inspire.v1.GetCouponsAllocatedByAppUserResponse) |  |
| UpdateCouponAllocated | [UpdateCouponAllocatedRequest](#cloud.hashing.inspire.v1.UpdateCouponAllocatedRequest) | [UpdateCouponAllocatedResponse](#cloud.hashing.inspire.v1.UpdateCouponAllocatedResponse) |  |
| CreateCouponPool | [CreateCouponPoolRequest](#cloud.hashing.inspire.v1.CreateCouponPoolRequest) | [CreateCouponPoolResponse](#cloud.hashing.inspire.v1.CreateCouponPoolResponse) |  |
| GetCouponPool | [GetCouponPoolRequest](#cloud.hashing.inspire.v1.GetCouponPoolRequest) | [GetCouponPoolResponse](#cloud.hashing.inspire.v1.GetCouponPoolResponse) |  |
| GetCouponPoolsByApp | [GetCouponPoolsByAppRequest](#cloud.hashing.inspire.v1.GetCouponPoolsByAppRequest) | [GetCouponPoolsByAppResponse](#cloud.hashing.inspire.v1.GetCouponPoolsByAppResponse) |  |
| GetCouponPoolsByAppReleaser | [GetCouponPoolsByAppReleaserRequest](#cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserRequest) | [GetCouponPoolsByAppReleaserResponse](#cloud.hashing.inspire.v1.GetCouponPoolsByAppReleaserResponse) |  |
| CreateAppCouponSetting | [CreateAppCouponSettingRequest](#cloud.hashing.inspire.v1.CreateAppCouponSettingRequest) | [CreateAppCouponSettingResponse](#cloud.hashing.inspire.v1.CreateAppCouponSettingResponse) |  |
| GetAppCouponSetting | [GetAppCouponSettingRequest](#cloud.hashing.inspire.v1.GetAppCouponSettingRequest) | [GetAppCouponSettingResponse](#cloud.hashing.inspire.v1.GetAppCouponSettingResponse) |  |
| GetAppCouponSettingByApp | [GetAppCouponSettingByAppRequest](#cloud.hashing.inspire.v1.GetAppCouponSettingByAppRequest) | [GetAppCouponSettingByAppResponse](#cloud.hashing.inspire.v1.GetAppCouponSettingByAppResponse) |  |
| UpdateAppCouponSetting | [UpdateAppCouponSettingRequest](#cloud.hashing.inspire.v1.UpdateAppCouponSettingRequest) | [UpdateAppCouponSettingResponse](#cloud.hashing.inspire.v1.UpdateAppCouponSettingResponse) |  |

 



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

