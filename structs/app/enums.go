package app

const (
	StatusActive   = "ACTIVE"
	StatusInactive = "INACTIVE"
)

// See https://developer.okta.com/docs/api/resources/apps.html#features
const (
	FeaturesImportNewUsers       = "IMPORT_NEW_USERS"
	FeaturesImportProfileUpdates = "IMPORT_PROFILE_UPDATES"
	FeaturesProfileMastering     = "PROFILE_MASTERING"
	FeaturesImportUserSchema     = "IMPORT_USER_SCHEMA"
	FeaturesPushNewUsers         = "PUSH_NEW_USERS"
	FeaturesProfileUpdates       = "PUSH_PROFILE_UPDATES"
	FeaturesUserDeactivation     = "PUSH_USER_DEACTIVATION"
	FeaturesReactivateUsers      = "REACTIVATE_USERS"
	FeaturesPushPasswordUpdates  = "PUSH_PASSWORD_UPDATES"
	FeaturesGroupPush            = "GROUP_PUSH"
)

// See https://developer.okta.com/docs/api/resources/apps.html#signon-modes
const (
	SignOnModesBookmark            = "BOOKMARK"
	SignOnModesBasicAuth           = "BASIC_AUTH"
	SignOnModesBrowserPlugin       = "BROWSER_PLUGIN"
	SignOnModesSecurePasswordStore = "SECURE_PASSWORD_STORE"
	SignOnModesSAML2               = "SAML_2_0"
	SignOnModesWSFederation        = "WS_FEDERATION"
	SignOnModesAutoLogin           = "AUTO_LOGIN"
	SignOnModesOpenIDConnect       = "OPENID_CONNECT"
	SignOnModesCustom              = "Custom"
)

// See https://developer.okta.com/docs/api/resources/apps.html#authentication-schemes
const (
	SchemeSharedUsernameAndPassword = "SHARED_USERNAME_AND_PASSWORD"
	SchemeExternalPasswordSync      = "EXTERNAL_PASSWORD_SYNC"
	SchemeEditUsernameAndPassword   = "EDIT_USERNAME_AND_PASSWORD"
	SchemeEditPasswordOnly          = "EDIT_PASSWORD_ONLY"
	SchemeAdminSetsCredentials      = "ADMIN_SETS_CREDENTIALS"
)
