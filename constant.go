package errors

type ErrorType uint8
type Locale string
type Status uint16

type Errors struct {
	Locale    Locale
	ErrorList []ErrorDictionary
}
type ErrorDictionary struct {
	key   string
	value string
}

const RootPath = "locale"
const (
	Default = iota + 1
	Info
	Warning
	Error
	Critical
)
const (
	EN_US = "en_us"
	VI_VN = "vi_vn"
)

func ConvertLocale(s string) Locale {
	switch s {
	case "en_us":
		return EN_US
	case "vi_vn":
		return VI_VN
	default:
		return EN_US
	}
}
func (s Locale) String() string {
	switch s {
	case EN_US:
		return "en_us"
	case VI_VN:
		return "vi_vn"
	default:
		return "en_us"
	}
}
func (s ErrorType) String() string {
	switch s {
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Critical:
		return "Critical"
	default:
		return "Default"
	}
}

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 7231, 6.3.1
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	_                       = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var statusText = map[int]string{
	StatusContinue:           "Continue",
	StatusSwitchingProtocols: "Switching Protocols",
	StatusProcessing:         "Processing",
	StatusEarlyHints:         "Early Hints",

	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthRequired:            "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusPreconditionFailed:           "Precondition Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	StatusExpectationFailed:            "Expectation Failed",
	StatusTeapot:                       "I'm a teapot",
	StatusMisdirectedRequest:           "Misdirected Request",
	StatusUnprocessableEntity:          "Unprocessable Entity",
	StatusLocked:                       "Locked",
	StatusFailedDependency:             "Failed Dependency",
	StatusTooEarly:                     "Too Early",
	StatusUpgradeRequired:              "Upgrade Required",
	StatusPreconditionRequired:         "Precondition Required",
	StatusTooManyRequests:              "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

func (s Status) String() string {
	switch s {
	case StatusContinue:
		return "StatusContinue"
	case StatusSwitchingProtocols:
		return "StatusSwitchingProtocols"
	case StatusProcessing:
		return "StatusProcessing"
	case StatusEarlyHints:
		return "StatusEarlyHints"
	case StatusOK:
		return "StatusOK"
	case StatusCreated:
		return "StatusCreated"
	case StatusAccepted:
		return "StatusAccepted"
	case StatusNonAuthoritativeInfo:
		return "StatusNonAuthoritativeInfo"
	case StatusNoContent:
		return "StatusNoContent"
	case StatusResetContent:
		return "StatusResetContent"
	case StatusPartialContent:
		return "StatusPartialContent"
	case StatusMultiStatus:
		return "StatusMultiStatus"
	case StatusAlreadyReported:
		return "StatusAlreadyReported"
	case StatusIMUsed:
		return "StatusIMUsed"
	case StatusMultipleChoices:
		return "StatusMultipleChoices"
	case StatusMovedPermanently:
		return "StatusMovedPermanently"
	case StatusFound:
		return "StatusFound"
	case StatusSeeOther:
		return "StatusSeeOther"
	case StatusNotModified:
		return "StatusNotModified"
	case StatusUseProxy:
		return "StatusUseProxy"
	case StatusTemporaryRedirect:
		return "StatusTemporaryRedirect"
	case StatusPermanentRedirect:
		return "StatusPermanentRedirect"
	case StatusBadRequest:
		return "StatusBadRequest"
	case StatusUnauthorized:
		return "StatusUnauthorized"
	case StatusPaymentRequired:
		return "StatusPaymentRequired"
	case StatusForbidden:
		return "StatusForbidden"
	case StatusNotFound:
		return "StatusNotFound"
	case StatusMethodNotAllowed:
		return "StatusMethodNotAllowed"
	case StatusNotAcceptable:
		return "StatusNotAcceptable"
	case StatusProxyAuthRequired:
		return "StatusProxyAuthRequired"
	case StatusRequestTimeout:
		return "StatusRequestTimeout"
	case StatusConflict:
		return "StatusConflict"
	case StatusGone:
		return "StatusGone"
	case StatusLengthRequired:
		return "StatusLengthRequired"
	case StatusPreconditionFailed:
		return "StatusPreconditionFailed"
	case StatusRequestEntityTooLarge:
		return "StatusRequestEntityTooLarge"
	case StatusRequestURITooLong:
		return "StatusRequestURITooLong"
	case StatusUnsupportedMediaType:
		return "StatusUnsupportedMediaType"
	case StatusRequestedRangeNotSatisfiable:
		return "StatusRequestedRangeNotSatisfiable"
	case StatusExpectationFailed:
		return "StatusExpectationFailed"
	case StatusTeapot:
		return "StatusTeapot"
	case StatusMisdirectedRequest:
		return "StatusMisdirectedRequest"
	case StatusUnprocessableEntity:
		return "StatusUnprocessableEntity"
	case StatusLocked:
		return "StatusLocked"
	case StatusFailedDependency:
		return "StatusFailedDependency"
	case StatusTooEarly:
		return "StatusTooEarly"
	case StatusUpgradeRequired:
		return "StatusUpgradeRequired"
	case StatusPreconditionRequired:
		return "StatusPreconditionRequired"
	case StatusTooManyRequests:
		return "StatusTooManyRequests"
	case StatusRequestHeaderFieldsTooLarge:
		return "StatusRequestHeaderFieldsTooLarge"
	case StatusUnavailableForLegalReasons:
		return "StatusUnavailableForLegalReasons"
	case StatusInternalServerError:
		return "StatusInternalServerError"
	case StatusNotImplemented:
		return "StatusNotImplemented"
	case StatusBadGateway:
		return "StatusBadGateway"
	case StatusServiceUnavailable:
		return "StatusServiceUnavailable"
	case StatusGatewayTimeout:
		return "StatusGatewayTimeout"
	case StatusHTTPVersionNotSupported:
		return "StatusHTTPVersionNotSupported"
	case StatusVariantAlsoNegotiates:
		return "StatusVariantAlsoNegotiates"
	case StatusInsufficientStorage:
		return "StatusInsufficientStorage"
	case StatusLoopDetected:
		return "StatusLoopDetected"
	case StatusNotExtended:
		return "StatusNotExtended"
	case StatusNetworkAuthenticationRequired:
		return "StatusNetworkAuthenticationRequired"
	default:
		return "StatusBadRequest"
	}
}
