package settings

import (
	"k8s.io/utils/pointer"

	"github.com/seal-io/seal/pkg/casdoor"
	"github.com/seal-io/seal/pkg/dao/model"
)

// the built-in settings for system.
var (
	// HttpProxy indicates the address for proxying none SSL http outbound traffic,
	// it's in form of http(s)://[user:password@]address[:port].
	HttpProxy = newValue("HttpProxy", editable, initializeFromSpecifiedEnv("HTTP_PROXY", ""), modifyWith(httpUrl))
	// HttpsProxy indicates the address for proxying SSL http outbound traffic,
	// it's in form of http(s)://[user:password@]address[:port].
	HttpsProxy = newValue("HttpsProxy", editable, initializeFromSpecifiedEnv("HTTPS_PROXY", ""), modifyWith(httpUrl))
	// AllProxy indicates the address for proxying outbound traffic,
	// it's in form of scheme://[user:password@]address[:port].
	AllProxy = newValue("AllProxy", editable, initializeFromSpecifiedEnv("ALL_PROXY", ""), modifyWith(sockUrl))
	// NoProxy indicates the host exclusion list when proxying outbound traffic,
	// it's a comma-separated string.
	NoProxy = newValue("NoProxy", editable, initializeFromSpecifiedEnv("NO_PROXY", ""), nil)
)

// the built-in settings for server.
var (
	// FirstLogin indicates whether it's the first time to login.
	FirstLogin = newValue("FirstLogin", hidden, initializeFrom("true"), nil)
	// CasdoorCred keeps the AK/SK for accessing Casdoor server.
	CasdoorCred = newValue("CasdoorCred", private, initializeFromJSON(casdoor.ApplicationCredential{}), modifyWith(once))
	// PrivilegeApiToken keeps the token for accessing server APIs.
	PrivilegeApiToken = newValue("PrivilegeApiToken", private, nil, nil)
	// ServeUrl keeps the URL for accessing server.
	ServeUrl = newValue("ServeUrl", editable, nil, modifyWith(httpUrl))
	// ServeUiIndex keeps the address for serving UI.
	ServeUiIndex = newValue("ServeUiIndex", editable|hidden, initializeFrom("https://seal-ui-1303613262.cos.ap-guangzhou.myqcloud.com/latest/index.html"), modifyWith(anyUrl))
	// TerraformDeployerImage indicates the image for terraform deployment.
	TerraformDeployerImage = newValue("TerraformDeployerImage", editable, initializeFrom("sealio/terraform-deployer:v0.1.0"), modifyWith(notBlank, containerImageReference))
	// DataEncryptionSentry keeps the sentry for indicating whether enables data encryption.
	DataEncryptionSentry = newValue("DataEncryptionSentry", private, initializeFrom(""), modifyWith(notBlank))
	// OpenAiApiToken keeps the openAI API token for generating module completions.
	// TODO protect the stored token.
	OpenAiApiToken = newValue("openAiApiToken", editable, nil, nil)
)

// the built-in settings for server cron jobs.
var (
	// CostCollectCronExpr indicates the cron expression of collect cost data,
	// default cron expression means executing collection per hour,
	// the cron expression is in form of `Seconds Minutes Hours DayOfMonth Month DayOfWeek`.
	CostCollectCronExpr = newValue("CostCollectCronExpr", editable, initializeFrom("0 0 * ? * *"), modifyWith(notBlank, cronExpression))
	// ConnectorCheckCronExpr indicates the cron expression of check connector status,
	// default cron expression means executing check every 5 minutes.
	ConnectorCheckCronExpr = newValue("ConnectorCheckCronExpr", editable, initializeFrom("0 */5 * ? * *"), modifyWith(notBlank, cronExpression))
	// ResourceStatusCheckCronExpr indicates the cron expression of state application resource,
	// default cron expression means stating every 1 minute.
	ResourceStatusCheckCronExpr = newValue("ResourceStatusCheckCronExpr", editable, initializeFrom("0 */1 * ? * *"), modifyWith(notBlank, cronExpression))
	// ResourceLabelApplyCronExpr indicates the cron expression of set labels to application resource linked kubernetes resource,
	// default cron expression means stating every 2 minute.
	ResourceLabelApplyCronExpr = newValue("ResourceLabelApplyCronExpr", editable, initializeFrom("0 */2 * ? * *"), modifyWith(notBlank, cronExpression))
)

// setting property list.
const (
	_default uint8 = 0
	hidden         = 1 << (iota - 1)
	editable
	private
)

var valuesOrder []string
var valuesIndex = map[string]value{}

// newValue creates a value with the given name and modifier,
// then indexes the new value by its name.
func newValue(name string, property uint8, initialize initializer, modify modifier) value {
	if modify == nil {
		modify = modifyWith(many)
	}
	if initialize == nil {
		initialize = initializeFromEnv("")
		if property&private == private {
			initialize = initializeFrom("")
		}
	}
	var v = value{
		refer: model.Setting{
			Name:     name,
			Value:    initialize(name),
			Hidden:   pointer.Bool(property&hidden == hidden),
			Editable: pointer.Bool(property&editable == editable),
			Private:  pointer.Bool(property&private == private),
		},
		modify: modify,
	}
	valuesOrder = append(valuesOrder, name)
	valuesIndex[name] = v
	return v
}

// ForEach iterates each setting in input.
func ForEach(input func(setting model.Setting) error) error {
	if input == nil {
		return nil
	}

	for _, n := range valuesOrder {
		var err = input(valuesIndex[n].refer)
		if err != nil {
			return err
		}
	}
	return nil
}

// All returns all settings.
func All() (r model.Settings) {
	_ = ForEach(func(s model.Setting) error {
		r = append(r, &s)
		return nil
	})
	return
}

// Index returns the setting with the given name.
func Index(name string) Value {
	return valuesIndex[name]
}
