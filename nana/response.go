package nana

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// isSuccess checks if status code is 2xx: The action was successfully received,
// understood, and accepted.
func isSuccess(code int) bool {
	return code/100 == 2
}

func checkResponse(res *http.Response) error {
	if isSuccess(res.StatusCode) {
		return nil
	}
	decoder := json.NewDecoder(res.Body)
	result := ErrorResponse{}
	if err := decoder.Decode(&result); err != nil {
		byteArray, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(byteArray))
		return &APIError{
			Code: res.StatusCode,
		}
	}
	return &APIError{
		Code:     res.StatusCode,
		Response: &result,
	}
}

// closeResponse ...
func closeResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(ioutil.Discard, res.Body)
	return err
}

// decodeCreateResponse ...
func decodeCreateResponse(res *http.Response) (*CreateResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &CreateResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodeLoginResponse ...
func decodeLoginResponse(res *http.Response) (*LoginResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &LoginResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodeGetMyInfoResponse
func decodeGetMyInfoResponse(res *http.Response) (*GetMyInfoResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &GetMyInfoResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodeUserFollowResponse
func decodeUserFollowResponse(res *http.Response) (*UserFollowResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &UserFollowResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodePostPlayResponse
func decodePostPlayResponse(res *http.Response) (*PostPlayResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &PostPlayResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodePostApplauseResponse
func decodePostApplauseResponse(res *http.Response) (*PostApplauseResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &PostApplauseResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodePostCommentResponse
func decodePostCommentResponse(res *http.Response) (*PostCommentResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &PostCommentResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// decodeCommunityJoinResponse
func decodeCommunityJoinResponse(res *http.Response) (*CommunityJoinResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &CommunityJoinResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
