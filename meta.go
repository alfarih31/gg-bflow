package app

type meta struct {
	AppName        string `json:"app_name"`
	AppVersion     string `json:"app_version,omitempty"`
	AppDescription string `json:"app_description"`
}

var Meta = meta{
	AppName:        "GG Bflow",
	AppVersion:     Version,
	AppDescription: "Go GRPC Buffer Flow",
}
