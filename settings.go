// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 17:50:57 863577                               [zr/settings.go]
// -----------------------------------------------------------------------------

package zr

// # CONTENTS:
//   SettingsAccessor interface
//   Settings struct
//
// # Methods
//   (ob *Settings) GetSetting(name string) string
//   (ob *Settings) HasSetting(name string) bool
//   (ob *Settings) SetSetting(name string, val interface{})
//
// # Extenders
//   (ob *Settings) ExtendGet(
//                      handler func(name, val string, exists bool) string,
//                  )
//   (ob *Settings) ExtendHas(
//                      handler func(name, val string, exists bool) bool,
//                  )
//   (ob *Settings) ExtendSet(
//                      handler func(name string, old, val interface{}) *string,
//                  )

// SettingsAccessor __
type SettingsAccessor interface {
	GetSetting(name string) string
	HasSetting(name string) bool
	SetSetting(name string, val interface{})
} //                                                            SettingsAccessor

// Settings __
type Settings struct {
	m         map[string]string
	extendGet func(name, val string, exists bool) string
	extendHas func(name, val string, exists bool) bool
	extendSet func(name string, old, val interface{}) *string
} //                                                                    Settings

// -----------------------------------------------------------------------------
// # Methods

// GetSetting __
func (ob *Settings) GetSetting(name string) string {
	const erv = ""
	if ob == nil {
		mod.Error(ENilReceiver)
		return erv
	}
	name = str.ToLower(str.Trim(name, SPACES))
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return erv
	}
	var ret, exists = ob.m[name]
	if ob.extendGet != nil {
		ret = ob.extendGet(name, ret, exists)
	}
	return ret
} //                                                                  GetSetting

// HasSetting __
func (ob *Settings) HasSetting(name string) bool {
	const erv = false
	if ob == nil {
		mod.Error(ENilReceiver)
		return erv
	}
	name = str.ToLower(str.Trim(name, SPACES))
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return erv
	}
	var val, exists = ob.m[name]
	if ob.extendSet != nil {
		return ob.extendHas(name, val, exists)
	}
	return exists
} //                                                                  HasSetting

// SetSetting __
func (ob *Settings) SetSetting(name string, val interface{}) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	name = str.ToLower(str.Trim(name, SPACES))
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return
	}
	if ob.m == nil {
		ob.m = map[string]string{}
	}
	var s = String(val)
	if ob.extendSet != nil {
		var old = String(ob.m[name])
		var result = ob.extendSet(name, old, s)
		if result == nil {
			return
		}
		s = *result
	}
	ob.m[name] = s
} //                                                                  SetSetting

// -----------------------------------------------------------------------------
// # Extenders

// ExtendGet makes 'handler' process every call to GetSetting()
func (ob *Settings) ExtendGet(
	handler func(name, val string, exists bool) string,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendGet = handler
} //                                                                   ExtendGet

// ExtendHas makes 'handler' process every call to HasSetting()
func (ob *Settings) ExtendHas(
	handler func(name, val string, exists bool) bool,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendHas = handler
} //                                                                   ExtendHas

// ExtendSet makes 'handle' process every call to SetSetting()
func (ob *Settings) ExtendSet(
	handler func(name string, old, val interface{}) *string,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendSet = handler
} //                                                                   ExtendSet

//end
