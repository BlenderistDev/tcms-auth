package errors

import "errors"

// ErrNoUser user not found
var ErrNoUser = errors.New("no user")

// ErrWrongPassword password is incorrect
var ErrWrongPassword = errors.New("wrong password")

// ErrNoAuth user is not authed
var ErrNoAuth = errors.New("user is not authed")

// ErrNoUserId error if no user id in meta data
var ErrNoUserId = errors.New("no user id")
