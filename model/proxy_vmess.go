package model

type HTTPOptions struct {
	Method  string              `proxy:"method,omitempty"`
	Path    []string            `proxy:"path,omitempty"`
	Headers map[string][]string `proxy:"headers,omitempty"`
}

type HTTP2Options struct {
	Host []string `proxy:"host,omitempty"`
	Path string   `proxy:"path,omitempty"`
}

type GrpcOptions struct {
	GrpcServiceName string `proxy:"grpc-service-name,omitempty"`
}

type RealityOptions struct {
	PublicKey string `proxy:"public-key"`
	ShortID   string `proxy:"short-id"`
}

type WSOptions struct {
	Path                string            `proxy:"path,omitempty"`
	Headers             map[string]string `proxy:"headers,omitempty"`
	MaxEarlyData        int               `proxy:"max-early-data,omitempty"`
	EarlyDataHeaderName string            `proxy:"early-data-header-name,omitempty"`
}

type Vmess struct {
	Type                string         `yaml:"type"`
	Name                string         `yaml:"name"`
	Server              string         `yaml:"server"`
	Port                int            `yaml:"port"`
	UUID                string         `yaml:"uuid"`
	AlterID             int            `yaml:"alterId"`
	Cipher              string         `yaml:"cipher"`
	UDP                 bool           `yaml:"udp,omitempty"`
	Network             string         `yaml:"network,omitempty"`
	TLS                 bool           `yaml:"tls,omitempty"`
	ALPN                []string       `yaml:"alpn,omitempty"`
	SkipCertVerify      bool           `yaml:"skip-cert-verify,omitempty"`
	Fingerprint         string         `yaml:"fingerprint,omitempty"`
	ServerName          string         `yaml:"servername,omitempty"`
	RealityOpts         RealityOptions `yaml:"reality-opts,omitempty"`
	HTTPOpts            HTTPOptions    `yaml:"http-opts,omitempty"`
	HTTP2Opts           HTTP2Options   `yaml:"h2-opts,omitempty"`
	GrpcOpts            GrpcOptions    `yaml:"grpc-opts,omitempty"`
	WSOpts              WSOptions      `yaml:"ws-opts,omitempty"`
	PacketAddr          bool           `yaml:"packet-addr,omitempty"`
	XUDP                bool           `yaml:"xudp,omitempty"`
	PacketEncoding      string         `yaml:"packet-encoding,omitempty"`
	GlobalPadding       bool           `yaml:"global-padding,omitempty"`
	AuthenticatedLength bool           `yaml:"authenticated-length,omitempty"`
	ClientFingerprint   string         `yaml:"client-fingerprint,omitempty"`
}

func ProxyToVmess(p Proxy) Vmess {
	return Vmess{
		Type:                "vmess",
		Name:                p.Name,
		Server:              p.Server,
		Port:                p.Port,
		UUID:                p.UUID,
		AlterID:             p.AlterID,
		Cipher:              p.Cipher,
		UDP:                 p.UDP,
		Network:             p.Network,
		TLS:                 p.TLS,
		ALPN:                p.Alpn,
		SkipCertVerify:      p.SkipCertVerify,
		Fingerprint:         p.Fingerprint,
		ServerName:          p.Servername,
		RealityOpts:         p.RealityOpts,
		HTTPOpts:            p.HTTPOpts,
		HTTP2Opts:           p.HTTP2Opts,
		GrpcOpts:            p.GrpcOpts,
		WSOpts:              p.WSOpts,
		PacketAddr:          p.PacketAddr,
		XUDP:                p.XUDP,
		PacketEncoding:      p.PacketEncoding,
		GlobalPadding:       p.GlobalPadding,
		AuthenticatedLength: p.AuthenticatedLength,
		ClientFingerprint:   p.ClientFingerprint,
	}
}
