package types

type AnyRequest struct {
}

type Category struct {
	Model
	CID   string `json:"cid,optional" gorm:"column:cid"`
	Depth int    `json:"depth,optional"`
	Title string `json:"title"`
	Order int    `json:"order,optional"`
	Path  string `json:"path,optional" gorm:"column:path;unique"`
}

type Collection struct {
	Model
	CID         string    `json:"cid,optional" gorm:"column:cid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Order       int       `json:"order,optional" gorm:"column:order"`
	Path        string    `json:"path,optional" gorm:"column:path;unique"`
	Proxy       bool      `json:"proxy,optional" gorm:"column:proxy"`
	CategoryID  string    `json:"-" gorm:"column:category_id;index"` // 外键字段
	Category    *Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Description string    `json:"description,optional"`
	Thumbnail   string    `json:"thumbnail,optional"`
	Tags        []string  `json:"tags,optional" gorm:"type:json"`
	View        int       `json:"view,optional"`
}

type CollectionsData struct {
	Category *Category `json:"category"` // 使用 inline 标签将 Category 字段内联
	Groups   []*Group  `json:"groups"`
}

type CollectionsResponse struct {
	Datas []*CollectionsData `json:"datas"`
}

type Comment struct {
	ID           string `json:"id"`
	CollectionId string `json:"collection_id"`
	Nickname     string `json:"nickname"`
	Message      string `json:"message"`
}

type CommentRequest struct {
	CollectionId string `json:"collection_id"`
	Nickname     string `json:"nickname"`
	Message      string `json:"message"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type Group struct {
	Category    *Category     `json:"category" yaml:",inline"` // 使用 inline 标签将 Category 字段内联
	Collections []*Collection `json:"collections"`
}

type IDRequest struct {
	ID int `json:"id"`
}

type IDResponse struct {
	ID int `json:"id"`
}

type Model struct {
	ID        uint   `json:"-" gorm:"primarykey"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type Record struct {
	Model
	IP string
}

type TagRequest struct {
	Tags []string `json:"tags"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}
