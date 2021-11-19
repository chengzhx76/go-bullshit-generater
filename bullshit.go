package bullshit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func Generator(title string, length int) string {
	fileContent := readFile("./data.json")

	data := make(map[string]interface{})
	err := json.Unmarshal(fileContent, &data)
	if err != nil {
		panic(err)
	}

	famous := data["famous"].([]interface{})
	befores := data["before"].([]interface{})
	afters := data["after"].([]interface{})
	boshs := data["bosh"].([]interface{})

	article := ""
	for len(article) < length {
		randNum := rand.Intn(100)
		if randNum < 10 {
			article += "\r\n"
		} else if randNum < 20 {
			famou := famous[rand.Intn(len(famous))].(string)
			before := befores[rand.Intn(len(befores))].(string)
			after := afters[rand.Intn(len(afters))].(string)

			content := strings.Replace(famou, "a", before, 1)
			content = strings.Replace(content, "b", after, 1)
			article += content
		} else {
			bosh := boshs[rand.Intn(len(boshs))].(string)
			article += bosh
			article = strings.Replace(article, "x", title, 1)
		}
	}

	return article
}

func readFile(path string) []byte {
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		strByte, _ := ioutil.ReadAll(file)
		return strByte
	} else {
		panic(err)
	}

}

func main() {
	fmt.Println(Generator("哈哈", 6000))
}
