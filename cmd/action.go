package cmd

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type action interface {
	String() string
}

type typeAction struct {
	Type  string `mapstructure:"type"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

type keyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

type sleepAction struct {
	Sleep int `mapstructure:"sleep"`
}

type pauseAction struct{}

type ctrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

func (action *typeAction) String() string {
	return fmt.Sprintf("Type: %s", truncateString(action.Type, 37))
}

func (action *keyAction) String() string {
	return fmt.Sprintf("Key: %s", action.Key)
}

func (action *sleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Sleep)
}

func (action *pauseAction) String() string {
	return "Pause: Press enter to continue"
}

func (action *ctrlAction) String() string {
	return fmt.Sprintf("Ctrl+%s", action.Ctrl)
}

func parseAction(settings *settings, v interface{}) (action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &pauseAction{}, nil
		}
	case map[string]interface{}:
		if _, ok := v["pause"]; ok {
			return parsePauseAction(settings, v)
		}
		if _, ok := v["type"]; ok {
			return parseTypeAction(settings, v)
		}
		if _, ok := v["key"]; ok {
			return parseKeyAction(settings, v)
		}
		if _, ok := v["sleep"]; ok {
			return parseSleepAction(settings, v)
		}
		if _, ok := v["ctrl"]; ok {
			return parseCtrlAction(settings, v)
		}
	}

	return nil, fmt.Errorf("invalid action: %#v", v)
}

func parseTypeAction(settings *settings, m map[string]interface{}) (*typeAction, error) {
	action := typeAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseKeyAction(settings *settings, m map[string]interface{}) (*keyAction, error) {
	action := keyAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseSleepAction(settings *settings, m map[string]interface{}) (*sleepAction, error) {
	var action sleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parsePauseAction(settings *settings, m map[string]interface{}) (*pauseAction, error) {
	var action pauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseCtrlAction(settings *settings, m map[string]interface{}) (*ctrlAction, error) {
	action := ctrlAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}
