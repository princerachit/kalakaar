package plans

// All the plans except the CUSTOM plan should have default quotas defined
// in some configuration file that lives at one place and acts as a source of truth
// Ideally this configuration should live in a separate repository
const (
	FREE = iota
	BASIC
	PROFESSIONAL
	ENTERPRISE
	PARTNER
	STARTUP
	CUSTOM
)
