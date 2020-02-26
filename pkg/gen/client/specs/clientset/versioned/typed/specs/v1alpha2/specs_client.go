/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/deislabs/smi-sdk-go/pkg/apis/specs/v1alpha2"
	"github.com/deislabs/smi-sdk-go/pkg/gen/client/specs/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SpecsV1alpha2Interface interface {
	RESTClient() rest.Interface
	HTTPRouteGroupsGetter
	TCPRoutesGetter
}

// SpecsV1alpha2Client is used to interact with features provided by the specs.smi-spec.io group.
type SpecsV1alpha2Client struct {
	restClient rest.Interface
}

func (c *SpecsV1alpha2Client) HTTPRouteGroups(namespace string) HTTPRouteGroupInterface {
	return newHTTPRouteGroups(c, namespace)
}

func (c *SpecsV1alpha2Client) TCPRoutes(namespace string) TCPRouteInterface {
	return newTCPRoutes(c, namespace)
}

// NewForConfig creates a new SpecsV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*SpecsV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SpecsV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new SpecsV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SpecsV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SpecsV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *SpecsV1alpha2Client {
	return &SpecsV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SpecsV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
