package errors

import goErrors "errors"

var ErrNoRows = goErrors.New("no rows in result set")
var UnauthorizedError = goErrors.New("unauthorized")
var InvalidLoginPasswordError = goErrors.New("invalid login/password")
var NotFoundError = goErrors.New("not found")
