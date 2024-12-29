// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure GodotenvProvider satisfies various provider interfaces.
var _ provider.Provider = &GodotenvProvider{}
var _ provider.ProviderWithFunctions = &GodotenvProvider{}

type GodotenvProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type GodotenvProviderModel struct{}

func (p *GodotenvProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "godotenv"
	resp.Version = p.version
}

func (p *GodotenvProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Unmarshal .env files into Terraform using godotenv - A Go port of Ruby's dotenv library.`,
		MarkdownDescription: `Unmarshal .env files into Terraform using [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library.

Say yes to parsing of diverse ` + "`.env`" + `formats including;

* unquoted values
* multiline values
* comments
* export declarations

##### example .env file
` + "```sh" + `
# comments are supported
PROFILES"local dev" # and trailing comments

# so are unquoted values
FOO=fizz buzz

# and export declarations
export BAR="foobar"

# multiline values too (both inline and newline separated)!
STOOGES="larry\ncurly\nmoe"
PGP_LINUS="-----BEGIN PGP PUBLIC KEY BLOCK-----
...lots of base64 encoded text...
-----END PGP PUBLIC KEY BLOCK-----"
` + "```",
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *GodotenvProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data GodotenvProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *GodotenvProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewReadFunction,
	}
}

func (p *GodotenvProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *GodotenvProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GodotenvProvider{
			version: version,
		}
	}
}
