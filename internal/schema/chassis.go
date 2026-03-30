package schema

type CreateChassisInput struct {
	FrameID       string `json:"frame_id" binding:"required"`
	SuspensionID  string `json:"suspension_id" binding:"required"`
	BreakSystemID string `json:"break_system_id" binding:"required"`
}

type CreateFrameInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateSuspensionInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateBreakSystemInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}
