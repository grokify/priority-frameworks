package priorityframeworks

// FrameworkID constants for built-in frameworks.
const (
	IDSeverity         = "severity"
	IDPriority         = "priority"
	IDIETF             = "ietf"
	IDIETFProhibitions = "ietf-prohibitions"
	IDMoSCoW           = "moscow"
	IDGeneral          = "general"
)

// AllBuiltinIDs returns all built-in framework IDs.
func AllBuiltinIDs() []string {
	return []string{
		IDSeverity,
		IDPriority,
		IDIETF,
		IDIETFProhibitions,
		IDMoSCoW,
		IDGeneral,
	}
}

// Get returns a built-in framework by ID.
// Returns nil if not found.
func Get(id string) *Framework {
	switch id {
	case IDSeverity:
		return Severity()
	case IDPriority:
		return Priority()
	case IDIETF:
		return IETF()
	case IDIETFProhibitions:
		return IETFProhibitions()
	case IDMoSCoW:
		return MoSCoW()
	case IDGeneral:
		return General()
	default:
		return nil
	}
}

// All returns all built-in frameworks.
func All() []*Framework {
	return []*Framework{
		Severity(),
		Priority(),
		IETF(),
		IETFProhibitions(),
		MoSCoW(),
		General(),
	}
}

// Severity returns the Severity framework (Critical/High/Medium/Low/Informational).
// Common in security, incident response, and bug tracking.
func Severity() *Framework {
	return &Framework{
		ID:          IDSeverity,
		Name:        "Severity",
		Description: "Security and incident severity levels. Common in vulnerability assessment and bug tracking.",
		Levels: []Level{
			{ID: "critical", Name: "Critical", Aliases: []string{"CRITICAL", "Crit", "S1"}, Actionable: true, Color: "#7f1d1d"},
			{ID: "high", Name: "High", Aliases: []string{"HIGH", "S2"}, Actionable: true, Color: "#dc2626"},
			{ID: "medium", Name: "Medium", Aliases: []string{"MEDIUM", "Med", "S3"}, Actionable: true, Color: "#ea580c"},
			{ID: "low", Name: "Low", Aliases: []string{"LOW", "S4"}, Actionable: true, Color: "#ca8a04"},
			{ID: "informational", Name: "Informational", Aliases: []string{"INFO", "Info", "S5"}, Actionable: false, Color: "#6b7280"},
		},
	}
}

// Priority returns the Priority framework (P0/P1/P2/P3/P4).
// Common in engineering teams for work prioritization.
func Priority() *Framework {
	return &Framework{
		ID:          IDPriority,
		Name:        "Priority (P#)",
		Description: "Engineering priority levels. P0 is highest urgency, P4 is lowest.",
		Levels: []Level{
			{ID: "p0", Name: "P0", Aliases: []string{"p0", "P-0"}, Actionable: true, Color: "#7f1d1d"},
			{ID: "p1", Name: "P1", Aliases: []string{"p1", "P-1"}, Actionable: true, Color: "#dc2626"},
			{ID: "p2", Name: "P2", Aliases: []string{"p2", "P-2"}, Actionable: true, Color: "#ea580c"},
			{ID: "p3", Name: "P3", Aliases: []string{"p3", "P-3"}, Actionable: true, Color: "#ca8a04"},
			{ID: "p4", Name: "P4", Aliases: []string{"p4", "P-4"}, Actionable: false, Color: "#16a34a"},
		},
	}
}

// IETF returns the IETF RFC 2119 requirements framework (MUST/SHOULD/MAY).
// Use this for prioritizing requirements - what to implement.
// For prohibitions (MUST NOT/SHOULD NOT), use IETFProhibitions().
func IETF() *Framework {
	return &Framework{
		ID:          IDIETF,
		Name:        "IETF RFC 2119",
		Description: "Requirement levels from RFC 2119 for prioritizing what to implement.",
		Levels: []Level{
			{ID: "must", Name: "MUST", Aliases: []string{"REQUIRED", "SHALL"}, Actionable: true, Color: "#dc2626"},
			{ID: "should", Name: "SHOULD", Aliases: []string{"RECOMMENDED"}, Actionable: true, Color: "#ea580c"},
			{ID: "may", Name: "MAY", Aliases: []string{"OPTIONAL"}, Actionable: false, Color: "#16a34a"},
		},
	}
}

// IETFProhibitions returns the IETF RFC 2119 prohibitions framework (MUST NOT/SHOULD NOT).
// Use this for compliance checking - constraints to validate against.
// For requirements (MUST/SHOULD/MAY), use IETF().
func IETFProhibitions() *Framework {
	return &Framework{
		ID:          IDIETFProhibitions,
		Name:        "IETF RFC 2119 Prohibitions",
		Description: "Prohibition levels from RFC 2119 for compliance validation.",
		Levels: []Level{
			{ID: "must-not", Name: "MUST NOT", Aliases: []string{"SHALL NOT"}, Actionable: true, Color: "#7f1d1d"},
			{ID: "should-not", Name: "SHOULD NOT", Aliases: []string{"NOT RECOMMENDED"}, Actionable: true, Color: "#ca8a04"},
		},
	}
}

// MoSCoW returns the MoSCoW framework (Must/Should/Could/Won't).
// Common in agile and product management.
func MoSCoW() *Framework {
	return &Framework{
		ID:          IDMoSCoW,
		Name:        "MoSCoW",
		Description: "Prioritization method for requirements. Must/Should/Could/Won't have.",
		Levels: []Level{
			{ID: "must", Name: "Must have", Aliases: []string{"Must", "M"}, Actionable: true, Color: "#dc2626"},
			{ID: "should", Name: "Should have", Aliases: []string{"Should", "S"}, Actionable: true, Color: "#ea580c"},
			{ID: "could", Name: "Could have", Aliases: []string{"Could", "C"}, Actionable: true, Color: "#ca8a04"},
			{ID: "wont", Name: "Won't have", Aliases: []string{"Wont", "Won't", "W"}, Actionable: false, Color: "#6b7280"},
		},
	}
}

// General returns a general-purpose requirement framework.
// Simple Required/Recommended/Optional/Avoid levels.
func General() *Framework {
	return &Framework{
		ID:          IDGeneral,
		Name:        "General",
		Description: "General-purpose requirement levels. Simple and widely applicable.",
		Levels: []Level{
			{ID: "required", Name: "Required", Aliases: []string{"Req", "R"}, Actionable: true, Color: "#dc2626"},
			{ID: "recommended", Name: "Recommended", Aliases: []string{"Rec"}, Actionable: true, Color: "#2563eb"},
			{ID: "optional", Name: "Optional", Aliases: []string{"Opt", "O"}, Actionable: false, Color: "#16a34a"},
			{ID: "avoid", Name: "Avoid", Aliases: []string{"Deprecated", "Dep"}, Actionable: false, Color: "#6b7280"},
		},
	}
}
