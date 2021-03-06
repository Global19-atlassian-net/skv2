package codegen_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"

	. "github.com/solo-io/skv2/codegen"
)

var _ = Describe("Cmd", func() {
	It("generates controller code and manifests for a proto file", func() {

		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2_anyvendor.CreateDefaultMatchOptions(
				[]string{"codegen/test/*.proto"},
			),
			RenderProtos: true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Image: Image{
								Tag:        "v0.0.0",
								Repository: "painter",
								Registry:   "quay.io/solo-io",
								PullPolicy: "IfNotPresent",
							},
						},
						Args: []string{"foo"},
					},
				},
				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		err = exec.Command("goimports", "-w", "codegen/test/api").Run()
		Expect(err).NotTo(HaveOccurred())
	})

	It("generates CRD with validaton schema for a proto file", func() {

		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:         true,
					RenderValidationSchemas: true,
					ApiRoot:                 "codegen/test/api",
				},
			},
			AnyVendorConfig: skv2_anyvendor.CreateDefaultMatchOptions(
				[]string{"codegen/test/*.proto"},
			),
			RenderProtos: true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Image: Image{
								Tag:        "v0.0.0",
								Repository: "painter",
								Registry:   "quay.io/solo-io",
								PullPolicy: "IfNotPresent",
							},
						},
						Args: []string{"foo"},
					},
				},
				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())
	})
})
