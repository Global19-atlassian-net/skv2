package v1

import (
	"context"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "k8s.io/api/rbac/v1"
)

// clienset for the rbac.authorization.k8s.io/v1 APIs
type Clientset interface {
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	Roles() RoleClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	RoleBindings() RoleBindingClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	ClusterRoles() ClusterRoleClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	ClusterRoleBindings() ClusterRoleBindingClient
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

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) Roles() RoleClient {
	return NewRoleClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) RoleBindings() RoleBindingClient {
	return NewRoleBindingClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) ClusterRoles() ClusterRoleClient {
	return NewClusterRoleClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) ClusterRoleBindings() ClusterRoleBindingClient {
	return NewClusterRoleBindingClient(c.client)
}

// Reader knows how to read and list Roles.
type RoleReader interface {
	// Get retrieves a Role for the given object key
	GetRole(ctx context.Context, key client.ObjectKey) (*Role, error)

	// List retrieves list of Roles for a given namespace and list options.
	ListRole(ctx context.Context, opts ...client.ListOption) (*RoleList, error)
}

// Writer knows how to create, delete, and update Roles.
type RoleWriter interface {
	// Create saves the Role object.
	CreateRole(ctx context.Context, obj *Role, opts ...client.CreateOption) error

	// Delete deletes the Role object.
	DeleteRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Role object.
	UpdateRole(ctx context.Context, obj *Role, opts ...client.UpdateOption) error

	// Patch patches the given Role object.
	PatchRole(ctx context.Context, obj *Role, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Role objects matching the given options.
	DeleteAllOfRole(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a Role object.
type RoleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Role object.
	UpdateRoleStatus(ctx context.Context, obj *Role, opts ...client.UpdateOption) error

	// Patch patches the given Role object's subresource.
	PatchRoleStatus(ctx context.Context, obj *Role, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Roles.
type RoleClient interface {
	RoleReader
	RoleWriter
	RoleStatusWriter
}

type roleClient struct {
	client client.Client
}

func NewRoleClient(client client.Client) *roleClient {
	return &roleClient{client: client}
}

func (c *roleClient) GetRole(ctx context.Context, key client.ObjectKey) (*Role, error) {
	obj := &Role{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *roleClient) ListRole(ctx context.Context, opts ...client.ListOption) (*RoleList, error) {
	list := &RoleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *roleClient) CreateRole(ctx context.Context, obj *Role, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *roleClient) DeleteRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Role{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *roleClient) UpdateRole(ctx context.Context, obj *Role, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *roleClient) PatchRole(ctx context.Context, obj *Role, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *roleClient) DeleteAllOfRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Role{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *roleClient) UpdateRoleStatus(ctx context.Context, obj *Role, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *roleClient) PatchRoleStatus(ctx context.Context, obj *Role, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list RoleBindings.
type RoleBindingReader interface {
	// Get retrieves a RoleBinding for the given object key
	GetRoleBinding(ctx context.Context, key client.ObjectKey) (*RoleBinding, error)

	// List retrieves list of RoleBindings for a given namespace and list options.
	ListRoleBinding(ctx context.Context, opts ...client.ListOption) (*RoleBindingList, error)
}

// Writer knows how to create, delete, and update RoleBindings.
type RoleBindingWriter interface {
	// Create saves the RoleBinding object.
	CreateRoleBinding(ctx context.Context, obj *RoleBinding, opts ...client.CreateOption) error

	// Delete deletes the RoleBinding object.
	DeleteRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given RoleBinding object.
	UpdateRoleBinding(ctx context.Context, obj *RoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given RoleBinding object.
	PatchRoleBinding(ctx context.Context, obj *RoleBinding, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all RoleBinding objects matching the given options.
	DeleteAllOfRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a RoleBinding object.
type RoleBindingStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given RoleBinding object.
	UpdateRoleBindingStatus(ctx context.Context, obj *RoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given RoleBinding object's subresource.
	PatchRoleBindingStatus(ctx context.Context, obj *RoleBinding, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on RoleBindings.
type RoleBindingClient interface {
	RoleBindingReader
	RoleBindingWriter
	RoleBindingStatusWriter
}

type roleBindingClient struct {
	client client.Client
}

func NewRoleBindingClient(client client.Client) *roleBindingClient {
	return &roleBindingClient{client: client}
}

func (c *roleBindingClient) GetRoleBinding(ctx context.Context, key client.ObjectKey) (*RoleBinding, error) {
	obj := &RoleBinding{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *roleBindingClient) ListRoleBinding(ctx context.Context, opts ...client.ListOption) (*RoleBindingList, error) {
	list := &RoleBindingList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *roleBindingClient) CreateRoleBinding(ctx context.Context, obj *RoleBinding, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *roleBindingClient) DeleteRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &RoleBinding{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *roleBindingClient) UpdateRoleBinding(ctx context.Context, obj *RoleBinding, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *roleBindingClient) PatchRoleBinding(ctx context.Context, obj *RoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *roleBindingClient) DeleteAllOfRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &RoleBinding{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *roleBindingClient) UpdateRoleBindingStatus(ctx context.Context, obj *RoleBinding, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *roleBindingClient) PatchRoleBindingStatus(ctx context.Context, obj *RoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ClusterRoles.
type ClusterRoleReader interface {
	// Get retrieves a ClusterRole for the given object key
	GetClusterRole(ctx context.Context, name string) (*ClusterRole, error)

	// List retrieves list of ClusterRoles for a given namespace and list options.
	ListClusterRole(ctx context.Context, opts ...client.ListOption) (*ClusterRoleList, error)
}

// Writer knows how to create, delete, and update ClusterRoles.
type ClusterRoleWriter interface {
	// Create saves the ClusterRole object.
	CreateClusterRole(ctx context.Context, obj *ClusterRole, opts ...client.CreateOption) error

	// Delete deletes the ClusterRole object.
	DeleteClusterRole(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given ClusterRole object.
	UpdateClusterRole(ctx context.Context, obj *ClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRole object.
	PatchClusterRole(ctx context.Context, obj *ClusterRole, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ClusterRole objects matching the given options.
	DeleteAllOfClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a ClusterRole object.
type ClusterRoleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ClusterRole object.
	UpdateClusterRoleStatus(ctx context.Context, obj *ClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRole object's subresource.
	PatchClusterRoleStatus(ctx context.Context, obj *ClusterRole, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ClusterRoles.
type ClusterRoleClient interface {
	ClusterRoleReader
	ClusterRoleWriter
	ClusterRoleStatusWriter
}

type clusterRoleClient struct {
	client client.Client
}

func NewClusterRoleClient(client client.Client) *clusterRoleClient {
	return &clusterRoleClient{client: client}
}

func (c *clusterRoleClient) GetClusterRole(ctx context.Context, name string) (*ClusterRole, error) {
	obj := &ClusterRole{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *clusterRoleClient) ListClusterRole(ctx context.Context, opts ...client.ListOption) (*ClusterRoleList, error) {
	list := &ClusterRoleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clusterRoleClient) CreateClusterRole(ctx context.Context, obj *ClusterRole, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *clusterRoleClient) DeleteClusterRole(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &ClusterRole{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *clusterRoleClient) UpdateClusterRole(ctx context.Context, obj *ClusterRole, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *clusterRoleClient) PatchClusterRole(ctx context.Context, obj *ClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *clusterRoleClient) DeleteAllOfClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ClusterRole{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *clusterRoleClient) UpdateClusterRoleStatus(ctx context.Context, obj *ClusterRole, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *clusterRoleClient) PatchClusterRoleStatus(ctx context.Context, obj *ClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ClusterRoleBindings.
type ClusterRoleBindingReader interface {
	// Get retrieves a ClusterRoleBinding for the given object key
	GetClusterRoleBinding(ctx context.Context, name string) (*ClusterRoleBinding, error)

	// List retrieves list of ClusterRoleBindings for a given namespace and list options.
	ListClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*ClusterRoleBindingList, error)
}

// Writer knows how to create, delete, and update ClusterRoleBindings.
type ClusterRoleBindingWriter interface {
	// Create saves the ClusterRoleBinding object.
	CreateClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, opts ...client.CreateOption) error

	// Delete deletes the ClusterRoleBinding object.
	DeleteClusterRoleBinding(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given ClusterRoleBinding object.
	UpdateClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRoleBinding object.
	PatchClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ClusterRoleBinding objects matching the given options.
	DeleteAllOfClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a ClusterRoleBinding object.
type ClusterRoleBindingStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ClusterRoleBinding object.
	UpdateClusterRoleBindingStatus(ctx context.Context, obj *ClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRoleBinding object's subresource.
	PatchClusterRoleBindingStatus(ctx context.Context, obj *ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ClusterRoleBindings.
type ClusterRoleBindingClient interface {
	ClusterRoleBindingReader
	ClusterRoleBindingWriter
	ClusterRoleBindingStatusWriter
}

type clusterRoleBindingClient struct {
	client client.Client
}

func NewClusterRoleBindingClient(client client.Client) *clusterRoleBindingClient {
	return &clusterRoleBindingClient{client: client}
}

func (c *clusterRoleBindingClient) GetClusterRoleBinding(ctx context.Context, name string) (*ClusterRoleBinding, error) {
	obj := &ClusterRoleBinding{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *clusterRoleBindingClient) ListClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*ClusterRoleBindingList, error) {
	list := &ClusterRoleBindingList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clusterRoleBindingClient) CreateClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) DeleteClusterRoleBinding(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &ClusterRoleBinding{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) UpdateClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) PatchClusterRoleBinding(ctx context.Context, obj *ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *clusterRoleBindingClient) DeleteAllOfClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ClusterRoleBinding{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) UpdateClusterRoleBindingStatus(ctx context.Context, obj *ClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) PatchClusterRoleBindingStatus(ctx context.Context, obj *ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
