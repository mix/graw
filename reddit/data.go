package reddit

import "strings"

// Comment represents a comment on Reddit (Reddit type t1_).
// https://github.com/reddit/reddit/wiki/JSON#comment-implements-votable--created
type Comment struct {
	ID        string `mapstructure:"id"`
	Name      string `mapstructure:"name"`
	Permalink string `mapstructure:"permalink"`

	CreatedUTC uint64 `mapstructure:"created_utc"`
	Edited     uint64 `mapstructure:"edited"`
	Deleted    bool   `mapstructure:"deleted"`

	Ups   int32 `mapstructure:"ups"`
	Downs int32 `mapstructure:"downs"`
	Likes bool  `mapstructure:"likes"`

	Author              string `mapstructure:"author"`
	AuthorFlairCSSClass string `mapstructure:"author_flair_css_class"`
	AuthorFlairText     string `mapstructure:"author_flair_text"`

	LinkAuthor string `mapstructure:"link_author"`
	LinkURL    string `mapstructure:"link_url"`
	LinkTitle  string `mapstructure:"link_title"`

	Subreddit   string `mapstructure:"subreddit"`
	SubredditID string `mapstructure:"subreddit_id"`

	Body     string `mapstructure:"body"`
	BodyHTML string `mapstructure:"body_html"`

	ParentID string     `mapstructure:"parent_id"`
	Replies  []*Comment `mapstructure:"reply_tree"`
	More     *More

	Gilded        int32  `mapstructure:"gilded"`
	Distinguished string `mapstructure:"distinguished"`
}

// IsTopLevel is true when the comment is a top level comment.
func (c *Comment) IsTopLevel() bool {
	parentType := strings.Split(c.ParentID, "_")[0]
	return parentType == postKind
}

// Media represents a subfield in the response about posts
type Media struct {
	Type   string `mapstructure:"type"`
	OEmbed struct {
		ProviderURL     string `mapstructure:"provider_url"`
		Description     string `mapstructure:"description"`
		Title           string `mapstructure:"title"`
		ThumbnailWidth  int    `mapstructure:"thumbnail_width"`
		Height          int    `mapstructure:"height"`
		Width           int    `mapstructure:"width"`
		HTML            string `mapstructure:"html"`
		Version         string `mapstructure:"version"`
		ProviderName    string `mapstructure:"provider_name"`
		ThumbnailURL    string `mapstructure:"thumbnail_url"`
		Type            string `mapstructure:"type"`
		ThumbnailHeight int    `mapstructure:"thumbnail_height"`
	} `mapstructure:"oembed"`
	RedditVideo struct {
		FallbackURL       string `mapstructure:"fallback_url"`
		Height            int    `mapstructure:"height"`
		Width             int    `mapstructure:"width"`
		ScrubberMediaURL  string `mapstructure:"scrubber_media_url"`
		DashURL           string `mapstructure:"dash_url"`
		Duration          int    `mapstructure:"duration"`
		HLSURL            string `mapstructure:"hls_url"`
		IsGIF             bool   `mapstructure:"is_gif"`
		TranscodingStatus string `mapstructure:"transcoding_status"`
	} `mapstructure:"reddit_video"`
}

// Post represents posts on Reddit (Reddit type t3_).
// https://github.com/reddit/reddit/wiki/JSON#link-implements-votable--created
type Post struct {
	ID        string `mapstructure:"id"`
	Name      string `mapstructure:"name"`
	Permalink string `mapstructure:"permalink"`

	CreatedUTC uint64 `mapstructure:"created_utc"`
	Deleted    bool   `mapstructure:"deleted"`

	Ups   int32 `mapstructure:"ups"`
	Downs int32 `mapstructure:"downs"`
	Likes bool  `mapstructure:"likes"`

	Author              string `mapstructure:"author"`
	AuthorFlairCSSClass string `mapstructure:"author_flair_css_class"`
	AuthorFlairText     string `mapstructure:"author_flair_text"`

	Title  string `mapstructure:"title"`
	Score  int32  `mapstructure:"score"`
	URL    string `mapstructure:"url"`
	Domain string `mapstructure:"domain"`
	NSFW   bool   `mapstructure:"over_18"`

	Subreddit   string `mapstructure:"subreddit"`
	SubredditID string `mapstructure:"subreddit_id"`

	IsSelf       bool   `mapstructure:"is_self"`
	SelfText     string `mapstructure:"selftext"`
	SelfTextHTML string `mapstructure:"selftext_html"`

	Replies []*Comment `mapstructure:"reply_tree"`
	More    *More

	Hidden            bool   `mapstructure:"hidden"`
	LinkFlairCSSClass string `mapstructure:"link_flair_css_class"`
	LinkFlairText     string `mapstructure:"link_flair_text"`
	RemovedByCategory string `mapstructure:"removed_by_category"`

	NumComments int32  `mapstructure:"num_comments"`
	Locked      bool   `mapstructure:"locked"`
	Thumbnail   string `mapstructure:"thumbnail"`

	Gilded        int32  `mapstructure:"gilded"`
	Distinguished string `mapstructure:"distinguished"`
	Stickied      bool   `mapstructure:"stickied"`

	IsRedditMediaDomain bool  `mapstructure:"is_reddit_media_domain"`
	Media               Media `mapstructure:"media"`
	SecureMedia         Media `mapstructure:"secure_media"`

	Preview PostPreview `mapstructure:"preview"`

	GalleryData PostGallery `mapstructure:"gallery_data"`

	MediaMetadata map[string]MediaMetadata `mapstructure:"media_metadata"`

	CrosspostParentList []Post `mapstructure:"crosspost_parent_list"`
}
type PostPreview struct {
	Images             []*PostPreviewImageSet         `mapstructure:"images"`
	RedditVideoPreview *PostPreviewRedditVideoPreview `mapstructure:"reddit_video_preview"`
}
type PostPreviewImageSet struct {
	Source PostPreviewImageMedia `mapstructure:"source"`
}
type PostPreviewImageMedia struct {
	URL    string `mapstructure:"url"`
	Height int    `mapstructure:"height"`
	Width  int    `mapstructure:"width"`
}
type PostPreviewRedditVideoPreview struct {
	BitrateKPBS       int    `mapstructure:"bitrate_kbps"`
	FallbackURL       string `mapstructure:"fallback_url"`
	Height            int    `mapstructure:"height"`
	Width             int    `mapstructure:"width"`
	ScrubberMediaURL  string `mapstructure:"scrubber_media_url"`
	DashURL           string `mapstructure:"dash_url"`
	Duration          int    `mapstructure:"duration"`
	HLSURL            string `mapstructure:"hls_url"`
	IsGIF             bool   `mapstructure:"is_gif"`
	TranscodingStatus string `mapstructure:"transcoding_status"`
}

type PostGallery struct {
	Items []*GalleryItem `mapstructure:"items"`
}

type GalleryItem struct {
	ID      int    `mapstructure:"id"`
	MediaId string `mapstructure:"media_id"`
}

type MediaMetadata struct {
	Status string `mapstructure:"status"`
	E      string `mapstructure:"e"`
	M      string `mapstructure:"m"`  // mimetype like 'image/jpg'
	ID     string `mapstructure:"id"` //id should make MediaId in GalleryItem
}

// Message represents messages on Reddit (Reddit type t4_).
// https://github.com/reddit/reddit/wiki/JSON#message-implements-created
type Message struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`

	CreatedUTC uint64 `mapstructure:"created_utc"`

	Author   string `mapstructure:"author"`
	Subject  string `mapstructure:"subject"`
	Body     string `mapstructure:"body"`
	BodyHTML string `mapstructure:"body_html"`

	Context          string `mapstructure:"context"`
	FirstMessageName string `mapstructure:"first_message_name"`
	Likes            bool   `mapstructure:"likes"`
	LinkTitle        string `mapstructure:"link_title"`

	New      bool   `mapstructure:"new"`
	ParentID string `mapstructure:"parent_id"`

	Subreddit  string `mapstructure:"subreddit"`
	WasComment bool   `mapstructure:"was_comment"`
}

// More represents a more comments list on Reddit
// https://github.com/reddit-archive/reddit/wiki/JSON#more
type More struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`

	Count    int    `mapstructure:"count"`
	Depth    int    `mapstructure:"depth"`
	ParentID string `mapstructure:"parent_id"`

	Children []string `mapstructure:"children"`
}

// Harvest is a set of all possible elements that Reddit could return in a
// listing.
//
// Typically the items returned in this harvest are flat. You will not find
// the comment tree for each post in their `Replies` field. This is because
// that can be a lot of data. If you want a full comment tree, use the
// `Thread` method on the bot with the post's `Permalink`.
type Harvest struct {
	Comments []*Comment
	Posts    []*Post
	Messages []*Message
	Mores    []*More
}

type Submission struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}
