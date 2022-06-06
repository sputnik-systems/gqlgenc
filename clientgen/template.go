package clientgen

import (
	"fmt"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	gqlgencConfig "github.com/sputnik-systems/gqlgenc/config"
)

func RenderTemplate(cfg *config.Config, query *Query, mutation *Mutation, fragments []*Fragment, operations []*Operation, operationResponses []*OperationResponse, generateCfg *gqlgencConfig.GenerateConfig, client config.PackageConfig) error {
	if err := templates.Render(templates.Options{
		PackageName: client.Package,
		Filename:    client.Filename,
		Data: map[string]interface{}{
			"Query":               query,
			"Mutation":            mutation,
			"Fragment":            fragments,
			"Operation":           operations,
			"OperationResponse":   operationResponses,
			"GenerateClient":      generateCfg.ShouldGenerateClient(),
			"ClientInterfaceName": generateCfg.ClientInterfaceName,
		},
		Packages:   cfg.Packages,
		PackageDoc: "// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.\n",
	}); err != nil {
		return fmt.Errorf("%s generating failed: %w", client.Filename, err)
	}

	return nil
}
