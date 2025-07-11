package models

type Comment struct {
	CommentId int    `json:"commentId"`
	ProfileId int    `json:"profileId"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}
