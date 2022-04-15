package server

type NetCmd string

const (
	NetCmdSignupRequest         NetCmd = "signup_request"
	NetCmdSignupFailedResponse  NetCmd = "signup_failed_response"
	NetCmdSignupSucceedResponse NetCmd = "signup_succeed_response"

	NetCmdSigninRequest         NetCmd = "signin_request"
	NetCmdSigninFailedResponse  NetCmd = "signin_failed_response"
	NetCmdSigninSucceedResponse NetCmd = "signin_succeed_response"

	NetCmdGetGameListRequest         NetCmd = "GetGameList_request"
	NetCmdGetGameListFailedResponse  NetCmd = "GetGameList_failed_response"
	NetCmdGetGameListSucceedResponse NetCmd = "GetGameList_succeed_response"

	NetCmdGetGameRequest         NetCmd = "GetGame_request"
	NetCmdGetGameFailedResponse  NetCmd = "GetGame_failed_response"
	NetCmdGetGameSucceedResponse NetCmd = "GetGame_succeed_response"

	NetCmdEnterGameRequest         NetCmd = "EnterGame_request"
	NetCmdEnterGameFailedResponse  NetCmd = "EnterGame_failed_response"
	NetCmdEnterGameSucceedResponse NetCmd = "EnterGame_succeed_response"
)
