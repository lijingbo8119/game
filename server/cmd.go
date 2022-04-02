package server

type DataCmd string

const (
	CmdSignupRequest         DataCmd = "signup_request"
	CmdSignupFailedResponse  DataCmd = "signup_failed_response"
	CmdSignupSucceedResponse DataCmd = "signup_succeed_response"

	CmdSigninRequest         DataCmd = "signin_request"
	CmdSigninFailedResponse  DataCmd = "signin_failed_response"
	CmdSigninSucceedResponse DataCmd = "signin_succeed_response"
)
