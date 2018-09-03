package device

type config struct {
	Instance   int        `json:"instance"`
	Latency    int        `json:"latency"`
	Name       string     `json:"name"`
	Volume     volume     `json:"volume"`
	Connected  bool       `json:"connected"`
	Host       host       `json:"host"`
	LastSeen   lastSeen   `json:"lastSeen"`
	Snapclient snapclient `json:"sanpclient"`
}
type client struct {
	Config config `json:"config"`
}
type group struct {
	Clients []client `json:"clients"`
}
type volume struct {
	Muted    bool `json:"muted"`
	Perecent int  `json:"percent"`
}

type host struct {
	Arch string `json:"arch"`
	Ip   string `json:"ip"`
	Mac  string `json:"mac"`
	Name string `json:"name"`
	Os   string `json:"os"`
}

type lastSeen struct {
	Sec  int `json:"sec"`
	Usec int `json:"usec"`
}
type snapclient struct {
	Name            string `json:"name"`
	ProtocolVersion int16  `json:"protocolVersion"`
	Version         string `json:"version"`
}

type stream struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Uri    uri    `json:"uri"`
}
type uri struct {
	Fragment string `json:"fragment"`
	Host     string `json:"host"`
	Path     string `json:"path"`
	Query    query  `json:"query"`
	Raw      string `json:"raw"`
	Scheme   string `json:"scheme"`
}

type query struct {
	BufferMs     string `json:"buffer_ms"`
	Codec        string `json:"codec"`
	Name         string `json:"name"`
	Sampleformat string `json:"sampleformat"`
}
