package test

var (
	// JSON ...
	JSON = []byte(`
		{
			"bundles": [{
				"path": "~/testing_gis",
				"zones": [{
						"user": "hendricius",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"pizza-dough",
							"the-bread-code"
						]
					},
					{
						"user": "cocktails-for-programmers",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"cocktails-for-programmers"
						]
					},
					{
						"user": "rochacbruno",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"vegan_recipes"
						]
					},
					{
						"user": "niw",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"ramen"
						]
					}
				]
			}]
		}
`)
)
