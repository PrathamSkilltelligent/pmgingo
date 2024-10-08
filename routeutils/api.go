package routeutils

import (
	"net/http"

	"github.com/PrathamSkilltelligent/pmgingo/errors"
	"github.com/PrathamSkilltelligent/pmgingo/request"
	"github.com/PrathamSkilltelligent/pmgingo/types"
	"github.com/PrathamSkilltelligent/pmgo/fault"
	"github.com/PrathamSkilltelligent/pmgo/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/mo"
)

type RequestCtx struct {
	GinCtx *gin.Context
	IP     types.Ip
	CallId types.CallId
	UserId types.UserId
	OrgIds []types.OrgId
}

func NewRequestCtx(
	ginCtx *gin.Context,
) *RequestCtx {
	ip := ginCtx.ClientIP()

	// var callId *types.CallId
	// callId, f := GetCallerId(ginCtx).Get()
	// if f != nil {
	// 	callId = (*types.CallId)(utils.ToPtr(uuid.New()))
	// }

	// userId, f := GetUserId(ginCtx).Get()
	// if f != nil {
	// 	userId = (*types.UserId)(utils.ToPtr(uuid.New()))
	// }
	// var orgIds []types.OrgId
	// if ids, _ := GetOrgIds(ginCtx).Get(); ids != nil && len(*ids) > 0 {
	// 	orgIds = *ids
	// 	fmt.Println("orgIDs", uuid.UUID(orgIds[0]))
	// }
	return &RequestCtx{
		GinCtx: ginCtx,
		IP:     types.Ip(ip),
		// CallId: *callId,
		// UserId: *userId,
		// OrgIds: orgIds,
	}
}

type ApplicationContext interface {
	// SetFaultBundle(*i18n.Bundle)
	// GetFaultBundle() *i18n.Bundle
	// SetLogger(*logger.Logger)
	// GetLogger() *logger.Logger
	IsApplicationContext() //marker method
}

type ApiRequestHandler[C ApplicationContext, T any] func(C, *RequestCtx) mo.Result[*T]

func HandleRequest[C ApplicationContext, T any](ctx C, handler ApiRequestHandler[C, T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res mo.Result[*T]
		defer func() {
			switch exception := recover(); exception {
			case nil:
				if res.IsError() {
					_, f := res.Get()
					originalErr, _ := f.(fault.Fault) //No need to check for type assertion success, since we know that upstream will always provide fault.Fault
					status := getStatusCode(originalErr.ResponseErrType())
					// res := getErrorResponse(originalErr, ctx.GetFaultBundle())
					c.JSON(status, res)
					c.Request.Body.Close() // #nosec G104
				} else {
					responseData, _ := res.Get()
					// c.JSON(200, *responseData)
					if responseData == nil {
						c.JSON(200, map[string]any{
							"data": nil,
						})
					} else {
						c.JSON(200, map[string]any{
							"data": *responseData,
						})
					}
					c.Request.Body.Close() // #nosec G104
				}
			default:
				// f := errors.InternalServerError(fmt.Errorf("error %+v", exception))
				// ctx.GetLogger().Error(f.ToMessageAwareFault(ctx.GetFaultBundle()).Message("en"), utils.ToPtr(f.Code()))
				// stack := debug.Stack()
				// ctx.GetLogger().Error(fmt.Sprintf("Stack trace for exception : %v \n %v", exception, string(stack)), nil, nil, nil, nil, nil)
				// if _, file, line, ok := runtime.Caller(1); ok {
				// 	ctx.GetLogger().Error(fmt.Sprintf("Recovered from panic in file %s at line %d: %v\n", file, line, exception), nil, nil, nil, nil, nil)
				// } else {
				// 	ctx.GetLogger().Error(fmt.Sprintln("Recovered from panic but couldn't retrieve file name and line number"), nil, nil, nil, nil, nil)
				// }
				c.AbortWithStatusJSON(500, gin.H{
					"Message": "Internal Server Error. Please Contact Admin.",
				})
				c.Request.Body.Close() // #nosec G104
			}
		}()

		reqCtx := NewRequestCtx(c)
		res = handler(ctx, reqCtx)
	}
}

func getStatusCode(response fault.ResponseErrType) int {
	switch response {
	case errors.BadRequest:
		return http.StatusBadRequest
	case errors.Unauthorized:
		return http.StatusUnauthorized
	case errors.Forbidden:
		return http.StatusForbidden
	case errors.NotFound:
		return http.StatusNotFound
	case errors.AlreadyExists:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

// func getErrorResponse(f fault.Fault) map[string]any {
// 	var errors []string
// 	for _, cause := range f.Causes() {
// 		errors = append(errors, cause.Error())
// 	}
// 	return map[string]any{
// 		"errors": map[string]any{
// 			// "message":     f.ToMessageAwareFault(fb).Message("en"),
// 			"errorCode":   f.Code().String(),
// 			"otherErrors": errors,
// 		},
// 	}
// }

type ApiMiddlewareHandler[C ApplicationContext] func(C, *gin.Context) mo.Result[*bool]

func HandleMiddleware[C ApplicationContext](ctx C, middlewareHandler ApiMiddlewareHandler[C]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res mo.Result[*bool]
		defer func() {
			switch exception := recover(); exception {
			case nil:
				if res.IsError() {
					_, f := res.Get()
					originalF, _ := f.(fault.Fault)
					status := getStatusCode(originalF.ResponseErrType())
					// res := getErrorResponse(originalF, ctx.GetFaultBundle())
					c.AbortWithStatusJSON(status, res)
					c.Request.Body.Close() // #nosec G104
				} else {
					c.Next()
				}
			default:
				// f := errors.InternalServerError(fmt.Errorf("error %+v", exception))
				// ctx.GetLogger().Error(f.ToMessageAwareFault(ctx.GetFaultBundle()).Message("en"), utils.ToPtr(f.Code()))
				// stack := debug.Stack()
				// ctx.GetLogger().Error(fmt.Sprintf("Stack trace for exception : %v \n %v", exception, string(stack)), nil, nil, nil, nil, nil)
				// if _, file, line, ok := runtime.Caller(1); ok {
				// 	ctx.GetLogger().Error(fmt.Sprintf("Recovered from panic in file %s at line %d: %v\n", file, line, exception), nil, nil, nil, nil, nil)
				// } else {
				// 	ctx.GetLogger().Error(fmt.Sprintln("Recovered from panic but couldn't retrieve file name and line number"), nil, nil, nil, nil, nil)
				// }
				c.AbortWithStatusJSON(500, gin.H{
					"Message": "Internal Server Error. Please Contact Admin.",
				})
				c.Request.Body.Close() // #nosec G104
			}
		}()
		res = middlewareHandler(ctx, c)
	}
}

func GetDataFromRequestBody[T any](c *gin.Context) mo.Result[*T] {
	var t T
	if err := c.ShouldBindJSON(&t); err != nil {
		return mo.Err[*T](errors.GetRequestDataError(err))
	}
	return mo.Ok(&t)
}

func GetDataFromFormRequestBody[T any](c *gin.Context) mo.Result[*T] {
	var t T
	if err := c.ShouldBind(&t); err != nil {
		return mo.Err[*T](errors.GetRequestDataError(err))
	}
	return mo.Ok(&t)
}

func GetUnixTimeFromParam(paramName string, isMandatory bool, c *gin.Context) mo.Result[*types.Milliseconds] {
	timeParamResult := request.GetIntegerParam(c, paramName, request.QueryParameter, isMandatory)
	if timeParamResult.IsError() {
		return mo.Err[*types.Milliseconds](errors.GetUnixTimeFromQueryParamError())
	}
	timeParam, _ := timeParamResult.Get()
	if timeParam == nil {
		return mo.Ok[*types.Milliseconds](nil)
	}
	unixMilli := types.Milliseconds(*timeParam) // #nosec G115
	return mo.Ok(&unixMilli)
}

func GetCallerId(c *gin.Context) mo.Result[*types.CallId] {
	callerIdResult := request.GetUuidParam(c, "call-id", request.HttpHeader, true)
	if callerIdResult.IsError() {
		callerIdResult = request.GetUuidParam(c, "x-call-id", request.HttpHeader, true)
		if callerIdResult.IsError() {
			_, err := callerIdResult.Get()
			originalErr, _ := err.(fault.Fault)
			return mo.Err[*types.CallId](errors.GetCallerIdError("x-call-id or call-id", originalErr.Cause()))
		}
	}
	callId, _ := callerIdResult.Get()
	return mo.Ok(utils.ToPtr(types.CallId(*callId)))
}

func GetUserId(c *gin.Context) mo.Result[*types.UserId] {
	userIdResult := request.GetValueFromGinContext[types.UserId](c, "user_id")
	if userIdResult.IsError() {
		_, err := userIdResult.Get()
		originalErr, _ := err.(fault.Fault)
		return mo.Err[*types.UserId](errors.GetUserIdError("user_id", originalErr.Cause()))
	}
	return userIdResult
}

func GetOrgIds(c *gin.Context) mo.Result[*[]types.OrgId] {
	orgIds, f := request.GetValueFromGinContext[[]types.OrgId](c, "org_ids").Get()
	if f != nil {
		return mo.Err[*[]types.OrgId](f)
	}
	return mo.Ok(orgIds)
}

func GetOrgIdFromParam(c *gin.Context) mo.Result[*types.OrgId] {
	orgIdResult := request.GetUuidParam(c, "orgid", request.PathParameter, true)
	if orgIdResult.IsError() {
		_, originalErr := orgIdResult.Get()
		err, _ := originalErr.(fault.Fault)
		return mo.Err[*types.OrgId](errors.GetOrgIdFromParamError(err.Cause()))
	}
	orgId, _ := orgIdResult.Get()
	return mo.Ok((*types.OrgId)(orgId))
}
