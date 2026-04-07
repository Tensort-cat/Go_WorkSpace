package request

type EnterGroupDirectlyRequest struct {
	GroupId   string `json:"group_id"`
	ContactId string `json:"contact_id"`
}
