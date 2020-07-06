package conf

import "os"

// UserSvcAddress returns configured address of User Service
func UserSvcAddress() string {
	return os.Getenv("USER_SVC_ADDRESS")
}

// AuthDomain returns configured auth domain
func AuthDomain() string {
	return os.Getenv("AUTH_DOMAIN")
}
