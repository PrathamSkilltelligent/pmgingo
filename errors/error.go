package errors

import (
	"github.com/PrathamSkilltelligent/pmgingo/types"
	"github.com/PrathamSkilltelligent/pmgo/fault"
	"github.com/google/uuid"
)

/** Error Component Constants **/
const (
	ErrService     fault.ErrComponent = "service"
	ErrRepo                           = "repository"
	ErrLib                            = "library"
	ErrController  fault.ErrComponent = "controller"
	ErrApplication fault.ErrComponent = "application"
)

/** Response Error Type Constants **/
const (
	BadRequest     fault.ResponseErrType = "BadRequest"
	Forbidden      fault.ResponseErrType = "Forbidden"
	NotFound       fault.ResponseErrType = "NotFound"
	AlreadyExists  fault.ResponseErrType = "AlreadyExists"
	InternalServer fault.ResponseErrType = "InternalServerError"
	Unauthorized   fault.ResponseErrType = "Unauthorized"
)

/** Error Code Constants **/
const (
	ErrParamNotFound          fault.ErrorCode = "err-parameter-not-provided"
	ErrParamInvalid           fault.ErrorCode = "err-invalid-parameter"
	ErrParamSourceInvalid     fault.ErrorCode = "err-parameter-source-is-invalid"
	ErrTypeCast               fault.ErrorCode = "err-type-cast-failed"
	ErrGetValFromGinCtxFailed fault.ErrorCode = "err-get-val-from-gin-ctx-failed"

	// configuration error codes
	ErrAppConfigError        fault.ErrorCode = "CNF0000000000"
	ErrInternalServerError   fault.ErrorCode = "CNF0000000010"
	ErrDatabaseInternalError fault.ErrorCode = "CNF0000000020"

	// request error codes
	ErrFailedToExtractDataFromRequest fault.ErrorCode = "REQ0000000000"
	ErrGetCallerIdFromHeader          fault.ErrorCode = "REQ0000000010"
	ErrGetUserIdFromGinCtx            fault.ErrorCode = "REQ0000000020"
	ErrGetUnixTimeFromQueryParam      fault.ErrorCode = "REQ0000000030"
	ErrGetOrgIdFromPathParam          fault.ErrorCode = "REQ0000000040"
	ErrInvalidRequestBody             fault.ErrorCode = "REQ0000000060"
	ErrGeneratePostRequest            fault.ErrorCode = "REQ0000000070"
	ErrGenerateGetRequest             fault.ErrorCode = "REQ0000000080"
	ErrExecutingRequest               fault.ErrorCode = "REQ0000000090"
	ErrReadingRespBody                fault.ErrorCode = "REQ0000000100"
	ErrDecodingResponseBody           fault.ErrorCode = "REQ0000000110"
	ErrUnmarshalResponse              fault.ErrorCode = "REQ0000000120"
	ErrAuthTokenNotFound              fault.ErrorCode = "REQ0000000130" // #nosec G101
	ErrInvalidAuthToken               fault.ErrorCode = "REQ0000000140" // #nosec G101

	// db error codes
	ErrRecordNotFound fault.ErrorCode = "REPO0000000000"

	// user error codes
	ErrUserNotFound fault.ErrorCode = "USR0000000000"

	// org error codes
	ErrOrgNotFound fault.ErrorCode = "ORG0000000000"
)

// Initialize your basicfaultcache here
func buildBasicFaults() map[fault.ErrorCode]fault.BasicFault {
	var localBasicFaults = map[fault.ErrorCode]fault.BasicFault{}
	localBasicFaults[ErrParamNotFound] = fault.NewBasicFault(ErrParamNotFound).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrParamInvalid] = fault.NewBasicFault(ErrParamInvalid).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrParamSourceInvalid] = fault.NewBasicFault(ErrParamSourceInvalid).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrTypeCast] = fault.NewBasicFault(ErrTypeCast).SetComponent(ErrApplication).SetResponseType(BadRequest)
	localBasicFaults[ErrGetValFromGinCtxFailed] = fault.NewBasicFault(ErrGetValFromGinCtxFailed).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrAppConfigError] = fault.NewBasicFault(ErrAppConfigError).SetComponent(ErrApplication).SetResponseType(InternalServer)
	localBasicFaults[ErrInternalServerError] = fault.NewBasicFault(ErrInternalServerError).SetComponent(ErrApplication).SetResponseType(InternalServer)
	localBasicFaults[ErrDatabaseInternalError] = fault.NewBasicFault(ErrDatabaseInternalError).SetComponent(ErrRepo).SetResponseType(InternalServer)
	localBasicFaults[ErrFailedToExtractDataFromRequest] = fault.NewBasicFault(ErrFailedToExtractDataFromRequest).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrGetCallerIdFromHeader] = fault.NewBasicFault(ErrGetCallerIdFromHeader).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrGetUserIdFromGinCtx] = fault.NewBasicFault(ErrGetUserIdFromGinCtx).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrGetUnixTimeFromQueryParam] = fault.NewBasicFault(ErrGetUnixTimeFromQueryParam).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrInvalidRequestBody] = fault.NewBasicFault(ErrInvalidRequestBody).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrGeneratePostRequest] = fault.NewBasicFault(ErrGeneratePostRequest).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrGenerateGetRequest] = fault.NewBasicFault(ErrGenerateGetRequest).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrExecutingRequest] = fault.NewBasicFault(ErrExecutingRequest).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrReadingRespBody] = fault.NewBasicFault(ErrReadingRespBody).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrDecodingResponseBody] = fault.NewBasicFault(ErrDecodingResponseBody).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrUnmarshalResponse] = fault.NewBasicFault(ErrUnmarshalResponse).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrRecordNotFound] = fault.NewBasicFault(ErrRecordNotFound).SetComponent(ErrController).SetResponseType(InternalServer)
	localBasicFaults[ErrUserNotFound] = fault.NewBasicFault(ErrUserNotFound).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrOrgNotFound] = fault.NewBasicFault(ErrOrgNotFound).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrGetOrgIdFromPathParam] = fault.NewBasicFault(ErrGetOrgIdFromPathParam).SetComponent(ErrController).SetResponseType(BadRequest)
	localBasicFaults[ErrAuthTokenNotFound] = fault.NewBasicFault(ErrAuthTokenNotFound).SetComponent(ErrController).SetResponseType(Unauthorized)
	localBasicFaults[ErrInvalidAuthToken] = fault.NewBasicFault(ErrInvalidAuthToken).SetComponent(ErrController).SetResponseType(Unauthorized)

	return localBasicFaults
}

var localFaultCache = fault.NewBasicFaultCache(buildBasicFaults())

// var FaultBundle = faultBundle()

// //TODO Expose the full I18N bundle

// func faultBundle() *i18n.Bundle {
// 	ginFaultMsgs := []*i18n.Message{
// 		{
// 			ID: ErrParamSourceInvalid.String(),

// 			One:   "Error invalid paramter scource {{.source}}",
// 			Other: "Error invalid paramter scource {{.source}}",
// 		},
// 		{
// 			ID: ErrParamInvalid.String(),

// 			One:   "Error invalid paramter {{.name}}",
// 			Other: "Error invalid paramter {{.name}}",
// 		},
// 		{
// 			ID: ErrParamNotFound.String(),

// 			One:   "Error paramter {{.name}} not found",
// 			Other: "Error paramter {{.name}} not found",
// 		},
// 		{
// 			ID: ErrTypeCast.String(),

// 			One:   "Error failed to cast {{.name}} having value {{.val}} to {{.datatype}}",
// 			Other: "Error failed to cast {{.name}} having value {{.val}} to {{.datatype}}",
// 		},
// 		{
// 			ID: ErrGetValFromGinCtxFailed.String(),

// 			One:   "Error failed to get value from gin context for key {{.name}}",
// 			Other: "Error failed to get value from gin context for key {{.name}}",
// 		},
// 	}
// 	b := i18n.NewBundle(language.English)
// 	b.MustAddMessages(language.English, ginFaultMsgs...)
// 	return b
// }

//TODO Now start defining your Fault constructors as Closures

func ErrInvalidParameterSource(source string) fault.Fault {
	data := map[string]any{
		"source": source,
	}
	return fault.NewBasicFault(ErrParamSourceInvalid).
		SetComponent(ErrLib).SetResponseType(NotFound).
		ToFault(data, nil)
}

func ErrInvalidParameter(name string) fault.Fault {
	data := map[string]any{
		"name": name,
	}
	return fault.NewBasicFault(ErrParamInvalid).
		SetComponent(ErrLib).SetResponseType(NotFound).
		ToFault(data, nil)
}

func ErrParameterNotFound(name string) fault.Fault {
	data := map[string]any{
		"name": name,
	}
	return fault.NewBasicFault(ErrParamNotFound).
		SetComponent(ErrLib).SetResponseType(NotFound).
		ToFault(data, nil)
}

func ErrTypeCastFailed(name string, val string, datatype string, cause error) fault.Fault {
	data := map[string]any{
		"name":     name,
		"val":      val,
		"datatype": datatype,
	}
	return fault.NewBasicFault(ErrTypeCast).
		SetComponent(ErrLib).SetResponseType(NotFound).
		ToFault(data, nil)
}

func ErrGetValFromGinCtx(name string, cause error) fault.Fault {
	data := map[string]any{
		"name": name,
	}
	return fault.NewBasicFault(ErrGetValFromGinCtxFailed).
		SetComponent(ErrLib).SetResponseType(NotFound).
		ToFault(data, nil)
}

func _InternalServerError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrInternalServerError).ToFault(nil, cause)
	}
	return returnFn
}

var InternalServerError = _InternalServerError(&localFaultCache)

func _ConfigError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrAppConfigError).ToFault(nil, cause)
	}
	return returnFn
}

var ConfigError = _ConfigError(&localFaultCache)

func _DBError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrDatabaseInternalError).ToFault(nil, cause)
	}
	return returnFn
}

var DBError = _DBError(&localFaultCache)

func _RecordNotFoundError(basicFaultCache *fault.BasicFaultsCache) func(string, error) fault.Fault {
	var returnFn = func(id string, cause error) fault.Fault {
		data := map[string]any{
			"id": id,
		}
		return basicFaultCache.GetBasicFault(ErrRecordNotFound).ToFault(data, cause)
	}
	return returnFn
}

var RecordNotFoundError = _RecordNotFoundError(&localFaultCache)

func _GetRequestDataError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrFailedToExtractDataFromRequest).ToFault(nil, cause)
	}
	return returnFn
}

var GetRequestDataError = _GetRequestDataError(&localFaultCache)

func _GetUnixTimeFromQueryParamError(basicFaultCache *fault.BasicFaultsCache) func() fault.Fault {
	var returnFn = func() fault.Fault {
		return basicFaultCache.GetBasicFault(ErrGetUnixTimeFromQueryParam).ToFault(nil, nil)
	}
	return returnFn
}

var GetUnixTimeFromQueryParamError = _GetUnixTimeFromQueryParamError(&localFaultCache)

func _GetCallerIdError(basicFaultCache *fault.BasicFaultsCache) func(string, error) fault.Fault {
	var returnFn = func(name string, cause error) fault.Fault {
		data := map[string]any{
			"name": name,
		}
		return basicFaultCache.GetBasicFault(ErrGetCallerIdFromHeader).ToFault(data, cause)
	}
	return returnFn
}

var GetCallerIdError = _GetCallerIdError(&localFaultCache)

func _GetUserIdError(basicFaultCache *fault.BasicFaultsCache) func(string, error) fault.Fault {
	var returnFn = func(id string, cause error) fault.Fault {
		data := map[string]any{
			"name": id,
		}
		return basicFaultCache.GetBasicFault(ErrGetUserIdFromGinCtx).ToFault(data, cause)
	}
	return returnFn
}

var GetUserIdError = _GetUserIdError(&localFaultCache)

func _GeneratePostRequestError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrGeneratePostRequest).ToFault(nil, cause)
	}
	return returnFn
}

var GeneratePostRequestError = _GeneratePostRequestError(&localFaultCache)

func _ReadingResponseBodyError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrReadingRespBody).ToFault(nil, cause)
	}
	return returnFn
}

var ReadingResponseBodyError = _ReadingResponseBodyError(&localFaultCache)

func _GenerateGetRequestError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrGenerateGetRequest).ToFault(nil, cause)
	}
	return returnFn
}

var GenerateGetRequestError = _GenerateGetRequestError(&localFaultCache)

func _ExecutingRequestError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrExecutingRequest).ToFault(nil, cause)
	}
	return returnFn
}

var ExecutingGetRequestError = _ExecutingRequestError(&localFaultCache)

func _DecodingResponseBodyError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrDecodingResponseBody).ToFault(nil, cause)
	}
	return returnFn
}

var DecodingResponseBodyError = _DecodingResponseBodyError(&localFaultCache)

func _UnmarshalResponseError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrUnmarshalResponse).ToFault(nil, cause)
	}
	return returnFn
}

var UnmarshalResponseError = _UnmarshalResponseError(&localFaultCache)

func _InvalidRequestError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	var returnFn = func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrInvalidRequestBody).ToFault(nil, cause)
	}
	return returnFn
}

var InvalidRequestError = _InvalidRequestError(&localFaultCache)

func _OrgNotFoundError(basicFaultCache *fault.BasicFaultsCache) func(error, types.OrgId) fault.Fault {
	var returnFn = func(cause error, orgId types.OrgId) fault.Fault {
		data := map[string]any{
			"org_id": uuid.UUID(orgId),
		}
		return basicFaultCache.GetBasicFault(ErrOrgNotFound).ToFault(data, cause)
	}
	return returnFn
}

var OrgNotFoundError = _OrgNotFoundError(&localFaultCache)

func _UserNotFoundError(basicFaultCache *fault.BasicFaultsCache) func(error, types.UserId) fault.Fault {
	var returnFn = func(cause error, userId types.UserId) fault.Fault {
		data := map[string]any{
			"user_id": uuid.UUID(userId),
		}
		return basicFaultCache.GetBasicFault(ErrUserNotFound).ToFault(data, cause)
	}
	return returnFn
}

var UserNotFoundError = _UserNotFoundError(&localFaultCache)

func _GetOrgIdFromParamError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrGetOrgIdFromPathParam).ToFault(nil, cause)
	}
}

var GetOrgIdFromParamError = _GetOrgIdFromParamError(&localFaultCache)

func _AuthTokenNotFoundError(basicFaultCache *fault.BasicFaultsCache) func() fault.Fault {
	return func() fault.Fault {
		return basicFaultCache.GetBasicFault(ErrAuthTokenNotFound).ToFault(nil, nil)
	}
}

var AuthTokenNotFoundError = _AuthTokenNotFoundError(&localFaultCache)

func _AuthTokenInvalidError(basicFaultCache *fault.BasicFaultsCache) func(error) fault.Fault {
	return func(cause error) fault.Fault {
		return basicFaultCache.GetBasicFault(ErrInvalidAuthToken).ToFault(nil, cause)
	}
}

var AuthTokenInvalidError = _AuthTokenInvalidError(&localFaultCache)
