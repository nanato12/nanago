package nana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetMyInfo ...
func (n *Nana) GetMyInfo() (*GetMyInfoResponse, error) {
	endpoint := APILatestVersion + APIEndpointGetMyInfo
	res, err := n.get(nil, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeGetMyInfoResponse(res)
}

// Follow ...
func (n *Nana) Follow(AccountID string) (*UserFollowResponse, error) {
	endpoint := APILatestVersion + fmt.Sprintf(APIEndpointUserFollow, AccountID)
	res, err := n.post(nil, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeUserFollowResponse(res)
}

// PlayPost ...
func (n *Nana) PlayPost(PostID string) (*PostPlayResponse, error) {
	postID, err := strconv.ParseInt(PostID, 16, 0)
	if err != nil {
		return nil, err
	}
	endpoint := APILegacyVersion + fmt.Sprintf(APIEndpointPostPlay, postID)
	res, err := n.post(nil, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodePostPlayResponse(res)
}

// ApplausePost ...
func (n *Nana) ApplausePost(PostID string) (*PostApplauseResponse, error) {
	postID, err := strconv.ParseInt(PostID, 16, 0)
	if err != nil {
		return nil, err
	}
	endpoint := APILegacyVersion + fmt.Sprintf(APIEndpointPostApplause, postID)
	res, err := n.post(nil, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodePostApplauseResponse(res)
}

// CommentPost ...
func (n *Nana) CommentPost(PostID string, comment string) (*PostCommentResponse, error) {
	postID, err := strconv.ParseInt(PostID, 16, 0)
	if err != nil {
		return nil, err
	}
	endpoint := APILegacyVersion + fmt.Sprintf(APIEndpointPostComment, postID)
	params := struct {
		Body string `json:"body"`
	}{
		Body: comment,
	}
	paramsJSON, _ := json.Marshal(params)
	res, err := n.post(nil, endpoint, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodePostCommentResponse(res)
}

// JoinCommunity ...
func (n *Nana) JoinCommunity(CommunityID string) (*PostApplauseResponse, error) {
	endpoint := APILegacyVersion + fmt.Sprintf(APIEndpointCommunityJoin, CommunityID)
	res, err := n.post(nil, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodePostApplauseResponse(res)
}
