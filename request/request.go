package request

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/PrathamSkilltelligent/pmgingo/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/mo"
	"golang.org/x/exp/constraints"
)

type ParamValidatorFn[T ParamType] func(T) bool
type ParamValueConverterFn[T ParamType] func(string) mo.Result[*T]

type ParamType interface {
	constraints.Integer | bool | string | uuid.UUID
}

type ParameterSource int

const (
	Undefined ParameterSource = iota
	PathParameter
	QueryParameter
	HttpHeader
)

func (p ParameterSource) String() string {
	switch p {
	case PathParameter:
		return "path"
	case QueryParameter:
		return "query"
	case HttpHeader:
		return "header"
	default:
		return "unknown"
	}
}

/** Path Parameter Functions **/

func getParamFrom(
	c *gin.Context,
	name string,
	source ParameterSource,
) mo.Result[string] {
	switch source {
	case PathParameter:
		val := c.Param(name)
		return mo.Ok[string](val)
	case QueryParameter:
		val, _ := c.GetQuery(name)
		return mo.Ok[string](val)
	case HttpHeader:
		val := c.GetHeader(name)
		return mo.Ok[string](val)
	}
	return mo.Err[string](errors.ErrInvalidParameterSource(source.String()))
}

func getParam[T ParamType](
	c *gin.Context,
	name string,
	isMandatory bool,
	source ParameterSource,
	converter ParamValueConverterFn[T],
	validator *ParamValidatorFn[T],
) mo.Result[*T] {
	paramResult := getParamFrom(c, name, source)

	paramVal, err := paramResult.Get()
	if err != nil {
		return mo.Err[*T](err)
	}
	if paramVal == "" {
		if isMandatory {
			return mo.Err[*T](errors.ErrParameterNotFound(name))
		} else {
			//optional and hence return nil
			return mo.Ok[*T](nil)
		}
	} else {
		convertResult := converter(paramVal)
		convertPtr, err := convertResult.Get()
		if err != nil {
			return mo.Err[*T](errors.ErrInvalidParameter(name))
		}
		converted := *convertPtr
		if validator != nil {
			validatorFn := *validator
			validated := validatorFn(converted)
			if validated {
				return mo.Ok[*T](&converted)
			} else {
				return mo.Err[*T](errors.ErrInvalidParameter(name))
			}
		} else {
			return mo.Ok[*T](&converted)
		}
	}

}

func GetIntegerParam(
	c *gin.Context,
	name string,
	source ParameterSource,
	isMandatory bool,
) mo.Result[*uint64] {
	valResult := getParam(c, name, isMandatory, source, func(v string) mo.Result[*uint64] {
		val, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return mo.Err[*uint64](errors.ErrTypeCastFailed(name, v, "uint64", err))
		} else {
			return mo.Ok(&val)
		}
	}, nil)
	val, err := valResult.Get()
	if err != nil {
		return mo.Err[*uint64](err)
	} else {
		return mo.Ok(val)
	}
}

func GetStringParam(
	c *gin.Context,
	name string,
	source ParameterSource,
	isMandatory bool,
) mo.Result[*string] {
	valResult := getParam(c, name, isMandatory, source, func(v string) mo.Result[*string] {
		return mo.Ok(&v)
	}, nil)
	val, err := valResult.Get()
	if err != nil {
		return mo.Err[*string](err)
	} else {
		return mo.Ok(val)
	}
}

func GetValidatedStringParam(
	c *gin.Context,
	name string,
	source ParameterSource,
	isMandatory bool,
	validatorFn ParamValidatorFn[string],
) mo.Result[*string] {
	valResult := getParam(c, name, isMandatory, source, func(v string) mo.Result[*string] {
		return mo.Ok(&v)
	}, &validatorFn)
	val, err := valResult.Get()
	if err != nil {
		return mo.Err[*string](err)
	} else {
		return mo.Ok(val)
	}
}

func GetUuidParam(
	c *gin.Context,
	name string,
	source ParameterSource,
	isMandatory bool,
) mo.Result[*uuid.UUID] {
	valResult := getParam(c, name, isMandatory, source, func(v string) mo.Result[*uuid.UUID] {
		val, err := uuid.Parse(v)
		log.Println(v, val, err)
		if err != nil {
			return mo.Err[*uuid.UUID](errors.ErrTypeCastFailed(name, v, "uuid", err))
		} else {
			return mo.Ok(&val)
		}
	}, nil)
	val, err := valResult.Get()
	if err != nil {
		return mo.Err[*uuid.UUID](err)
	} else {
		return mo.Ok(val)
	}

}

func GetBooleanParam(
	c *gin.Context,
	name string,
	source ParameterSource,
	isMandatory bool,
) mo.Result[*bool] {
	valResult := getParam(c, name, isMandatory, source, func(v string) mo.Result[*bool] {
		val, err := strconv.ParseBool(v)
		if err != nil {
			return mo.Err[*bool](errors.ErrTypeCastFailed(name, v, "bool", err))
		} else {
			return mo.Ok(&val)
		}
	}, nil)
	val, err := valResult.Get()
	if err != nil {
		return mo.Err[*bool](err)
	} else {
		return mo.Ok(val)
	}

}

func GetValueFromGinContext[T any](c *gin.Context, name string) mo.Result[*T] {
	val, valExist := c.Get(name)
	if valExist {
		t, ok := val.(*T)
		if !ok {
			bval, _ := json.Marshal(val)
			return mo.Err[*T](errors.ErrTypeCastFailed(name, string(bval), "bool", nil))
		}
		return mo.Ok(t)
	}
	return mo.Err[*T](errors.ErrGetValFromGinCtx(name, nil))
}
