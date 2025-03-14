// Create a new dashboard with hostmap widget

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.Dashboard{
		Title:       "Example-Dashboard",
		Description: *datadog.NewNullableString(nil),
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 22,
				},
				Definition: datadogV1.WidgetDefinition{
					HostMapWidgetDefinition: &datadogV1.HostMapWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP,
						Requests: datadogV1.HostMapWidgetDefinitionRequests{
							Fill: &datadogV1.HostMapRequest{
								Q: datadog.PtrString("avg:system.cpu.user{*} by {host}"),
							},
						},
						NodeType:      datadogV1.WIDGETNODETYPE_HOST.Ptr(),
						NoMetricHosts: datadog.PtrBool(true),
						NoGroupHosts:  datadog.PtrBool(true),
						Style: &datadogV1.HostMapWidgetDefinitionStyle{
							Palette:     datadog.PtrString("green_to_orange"),
							PaletteFlip: datadog.PtrBool(false),
						},
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
		NotifyList:        *datadog.NewNullableList(&[]string{}),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
