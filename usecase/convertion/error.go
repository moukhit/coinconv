package convertion

import "errors"

var ErrInvalidRequest = errors.New("invalid request for currency conversion")
var ErrGettingResponse = errors.New("error has occurred while getting response from service")
