package goflat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUMapWithNil(t *testing.T) {
	result := UMap(nil, &Options{})
	assert.Nil(t, result)
}

func TestUListWithNil(t *testing.T) {
	result := UMap(nil, &Options{})
	assert.Nil(t, result)
}

func TestUMapWithEmptyDelimiter(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, &Options{}).(map[string]interface{})

	assert.Equal(t, "MapSingleField", result["Field"])
}

func TestUListWithEmptyDelimiter(t *testing.T) {
	flat := Map(readList(datasetList), nil)
	result := UMap(flat, &Options{})
	assert.Equal(t, "0", result.([]interface{})[0].([]interface{})[0])
}

func TestUnflatMapSingleField(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, nil).(map[string]interface{})

	assert.Equal(t, "MapSingleField", result["Field"])
}

func TestUnflatMapNestedFields(t *testing.T) {
	opts := &Options{
		Delimiter: "<>",
		Fold:      LowerCaseFold,
	}

	flat := Map(readMap(datasetMap), opts)
	result := UMap(flat, opts).(map[string]interface{})

	assert.Equal(t, "MapNestedField", result["nested"].(map[string]interface{})["nested"].(map[string]interface{})["field"])
	assert.Equal(t, "AnotherValue", result["nested"].(map[string]interface{})["nested"].(map[string]interface{})["antotherfield"])
	assert.Nil(t, result["nested"].(map[string]interface{})["emptyobject"])
}

func TestUnflatMapHazard(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, nil).(map[string]interface{})

	assert.Equal(t, "0", result["Hazard"].(map[string]interface{})["0"])
	assert.Equal(t, "1", result["Hazard"].(map[string]interface{})["1"])
	assert.Equal(t, "Value", result["Hazard"].(map[string]interface{})["Nested"].(map[string]interface{})["Field"])
}

func TestUnflatMapList(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, nil).(map[string]interface{})

	assert.Equal(t, "0", result["List"].([]interface{})[0])
	assert.Equal(t, "1", result["List"].([]interface{})[1])
	assert.Equal(t, "2", result["List"].([]interface{})[2])
}

func TestUnflatMapListVariousTypes(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, nil).(map[string]interface{})

	assert.Equal(t, "0", result["ListVariousTypes"].([]interface{})[0])
	assert.Equal(t, 1, int(result["ListVariousTypes"].([]interface{})[1].(float64)))
	assert.Equal(t, 2.2, result["ListVariousTypes"].([]interface{})[2])
	assert.Equal(t, false, result["ListVariousTypes"].([]interface{})[3])
}

func TestUnflatMapNestedList(t *testing.T) {
	flat := Map(readMap(datasetMap), nil)
	result := UMap(flat, nil).(map[string]interface{})

	assert.Equal(t, "0", result["NestedList"].([]interface{})[0].([]interface{})[0])
	assert.Equal(t, "0", result["NestedList"].([]interface{})[1].([]interface{})[0])
	assert.Equal(t, "1", result["NestedList"].([]interface{})[1].([]interface{})[1])
	assert.Equal(t, "0", result["NestedList"].([]interface{})[2].([]interface{})[0])
	assert.Equal(t, "1", result["NestedList"].([]interface{})[2].([]interface{})[1])
	assert.Equal(t, "2", result["NestedList"].([]interface{})[2].([]interface{})[2])
}

func TestUnflatListNestedListsWithObjects(t *testing.T) {
	flat := Map(readList(datasetList), nil)
	result := UMap(flat, nil).([]interface{})

	assert.Equal(t, 0, int(result[2].(map[string]interface{})["List"].([]interface{})[0].(float64)))
	assert.Equal(t, false, result[2].(map[string]interface{})["List"].([]interface{})[1])
	assert.Equal(t, "Value", result[2].(map[string]interface{})["List"].([]interface{})[2].([]interface{})[0])
	assert.Equal(t, "Value", result[2].(map[string]interface{})["List"].([]interface{})[2].([]interface{})[1].(map[string]interface{})["Field"])
}

func TestUnflatListNestedObjects(t *testing.T) {
	flat := Map(readList(datasetList), nil)
	result := UMap(flat, nil).([]interface{})

	assert.Equal(t, "Value", result[2].(map[string]interface{})["Field"])
	assert.Equal(t, "Value", result[2].(map[string]interface{})["Nested"].(map[string]interface{})["Nested"].(map[string]interface{})["Field"])
}

func TestUnflatListNestedLists(t *testing.T) {
	flat := Map(readList(datasetList), nil)
	result := UMap(flat, nil).([]interface{})

	assert.Equal(t, "0", result[0].([]interface{})[0])
	assert.Equal(t, "0", result[1].([]interface{})[0])
	assert.Equal(t, "1", result[1].([]interface{})[1])
}
