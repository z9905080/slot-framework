package entity

// SlotGameInput GameInput
type SlotGameInput struct {
	PlayLine      uint   `json:"play_line"`       // 幾線
	WinLineCount  uint   `json:"win_line_count"`  // 贏分線數量
	ReelPosition  []uint `json:"reel_position"`   // 滾輪位置
	UserID        int    `json:"user_id"`         // 使用者ID
	FreeGameTimes int    `json:"free_game_times"` // 目前FreeGame的場次
	Demo          bool   `json:"-"`               // 是否為Demo
}
