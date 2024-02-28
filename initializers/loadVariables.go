package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var AppProperties map[string]string

func LoadVariables() {	
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		LoadAppProperties()
	}
}

func LoadAppProperties() {
	file, err := os.Open("/opt/myapp/app.properties")
	if err != nil {
		log.Fatal("Error opening app.properties file:", err)
	}
	defer file.Close()

	// The make function is used to initialize maps, slices, and channels.
	AppProperties = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			continue
		}
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		AppProperties[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading app.properties file:", err)
	}
}