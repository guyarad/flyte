// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package update

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (ProjectConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (ProjectConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in ProjectConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg ProjectConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("ProjectConfig", pflag.ExitOnError)
	cmdFlags.BoolVarP(&(projectConfig.ActivateProject), fmt.Sprintf("%v%v", prefix, "activateProject"), "t", *new(bool), "Activates the project specified as argument.")
	cmdFlags.BoolVarP(&(projectConfig.ArchiveProject), fmt.Sprintf("%v%v", prefix, "archiveProject"), "a", *new(bool), "Archives the project specified as argument.")
	return cmdFlags
}
