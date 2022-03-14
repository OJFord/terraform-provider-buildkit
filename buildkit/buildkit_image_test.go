package buildkit

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

func TestMe(t *testing.T) {
	cli, _ := client.NewClientWithOpts()
	result, _, _ := cli.ImageInspectWithRaw(context.TODO(), "sha256:3714362ff22a")
	print(result.ID)

}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestAccImage_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"buildkit": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: step1(),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func step1() string {
	return `
		provider buildkit {
		}
		resource buildkit_image this {
			context = "../examples/basic"
			dockerfile = "./Dockerfile"
			labels = {
				"paul" = "love"
			}
		}
	`
}
