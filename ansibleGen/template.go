package ansibleGen

import "time"

type projectTemplate struct {
	url         string
	lastUpdated time.Time
}

var templates []projectTemplate

func loadTemplatesFromCache() {

}

func updateTemplatesCache() {

}

func saveToTemplatesCache() {

}
