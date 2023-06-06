package entity

import "github.com/shopspring/decimal"

// SlotWinLine 贏分線資訊
type SlotWinLine struct {
	WinType       int             `json:"win_type"`               // 贏分線的種類
	LineNo        int             `json:"line_no"`                // 第幾線
	Credit        uint            `json:"credit"`                 // 押注
	Multiply      uint            `json:"multiply"`               // 倍率
	SymbolID      uint            `json:"symbol_id"`              // Symbol編號
	Count         uint            `json:"count"`                  // 幾連線
	IsWinPosition [][]bool        `json:"is_win_position"`        // 這個位置是否有贏分線
	PayComboID    uint            `json:"pay_combo_id,omitempty"` // 賠付ID
	WinPoint      decimal.Decimal `json:"win_point"`              // 每條線實際贏分
	WinLineCount  int             `json:"win_line_count"`         // 贏分線數量(位的遊戲才會大於1)
}
