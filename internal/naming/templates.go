package naming

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO unit test for this shit
// template for bitflow engine name
func Name(index int64) string {
	return fmt.Sprintf("%s-%d", DefaultEngineName, index)
}

// template for bitflow engine topic
func Topic(index int64, topicStub string) string {
	if index != -1 {
		return fmt.Sprintf("%s/%s/%d/%s", Bitflow, DefaultEngineName, index, topicStub)
	} else {
		return fmt.Sprintf("%s/%s/%s/%s", Bitflow, DefaultEngineName, "+", topicStub)
	}
}

// template for various names across EdgeX, most prominently name of Export Client
func ExportName(index int64, topicStub string) string {
	topicStub = strings.Title(topicStub)
	capBitflow := strings.Title(Bitflow)
	capDefaultEngineName := strings.Title(DefaultEngineName)

	if index != -1 {
		return fmt.Sprintf("%s%s%d%s%s",
			capBitflow, capDefaultEngineName, index, topicStub, "Topic")
	} else {
		return fmt.Sprintf("%s-%s-%s-%s-%s",
			capBitflow, capDefaultEngineName, "Any", topicStub, "Topic")
	}
}

// template for bitflow engine role string
// e.g. bitflow-engine-0-source-publisher
func Role(index int64, topicStub string, role string) string {
	if index != -1 {
		return fmt.Sprintf("%s-%s-%d-%s-%s",
			Bitflow, DefaultEngineName, index, topicStub, role)
	} else {
		return fmt.Sprintf("%s-%s-%s-%s-%s",
			Bitflow, DefaultEngineName, "any", topicStub, role)
	}
}

// template for bitflow engine publisher string
func Publisher(index int64, topicStub string) string {
	return Role(index, topicStub, "publisher")
}

// template for bitflow engine subscriber string
func Subscriber(index int64, topicStub string) string {
	return Role(index, topicStub, "subscriber")
}

// extract index from arbitrary strings that are splittable by separator
func ExtractIndex(source string, separator string, position int64) (int64, error) {
	indexString := strings.Split(source, separator)[position]
	index, err := strconv.ParseInt(indexString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("index could not be extracted from source string %s", source)
	}

	return index, nil
}

