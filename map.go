package termloop

import (
	"encoding/json"
	"io/ioutil"
)

type levelMap struct {
	Type string
	Data map[string]interface{}
}

// An EntityParser is a function which composes an object
// from data that has been parsed from a JSON file.
// Returns a Drawable
type EntityParser func(map[string]interface{}) Drawable

func parseRectangle(data map[string]interface{}) *Rectangle {
	return NewRectangle(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		int(data["width"].(float64)),
		int(data["height"].(float64)),
		Attr(data["color"].(float64)),
	)
}

func parseText(data map[string]interface{}) *Text {
	return NewText(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		data["text"].(string),
		Attr(data["fg"].(float64)),
		Attr(data["bg"].(float64)),
	)
}

func parseEntity(data map[string]interface{}) (*Entity, error) {
	filename := data["text"].(string)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	e := NewEntityFromCanvas(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		CanvasFromString(string(text)),
	)
	bgfile := data["bg"].(string)
	if bgfile != "" {
		e.ApplyCanvas(BackgroundCanvasFromFile(bgfile))
	}
	fgfile := data["fg"].(string)
	if fgfile != "" {
		e.ApplyCanvas(ForegroundCanvasFromFile(fgfile))
	}
	return e, nil
}

// LoadLevelFromMap can be used to populate a Level with entities, given
// a JSON string to read from (jsonMap).
//
// The map 'parsers' is a map of entity names to EntityParser functions. This can
// be used to define parsers for objects that are not Termloop builtins.
//
// The JSON string should take the format of an array of objects, like so:
// [ {"type": "Rectangle", "data": {"x: 12 ...}}, ... ]
// For Rectangles and Text, all attributes must be defined in the JSON. For an Entity,
// fg and bg can be left as empty strings if you do not wish to color the Entity with
// an image, but the keys must still be present or the parsing will break.
//
// There is an example of how to use this method at _examples/levelmap.go.
//
// LoadLevelFromMap returns an error, or nil if all is well.
func LoadLevelFromMap(jsonMap string, parsers map[string]EntityParser, l Level) error {
	parsedMap := []levelMap{}
	if err := json.Unmarshal([]byte(jsonMap), &parsedMap); err != nil {
		return err
	}
	for _, lm := range parsedMap {
		var entity Drawable
		var err error
		switch lm.Type {
		case "Rectangle":
			entity = parseRectangle(lm.Data)
		case "Text":
			entity = parseText(lm.Data)
		case "Entity":
			entity, err = parseEntity(lm.Data)
			if err != nil {
				return err
			}
		default:
			entity = parsers[lm.Type](lm.Data)
		}
		l.AddEntity(entity)
	}
	return nil
}
