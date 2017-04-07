package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"reflect"
	"strings"
)

func randomize(source interface{}, randomSentence *RandomSentence) interface{} {
	switch source.(type) {
	case map[string]interface{}:
		s := source.(map[string]interface{})
		for k, v := range s {
			s[k] = randomize(v, randomSentence)
		}
	case []interface{}:
		s := source.([]interface{})
		for i, v := range s {
			s[i] = randomize(v, randomSentence)
		}
	case string:
		source = randomSentence.Generate()
	case bool:
		source = rand.Intn(1) == 0
	case int:
		n := source.(int)
		if n > 0 {
			source = rand.Intn(n)
		} else {
			source = rand.Int()
		}
	case float32:
		n := int(source.(float32))
		source = rand.Float32() + float32(n)
	case float64:
		n := int(source.(float64))
		source = rand.Float64() + float64(n)
	default:
		log.Fatalln(reflect.TypeOf(source))
	}
	return source
}

func parseJson(m interface{}, jsonPath string) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(data, m); err != nil {
		log.Fatalln(err)
	}
}

func parseDictionary(dictionaryPath string, lines *[]string) {
	data, err := ioutil.ReadFile(dictionaryPath)
	if err != nil {
		log.Fatalln(err)
	}
	*lines = strings.Split(string(data), "\n")
}

func main() {
	var (
		templateJsonPath string
		dictionaryPath   string
		generateCount    int
		repeatCount      int
		sentences        []string
	)

	flag.StringVar(&dictionaryPath, "d", "", "")
	flag.IntVar(&generateCount, "c", 10, "")
	flag.IntVar(&repeatCount, "r", 10, "")
	flag.Parse()

	if len(flag.Args()) > 0 {
		templateJsonPath = flag.Arg(0)
	}

	if dictionaryPath != "" {
		parseDictionary(dictionaryPath, &sentences)
	}
	var randomSentence = NewRandomSentence(sentences, generateCount)

	var results []interface{}
	var template interface{}

	parseJson(&template, templateJsonPath)
	for i := 0; i < repeatCount; i++ {
		results = append(results, randomize(template, randomSentence))
	}

	b, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
