package PlentymarketsAPI

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Web-Softwareentwicklung-Florian-Hame/Plentymarkets-API/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/kardianos/osext"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type apiClient struct {
}

type ApiConfig struct {
	Host          string
	Scheme        string
	Debug         bool
	DefaultHeader map[string]string
	Username      string
	Password      string
}

type ApiClientConfig struct {
	StoreTokenDataInFile bool
	TokenDataFilePath    string
	ApiConfig
}

const (
	longPeriodCallsLeftHeader = "X-Plenty-Global-Long-Period-Calls-Left"
	shortPeriodCallsLeftHeader = "X-Plenty-Global-Short-Period-Calls-Left"
	shortPeriodSecondsToWaitHeader = "X-Plenty-Global-Short-Period-Decay"
)

func NewApiClient(cfg ApiClientConfig) (*api.APIClient, error) {
	client := apiClient{}
	var storedToken *oauth2.Token
	var err error

	if cfg.StoreTokenDataInFile {
		storedToken, err = client.loadTokenFromFile(cfg.TokenDataFilePath)
		if err != nil {
			return nil, err
		}
	}

	if storedToken == nil {
		return nil, fmt.Errorf("please provide one of the possible token data storage options in config")
	}

	tokenSource, err := client.getTokenSource(storedToken, cfg.ApiConfig.Username, cfg.ApiConfig.Password)
	if err != nil {
		return nil, err
	}

	currentToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	err = client.storeTokenInFile(cfg.TokenDataFilePath, currentToken, storedToken)
	if err != nil {
		return nil, err
	}

	defaultHeader := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", currentToken.AccessToken),
	}

	apiConfig := api.NewConfiguration()
	apiConfig.Host = cfg.ApiConfig.Host
	apiConfig.Scheme = cfg.ApiConfig.Scheme
	apiConfig.Debug = cfg.ApiConfig.Debug
	apiConfig.DefaultHeader = defaultHeader
	apiConfig.HTTPClient = oauth2.NewClient(context.Background(), tokenSource)

	return api.NewAPIClient(apiConfig), nil
}

func (c apiClient) storeTokenInFile(tokenDataFilePath string, newToken, storedToken *oauth2.Token) error {
	if newToken.AccessToken == storedToken.AccessToken {
		return nil
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	tokenPath := fmt.Sprintf("%s/%s", folderPath, tokenDataFilePath)

	tokenFile, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer tokenFile.Close()
	jsonContent, err := json.Marshal(newToken)
	if err != nil {
		return err
	}

	_, err = tokenFile.Write(jsonContent)
	if err != nil {
		return err
	}

	return nil
}

func (c apiClient) loadTokenFromFile(tokenFilePath string) (*oauth2.Token, error) {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return nil, err
	}
	tokenPath := fmt.Sprintf("%s/%s", folderPath, tokenFilePath)

	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		return &oauth2.Token{}, nil
	}

	f, err := os.Open(tokenPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	err = json.Unmarshal(fileContent, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (c apiClient) getTokenSource(storedToken *oauth2.Token, username, password string) (oauth2.TokenSource, error) {
	var err error
	conf := oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://lichtyam.plentymarkets-cloud01.com/rest/login",
			TokenURL: "https://lichtyam.plentymarkets-cloud01.com/rest/login",
		},
	}

	if storedToken.Expiry.UTC().Before(time.Now().Add(5*time.Minute).UTC()) {
		storedToken, err = conf.PasswordCredentialsToken(context.Background(), username, password)
		if err != nil {
			return nil, err
		}
	}

	tokenSource := conf.TokenSource(context.Background(), storedToken)
	return tokenSource, nil
}

func (c apiClient) getTokenExpiry(tokenString string) (*time.Time, error) {
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, claims)

	if err != nil {
		return nil, err
	}

	var expiresAt time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		expiresAt = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		expiresAt = time.Unix(v, 0)
	}

	return &expiresAt, nil
}


// EvaluateApiLimits evaluates the plentymarkets api limits from given http response header and
// returns whether there are free limits left in long period, how long it's needed to wait until short period limit is reset when there is no call left in short period
func EvaluateApiLimits(responseHeader http.Header) (bool, *time.Duration, error) {
	longPeriodCallsLeft := responseHeader.Get(longPeriodCallsLeftHeader)
	if longPeriodCallsLeft == "0" {
		// no free limits left for long period
		return false, nil, nil
	}

	shortPeriodCallsLeft := responseHeader.Get(shortPeriodCallsLeftHeader)
	shortPeriodSecondsToWait, err := strconv.Atoi(responseHeader.Get(shortPeriodSecondsToWaitHeader))
	if err != nil {
		return false, nil, err
	}

	if shortPeriodCallsLeft == "0" {
		// wait until short period limit resets
		secondsToWait := time.Second * time.Duration(shortPeriodSecondsToWait)
		return true, &secondsToWait, nil
	}

	return true, nil, nil
}