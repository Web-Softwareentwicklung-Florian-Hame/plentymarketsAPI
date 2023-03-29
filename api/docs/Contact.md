# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**TypeId** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**FormOfAddress** | Pointer to **string** |  | [optional] 
**NewsletterAllowanceAt** | Pointer to **string** |  | [optional] 
**ClassId** | Pointer to **int32** |  | [optional] 
**Blocked** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **string** |  | [optional] 
**BookAccount** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**ReferrerId** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**BirthdayAt** | Pointer to **string** |  | [optional] 
**LastLoginAt** | Pointer to **string** |  | [optional] 
**LastOrderAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Valuta** | Pointer to **string** |  | [optional] 
**IsLead** | Pointer to **bool** |  | [optional] 
**LeadStatusKey** | Pointer to **string** |  | [optional] 
**PlentyId** | Pointer to **int32** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EbayName** | Pointer to **string** |  | [optional] 
**PrivatePhone** | Pointer to **string** |  | [optional] 
**PrivateFax** | Pointer to **string** |  | [optional] 
**PrivateMobile** | Pointer to **string** |  | [optional] 
**PaypalEmail** | Pointer to **string** |  | [optional] 
**PaypalPayerId** | Pointer to **string** |  | [optional] 
**KlarnaPersonalId** | Pointer to **string** |  | [optional] 
**DhlPostIdent** | Pointer to **string** |  | [optional] 
**SingleAccess** | Pointer to **string** |  | [optional] 
**ContactPerson** | Pointer to **string** |  | [optional] 
**MarketplacePartner** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**InLeadStatusSince** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]ContactOption**](ContactOption.md) |  | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *Contact) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Contact) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Contact) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Contact) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExternalId

`func (o *Contact) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Contact) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Contact) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Contact) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTypeId

`func (o *Contact) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Contact) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Contact) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Contact) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetGender

`func (o *Contact) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Contact) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Contact) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Contact) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetTitle

`func (o *Contact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Contact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Contact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Contact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFormOfAddress

`func (o *Contact) GetFormOfAddress() string`

GetFormOfAddress returns the FormOfAddress field if non-nil, zero value otherwise.

### GetFormOfAddressOk

`func (o *Contact) GetFormOfAddressOk() (*string, bool)`

GetFormOfAddressOk returns a tuple with the FormOfAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOfAddress

`func (o *Contact) SetFormOfAddress(v string)`

SetFormOfAddress sets FormOfAddress field to given value.

### HasFormOfAddress

`func (o *Contact) HasFormOfAddress() bool`

HasFormOfAddress returns a boolean if a field has been set.

### GetNewsletterAllowanceAt

`func (o *Contact) GetNewsletterAllowanceAt() string`

GetNewsletterAllowanceAt returns the NewsletterAllowanceAt field if non-nil, zero value otherwise.

### GetNewsletterAllowanceAtOk

`func (o *Contact) GetNewsletterAllowanceAtOk() (*string, bool)`

GetNewsletterAllowanceAtOk returns a tuple with the NewsletterAllowanceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletterAllowanceAt

`func (o *Contact) SetNewsletterAllowanceAt(v string)`

SetNewsletterAllowanceAt sets NewsletterAllowanceAt field to given value.

### HasNewsletterAllowanceAt

`func (o *Contact) HasNewsletterAllowanceAt() bool`

HasNewsletterAllowanceAt returns a boolean if a field has been set.

### GetClassId

`func (o *Contact) GetClassId() int32`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *Contact) GetClassIdOk() (*int32, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *Contact) SetClassId(v int32)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *Contact) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetBlocked

`func (o *Contact) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Contact) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Contact) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *Contact) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetRating

`func (o *Contact) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Contact) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Contact) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Contact) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetBookAccount

`func (o *Contact) GetBookAccount() string`

GetBookAccount returns the BookAccount field if non-nil, zero value otherwise.

### GetBookAccountOk

`func (o *Contact) GetBookAccountOk() (*string, bool)`

GetBookAccountOk returns a tuple with the BookAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookAccount

`func (o *Contact) SetBookAccount(v string)`

SetBookAccount sets BookAccount field to given value.

### HasBookAccount

`func (o *Contact) HasBookAccount() bool`

HasBookAccount returns a boolean if a field has been set.

### GetLang

`func (o *Contact) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Contact) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Contact) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Contact) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetReferrerId

`func (o *Contact) GetReferrerId() float32`

GetReferrerId returns the ReferrerId field if non-nil, zero value otherwise.

### GetReferrerIdOk

`func (o *Contact) GetReferrerIdOk() (*float32, bool)`

GetReferrerIdOk returns a tuple with the ReferrerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerId

`func (o *Contact) SetReferrerId(v float32)`

SetReferrerId sets ReferrerId field to given value.

### HasReferrerId

`func (o *Contact) HasReferrerId() bool`

HasReferrerId returns a boolean if a field has been set.

### GetUserId

`func (o *Contact) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Contact) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Contact) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Contact) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetBirthdayAt

`func (o *Contact) GetBirthdayAt() string`

GetBirthdayAt returns the BirthdayAt field if non-nil, zero value otherwise.

### GetBirthdayAtOk

`func (o *Contact) GetBirthdayAtOk() (*string, bool)`

GetBirthdayAtOk returns a tuple with the BirthdayAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdayAt

`func (o *Contact) SetBirthdayAt(v string)`

SetBirthdayAt sets BirthdayAt field to given value.

### HasBirthdayAt

`func (o *Contact) HasBirthdayAt() bool`

HasBirthdayAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *Contact) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *Contact) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *Contact) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *Contact) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastOrderAt

`func (o *Contact) GetLastOrderAt() string`

GetLastOrderAt returns the LastOrderAt field if non-nil, zero value otherwise.

### GetLastOrderAtOk

`func (o *Contact) GetLastOrderAtOk() (*string, bool)`

GetLastOrderAtOk returns a tuple with the LastOrderAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOrderAt

`func (o *Contact) SetLastOrderAt(v string)`

SetLastOrderAt sets LastOrderAt field to given value.

### HasLastOrderAt

`func (o *Contact) HasLastOrderAt() bool`

HasLastOrderAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Contact) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Contact) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Contact) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Contact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Contact) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Contact) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Contact) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Contact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValuta

`func (o *Contact) GetValuta() string`

GetValuta returns the Valuta field if non-nil, zero value otherwise.

### GetValutaOk

`func (o *Contact) GetValutaOk() (*string, bool)`

GetValutaOk returns a tuple with the Valuta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuta

`func (o *Contact) SetValuta(v string)`

SetValuta sets Valuta field to given value.

### HasValuta

`func (o *Contact) HasValuta() bool`

HasValuta returns a boolean if a field has been set.

### GetIsLead

`func (o *Contact) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *Contact) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *Contact) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *Contact) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.

### GetLeadStatusKey

`func (o *Contact) GetLeadStatusKey() string`

GetLeadStatusKey returns the LeadStatusKey field if non-nil, zero value otherwise.

### GetLeadStatusKeyOk

`func (o *Contact) GetLeadStatusKeyOk() (*string, bool)`

GetLeadStatusKeyOk returns a tuple with the LeadStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadStatusKey

`func (o *Contact) SetLeadStatusKey(v string)`

SetLeadStatusKey sets LeadStatusKey field to given value.

### HasLeadStatusKey

`func (o *Contact) HasLeadStatusKey() bool`

HasLeadStatusKey returns a boolean if a field has been set.

### GetPlentyId

`func (o *Contact) GetPlentyId() int32`

GetPlentyId returns the PlentyId field if non-nil, zero value otherwise.

### GetPlentyIdOk

`func (o *Contact) GetPlentyIdOk() (*int32, bool)`

GetPlentyIdOk returns a tuple with the PlentyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlentyId

`func (o *Contact) SetPlentyId(v int32)`

SetPlentyId sets PlentyId field to given value.

### HasPlentyId

`func (o *Contact) HasPlentyId() bool`

HasPlentyId returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEbayName

`func (o *Contact) GetEbayName() string`

GetEbayName returns the EbayName field if non-nil, zero value otherwise.

### GetEbayNameOk

`func (o *Contact) GetEbayNameOk() (*string, bool)`

GetEbayNameOk returns a tuple with the EbayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayName

`func (o *Contact) SetEbayName(v string)`

SetEbayName sets EbayName field to given value.

### HasEbayName

`func (o *Contact) HasEbayName() bool`

HasEbayName returns a boolean if a field has been set.

### GetPrivatePhone

`func (o *Contact) GetPrivatePhone() string`

GetPrivatePhone returns the PrivatePhone field if non-nil, zero value otherwise.

### GetPrivatePhoneOk

`func (o *Contact) GetPrivatePhoneOk() (*string, bool)`

GetPrivatePhoneOk returns a tuple with the PrivatePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatePhone

`func (o *Contact) SetPrivatePhone(v string)`

SetPrivatePhone sets PrivatePhone field to given value.

### HasPrivatePhone

`func (o *Contact) HasPrivatePhone() bool`

HasPrivatePhone returns a boolean if a field has been set.

### GetPrivateFax

`func (o *Contact) GetPrivateFax() string`

GetPrivateFax returns the PrivateFax field if non-nil, zero value otherwise.

### GetPrivateFaxOk

`func (o *Contact) GetPrivateFaxOk() (*string, bool)`

GetPrivateFaxOk returns a tuple with the PrivateFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateFax

`func (o *Contact) SetPrivateFax(v string)`

SetPrivateFax sets PrivateFax field to given value.

### HasPrivateFax

`func (o *Contact) HasPrivateFax() bool`

HasPrivateFax returns a boolean if a field has been set.

### GetPrivateMobile

`func (o *Contact) GetPrivateMobile() string`

GetPrivateMobile returns the PrivateMobile field if non-nil, zero value otherwise.

### GetPrivateMobileOk

`func (o *Contact) GetPrivateMobileOk() (*string, bool)`

GetPrivateMobileOk returns a tuple with the PrivateMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMobile

`func (o *Contact) SetPrivateMobile(v string)`

SetPrivateMobile sets PrivateMobile field to given value.

### HasPrivateMobile

`func (o *Contact) HasPrivateMobile() bool`

HasPrivateMobile returns a boolean if a field has been set.

### GetPaypalEmail

`func (o *Contact) GetPaypalEmail() string`

GetPaypalEmail returns the PaypalEmail field if non-nil, zero value otherwise.

### GetPaypalEmailOk

`func (o *Contact) GetPaypalEmailOk() (*string, bool)`

GetPaypalEmailOk returns a tuple with the PaypalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmail

`func (o *Contact) SetPaypalEmail(v string)`

SetPaypalEmail sets PaypalEmail field to given value.

### HasPaypalEmail

`func (o *Contact) HasPaypalEmail() bool`

HasPaypalEmail returns a boolean if a field has been set.

### GetPaypalPayerId

`func (o *Contact) GetPaypalPayerId() string`

GetPaypalPayerId returns the PaypalPayerId field if non-nil, zero value otherwise.

### GetPaypalPayerIdOk

`func (o *Contact) GetPaypalPayerIdOk() (*string, bool)`

GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerId

`func (o *Contact) SetPaypalPayerId(v string)`

SetPaypalPayerId sets PaypalPayerId field to given value.

### HasPaypalPayerId

`func (o *Contact) HasPaypalPayerId() bool`

HasPaypalPayerId returns a boolean if a field has been set.

### GetKlarnaPersonalId

`func (o *Contact) GetKlarnaPersonalId() string`

GetKlarnaPersonalId returns the KlarnaPersonalId field if non-nil, zero value otherwise.

### GetKlarnaPersonalIdOk

`func (o *Contact) GetKlarnaPersonalIdOk() (*string, bool)`

GetKlarnaPersonalIdOk returns a tuple with the KlarnaPersonalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKlarnaPersonalId

`func (o *Contact) SetKlarnaPersonalId(v string)`

SetKlarnaPersonalId sets KlarnaPersonalId field to given value.

### HasKlarnaPersonalId

`func (o *Contact) HasKlarnaPersonalId() bool`

HasKlarnaPersonalId returns a boolean if a field has been set.

### GetDhlPostIdent

`func (o *Contact) GetDhlPostIdent() string`

GetDhlPostIdent returns the DhlPostIdent field if non-nil, zero value otherwise.

### GetDhlPostIdentOk

`func (o *Contact) GetDhlPostIdentOk() (*string, bool)`

GetDhlPostIdentOk returns a tuple with the DhlPostIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhlPostIdent

`func (o *Contact) SetDhlPostIdent(v string)`

SetDhlPostIdent sets DhlPostIdent field to given value.

### HasDhlPostIdent

`func (o *Contact) HasDhlPostIdent() bool`

HasDhlPostIdent returns a boolean if a field has been set.

### GetSingleAccess

`func (o *Contact) GetSingleAccess() string`

GetSingleAccess returns the SingleAccess field if non-nil, zero value otherwise.

### GetSingleAccessOk

`func (o *Contact) GetSingleAccessOk() (*string, bool)`

GetSingleAccessOk returns a tuple with the SingleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccess

`func (o *Contact) SetSingleAccess(v string)`

SetSingleAccess sets SingleAccess field to given value.

### HasSingleAccess

`func (o *Contact) HasSingleAccess() bool`

HasSingleAccess returns a boolean if a field has been set.

### GetContactPerson

`func (o *Contact) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *Contact) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *Contact) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *Contact) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetMarketplacePartner

`func (o *Contact) GetMarketplacePartner() string`

GetMarketplacePartner returns the MarketplacePartner field if non-nil, zero value otherwise.

### GetMarketplacePartnerOk

`func (o *Contact) GetMarketplacePartnerOk() (*string, bool)`

GetMarketplacePartnerOk returns a tuple with the MarketplacePartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplacePartner

`func (o *Contact) SetMarketplacePartner(v string)`

SetMarketplacePartner sets MarketplacePartner field to given value.

### HasMarketplacePartner

`func (o *Contact) HasMarketplacePartner() bool`

HasMarketplacePartner returns a boolean if a field has been set.

### GetFullName

`func (o *Contact) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Contact) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Contact) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Contact) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetInLeadStatusSince

`func (o *Contact) GetInLeadStatusSince() int32`

GetInLeadStatusSince returns the InLeadStatusSince field if non-nil, zero value otherwise.

### GetInLeadStatusSinceOk

`func (o *Contact) GetInLeadStatusSinceOk() (*int32, bool)`

GetInLeadStatusSinceOk returns a tuple with the InLeadStatusSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInLeadStatusSince

`func (o *Contact) SetInLeadStatusSince(v int32)`

SetInLeadStatusSince sets InLeadStatusSince field to given value.

### HasInLeadStatusSince

`func (o *Contact) HasInLeadStatusSince() bool`

HasInLeadStatusSince returns a boolean if a field has been set.

### GetOptions

`func (o *Contact) GetOptions() []ContactOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Contact) GetOptionsOk() (*[]ContactOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Contact) SetOptions(v []ContactOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Contact) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


