package flags_searcher

type FlagFound struct {
	Flag   string `json:"flag"`
	Founds []File `json:"founds"`
}

func Run(projectPath string, apiKey string) int {
	options := KPLOptions{BaseUri: "http://localhost.charlesproxy.com:8080", ApiKey: apiKey}
	client := NewClient(options)

	flags, err := client.GetListFlags()
	if err != nil {
		return 1
	}

	founds := make([]FlagFound, 0)

	for _, flag := range flags {
		flagFounds, err := FileSearcher(projectPath, flag)
		if err == 1 {
			return 1
		}

		founds = append(founds, FlagFound{
			Flag:   flag,
			Founds: flagFounds,
		})
	}

	return client.SaveFlagsInformation(founds)
}
