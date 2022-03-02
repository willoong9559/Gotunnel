package gotunnel

const (
	MaxPoolCap    = 10
	PoolTimeoutMS = 150000
)

type TunnelClient struct {
	listenAddr string
	serverAddr string
	obfsType   string
	obfsParam  string
	psk        string
}

func (s *TunnelClient) Close() {
	s.Close()
}

func newClient(listenAddr, serverAddr, obfsType, obfsParam, psk string) (*TunnelClient, error) {

	if obfsParam == "" {
		obfsParam = "www.bing.com"
	}

	//cipher := aead.NewPskCipher([]byte(psk))
	s := &TunnelClient{
		listenAddr: listenAddr,
		serverAddr: serverAddr,
		obfsType:   obfsType,
		obfsParam:  obfsParam,
		psk:        psk,
	}

	//p, err := newPool(MaxPoolCap, PoolTimeoutMS, sc.newSession)

	return s, nil
}
