// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = ReadFunction{}
)

func NewReadFunction() function.Function {
	return ReadFunction{}
}

type ReadFunction struct{}

func (r ReadFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "read"
}

func (r ReadFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Read .env file(s) and unmarshal into a map",
		Description: "Read .env file(s) and unmarshal into a map",
		Parameters:  []function.Parameter{},
		VariadicParameter: function.StringParameter{
			Name:        "path",
			Description: "Path of .env file(s)",
		},
		Return: function.MapReturn{
			ElementType: types.StringType,
		},
	}
}

func (r ReadFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var paths []string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &paths))

	if resp.Error != nil {
		return
	}

	myEnv, err := godotenv.Read(paths...)

	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("Failed to read paths, error:\n %v", err))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, myEnv))
}
