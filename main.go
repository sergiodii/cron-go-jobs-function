package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	//##localToExecution##
}

// ========================================================
// WRITE THE FUNCTIONS BELOW
// ========================================================

func UpdateReditArtificialHot() {
	responseString := HttpGet("https://api.reddit.com/r/artificial/hot")

	if len(responseString) <= 0 {
		log.Fatal("Error on get url")
	}

	var SubreditResponse SubreditArtificialHotResponse
	err := json.Unmarshal([]byte(responseString), &SubreditResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range SubreditResponse.Data.Children {
		var dataToSend SubreditArtificialHotDataToSend
		dataToSend.Author = item.Data.AuthorFullname
		dataToSend.Title = item.Data.Title
		dataToSend.Ups = item.Data.Ups
		dataToSend.NumComments = item.Data.NumComments
		dataToSend.CreationData = item.Data.Created.integer
		HttpPost("http://localhost:8080/posts", dataToSend)
	}
}

func HttpPost(url string, data SubreditArtificialHotDataToSend) {
	resp, err := http.Post(url, "aplication/json", bytes.NewBuffer([]byte(data.ToJson())))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func HttpGet(url string) string {
	// return mocks.GET_FROM_HT
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	log.Fatal("Error:", resp.Status, resp.StatusCode)
	return ""
}

// ========================================================
// WRITE THE TYPES AND INTERFACES BELOW
// ========================================================

type SubreditArtificialHotDataToSend struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	Ups          int64  `json:"ups"`
	NumComments  int64  `json:"num_comments"`
	CreationData int64  `json:"created_data"`
}

func (sd *SubreditArtificialHotDataToSend) FromJson(dataJson string) *SubreditArtificialHotDataToSend {
	ndata := new(SubreditArtificialHotDataToSend)
	json.Unmarshal([]byte(dataJson), &ndata)
	return ndata
}

func (sd *SubreditArtificialHotDataToSend) ToJson() string {
	j, e := json.Marshal(sd)
	if e != nil {
		return `{"error":"error"}`
	}
	return string(j)
}

type TimesTamp struct {
	integer int64
}

func (t *TimesTamp) UnmarshalJSON(data []byte) error {
	// fmt.Println("sergio", string(data[:]))
	in, err := strconv.ParseInt(strings.ReplaceAll(string(data[:]), ".0", ""), 10, 64)
	if err != nil {
		return err
	}
	t.integer = in
	return nil
}

type SubreditArtificialHotResponse struct {
	Data struct {
		After    string      `json:"after"`
		Before   interface{} `json:"before"`
		Children []struct {
			Data struct {
				AllAwardings []struct {
					AwardSubType                     string      `json:"award_sub_type"`
					AwardType                        string      `json:"award_type"`
					AwardingsRequiredToGrantBenefits interface{} `json:"awardings_required_to_grant_benefits"`
					CoinPrice                        int64       `json:"coin_price"`
					CoinReward                       int64       `json:"coin_reward"`
					Count                            int64       `json:"count"`
					DaysOfDripExtension              int64       `json:"days_of_drip_extension"`
					DaysOfPremium                    int64       `json:"days_of_premium"`
					Description                      string      `json:"description"`
					EndDate                          interface{} `json:"end_date"`
					GiverCoinReward                  int64       `json:"giver_coin_reward"`
					IconFormat                       string      `json:"icon_format"`
					IconHeight                       int64       `json:"icon_height"`
					IconURL                          string      `json:"icon_url"`
					IconWidth                        int64       `json:"icon_width"`
					ID                               string      `json:"id"`
					IsEnabled                        bool        `json:"is_enabled"`
					IsNew                            bool        `json:"is_new"`
					Name                             string      `json:"name"`
					PennyDonate                      int64       `json:"penny_donate"`
					PennyPrice                       int64       `json:"penny_price"`
					ResizedIcons                     []struct {
						Height int64  `json:"height"`
						URL    string `json:"url"`
						Width  int64  `json:"width"`
					} `json:"resized_icons"`
					ResizedStaticIcons []struct {
						Height int64  `json:"height"`
						URL    string `json:"url"`
						Width  int64  `json:"width"`
					} `json:"resized_static_icons"`
					StartDate                interface{} `json:"start_date"`
					StaticIconHeight         int64       `json:"static_icon_height"`
					StaticIconURL            string      `json:"static_icon_url"`
					StaticIconWidth          int64       `json:"static_icon_width"`
					SubredditCoinReward      int64       `json:"subreddit_coin_reward"`
					SubredditID              interface{} `json:"subreddit_id"`
					TiersByRequiredAwardings interface{} `json:"tiers_by_required_awardings"`
				} `json:"all_awardings"`
				AllowLiveComments          bool          `json:"allow_live_comments"`
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				ApprovedBy                 interface{}   `json:"approved_by"`
				Archived                   bool          `json:"archived"`
				Author                     string        `json:"author"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				AuthorFlairCSSClass        interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext        []interface{} `json:"author_flair_richtext"`
				AuthorFlairTemplateID      string        `json:"author_flair_template_id"`
				AuthorFlairText            string        `json:"author_flair_text"`
				AuthorFlairTextColor       string        `json:"author_flair_text_color"`
				AuthorFlairType            string        `json:"author_flair_type"`
				AuthorFullname             string        `json:"author_fullname"`
				AuthorIsBlocked            bool          `json:"author_is_blocked"`
				AuthorPatreonFlair         bool          `json:"author_patreon_flair"`
				AuthorPremium              bool          `json:"author_premium"`
				Awarders                   []interface{} `json:"awarders"`
				BannedAtUtc                interface{}   `json:"banned_at_utc"`
				BannedBy                   interface{}   `json:"banned_by"`
				CanGild                    bool          `json:"can_gild"`
				CanModPost                 bool          `json:"can_mod_post"`
				Category                   interface{}   `json:"category"`
				Clicked                    bool          `json:"clicked"`
				ContentCategories          interface{}   `json:"content_categories"`
				ContestMode                bool          `json:"contest_mode"`
				Created                    TimesTamp     `json:"created"`
				CreatedUtc                 interface{}   `json:"created_utc"`
				DiscussionType             interface{}   `json:"discussion_type"`
				Distinguished              string        `json:"distinguished"`
				Domain                     string        `json:"domain"`
				Downs                      int64         `json:"downs"`
				Edited                     interface{}   `json:"edited"`
				GalleryData                struct {
					Items []struct {
						ID      int64  `json:"id"`
						MediaID string `json:"media_id"`
					} `json:"items"`
				} `json:"gallery_data"`
				Gilded   int64 `json:"gilded"`
				Gildings struct {
					Gid1 int64 `json:"gid_1"`
					Gid2 int64 `json:"gid_2"`
					Gid3 int64 `json:"gid_3"`
				} `json:"gildings"`
				Hidden                   bool          `json:"hidden"`
				HideScore                bool          `json:"hide_score"`
				ID                       string        `json:"id"`
				IsCreatedFromAdsUI       bool          `json:"is_created_from_ads_ui"`
				IsCrosspostable          bool          `json:"is_crosspostable"`
				IsGallery                bool          `json:"is_gallery"`
				IsMeta                   bool          `json:"is_meta"`
				IsOriginalContent        bool          `json:"is_original_content"`
				IsRedditMediaDomain      bool          `json:"is_reddit_media_domain"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				IsSelf                   bool          `json:"is_self"`
				IsVideo                  bool          `json:"is_video"`
				Likes                    interface{}   `json:"likes"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				LinkFlairCSSClass        string        `json:"link_flair_css_class"`
				LinkFlairRichtext        []interface{} `json:"link_flair_richtext"`
				LinkFlairTemplateID      string        `json:"link_flair_template_id"`
				LinkFlairText            string        `json:"link_flair_text"`
				LinkFlairTextColor       string        `json:"link_flair_text_color"`
				LinkFlairType            string        `json:"link_flair_type"`
				Locked                   bool          `json:"locked"`
				Media                    struct {
					Oembed struct {
						AuthorName      string `json:"author_name"`
						AuthorURL       string `json:"author_url"`
						Height          int64  `json:"height"`
						HTML            string `json:"html"`
						ProviderName    string `json:"provider_name"`
						ProviderURL     string `json:"provider_url"`
						ThumbnailHeight int64  `json:"thumbnail_height"`
						ThumbnailURL    string `json:"thumbnail_url"`
						ThumbnailWidth  int64  `json:"thumbnail_width"`
						Title           string `json:"title"`
						Type            string `json:"type"`
						Version         string `json:"version"`
						Width           int64  `json:"width"`
					} `json:"oembed"`
					Type string `json:"type"`
				} `json:"media"`
				MediaEmbed struct {
					Content   string `json:"content"`
					Height    int64  `json:"height"`
					Scrolling bool   `json:"scrolling"`
					Width     int64  `json:"width"`
				} `json:"media_embed"`
				MediaMetadata struct {
					Qd1iin7nccg71 struct {
						DashURL string `json:"dashUrl"`
						E       string `json:"e"`
						HlsURL  string `json:"hlsUrl"`
						ID      string `json:"id"`
						IsGif   bool   `json:"isGif"`
						Status  string `json:"status"`
						X       int64  `json:"x"`
						Y       int64  `json:"y"`
					} `json:"qd1iin7nccg71"`
					Wxl0qzrsp8g71 struct {
						DashURL string `json:"dashUrl"`
						E       string `json:"e"`
						HlsURL  string `json:"hlsUrl"`
						ID      string `json:"id"`
						IsGif   bool   `json:"isGif"`
						Status  string `json:"status"`
						X       int64  `json:"x"`
						Y       int64  `json:"y"`
					} `json:"wxl0qzrsp8g71"`
					X7igi2xrccg71 struct {
						E  string `json:"e"`
						ID string `json:"id"`
						M  string `json:"m"`
						P  []struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"p"`
						S struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"s"`
						Status string `json:"status"`
					} `json:"x7igi2xrccg71"`
					Yyd9sjy9w7g71 struct {
						E  string `json:"e"`
						ID string `json:"id"`
						M  string `json:"m"`
						P  []struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"p"`
						S struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"s"`
						Status string `json:"status"`
					} `json:"yyd9sjy9w7g71"`
					Zxbbtjy9w7g71 struct {
						E  string `json:"e"`
						ID string `json:"id"`
						M  string `json:"m"`
						P  []struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"p"`
						S struct {
							U string `json:"u"`
							X int64  `json:"x"`
							Y int64  `json:"y"`
						} `json:"s"`
						Status string `json:"status"`
					} `json:"zxbbtjy9w7g71"`
				} `json:"media_metadata"`
				MediaOnly             bool          `json:"media_only"`
				ModNote               interface{}   `json:"mod_note"`
				ModReasonBy           interface{}   `json:"mod_reason_by"`
				ModReasonTitle        interface{}   `json:"mod_reason_title"`
				ModReports            []interface{} `json:"mod_reports"`
				Name                  string        `json:"name"`
				NoFollow              bool          `json:"no_follow"`
				NumComments           int64         `json:"num_comments"`
				NumCrossposts         int64         `json:"num_crossposts"`
				NumReports            interface{}   `json:"num_reports"`
				Over18                bool          `json:"over_18"`
				ParentWhitelistStatus string        `json:"parent_whitelist_status"`
				Permalink             string        `json:"permalink"`
				Pinned                bool          `json:"pinned"`
				PostHint              string        `json:"post_hint"`
				Preview               struct {
					Enabled bool `json:"enabled"`
					Images  []struct {
						ID          string `json:"id"`
						Resolutions []struct {
							Height int64  `json:"height"`
							URL    string `json:"url"`
							Width  int64  `json:"width"`
						} `json:"resolutions"`
						Source struct {
							Height int64  `json:"height"`
							URL    string `json:"url"`
							Width  int64  `json:"width"`
						} `json:"source"`
						Variants struct{} `json:"variants"`
					} `json:"images"`
				} `json:"preview"`
				Pwls              int64       `json:"pwls"`
				Quarantine        bool        `json:"quarantine"`
				RemovalReason     interface{} `json:"removal_reason"`
				RemovedBy         interface{} `json:"removed_by"`
				RemovedByCategory interface{} `json:"removed_by_category"`
				ReportReasons     interface{} `json:"report_reasons"`
				Saved             bool        `json:"saved"`
				Score             int64       `json:"score"`
				SecureMedia       struct {
					Oembed struct {
						AuthorName      string `json:"author_name"`
						AuthorURL       string `json:"author_url"`
						Height          int64  `json:"height"`
						HTML            string `json:"html"`
						ProviderName    string `json:"provider_name"`
						ProviderURL     string `json:"provider_url"`
						ThumbnailHeight int64  `json:"thumbnail_height"`
						ThumbnailURL    string `json:"thumbnail_url"`
						ThumbnailWidth  int64  `json:"thumbnail_width"`
						Title           string `json:"title"`
						Type            string `json:"type"`
						Version         string `json:"version"`
						Width           int64  `json:"width"`
					} `json:"oembed"`
					Type string `json:"type"`
				} `json:"secure_media"`
				SecureMediaEmbed struct {
					Content        string `json:"content"`
					Height         int64  `json:"height"`
					MediaDomainURL string `json:"media_domain_url"`
					Scrolling      bool   `json:"scrolling"`
					Width          int64  `json:"width"`
				} `json:"secure_media_embed"`
				Selftext              string        `json:"selftext"`
				SelftextHTML          string        `json:"selftext_html"`
				SendReplies           bool          `json:"send_replies"`
				Spoiler               bool          `json:"spoiler"`
				Stickied              bool          `json:"stickied"`
				Subreddit             string        `json:"subreddit"`
				SubredditID           string        `json:"subreddit_id"`
				SubredditNamePrefixed string        `json:"subreddit_name_prefixed"`
				SubredditSubscribers  int64         `json:"subreddit_subscribers"`
				SubredditType         string        `json:"subreddit_type"`
				SuggestedSort         string        `json:"suggested_sort"`
				Thumbnail             string        `json:"thumbnail"`
				ThumbnailHeight       int64         `json:"thumbnail_height"`
				ThumbnailWidth        int64         `json:"thumbnail_width"`
				Title                 string        `json:"title"`
				TopAwardedType        interface{}   `json:"top_awarded_type"`
				TotalAwardsReceived   int64         `json:"total_awards_received"`
				TreatmentTags         []interface{} `json:"treatment_tags"`
				Ups                   int64         `json:"ups"`
				UpvoteRatio           float64       `json:"upvote_ratio"`
				URL                   string        `json:"url"`
				URLOverriddenByDest   string        `json:"url_overridden_by_dest"`
				UserReports           []interface{} `json:"user_reports"`
				ViewCount             interface{}   `json:"view_count"`
				Visited               bool          `json:"visited"`
				WhitelistStatus       string        `json:"whitelist_status"`
				Wls                   int64         `json:"wls"`
			} `json:"data"`
			Kind string `json:"kind"`
		} `json:"children"`
		Dist      int64       `json:"dist"`
		GeoFilter interface{} `json:"geo_filter"`
		Modhash   string      `json:"modhash"`
	} `json:"data"`
	Kind string `json:"kind"`
}
