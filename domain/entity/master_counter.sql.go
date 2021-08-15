package entity

type MasterCounter struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Counter   int32  `json:"counter"`
	Prefix    string `json:"prefix"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type GetCounterByCounterIDRow struct {
	Counter   int32  `json:"counter"`
	Prefix    string `json:"prefix"`
	UpdatedAt int64  `json:"updated_at"`
}

type UpdateCounterByCounterIDParams struct {
	Counter   int32 `json:"counter"`
	UpdatedAt int64 `json:"updated_at"`
	CounterID int32 `json:"counter_id"`
}
