package server

type NetCmd string

const (
	NetCmdSignupRequest         NetCmd = "signup_request"
	NetCmdSignupFailedResponse  NetCmd = "signup_failed_response"
	NetCmdSignupSucceedResponse NetCmd = "signup_succeed_response"

	NetCmdSigninRequest         NetCmd = "signin_request"
	NetCmdSigninFailedResponse  NetCmd = "signin_failed_response"
	NetCmdSigninSucceedResponse NetCmd = "signin_succeed_response"

	NetCmdGetRoomListRequest         NetCmd = "GetRoomList_request"
	NetCmdGetRoomListFailedResponse  NetCmd = "GetRoomList_failed_response"
	NetCmdGetRoomListSucceedResponse NetCmd = "GetRoomList_succeed_response"

	NetCmdGetRoomRequest         NetCmd = "GetRoom_request"
	NetCmdGetRoomFailedResponse  NetCmd = "GetRoom_failed_response"
	NetCmdGetRoomSucceedResponse NetCmd = "GetRoom_succeed_response"

	NetCmdEnterRoomRequest         NetCmd = "EnterRoom_request"
	NetCmdEnterRoomFailedResponse  NetCmd = "EnterRoom_failed_response"
	NetCmdEnterRoomSucceedResponse NetCmd = "EnterRoom_succeed_response"
)
