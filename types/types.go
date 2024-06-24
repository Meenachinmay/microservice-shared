package types

type EnquiryPayload struct {
	UserID           int32  `json:"user_id"`
	AvailabelTimings string `json:"available_timings"`
	Email            string `json:"email"`
	PreferredMethod  string `json:"preferred_method"`
	PropertyID       int32  `json:"property_id"`
	PropertyName     string `json:"property_name"`
	PropertyLocation string `json:"property_location"`
	FudousanID       int32  `json:"fudousan_id"`
}

type LogPayload struct {
	ServiceName string
	LogData     string
}
