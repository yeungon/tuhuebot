package database

func Register() {

}

// func Connect() {
// 	XATA_API_KEY := config.Get().XATA_API_KEY
// 	workspaceCli, err := xata.NewWorkspacesClient(xata.WithAPIKey(XATA_API_KEY))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := workspaceCli.List(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%#v\n", *resp.Workspaces[0])

// 	item := *resp.Workspaces[0]
// 	fmt.Printf("%s\n", item.Role.String())

// }
