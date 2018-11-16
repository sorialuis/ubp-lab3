package users

import (
    "time"
    "encoding/json"
)

// Email
type Email struct {
    User   string `json:"user"`
    Domain string `json:"domain"`
}

// EmailCheckResponse
type EmailCheckResponse struct {
    Email       Email `json:"email"`
    EmailExists bool  `json:"email_exists"`
}

// UserResponse representa la información de un usuario en Mercado Libre
type UserResponse struct {
    ID               int64     `json:"id"`
    Nickname         string    `json:"nickname"`
    RegistrationDate time.Time `json:"registration_date"`
    FirstName        string    `json:"first_name"`
    LastName         string    `json:"last_name"`
    Email            string    `json:"email"`
    SiteID           string    `json:"site_id"`
}

// ParseUser recibe un json de un usuario y retorna una nueva estructura
func ParseUser(jsonContent string) (*UserResponse, error) {
    var userResponse UserResponse
    if err := json.Unmarshal([]byte(jsonContent), &userResponse); err != nil {
        return nil, err
    }
    return &userResponse, nil
}

// ParseUser recibe un json de un usuario y retorna una nueva estructura
func ParseEmailCheckResponse(jsonContent string) (*EmailCheckResponse, error) {
    var emailCheckResponse EmailCheckResponse
    if err := json.Unmarshal([]byte(jsonContent), &emailCheckResponse); err != nil {
        return nil, err
    }
    return &emailCheckResponse, nil
}
