module models

go 1.13

require sample/models v0.0.0

replace sample/models => ./models

require sample/routes v0.0.0

replace sample/routes => ./routes

require (
	github.com/gorilla/mux v1.7.4
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
