// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

import "github.com/lib/pq"

type Announce struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content" gorm:"column:content;unique"`
	Date    string `json:"date"`
}

type AnnouncesData struct {
	Announces []*Announce `json:"announces"`
}

type AnyRequest struct {
}

type AuthResponse struct {
	Username string `json:"username" gorm:"column:username;unique"`
	Email    string `json:"email,optional"`
	Nickname string `json:"nickname,optional"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type Category struct {
	Model
	CID   string `json:"cid,optional" gorm:"column:cid"`
	Depth int    `json:"depth,optional,omitempty" yaml:",omitempty"`
	Title string `json:"title"`
	Order int    `json:"order,optional"`
	Path  string `json:"path,optional,omitempty" gorm:"column:path;unique"`
}

type Collection struct {
	Model
	CID         string    `json:"cid,optional,omitempty" gorm:"column:cid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description,optional"`
	Order       int       `json:"order,optional,omitempty" gorm:"column:order"`
	Path        string    `json:"path,optional,omitempty" gorm:"column:path;unique"`
	CategoryID  uint      `json:"-,omitempty" gorm:"column:category_id;index"` // 外键字段
	Category    *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Favicon     string    `json:"favicon,optional,omitempty"`
	Tags        pq.StringArray  `json:"tags,optional,omitempty" gorm:"type:text[]"`
	View        int       `json:"view,optional,omitempty"`
}

type CollectionFilter struct {
	Categories pq.StringArray `json:"categories,optional"`
	Page       int      `json:"page,optional"`
	Limit      int      `json:"limit,optional"`
}

type CollectionsData struct {
	Category *Category `json:"category"` // 使用 inline 标签将 Category 字段内联
	Groups   []*Group  `json:"groups"`
}

type CollectionsDataResponse struct {
	Datas []*CollectionsData `json:"datas"`
}

type CollectionsResponse struct {
	Collections []*Collection `json:"collections"`
}

type Comment struct {
	Model
	Nickname string     `json:"nickname"`
	Content  string     `json:"content" gorm:"column:content;unique"`
	Date     string     `json:"date,optional"`
	Replies  []*Comment `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"replies,optional"`
	ParentID *uint      `json:"parent_id,omitempty,optional"`
}

type CommentsResponse struct {
	Comments []*Comment `json:"comments"`
}

type Group struct {
	Category    *Category     `json:"category" yaml:",inline"` // 使用 inline 标签将 Category 字段内联
	Collections []*Collection `json:"collections"`
}

type IDRequest struct {
	ID string `json:"id"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type Model struct {
	ID        uint   `json:"-" gorm:"primarykey"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type SiteStats struct {
	LastDayVisitors int `json:"last_day_visitors"` // 昨日访问者数量
	TotalVisitors   int `json:"total_visitors"`    // 总访问者数量
	LinksCount      int `json:"links_count"`       // 链接数量
}

type Statistics struct {
	Model
	IP   string `json:"ip"`   // 访问者的 IP 地址
	Date string `json:"date"` // 访问日期和时间
}

type TagRequest struct {
	Tags pq.StringArray `json:"tags"`
}

type TagsResponse struct {
	Tags pq.StringArray `json:"tags"`
}

type User struct {
	Model
	Username string `json:"username" gorm:"column:username;unique"`
	Email    string `json:"email,optional"`
	Role     string `json:"role,optional"`
	Password string `json:"password"`
	Nickname string `json:"nickname,optional"`
}

type UserResponse struct {
	Users []User
}
