package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/mrf345/swaggomnia/models"
)

type Entity map[string][]models.Resource
type Entities map[string]Entity
type Security string

const (
	REQUEST_GROUP = "request_group"
	REQUEST       = "request"
	YAML_FORMAT   = "yaml"
	JSON_FORMAT   = "json"

	BasicAuth Security = "Basic Auth"
	Bearer    Security = "Bearer"
	ApiKey    Security = "API Key"
)

var groups map[string]string

type SortedEntity struct {
	Key    string
	Entity Entity
	Group  string
}

type Swagger struct {
	Config         SwaggerConfig
	Entities       Entities
	SortedEntities []SortedEntity

	IsBasic, IsBearer, IsApiKey bool
}

type SwaggerConfig struct {
	Title       string   `json:"title"`
	Version     string   `json:"version"`
	Host        string   `json:"host"`
	BasePath    string   `json:"basePath"`
	Schemes     string   `json:"schemes"`
	Description string   `json:"description"`
	Security    Security `json:"security,omitempty"`
}

func parse(insomnia models.Insomnia) Entities {
	groups = make(map[string]string)
	es := make(Entities)

	for _, r := range insomnia.Resources {
		if r.Type == REQUEST_GROUP {
			groups[r.ID] = r.Name
			es[r.ID] = make(map[string][]models.Resource, 0)
		}

		if r.Type == REQUEST {
			fetchVariables(&r)

			if es[r.ParentID] == nil {
				es[r.ParentID] = make(map[string][]models.Resource, 0)
			}

			es[r.ParentID][r.URL] = append(es[r.ParentID][r.URL], r)
		}
	}

	maps.DeleteFunc(es, func(k string, e Entity) bool {
		return len(e) == 0
	})

	return es
}

func readInsomniaExport(name string) (i models.Insomnia) {
	var raw []byte
	var err error

	if raw, err = os.ReadFile(name); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &i); err != nil {
		log.Fatal(err)
	}

	return
}

func readSwaggerConfig(name string) (cfg SwaggerConfig) {
	var raw []byte
	var err error

	if raw, err = os.ReadFile(name); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &cfg); err != nil {
		log.Fatal(err)
	}

	return
}

func (s *Swagger) Generate(insomniaFile, configFile, outputFormat string) {
	s.Config = readSwaggerConfig(configFile)
	s.Entities = parse(readInsomniaExport(insomniaFile))
	s.IsApiKey = s.Config.Security == ApiKey
	s.IsBasic = s.Config.Security == BasicAuth
	s.IsBearer = s.Config.Security == Bearer

	for k, v := range s.Entities {
		s.SortedEntities = append(s.SortedEntities, SortedEntity{
			Key: k, Entity: v, Group: groups[k],
		})
	}

	slices.SortFunc(s.SortedEntities, func(a, b SortedEntity) int {
		return strings.Compare(a.Group, b.Group)
	})

	switch outputFormat {
	case YAML_FORMAT:
		s.generateYAML()
	case JSON_FORMAT:
		s.generateJSON()
	default:
		log.Fatal("Only json or yaml formats are supported")
	}
}

func (s Swagger) initTemplate() (tpl *template.Template) {
	var err error
	var data []byte
	var mappers = template.FuncMap{
		"ToLower": strings.ToLower,
		"UnescapeHTML": func(s string) template.HTML {
			return template.HTML(strings.ReplaceAll(s, `"`, `\"`))
		},
		"RemovePathPrefix": func(path string) (rp string) {
			for _, p := range regexp.
				MustCompile("{{(.*?)}}").
				FindAllStringSubmatch(path, -1) {
				rp = strings.ReplaceAll(path, p[0], "")
			}

			return strings.ReplaceAll(rp, s.Config.BasePath, "")
		},
	}

	if data, err = Asset("tmpl/swagger.yaml"); err != nil {
		log.Fatal(err)
	}

	if tpl, err = template.
		New("swagger.yaml").
		Funcs(mappers).
		Parse(string(data)); err != nil {
		log.Fatal(err)
	}

	return
}

func (s Swagger) generateJSON() {
	var err error
	var tpl bytes.Buffer
	var data []byte

	if err = s.initTemplate().Execute(&tpl, s); err != nil {
		log.Fatal(err)
	}

	if data, err = yaml.YAMLToJSON(tpl.Bytes()); err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("swagger.json", data, 0644); err != nil {
		log.Fatal(err)
	}
}

func (s Swagger) generateYAML() {
	var err error
	var f *os.File

	if f, err = os.Create("swagger.yaml"); err != nil {
		log.Fatal(err)
	}

	if err = s.initTemplate().Execute(f, s); err != nil {
		log.Fatal(err)
	}
}

func fetchVariables(r *models.Resource) {
	for _, param := range regexp.
		MustCompile("/{{ (.*?) }}#*").
		FindAllStringSubmatch(r.URL, -1) {
		r.InsomniaParams = append(r.InsomniaParams, param[1])
	}
}
