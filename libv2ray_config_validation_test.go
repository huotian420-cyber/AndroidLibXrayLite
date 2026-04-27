package libv2ray

import (
	"fmt"
	"strings"
	"testing"

	core "github.com/xtls/xray-core/core"
	coreserial "github.com/xtls/xray-core/infra/conf/serial"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func TestVlessXhttpOutboundWithMlkemBuildsAndStarts(t *testing.T) {
	const clientKey = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

	configJSON := fmt.Sprintf(`{
  "log": {
    "loglevel": "warning"
  },
  "outbounds": [
    {
      "protocol": "vless",
      "tag": "proxy",
      "settings": {
        "vnext": [
          {
            "address": "example.com",
            "port": 443,
            "users": [
              {
                "id": "15a97905-a451-4c93-bd4c-e16885cbc807",
                "encryption": "mlkem768x25519plus.native.0rtt.%s"
              }
            ]
          }
        ]
      },
      "streamSettings": {
        "network": "xhttp",
        "security": "tls",
        "xhttpSettings": {
          "host": "example.com",
          "path": "/ej45ditxjo",
          "mode": "auto",
          "extra": {
            "xPaddingBytes": "1000-2000",
            "xmux": {
              "maxConcurrency": "16-32",
              "cMaxReuseTimes": 0,
              "hMaxRequestTimes": "600-900",
              "hMaxReusableSecs": "1800-3000",
              "hKeepAlivePeriod": 0
            }
          }
        },
        "tlsSettings": {
          "serverName": "example.com",
          "fingerprint": "chrome",
          "alpn": ["h2", "http/1.1"],
          "allowInsecure": false
        }
      }
    }
  ]
}`, clientKey)

	config, err := coreserial.LoadJSONConfig(strings.NewReader(configJSON))
	if err != nil {
		t.Fatalf("expected config to load, got %v", err)
	}

	instance, err := core.New(config)
	if err != nil {
		t.Fatalf("expected core.New to accept VLESS+XHTTP+MLKEM config, got %v", err)
	}
	defer instance.Close()

	if err := instance.Start(); err != nil {
		t.Fatalf("expected core to start with VLESS+XHTTP+MLKEM config, got %v", err)
	}
}
