package remote

type Host struct {
	Addr string `json:"addr"`
	Path string `json:"path"`
}

func NewHost() *Host {
	return &Host{}
}

func (h *Host) Parse(data map[string]any) {
	h.Addr = data["addr"].(string)
	h.Path = data["path"].(string)
}
