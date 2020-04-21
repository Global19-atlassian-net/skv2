package v1

import (
	"context"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "k8s.io/api/admissionregistration/v1"
)

// clienset for the admissionregistration.k8s.io/v1 APIs
type Clientset interface {
	// clienset for the admissionregistration.k8s.io/v1/v1 APIs
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the admissionregistration.k8s.io/v1/v1 APIs
func (c *clientSet) ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClient {
	return NewValidatingWebhookConfigurationClient(c.client)
}

// Reader knows how to read and list ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationReader interface {
	// Get retrieves a ValidatingWebhookConfiguration for the given object key
	GetValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey) (*ValidatingWebhookConfiguration, error)

	// List retrieves list of ValidatingWebhookConfigurations for a given namespace and list options.
	ListValidatingWebhookConfiguration(ctx context.Context, opts ...client.ListOption) (*ValidatingWebhookConfigurationList, error)
}

// Writer knows how to create, delete, and update ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationWriter interface {
	// Create saves the ValidatingWebhookConfiguration object.
	CreateValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.CreateOption) error

	// Delete deletes the ValidatingWebhookConfiguration object.
	DeleteValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ValidatingWebhookConfiguration object.
	UpdateValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.UpdateOption) error

	// Patch patches the given ValidatingWebhookConfiguration object.
	PatchValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ValidatingWebhookConfiguration objects matching the given options.
	DeleteAllOfValidatingWebhookConfiguration(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a ValidatingWebhookConfiguration object.
type ValidatingWebhookConfigurationStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ValidatingWebhookConfiguration object.
	UpdateValidatingWebhookConfigurationStatus(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.UpdateOption) error

	// Patch patches the given ValidatingWebhookConfiguration object's subresource.
	PatchValidatingWebhookConfigurationStatus(ctx context.Context, obj *ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationClient interface {
	ValidatingWebhookConfigurationReader
	ValidatingWebhookConfigurationWriter
	ValidatingWebhookConfigurationStatusWriter
}

type validatingWebhookConfigurationClient struct {
	client client.Client
}

func NewValidatingWebhookConfigurationClient(client client.Client) *validatingWebhookConfigurationClient {
	return &validatingWebhookConfigurationClient{client: client}
}

func (c *validatingWebhookConfigurationClient) GetValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey) (*ValidatingWebhookConfiguration, error) {
	obj := &ValidatingWebhookConfiguration{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *validatingWebhookConfigurationClient) ListValidatingWebhookConfiguration(ctx context.Context, opts ...client.ListOption) (*ValidatingWebhookConfigurationList, error) {
	list := &ValidatingWebhookConfigurationList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *validatingWebhookConfigurationClient) CreateValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) DeleteValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ValidatingWebhookConfiguration{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) UpdateValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) PatchValidatingWebhookConfiguration(ctx context.Context, obj *ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *validatingWebhookConfigurationClient) DeleteAllOfValidatingWebhookConfiguration(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ValidatingWebhookConfiguration{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) UpdateValidatingWebhookConfigurationStatus(ctx context.Context, obj *ValidatingWebhookConfiguration, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) PatchValidatingWebhookConfigurationStatus(ctx context.Context, obj *ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
