package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New()
}
