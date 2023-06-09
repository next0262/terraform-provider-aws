// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package meta

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// @FrameworkDataSource
func newDataSourceDefaultTags(context.Context) (datasource.DataSourceWithConfigure, error) {
	d := &dataSourceDefaultTags{}
	d.SetMigratedFromPluginSDK(true)

	return d, nil
}

type dataSourceDefaultTags struct {
	framework.DataSourceWithConfigure
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSourceDefaultTags) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	response.TypeName = "aws_default_tags"
}

// Schema returns the schema for this data source.
func (d *dataSourceDefaultTags) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"tags": tftags.TagsAttributeComputedOnly(),
		},
	}
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSourceDefaultTags) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data dataSourceDefaultTagsData

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	defaultTagsConfig := d.Meta().DefaultTagsConfig
	ignoreTagsConfig := d.Meta().IgnoreTagsConfig
	tags := defaultTagsConfig.GetTags()

	data.ID = types.StringValue(d.Meta().Partition)
	data.Tags = flex.FlattenFrameworkStringValueMapLegacy(ctx, tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map())

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type dataSourceDefaultTagsData struct {
	ID   types.String `tfsdk:"id"`
	Tags types.Map    `tfsdk:"tags"`
}
