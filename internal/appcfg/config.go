package appcfg

import (
	"github.com/vitlobo/gator/internal/gatorapi"
)

// Config -
type Config struct {
	GatorClient      gatorapi.Client
	DBURL            string `json:"db_url"`
	CurrentUserName  string `json:"current_user_name"`
}