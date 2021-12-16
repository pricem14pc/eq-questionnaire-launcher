package surveys

import (
	"encoding/json"
	"log"
	"regexp"

	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/AreaHQ/jsonhal"
	"github.com/ONSdigital/eq-questionnaire-launcher/clients"
	"github.com/ONSdigital/eq-questionnaire-launcher/settings"
)

// LauncherSchema is a representation of a schema in the Launcher
type LauncherSchema struct {
	Name       string
	SurveyType string
	URL        string
}

// RegisterResponse is the response from the eq-survey-register request
type RegisterResponse struct {
	jsonhal.Hal
}

// Schemas is a list of Schema
type Schemas []Schema

// Schema is an available schema
type Schema struct {
	jsonhal.Hal
	Name string `json:"name"`
}

var eqIDFormTypeRegex = regexp.MustCompile(`^(?P<eq_id>[a-z0-9]+)_(?P<form_type>\w+)`)

// LauncherSchemaFromFilename creates a LauncherSchema record from a schema filename
func LauncherSchemaFromFilename(filename string, surveyType string) LauncherSchema {
	return LauncherSchema{
		Name:       filename,
		SurveyType: surveyType,
	}
}

// GetAvailableSchemas Gets the list of static schemas an joins them with any schemas from the eq-survey-register if defined
func GetAvailableSchemas() map[string][]LauncherSchema {
	runnerSchemas := getAvailableSchemasFromRunner()
	registerSchemas := getAvailableSchemasFromRegister()

	allSchemas := []LauncherSchema{}
	allSchemas = append(runnerSchemas, registerSchemas...)

	sort.Sort(ByFilename(allSchemas))

	schemasBySurveyType := map[string][]LauncherSchema{}
	for _, schema := range allSchemas {
		schemasBySurveyType[strings.Title(schema.SurveyType)] = append(schemasBySurveyType[strings.Title(schema.SurveyType)], schema)
	}

	return schemasBySurveyType
}

// ByFilename implements sort.Interface based on the Name field.
type ByFilename []LauncherSchema

func (a ByFilename) Len() int           { return len(a) }
func (a ByFilename) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByFilename) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getAvailableSchemasFromRegister() []LauncherSchema {

	schemaList := []LauncherSchema{}

	if settings.Get("SURVEY_REGISTER_URL") != "" {
		resp, err := clients.GetHTTPClient().Get(settings.Get("SURVEY_REGISTER_URL"))
		if err != nil {
			log.Fatal("Do: ", err)
			return []LauncherSchema{}
		}

		responseBody, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return schemaList
		}

		var registerResponse RegisterResponse
		if err := json.Unmarshal(responseBody, &registerResponse); err != nil {
			log.Print(err)
			return schemaList
		}

		var schemas Schemas

		schemasJSON, _ := json.Marshal(registerResponse.Embedded["schemas"])

		if err := json.Unmarshal(schemasJSON, &schemas); err != nil {
			log.Println(err)
		}

		for _, schema := range schemas {
			url := schema.Links["self"]
			schemaList = append(schemaList, LauncherSchema{
				Name:       schema.Name,
				URL:        url.Href,
				SurveyType: "Other",
			})
		}
	}

	return schemaList
}

func getAvailableSchemasFromRunner() []LauncherSchema {

	schemaList := []LauncherSchema{}

	hostURL := settings.Get("SURVEY_RUNNER_SCHEMA_URL")

	log.Printf("Survey Runner Schema URL: %s", hostURL)

	url := fmt.Sprintf("%s/schemas", hostURL)

	resp, err := clients.GetHTTPClient().Get(url)

	if err != nil {
		return []LauncherSchema{}
	}

	if resp.StatusCode != 200 {
		return []LauncherSchema{}
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return []LauncherSchema{}
	}

	var schemaMapResponse = map[string][]string{}

	if err := json.Unmarshal(responseBody, &schemaMapResponse); err != nil {
		log.Print(err)
		return []LauncherSchema{}
	}

	for surveyType, schemas := range schemaMapResponse {
		for _, schemaName := range schemas {
			schemaList = append(schemaList, LauncherSchemaFromFilename(schemaName, surveyType))
		}
	}

	return schemaList
}

// FindSurveyByName Finds the schema in the list of available schemas
func FindSurveyByName(name string) LauncherSchema {
	availableSchemas := GetAvailableSchemas()

	for _, schemasBySurveyType := range availableSchemas {
		for _, schema := range schemasBySurveyType {
			if schema.Name == name {
				return schema
			}
		}
	}

	panic("Schema not found")
}
