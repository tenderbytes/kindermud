package v1

import (
	"net/http"

	"github.com/danielkrainas/gobag/api/describe"
)

var (
	VersionHeader = describe.Parameter{
		Name:        "Api-Version",
		Type:        "string",
		Description: "The build version of the server.",
		Format:      "<version>",
		Examples:    []string{"0.0.0-dev"},
	}

	hostHeader = describe.Parameter{
		Name:        "Host",
		Type:        "string",
		Description: "",
		Format:      "<hostname>",
		Examples:    []string{"kindermud.io"},
	}

	jsonContentLengthHeader = describe.Parameter{
		Name:        "Content-Length",
		Type:        "integer",
		Description: "Length of the JSON body.",
		Format:      "<length>",
	}
)

var (
	errorsBody = `{
	"errors:" [
	    {
            "code": <error code>,
            "message": <error message>,
            "detail": ...
        },
        ...
    ]
}`
)

var API = struct {
	Routes []describe.Route `json:"routes"`
}{
	Routes: routeDescriptors,
}

var routeDescriptors = []describe.Route{
	{
		Name:        RouteNameBase,
		Path:        "/v1",
		Entity:      "Base",
		Description: "Base V1 API route, can be used for lightweight health and version check.",
		Methods: []describe.Method{
			{
				Method:      "GET",
				Description: "Check that the server supports the V1 API.",
				Requests: []describe.Request{
					{
						Headers: []describe.Parameter{
							hostHeader,
						},

						Successes: []describe.Response{
							{
								Description: "The API implements the V1 protocol and is accessible.",
								StatusCode:  http.StatusOK,
								Headers: []describe.Parameter{
									jsonContentLengthHeader,
									VersionHeader,
								},
							},
						},

						Failures: []describe.Response{
							{
								Description: "The API does not support the V1 protocol.",
								StatusCode:  http.StatusNotFound,
								Headers: []describe.Parameter{
									VersionHeader,
								},
							},
						},
					},
				},
			},
		},
	},
}

var APIDescriptor map[string]describe.Route

func init() {
	APIDescriptor = make(map[string]describe.Route, len(routeDescriptors))
	for _, descriptor := range routeDescriptors {
		APIDescriptor[descriptor.Name] = descriptor
	}
}
