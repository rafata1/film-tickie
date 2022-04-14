package cinema

type BaseRes struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
