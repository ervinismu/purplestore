package reason

var (
	Unauthorized = "Unauthorized access"
)

// category
var (
	CategoryFailedGetList   = "failed get list category"
	CategoryFailedCreate    = "failed create category"
	CategoryFailedGetDetail = "failed get detail category"
	CategoryFailedDelete    = "failed delete category"
	CategoryFailedUpdate    = "failed update category"
	CategorySuccessCreate   = "success create category"
	CategorySuccessUpdate   = "failed update category"
	CategorySuccessDelete   = "success delete category"
)

// authentications
var (
	RegisterSuccess      = "success register"
	RegisterFailed       = "success register"
	SessionFailedRefresh = "failed refresh token"
	SessionSuccessLogout = "success logout"
	SessionFailedLogout  = "success logout"
	SessionFailedLogin   = "success logout"
)
