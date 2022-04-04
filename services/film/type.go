package film

type Film struct {
    Name string `json:"name"`
}

type BaseRes struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
