package senate

import (
	tls_client "github.com/bogdanfinn/tls-client"
)

var client *tls_client.HttpClient

func Init(cli *tls_client.HttpClient) {
	client = cli
}
