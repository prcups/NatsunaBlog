// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DbBlogPost is the golang structure for table DB_BLOG_POST.
type DbBlogPost struct {
	Id         int    `json:"id"          orm:"ID"          ` //
	Title      string `json:"title"       orm:"TITLE"       ` //
	Timeline   string `json:"timeline"    orm:"TIMELINE"    ` //
	Time       string `json:"time"        orm:"TIME"        ` //
	Author     string `json:"author"      orm:"AUTHOR"      ` //
	Hid        int    `json:"hid"         orm:"HID"         ` //
	Content    string `json:"content"     orm:"CONTENT"     ` //
	Classify   string `json:"classify"    orm:"CLASSIFY"    ` //
	Tag        string `json:"tag"         orm:"TAG"         ` //
	VisitTimes int    `json:"visit_times" orm:"VISIT_TIMES" ` //
	Ontop      int    `json:"ontop"       orm:"ONTOP"       ` //
}
