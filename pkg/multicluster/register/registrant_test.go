package register_test

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/rbac.authorization.k8s.io/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	mock_clientcmd "github.com/solo-io/skv2/pkg/generated/mocks/k8s/clientcmd"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	"github.com/solo-io/skv2/pkg/multicluster/register/internal"
	mock_internal "github.com/solo-io/skv2/pkg/multicluster/register/internal/mocks"
	k8s_core_types "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Registrant", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		clusterRBACBinder        *mock_internal.MockClusterRBACBinder
		clusterRbacBinderFactory internal.ClusterRBACBinderFactory
		secretClient             *mock_k8s_core_clients.MockSecretClient
		secretClientFactory      k8s_core_v1.SecretClientFromConfigFactory
		nsClient                 *mock_k8s_core_clients.MockNamespaceClient
		nsClientFactory          k8s_core_v1.NamespaceClientFromConfigFactory
		saClient                 *mock_k8s_core_clients.MockServiceAccountClient
		saClientFactory          k8s_core_v1.ServiceAccountClientFromConfigFactory
		clusterRoleClient        *mock_k8s_rbac_clients.MockClusterRoleClient
		clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory
		roleClient               *mock_k8s_rbac_clients.MockRoleClient
		roleClientFactory        rbac_v1.RoleClientFromConfigFactory
		clientConfig             *mock_clientcmd.MockClientConfig

		_, remoteCtx, clusterName, namespace = "cfg-path", "context", "cluster-name", "namespace"
		testErr                              = eris.New("hello")

		saName = "sa-name"

		saObjectKey = func() client.ObjectKey {
			return client.ObjectKey{
				Namespace: namespace,
				Name:      saName,
			}
		}
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		secretClient = mock_k8s_core_clients.NewMockSecretClient(ctrl)
		secretClientFactory = func(_ *rest.Config) (k8s_core_v1.SecretClient, error) {
			return secretClient, nil
		}
		nsClient = mock_k8s_core_clients.NewMockNamespaceClient(ctrl)
		nsClientFactory = func(_ *rest.Config) (k8s_core_v1.NamespaceClient, error) {
			return nsClient, nil
		}
		saClient = mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		saClientFactory = func(_ *rest.Config) (k8s_core_v1.ServiceAccountClient, error) {
			return saClient, nil
		}
		clusterRoleClient = mock_k8s_rbac_clients.NewMockClusterRoleClient(ctrl)
		clusterRoleClientFactory = func(_ *rest.Config) (rbac_v1.ClusterRoleClient, error) {
			return clusterRoleClient, nil
		}
		roleClient = mock_k8s_rbac_clients.NewMockRoleClient(ctrl)
		roleClientFactory = func(_ *rest.Config) (rbac_v1.RoleClient, error) {
			return roleClient, nil
		}

		clusterRBACBinder = mock_internal.NewMockClusterRBACBinder(ctrl)
		clusterRbacBinderFactory = func(_ clientcmd.ClientConfig) (internal.ClusterRBACBinder, error) {
			return clusterRBACBinder, nil
		}
		clientConfig = mock_clientcmd.NewMockClientConfig(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("EnsureRemoteServiceAccount", func() {

		It("will create if does not exist", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
			)

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.Options{
				ClusterName:     clusterName,
				Namespace:       namespace,
				RemoteNamespace: "remote",
				RemoteCtx:       remoteCtx,
			}

			saClient.EXPECT().
				GetServiceAccount(ctx, client.ObjectKey{
					Namespace: opts.RemoteNamespace,
					Name:      opts.ClusterName,
				}).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			expected := &k8s_core_types.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: opts.RemoteNamespace,
					Name:      opts.ClusterName,
				},
			}

			saClient.EXPECT().
				CreateServiceAccount(ctx, expected).
				Return(nil)

			sa, err := clusterRegistrant.EnsureRemoteServiceAccount(ctx, clientConfig, opts)
			Expect(err).NotTo(HaveOccurred())
			Expect(sa).To(Equal(expected))
		})

	})

	Context("CreateRemoteAccessToken", func() {

		It("can successfully upsert all roles, and cluster roles", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
			)

			// Set secret lookup opts to reduce testing time
			register.SecretLookupOpts = []retry.Option{
				retry.Delay(time.Nanosecond),
				retry.Attempts(2),
				retry.DelayType(retry.FixedDelay),
			}

			sa := saObjectKey()

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.RbacOptions{
				Options: register.Options{
					ClusterName: clusterName,
					Namespace:   namespace,
					RemoteCtx:   remoteCtx,
				},
				Roles: []*rbacv1.Role{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "r-1",
							Namespace: namespace,
						},
					},
				},
				ClusterRoles: []*rbacv1.ClusterRole{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "cr-1",
						},
					},
				},
				RoleBindings: []client.ObjectKey{
					{
						Namespace: "namespace",
						Name:      "rb-1",
					},
				},
				ClusterRoleBindings: []client.ObjectKey{
					{
						Namespace: "",
						Name:      "crb-1",
					},
				},
			}

			roleClient.EXPECT().
				UpsertRole(ctx, opts.Roles[0]).
				Return(nil)

			clusterRBACBinder.EXPECT().
				BindRoles(ctx, sa, append(opts.RoleBindings, client.ObjectKey{
					Name:      opts.Roles[0].GetName(),
					Namespace: opts.Roles[0].GetNamespace(),
				})).
				Return(nil)

			clusterRoleClient.EXPECT().
				UpsertClusterRole(ctx, opts.ClusterRoles[0]).
				Return(nil)

			clusterRBACBinder.EXPECT().
				BindClusterRoles(ctx, sa, append(opts.ClusterRoleBindings, client.ObjectKey{
					Name: opts.ClusterRoles[0].GetName(),
				})).Return(nil)

			token := "hello"
			saSecret := &k8s_core_types.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "sa-secret",
					Namespace: namespace,
				},
				Data: map[string][]byte{
					register.SecretTokenKey: []byte(token),
				},
			}
			expectedSa := &k8s_core_types.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sa.Name,
					Namespace: sa.Namespace,
				},
				Secrets: []k8s_core_types.ObjectReference{
					{
						Namespace: namespace,
						Name:      saSecret.GetName(),
					},
				},
			}

			saClient.EXPECT().
				GetServiceAccount(ctx, sa).
				Return(expectedSa, nil).
				Times(2)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: saSecret.GetNamespace(),
					Name:      saSecret.GetName(),
				}).
				Return(nil, testErr)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: saSecret.GetNamespace(),
					Name:      saSecret.GetName(),
				}).Return(saSecret, nil)

			returnedToken, err := clusterRegistrant.CreateRemoteAccessToken(ctx, clientConfig, sa, opts)
			Expect(err).NotTo(HaveOccurred())
			Expect(returnedToken).To(Equal(token))
		})

	})

	Context("RegisterClusterWithToken", func() {

		var (
			token = "token"
		)

		It("works", func() {

			clusterRegistrant := register.NewClusterRegistrant(
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   remoteCtx,
			}

			restCfg := &rest.Config{
				Host: "mock-host",
			}
			apiCfg := api.Config{
				Clusters: map[string]*api.Cluster{
					clusterName: {
						Server:                   "fake-server",
						CertificateAuthorityData: []byte("fake-ca-data"),
					},
				},
				Contexts: map[string]*api.Context{
					remoteCtx: {
						Cluster: clusterName,
					},
				},
				CurrentContext: remoteCtx,
			}

			clientConfig.EXPECT().
				ClientConfig().
				Return(restCfg, nil)

			clientConfig.EXPECT().
				RawConfig().
				Return(apiCfg, nil)

			nsClient.EXPECT().
				GetNamespace(ctx, namespace).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			nsClient.EXPECT().
				CreateNamespace(ctx, &k8s_core_types.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: namespace,
					},
				}).Return(nil)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: namespace,
					Name:      clusterName,
				}).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			secret, err := kubeconfig.ToSecret(namespace, clusterName, api.Config{
				Kind:        "Secret",
				APIVersion:  "kubernetes_core",
				Preferences: api.Preferences{},
				Clusters: map[string]*api.Cluster{
					clusterName: apiCfg.Clusters[clusterName],
				},
				AuthInfos: map[string]*api.AuthInfo{
					clusterName: {
						Token: token,
					},
				},
				Contexts: map[string]*api.Context{
					clusterName: {
						Cluster:  clusterName,
						AuthInfo: clusterName,
					},
				},
				CurrentContext: clusterName,
			})
			Expect(err).NotTo(HaveOccurred())

			secretClient.EXPECT().
				CreateSecret(ctx, secret).
				Return(nil)

			err = clusterRegistrant.RegisterClusterWithToken(ctx, clientConfig, token, opts)

			Expect(err).NotTo(HaveOccurred())
		})

		It("can override local cluster domain", func() {
			clusterDomainOverride := "test-override"

			clusterRegistrant := register.NewTestingRegistrant(
				clusterDomainOverride,
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   "kind-test",
			}

			restCfg := &rest.Config{
				Host: "mock-host",
			}
			apiCfg := api.Config{
				Clusters: map[string]*api.Cluster{
					clusterName: {
						Server:                   "http://localhost:9080",
						CertificateAuthorityData: []byte("fake-ca-data"),
					},
				},
				Contexts: map[string]*api.Context{
					opts.RemoteCtx: {
						Cluster: clusterName,
					},
				},
				CurrentContext: opts.RemoteCtx,
			}

			clientConfig.EXPECT().
				ClientConfig().
				Return(restCfg, nil)

			clientConfig.EXPECT().
				RawConfig().
				Return(apiCfg, nil)

			nsClient.EXPECT().
				GetNamespace(ctx, namespace).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			nsClient.EXPECT().
				CreateNamespace(ctx, &k8s_core_types.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: namespace,
					},
				}).Return(nil)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: namespace,
					Name:      clusterName,
				}).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			overwrittenApiConfig := apiCfg.DeepCopy()
			overwrittenApiConfig.Clusters[clusterName].Server = fmt.Sprintf("https://%s:9080", clusterDomainOverride)
			overwrittenApiConfig.Clusters[clusterName].InsecureSkipTLSVerify = true
			overwrittenApiConfig.Clusters[clusterName].CertificateAuthority = ""
			overwrittenApiConfig.Clusters[clusterName].CertificateAuthorityData = []byte("")

			secret, err := kubeconfig.ToSecret(namespace, clusterName, api.Config{
				Kind:        "Secret",
				APIVersion:  "kubernetes_core",
				Preferences: api.Preferences{},
				Clusters: map[string]*api.Cluster{
					clusterName: overwrittenApiConfig.Clusters[clusterName],
				},
				AuthInfos: map[string]*api.AuthInfo{
					clusterName: {
						Token: token,
					},
				},
				Contexts: map[string]*api.Context{
					clusterName: {
						Cluster:  clusterName,
						AuthInfo: clusterName,
					},
				},
				CurrentContext: clusterName,
			})
			Expect(err).NotTo(HaveOccurred())

			secretClient.EXPECT().
				CreateSecret(ctx, secret).
				Return(nil)
			err = clusterRegistrant.RegisterClusterWithToken(ctx, clientConfig, token, opts)

			Expect(err).NotTo(HaveOccurred())
		})

	})

})