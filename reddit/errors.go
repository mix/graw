package reddit

import (
	"fmt"
)

var (
	PermissionDeniedErr   = fmt.Errorf("unauthorized access to endpoint")
	NotFoundErr           = fmt.Errorf("404 not found err")
	InternalServerErr     = fmt.Errorf("500 internal server error")
	BusyErr               = fmt.Errorf("Reddit is busy right now")
	RateLimitErr          = fmt.Errorf("Reddit is rate limiting requests")
	GatewayErr            = fmt.Errorf("502 bad gateway code from Reddit")
	GatewayTimeoutErr     = fmt.Errorf("504 gateway timeout from Reddit")
	ThreadDoesNotExistErr = fmt.Errorf("The requested post does not exist.")
)
