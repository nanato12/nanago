package nana

// ErrorResponse ...
type ErrorResponse struct {
	Result string `json:"result"`
	Data   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"data"`
}

// CreateResponse ...
type CreateResponse struct {
	Result string `json:"result"`
	Data   struct {
		User struct {
			UserID         int           `json:"user_id"`
			ScreenName     string        `json:"screen_name"`
			PicURL         string        `json:"pic_url"`
			PicURLMedium   string        `json:"pic_url_medium"`
			PicURLLarge    string        `json:"pic_url_large"`
			IsOfficial     bool          `json:"is_official"`
			ProfileURL     string        `json:"profile_url"`
			CoverPicURL    string        `json:"cover_pic_url"`
			IsPremium      bool          `json:"is_premium"`
			Profile        string        `json:"profile"`
			FollowingCount int           `json:"following_count"`
			FollowerCount  int           `json:"follower_count"`
			SoundCount     int           `json:"sound_count"`
			ApplauseCount  int           `json:"applause_count"`
			PlaylistCount  int           `json:"playlist_count"`
			CommunityCount int           `json:"community_count"`
			IsFollowing    bool          `json:"is_following"`
			IsFollower     bool          `json:"is_follower"`
			IsMute         bool          `json:"is_mute"`
			IsFavorite     bool          `json:"is_favorite"`
			PinnedPost     interface{}   `json:"pinned_post"`
			PinnedPlaylist interface{}   `json:"pinned_playlist"`
			IsBlocking     bool          `json:"is_blocking"`
			IsBlocked      bool          `json:"is_blocked"`
			Supporters     []interface{} `json:"supporters"`
			TwitterURL     string        `json:"twitter_url"`
			FacebookURL    string        `json:"facebook_url"`
			Country        struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"user"`
		Token string `json:"token"`
	} `json:"data"`
}

// LoginResponse ...
type LoginResponse struct {
	Result string `json:"result"`
	Data   struct {
		User struct {
			UserID         int           `json:"user_id"`
			ScreenName     string        `json:"screen_name"`
			PicURL         string        `json:"pic_url"`
			PicURLMedium   string        `json:"pic_url_medium"`
			PicURLLarge    string        `json:"pic_url_large"`
			IsOfficial     bool          `json:"is_official"`
			ProfileURL     string        `json:"profile_url"`
			CoverPicURL    string        `json:"cover_pic_url"`
			IsPremium      bool          `json:"is_premium"`
			Profile        string        `json:"profile"`
			FollowingCount int           `json:"following_count"`
			FollowerCount  int           `json:"follower_count"`
			SoundCount     int           `json:"sound_count"`
			ApplauseCount  int           `json:"applause_count"`
			PlaylistCount  int           `json:"playlist_count"`
			CommunityCount int           `json:"community_count"`
			IsFollowing    bool          `json:"is_following"`
			IsFollower     bool          `json:"is_follower"`
			IsMute         bool          `json:"is_mute"`
			IsFavorite     bool          `json:"is_favorite"`
			PinnedPost     interface{}   `json:"pinned_post"`
			PinnedPlaylist interface{}   `json:"pinned_playlist"`
			IsBlocking     bool          `json:"is_blocking"`
			IsBlocked      bool          `json:"is_blocked"`
			Supporters     []interface{} `json:"supporters"`
			TwitterURL     string        `json:"twitter_url"`
			FacebookURL    string        `json:"facebook_url"`
			Country        struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"country"`
			IosPurchaseExpiresDate     interface{} `json:"ios_purchase_expires_date"`
			AndroidPurchaseExpiresDate interface{} `json:"android_purchase_expires_date"`
			PremiumTicketExpiresDate   interface{} `json:"premium_ticket_expires_date"`
			IsPremiumByTicket          bool        `json:"is_premium_by_ticket"`
			Gender                     interface{} `json:"gender"`
			Birthday                   interface{} `json:"birthday"`
			TrialPremiumFunctions      struct {
				SoundEffect        bool `json:"sound_effect"`
				SearchApplauseSort bool `json:"search_applause_sort"`
				PlayCountReport    bool `json:"play_count_report"`
			} `json:"trial_premium_functions"`
			FirebaseProperties struct {
				UserID         string `json:"user_id"`
				PremiumStatus  string `json:"premium_status"`
				TwitterAccount string `json:"twitter_account"`
			} `json:"firebase_properties"`
			AdServer string `json:"ad_server"`
		} `json:"user"`
		Token string `json:"token"`
	} `json:"data"`
}

// GetMyInfoResponse ...
type GetMyInfoResponse struct {
	UserID         int           `json:"user_id"`
	ScreenName     string        `json:"screen_name"`
	PicURL         string        `json:"pic_url"`
	PicURLMedium   string        `json:"pic_url_medium"`
	PicURLLarge    string        `json:"pic_url_large"`
	IsOfficial     bool          `json:"is_official"`
	ProfileURL     string        `json:"profile_url"`
	CoverPicURL    string        `json:"cover_pic_url"`
	IsPremium      bool          `json:"is_premium"`
	Profile        string        `json:"profile"`
	FollowingCount int           `json:"following_count"`
	FollowerCount  int           `json:"follower_count"`
	SoundCount     int           `json:"sound_count"`
	ApplauseCount  int           `json:"applause_count"`
	PlaylistCount  int           `json:"playlist_count"`
	CommunityCount int           `json:"community_count"`
	IsFollowing    bool          `json:"is_following"`
	IsFollower     bool          `json:"is_follower"`
	IsMute         bool          `json:"is_mute"`
	IsFavorite     bool          `json:"is_favorite"`
	PinnedPost     interface{}   `json:"pinned_post"`
	PinnedPlaylist interface{}   `json:"pinned_playlist"`
	IsBlocking     bool          `json:"is_blocking"`
	IsBlocked      bool          `json:"is_blocked"`
	Supporters     []interface{} `json:"supporters"`
	TwitterURL     string        `json:"twitter_url"`
	FacebookURL    string        `json:"facebook_url"`
	Country        struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
	IosPurchaseExpiresDate     interface{} `json:"ios_purchase_expires_date"`
	AndroidPurchaseExpiresDate interface{} `json:"android_purchase_expires_date"`
	PremiumTicketExpiresDate   interface{} `json:"premium_ticket_expires_date"`
	IsPremiumByTicket          bool        `json:"is_premium_by_ticket"`
	Gender                     interface{} `json:"gender"`
	Birthday                   interface{} `json:"birthday"`
	TrialPremiumFunctions      struct {
		SoundEffect        bool `json:"sound_effect"`
		SearchApplauseSort bool `json:"search_applause_sort"`
		PlayCountReport    bool `json:"play_count_report"`
	} `json:"trial_premium_functions"`
	FirebaseProperties struct {
		UserID         string `json:"user_id"`
		PremiumStatus  string `json:"premium_status"`
		TwitterAccount string `json:"twitter_account"`
	} `json:"firebase_properties"`
	AdServer string `json:"ad_server"`
}

// UserFollowResponse ...
type UserFollowResponse struct {
	Result string `json:"result"`
	Data   struct {
		UserID        int `json:"user_id"`
		FollowerCount int `json:"follower_count"`
	} `json:"data"`
}

// PostPlayResponse ...
type PostPlayResponse struct {
	Result string `json:"result"`
	Data   struct {
	} `json:"data"`
}

// PostApplauseResponse ...
type PostApplauseResponse struct {
	Result string `json:"result"`
	Data   struct {
		PostID    int    `json:"post_id"`
		CreatedAt string `json:"created_at"`
		User      struct {
			UserID       int    `json:"user_id"`
			ScreenName   string `json:"screen_name"`
			PicURL       string `json:"pic_url"`
			PicURLMedium string `json:"pic_url_medium"`
			PicURLLarge  string `json:"pic_url_large"`
			IsOfficial   bool   `json:"is_official"`
			ProfileURL   string `json:"profile_url"`
			CoverPicURL  string `json:"cover_pic_url"`
			IsPremium    bool   `json:"is_premium"`
			Profile      string `json:"profile"`
			IsBlocking   bool   `json:"is_blocking"`
			IsBlocked    bool   `json:"is_blocked"`
			IsFollowing  bool   `json:"is_following"`
		} `json:"user"`
		PartID              int           `json:"part_id"`
		Caption             string        `json:"caption"`
		Artist              string        `json:"artist"`
		Title               string        `json:"title"`
		Duration            int           `json:"duration"`
		SoundURL            string        `json:"sound_url"`
		ImageURLFhd         interface{}   `json:"image_url_fhd"`
		IsCollaboWaiting    bool          `json:"is_collabo_waiting"`
		Key                 string        `json:"key"`
		PlayCount           int           `json:"play_count"`
		ApplauseCount       int           `json:"applause_count"`
		CommentCount        int           `json:"comment_count"`
		CollaboCount        int           `json:"collabo_count"`
		Collabos            []interface{} `json:"collabos"`
		PlayerURL           string        `json:"player_url"`
		Private             bool          `json:"private"`
		SingleTrackURL      string        `json:"single_track_url"`
		IsMixed             bool          `json:"is_mixed"`
		IsCollaboLater      bool          `json:"is_collabo_later"`
		Supporters          []interface{} `json:"supporters"`
		TotalGiftPoints     int           `json:"total_gift_points"`
		FirstToArriveUserID int           `json:"first_to_arrive_user_id"`
		IsApplauded         bool          `json:"is_applauded"`
		IsReposted          bool          `json:"is_reposted"`
		Comments            []struct {
			CommentID int    `json:"comment_id"`
			CreatedAt string `json:"created_at"`
			Body      string `json:"body"`
			User      struct {
				UserID       int    `json:"user_id"`
				ScreenName   string `json:"screen_name"`
				PicURL       string `json:"pic_url"`
				PicURLMedium string `json:"pic_url_medium"`
				PicURLLarge  string `json:"pic_url_large"`
				IsOfficial   bool   `json:"is_official"`
				ProfileURL   string `json:"profile_url"`
				CoverPicURL  string `json:"cover_pic_url"`
				IsPremium    bool   `json:"is_premium"`
			} `json:"user"`
			ReplyTo struct {
				CommentID int    `json:"comment_id"`
				CreatedAt string `json:"created_at"`
				Body      string `json:"body"`
				User      struct {
					UserID       int    `json:"user_id"`
					ScreenName   string `json:"screen_name"`
					PicURL       string `json:"pic_url"`
					PicURLMedium string `json:"pic_url_medium"`
					PicURLLarge  string `json:"pic_url_large"`
					IsOfficial   bool   `json:"is_official"`
					ProfileURL   string `json:"profile_url"`
					CoverPicURL  string `json:"cover_pic_url"`
					IsPremium    bool   `json:"is_premium"`
				} `json:"user"`
			} `json:"reply_to"`
		} `json:"comments"`
		OverdubCount int `json:"overdub_count"`
		Genre        struct {
			GenreID int    `json:"genre_id"`
			Label   string `json:"label"`
		} `json:"genre"`
		MusicKey string `json:"music_key"`
		Acc      bool   `json:"acc"`
		OgpURL   string `json:"ogp_url"`
		Language string `json:"language"`
	} `json:"data"`
}

// PostCommentResponse ...
type PostCommentResponse struct {
	Result string `json:"result"`
	Data   struct {
		CommentList []struct {
			CommentID int    `json:"comment_id"`
			CreatedAt string `json:"created_at"`
			Body      string `json:"body"`
			User      struct {
				UserID       int    `json:"user_id"`
				ScreenName   string `json:"screen_name"`
				PicURL       string `json:"pic_url"`
				PicURLMedium string `json:"pic_url_medium"`
				PicURLLarge  string `json:"pic_url_large"`
				IsOfficial   bool   `json:"is_official"`
				ProfileURL   string `json:"profile_url"`
				CoverPicURL  string `json:"cover_pic_url"`
				IsPremium    bool   `json:"is_premium"`
			} `json:"user"`
			ReplyTo interface{} `json:"reply_to"`
		} `json:"comment_list"`
		CommentCount int `json:"comment_count"`
	} `json:"data"`
}

// CommunityJoinResponse ...
type CommunityJoinResponse struct {
	Result string `json:"result"`
	Data   struct {
	} `json:"data"`
}
