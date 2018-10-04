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

/*
	{
		"groups": [{
			"clients":[{
				"config":{
					"instance":2,
					"latency":6,
					"name":"123 456",
					"volume":{"muted":false,"percent":48}
				},
				"connected":true,
				"host":{
					"arch":"x86_64",
					"ip":"127.0.0.1",
					"mac":"00:21:6a:7d:74:fc",
					"name":"T400",
					"os":"Linux Mint 17.3 Rosa"
				},
				"id":"00:21:6a:7d:74:fc#2",
				"lastSeen":{"sec":1488025751,"usec":654777},"snapclient":{"name":"Snapclient","protocolVersion":2,"version":"0.10.0"}}],"id":"4dcc4e3b-c699-a04b-7f0c-8260d23c43e1","muted":false,"name":"","stream_id":"stream 2"}],"server":{"host":{"arch":"x86_64","ip":"","mac":"","name":"T400","os":"Linux Mint 17.3 Rosa"},"snapserver":{"controlProtocolVersion":1,"name":"Snapserver","protocolVersion":1,"version":"0.10.0"}},"streams":[{"id":"stream 1","status":"idle","uri":{"fragment":"","host":"","path":"/tmp/snapfifo","query":{"buffer_ms":"20","codec":"flac","name":"stream 1","sampleformat":"48000:16:2"},"raw":"pipe:///tmp/snapfifo?name=stream 1","scheme":"pipe"}},{"id":"stream 2","status":"idle","uri":{"fragment":"","host":"","path":"/tmp/snapfifo","query":{"buffer_ms":"20","codec":"flac","name":"stream 2","sampleformat":"48000:16:2"},"raw":"pipe:///tmp/snapfifo?name=stream 2","scheme":"pipe"}}]}
*/
