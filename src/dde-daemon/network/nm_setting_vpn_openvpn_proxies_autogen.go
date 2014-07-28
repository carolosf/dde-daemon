// This file is automatically generated, please don't edit manually.
package network

import (
	"fmt"
)

// Get key type
func getSettingVpnOpenvpnProxiesKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE:
		t = ktypeString
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER:
		t = ktypeString
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT:
		t = ktypeUint32
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME:
		t = ktypeString
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD:
		t = ktypeString
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting section
func isKeyInSettingVpnOpenvpnProxies(key string) bool {
	switch key {
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnOpenvpnProxiesDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE:
		value = "none"
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER:
		value = ""
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT:
		value = uint32(0)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY:
		value = false
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME:
		value = ""
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD:
		value = ""
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS:
		value = uint32(0)
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnOpenvpnProxiesKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnOpenvpnProxiesKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE:
		value = getSettingVpnOpenvpnKeyProxyTypeJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER:
		value = getSettingVpnOpenvpnKeyProxyServerJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT:
		value = getSettingVpnOpenvpnKeyProxyPortJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY:
		value = getSettingVpnOpenvpnKeyProxyRetryJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME:
		value = getSettingVpnOpenvpnKeyHttpProxyUsernameJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD:
		value = getSettingVpnOpenvpnKeyHttpProxyPasswordJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS:
		value = getSettingVpnOpenvpnKeyHttpProxyPasswordFlagsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnOpenvpnProxiesKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnOpenvpnProxiesKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE:
		err = logicSetSettingVpnOpenvpnKeyProxyTypeJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER:
		err = setSettingVpnOpenvpnKeyProxyServerJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT:
		err = setSettingVpnOpenvpnKeyProxyPortJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY:
		err = setSettingVpnOpenvpnKeyProxyRetryJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME:
		err = setSettingVpnOpenvpnKeyHttpProxyUsernameJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD:
		err = setSettingVpnOpenvpnKeyHttpProxyPasswordJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS:
		err = setSettingVpnOpenvpnKeyHttpProxyPasswordFlagsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnOpenvpnKeyProxyTypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE)
}
func isSettingVpnOpenvpnKeyProxyServerExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER)
}
func isSettingVpnOpenvpnKeyProxyPortExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT)
}
func isSettingVpnOpenvpnKeyProxyRetryExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY)
}
func isSettingVpnOpenvpnKeyHttpProxyUsernameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME)
}
func isSettingVpnOpenvpnKeyHttpProxyPasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD)
}
func isSettingVpnOpenvpnKeyHttpProxyPasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS)
}

// Ensure section and key exists and not empty
func ensureSectionSettingVpnOpenvpnProxiesExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME))
	}
}
func ensureSettingVpnOpenvpnKeyProxyTypeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyProxyTypeExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyProxyType(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyProxyServerNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyProxyServerExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyProxyServer(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyProxyPortNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyProxyPortExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyProxyRetryNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyProxyRetryExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyHttpProxyUsernameNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyHttpProxyUsernameExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyHttpProxyUsername(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyHttpProxyPasswordNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyHttpProxyPasswordExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyHttpProxyPassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyHttpProxyPasswordFlagsNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyHttpProxyPasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingVpnOpenvpnKeyProxyType(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnOpenvpnKeyProxyServer(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnOpenvpnKeyProxyPort(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnOpenvpnKeyProxyRetry(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnOpenvpnKeyHttpProxyUsername(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnOpenvpnKeyHttpProxyPassword(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnOpenvpnKeyHttpProxyPasswordFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS)
	value = interfaceToUint32(ivalue)
	return
}

// Setter
func setSettingVpnOpenvpnKeyProxyType(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE, value)
}
func setSettingVpnOpenvpnKeyProxyServer(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER, value)
}
func setSettingVpnOpenvpnKeyProxyPort(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT, value)
}
func setSettingVpnOpenvpnKeyProxyRetry(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY, value)
}
func setSettingVpnOpenvpnKeyHttpProxyUsername(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME, value)
}
func setSettingVpnOpenvpnKeyHttpProxyPassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD, value)
}
func setSettingVpnOpenvpnKeyHttpProxyPasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS, value)
}

// JSON Getter
func getSettingVpnOpenvpnKeyProxyTypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE))
	return
}
func getSettingVpnOpenvpnKeyProxyServerJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER))
	return
}
func getSettingVpnOpenvpnKeyProxyPortJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT))
	return
}
func getSettingVpnOpenvpnKeyProxyRetryJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY))
	return
}
func getSettingVpnOpenvpnKeyHttpProxyUsernameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME))
	return
}
func getSettingVpnOpenvpnKeyHttpProxyPasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD))
	return
}
func getSettingVpnOpenvpnKeyHttpProxyPasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS))
	return
}

// JSON Setter
func setSettingVpnOpenvpnKeyProxyTypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE))
}
func setSettingVpnOpenvpnKeyProxyServerJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER))
}
func setSettingVpnOpenvpnKeyProxyPortJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT))
}
func setSettingVpnOpenvpnKeyProxyRetryJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY))
}
func setSettingVpnOpenvpnKeyHttpProxyUsernameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME))
}
func setSettingVpnOpenvpnKeyHttpProxyPasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD))
}
func setSettingVpnOpenvpnKeyHttpProxyPasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS, valueJSON, getSettingVpnOpenvpnProxiesKeyType(NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS))
}

// Logic JSON Setter
func logicSetSettingVpnOpenvpnKeyProxyTypeJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingVpnOpenvpnKeyProxyTypeJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingVpnOpenvpnKeyProxyTypeExists(data) {
		value := getSettingVpnOpenvpnKeyProxyType(data)
		err = logicSetSettingVpnOpenvpnKeyProxyType(data, value)
	}
	return
}

// Remover
func removeSettingVpnOpenvpnKeyProxyType(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_TYPE)
}
func removeSettingVpnOpenvpnKeyProxyServer(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_SERVER)
}
func removeSettingVpnOpenvpnKeyProxyPort(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_PORT)
}
func removeSettingVpnOpenvpnKeyProxyRetry(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROXY_RETRY)
}
func removeSettingVpnOpenvpnKeyHttpProxyUsername(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_USERNAME)
}
func removeSettingVpnOpenvpnKeyHttpProxyPassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD)
}
func removeSettingVpnOpenvpnKeyHttpProxyPasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_PROXIES_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_HTTP_PROXY_PASSWORD_FLAGS)
}