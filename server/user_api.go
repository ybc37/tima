package server

import "net/http"

type UserAPI struct {
	db   *DB
	auth *Auth
}

func NewUserAPI(db *DB, auth *Auth) *UserAPI {
	return &UserAPI{db, auth}
}

func (userAPI *UserAPI) SigninHandler(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	var credentials UserCredentials
	err := unmarshalJSON(r.Body, &credentials)
	if err != nil {
		return nil, &HandlerError{err, err.Error(), http.StatusBadRequest}
	}

	user := userAPI.db.GetUserByName(credentials.Username)
	if user == nil {
		return nil, &HandlerError{err, "Invalid username/password", http.StatusBadRequest}
	}

	token, err := userAPI.auth.Authenticate(user, credentials.Password)
	if err != nil {
		return nil, &HandlerError{err, "Invalid username/password", http.StatusBadRequest}
	}

	return jsonResultString(token)
}

func (userAPI *UserAPI) IsSignedInHandler(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	signedIn := userAPI.auth.ValidateToken(r)
	return jsonResultBool(signedIn)
}

func (userAPI *UserAPI) AddUser(username string, pwd string, firstName string, lastName string, email string) (*User, error) {
	pwdHash, err := userAPI.auth.GeneratePasswordHash(pwd)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           -1,
		Username:     username,
		PasswordHash: pwdHash,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
	}

	err = userAPI.db.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
