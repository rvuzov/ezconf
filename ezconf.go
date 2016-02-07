package ezconf

import(
	"os"
	"bufio"
	"regexp"
	"strconv"
)

var configData map[string]map[string]string
var re *regexp.Regexp;

func init() {
	configData = make(map[string]map[string]string)
	re = regexp.MustCompile("^\\s*([\\w-]*)\\s*:\\s*(.*)\\s*")
}

func Get(setting string) string {
	namespaceMap := fetchNamespace(os.Args[1])
	val, _ := namespaceMap[setting]
	return val
}

func GetUint(setting string) uint64 {
	namespaceMap := fetchNamespace(os.Args[1])
	val, _ := namespaceMap[setting]
	parsedVal, _ := strconv.ParseUint(val, 10, 64)
	return parsedVal
}

func GetInt(setting string) int64 {
	namespaceMap := fetchNamespace(os.Args[1])
	val, _ := namespaceMap[setting]
	parsedVal, _ := strconv.ParseInt(val, 10, 64)
	return parsedVal
}

func GetFloat(setting string) float64 {
	namespaceMap := fetchNamespace(os.Args[1])
	val, _ := namespaceMap[setting]
	parsedVal, _ := strconv.ParseFloat(val, 64)
	return parsedVal
}

func GetBool(setting string) bool {
	namespaceMap := fetchNamespace(os.Args[1])
	val, _ := namespaceMap[setting]
	parsedVal, _ := strconv.ParseBool(val)
	return parsedVal
}

func Copy() map[string]string {
	namespaceMap := fetchNamespace(os.Args[1])
	mapCopy := make(map[string]string)
	for k,v := range namespaceMap {
	  mapCopy[k] = v
	}
	return mapCopy
}

func Set(setting string, value string) {
	namespaceMap := fetchNamespace(os.Args[1])
	namespaceMap[setting] = value
}

func fetchNamespace(namespace string) map[string]string {
	namespaceMap, ok := configData[namespace]
	if !ok {
		importSettingsFromFile(namespace)
		namespaceMap, _ = configData[namespace]
	}
	return namespaceMap
}

func importSettingsFromFile(namespace string) {
	configData[namespace] = make(map[string]string)
	file, err := os.Open("config/"+ namespace +".conf")
	defer file.Close()
	if err != nil {
		// if no config file, that is fine and dandy, can still use it without config files.
		return
	}
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := re.FindStringSubmatch(line)
		if(len(parsedLine) == 3) {
			configData[namespace][parsedLine[1]] = parsedLine[2]
		}
	}
}
