package types

type EnquiryPayload struct {
	UserID     int32  `json:"user_id"`
	PropertyID int32  `json:"property_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	FudousanID int32  `json:"fudousan_id"`
