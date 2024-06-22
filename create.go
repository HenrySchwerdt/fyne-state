package fyne_state

type StateActions struct {
	State   map[string]interface{}
	Actions map[string]func()
}

func Create(initializer func(set func(string, interface{}), get func(string) interface{}) StateActions) {
	actions := initializer(Set, func(key string) interface{} { return Get[interface{}](key) })
	state = actions.State
	for key := range actions.Actions {
		Set(key, actions.Actions[key])
	}
}
