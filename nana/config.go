package nana

// APIEndpoint constants
const (
	APIEndpointBase = "https://jackson.nana-music.com"

	APILatestVersion = "/v2.1"
	APILegacyVersion = "/v2"

	APIEndpointSignup    = "/signup"
	APIEndpointLogin     = "/login"
	APIEndpointLogout    = "/logout"
	APIEndpointGetMyInfo = "/my_page"
	// User
	APIEndpointUserFollow = "/users/%s/follow"
	// Post
	APIEndpointPostPlay     = "/posts/%d/play"
	APIEndpointPostApplause = "/posts/%d/applause"
	APIEndpointPostComment  = "/posts/%d/comments"
	// Community
	APIEndpointCommunityJoin = "/communities/%s/members"

	OSVersion  = "14.2"
	APPVersion = "3.10.2"
)
