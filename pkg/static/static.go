package static

type Route struct {
	Method string
	Path   string
}

var (
	Applications = map[string]string{
		"Admin panel": "admin_panel",
		"Doctor app":  "doctor_app",
		"Client app":  "client_app",
	}

	AplicationRoles = map[string]string{
		"Admin panel": "Admin",
		"Doctor app":  "Doctor",
		"Client app":  "User",
	}
)
