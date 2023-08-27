package entities

import "errors"

var ErrInvalidHTTPStatusValue = errors.New("invalid http status value")

type HTTPStatus uint16

const (
	HTTPStatusContinue           HTTPStatus = 100
	HTTPStatusSwitchingProtocols HTTPStatus = 101
	HTTPStatusProcessing         HTTPStatus = 102
	HTTPStatusEarlyHints         HTTPStatus = 103

	HTTPStatusOK                   HTTPStatus = 200
	HTTPStatusCreated              HTTPStatus = 201
	HTTPStatusAccepted             HTTPStatus = 202
	HTTPStatusNonAuthoritativeInfo HTTPStatus = 203
	HTTPStatusNoContent            HTTPStatus = 204
	HTTPStatusResetContent         HTTPStatus = 205
	HTTPStatusPartialContent       HTTPStatus = 206
	HTTPStatusMultiHTTPStatus      HTTPStatus = 207
	HTTPStatusAlreadyReported      HTTPStatus = 208
	HTTPStatusIMUsed               HTTPStatus = 226

	HTTPStatusMultipleChoices   HTTPStatus = 300
	HTTPStatusMovedPermanently  HTTPStatus = 301
	HTTPStatusFound             HTTPStatus = 302
	HTTPStatusSeeOther          HTTPStatus = 303
	HTTPStatusNotModified       HTTPStatus = 304
	HTTPStatusUseProxy          HTTPStatus = 305
	HTTPStatusTemporaryRedirect HTTPStatus = 307
	HTTPStatusPermanentRedirect HTTPStatus = 308

	HTTPStatusBadRequest                   HTTPStatus = 400
	HTTPStatusUnauthorized                 HTTPStatus = 401
	HTTPStatusPaymentRequired              HTTPStatus = 402
	HTTPStatusForbidden                    HTTPStatus = 403
	HTTPStatusNotFound                     HTTPStatus = 404
	HTTPStatusMethodNotAllowed             HTTPStatus = 405
	HTTPStatusNotAcceptable                HTTPStatus = 406
	HTTPStatusProxyAuthRequired            HTTPStatus = 407
	HTTPStatusRequestTimeout               HTTPStatus = 408
	HTTPStatusConflict                     HTTPStatus = 409
	HTTPStatusGone                         HTTPStatus = 410
	HTTPStatusLengthRequired               HTTPStatus = 411
	HTTPStatusPreconditionFailed           HTTPStatus = 412
	HTTPStatusRequestEntityTooLarge        HTTPStatus = 413
	HTTPStatusRequestURITooLong            HTTPStatus = 414
	HTTPStatusUnsupportedMediaType         HTTPStatus = 415
	HTTPStatusRequestedRangeNotSatisfiable HTTPStatus = 416
	HTTPStatusExpectationFailed            HTTPStatus = 417
	HTTPStatusTeapot                       HTTPStatus = 418
	HTTPStatusMisdirectedRequest           HTTPStatus = 421
	HTTPStatusUnprocessableEntity          HTTPStatus = 422
	HTTPStatusLocked                       HTTPStatus = 423
	HTTPStatusFailedDependency             HTTPStatus = 424
	HTTPStatusTooEarly                     HTTPStatus = 425
	HTTPStatusUpgradeRequired              HTTPStatus = 426
	HTTPStatusPreconditionRequired         HTTPStatus = 428
	HTTPStatusTooManyRequests              HTTPStatus = 429
	HTTPStatusRequestHeaderFieldsTooLarge  HTTPStatus = 431
	HTTPStatusUnavailableForLegalReasons   HTTPStatus = 451

	HTTPStatusInternalServerError           HTTPStatus = 500
	HTTPStatusNotImplemented                HTTPStatus = 501
	HTTPStatusBadGateway                    HTTPStatus = 502
	HTTPStatusServiceUnavailable            HTTPStatus = 503
	HTTPStatusGatewayTimeout                HTTPStatus = 504
	HTTPStatusHTTPVersionNotSupported       HTTPStatus = 505
	HTTPStatusVariantAlsoNegotiates         HTTPStatus = 506
	HTTPStatusInsufficientStorage           HTTPStatus = 507
	HTTPStatusLoopDetected                  HTTPStatus = 508
	HTTPStatusNotExtended                   HTTPStatus = 510
	HTTPStatusNetworkAuthenticationRequired HTTPStatus = 511
)

var httpStatusByInt = map[uint16]HTTPStatus{
	100: HTTPStatusContinue,
	101: HTTPStatusSwitchingProtocols,
	102: HTTPStatusProcessing,
	103: HTTPStatusEarlyHints,
	200: HTTPStatusOK,
	201: HTTPStatusCreated,
	202: HTTPStatusAccepted,
	203: HTTPStatusNonAuthoritativeInfo,
	204: HTTPStatusNoContent,
	205: HTTPStatusResetContent,
	206: HTTPStatusPartialContent,
	207: HTTPStatusMultiHTTPStatus,
	208: HTTPStatusAlreadyReported,
	226: HTTPStatusIMUsed,
	300: HTTPStatusMultipleChoices,
	301: HTTPStatusMovedPermanently,
	302: HTTPStatusFound,
	303: HTTPStatusSeeOther,
	304: HTTPStatusNotModified,
	305: HTTPStatusUseProxy,
	307: HTTPStatusTemporaryRedirect,
	308: HTTPStatusPermanentRedirect,
	400: HTTPStatusBadRequest,
	401: HTTPStatusUnauthorized,
	402: HTTPStatusPaymentRequired,
	403: HTTPStatusForbidden,
	404: HTTPStatusNotFound,
	405: HTTPStatusMethodNotAllowed,
	406: HTTPStatusNotAcceptable,
	407: HTTPStatusProxyAuthRequired,
	408: HTTPStatusRequestTimeout,
	409: HTTPStatusConflict,
	410: HTTPStatusGone,
	411: HTTPStatusLengthRequired,
	412: HTTPStatusPreconditionFailed,
	413: HTTPStatusRequestEntityTooLarge,
	414: HTTPStatusRequestURITooLong,
	415: HTTPStatusUnsupportedMediaType,
	416: HTTPStatusRequestedRangeNotSatisfiable,
	517: HTTPStatusExpectationFailed,
	418: HTTPStatusTeapot,
	421: HTTPStatusMisdirectedRequest,
	422: HTTPStatusUnprocessableEntity,
	423: HTTPStatusLocked,
	424: HTTPStatusFailedDependency,
	425: HTTPStatusTooEarly,
	426: HTTPStatusUpgradeRequired,
	428: HTTPStatusPreconditionRequired,
	429: HTTPStatusTooManyRequests,
	431: HTTPStatusRequestHeaderFieldsTooLarge,
	451: HTTPStatusUnavailableForLegalReasons,
	500: HTTPStatusInternalServerError,
	501: HTTPStatusNotImplemented,
	502: HTTPStatusBadGateway,
	503: HTTPStatusServiceUnavailable,
	504: HTTPStatusGatewayTimeout,
	505: HTTPStatusHTTPVersionNotSupported,
	506: HTTPStatusVariantAlsoNegotiates,
	507: HTTPStatusInsufficientStorage,
	508: HTTPStatusLoopDetected,
	510: HTTPStatusNotExtended,
	511: HTTPStatusNetworkAuthenticationRequired,
}

func HTTPStatusFromInt(i uint16) (HTTPStatus, error) {
	status, found := httpStatusByInt[i]
	if !found {
		return 0, ErrInvalidHTTPStatusValue
	}

	return status, nil
}
