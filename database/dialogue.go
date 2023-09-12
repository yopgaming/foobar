package database

var MessageType = map[string]int8{
	"User":         1,
	"Trader":       2,
	"Auction":      3,
	"Flea":         4,
	"Admin":        5,
	"Group":        6,
	"System":       7,
	"Insurance":    8,
	"Global":       9,
	"QuestStart":   10,
	"QuestFail":    11,
	"QuestSuccess": 12,
	"GiftMessage":  13,
	"Support":      14,
}

type Dialog struct {
	ID             string          `json:"_id"`
	Type           int8            `json:"type"`
	Messages       []DialogMessage `json:"messages"`
	Pinned         bool            `json:"pinned"`
	New            int8            `json:"new"`
	AttachmentsNew int8            `json:"attachmentsNew"`
}

type DialogMessage struct {
	ID                  string        `json:"_id"`
	UID                 string        `json:"uid"`
	Type                int8          `json:"type"`
	DT                  int64         `json:"dt"`
	Text                string        `json:"text"`
	TemplateID          string        `json:"templateId,omitempty"`
	HasRewards          bool          `json:"hasRewards"`
	RewardCollected     bool          `json:"rewardCollected"`
	SystemData          string        `json:"systemData,omitempty"`
	ProfileChangeEvents []interface{} `json:"profileChangeEvents,omitempty"`
}

type DialogueDetails struct {
	RecipientID                    string        `json:"recipientId"`
	Sender                         int8          `json:"sender"`
	DialogType                     int8          `json:"dialogType"`
	Trader                         string        `json:"trader"`
	TemplateID                     string        `json:"template"`
	Items                          []interface{} `json:"items,omitempty"`
	ItemsMaxStorageLifetimeSeconds int           `json:"itemsMaxStorageLifetimeSeconds,omitempty"`
}
