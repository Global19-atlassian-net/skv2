// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	discovery_multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/discovery.multicluster.solo.io/v1alpha1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for AwsDiscoveryClient from Clientset
func AwsDiscoveryClientFromClientsetProvider(clients discovery_multicluster_solo_io_v1alpha1.Clientset) discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryClient {
	return clients.AwsDiscoveries()
}

// Provider for AwsDiscovery Client from Client
func AwsDiscoveryClientProvider(client client.Client) discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryClient {
	return discovery_multicluster_solo_io_v1alpha1.NewAwsDiscoveryClient(client)
}

type AwsDiscoveryClientFactory func(client client.Client) discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryClient

func AwsDiscoveryClientFactoryProvider() AwsDiscoveryClientFactory {
	return AwsDiscoveryClientProvider
}

type AwsDiscoveryClientFromConfigFactory func(cfg *rest.Config) (discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryClient, error)

func AwsDiscoveryClientFromConfigFactoryProvider() AwsDiscoveryClientFromConfigFactory {
	return func(cfg *rest.Config) (discovery_multicluster_solo_io_v1alpha1.AwsDiscoveryClient, error) {
		clients, err := discovery_multicluster_solo_io_v1alpha1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AwsDiscoveries(), nil
	}
}
