package kubernetes

import (
	"io"

	"github.com/go-kit/log"

	"github.com/grafana/dskit/kv/codec"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	fake_rest "k8s.io/client-go/rest/fake"
)

func NewInMemoryClient(codec codec.Codec, logger log.Logger) (*Client, io.Closer) {
	client, err := newClient(&Config{}, codec, logger, nil, func() (string, kubernetes.Interface, rest.Interface, error) {
		return "", fake.NewSimpleClientset(), &fake_rest.RESTClient{}, nil
	})
	if err != nil {
		panic("error generating in memory client: " + err.Error())
	}

	return client, io.NopCloser(nil)
}
