package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	s3V2 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gorm.io/datatypes"
	"io"
	"net/http"
	"net/url"

	"os"

	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gobeam/stringy"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func ParseDBError(err error, ph string) (int, string) {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		switch pgError.Code {
		case pgerrcode.ForeignKeyViolation:
			return http.StatusBadRequest, fmt.Sprintf("Failed to create %s", ph)
		case pgerrcode.UniqueViolation:
			return http.StatusBadRequest, fmt.Sprintf("%s with this Info already exists", ph)
		case pgerrcode.CaseNotFound:
			return http.StatusNotFound, fmt.Sprintf("%s doesn't exist", ph)
		case pgerrcode.UndefinedTable:
			return http.StatusInternalServerError, fmt.Sprintf("%s table doesn't exist", ph)
		default:
			return http.StatusBadRequest, fmt.Sprintf("Something went wrong with %s", ph)
		}
	}

	return http.StatusBadRequest, "Something went wrong"
}

func GetUserAndWorkspaceIDFromContext(ctx *gin.Context) (int, int, string, string, string, string) {
	userID := ctx.Request.Header.Get("userID")
	workspaceID := ctx.Request.Header.Get("workspaceID")
	airbyteWorkspaceId := ctx.Request.Header.Get("airbyteWorkspaceID")
	domain := ctx.Request.Header.Get("domain")
	domainUsers := ctx.Request.Header.Get("domainUsers")
	userRoles := ctx.Request.Header.Get("userRoles")

	userId, _ := strconv.Atoi(userID)
	workspaceId, _ := strconv.Atoi(workspaceID)

	return userId, workspaceId, airbyteWorkspaceId, domain, domainUsers, userRoles
}

func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func ContainsInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func GenerateURL(relativePath string, queryParam string, scheme string, host string) *url.URL {
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     relativePath,
		RawQuery: queryParam,
	}
}

func ToSnakeCase(input string) string {
	str := stringy.New(input)
	snakeStr := str.SnakeCase()

	return snakeStr.ToLower()
}

func ToSnakeCaseWithPostfix(postfix uuid.UUID, input string) string {
	unique := strings.Replace(postfix.String(), "-", "", 4)
	final := input + " " + unique

	//return "_" + ToSnakeCase(final)
	return ToSnakeCase(final)
}

func DownloadFile(url, path string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return
}

func DifferenceSliceStr(a, b []string) []string {
	diff := make([]string, 0)

	exists := make(map[string]struct{}, len(b))
	for _, x := range b {
		exists[x] = struct{}{}
	}

	for _, x := range a {
		if _, found := exists[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

func CleanWebpageName(name string) string {
	replaceCharsMap := map[string]string{
		" ":  "_",
		"/":  "_",
		"\\": "_",
		":":  "_",
		"*":  "_",
		"?":  "_",
		"\"": "_",
		".":  "_",
	}

	cleanName := name
	for oldChar, newChar := range replaceCharsMap {
		cleanName = strings.ReplaceAll(cleanName, oldChar, newChar)
	}

	return regexp.MustCompile(`[^a-zA-Z0-9 _]+`).ReplaceAllString(cleanName, "")
}

func Unique(input []string) []string {
	if len(input) == 0 {
		return input
	}

	inResult := make(map[string]struct{})

	capacity := 100 % (len(input) + 10)

	output := make([]string, 0, capacity)

	for _, str := range input {
		if _, exists := inResult[str]; !exists {
			inResult[str] = struct{}{}

			output = append(output, str)
		}
	}

	return output
}

func FilterNonEmptyStrings(strSlice []string) []string {
	var nonEmptyStrings []string

	for _, str := range strSlice {
		if str != "" {
			nonEmptyStrings = append(nonEmptyStrings, str)
		}
	}

	return nonEmptyStrings
}

func IsJsonLFileFormat(fileName string) bool {
	return strings.Contains(fileName, FileTypeJSONL)
}

func JSONLToMap(rawObject *s3V2.GetObjectOutput) ([]map[string]interface{}, error) {
	reader := bufio.NewReader(rawObject.Body)

	var rows []map[string]interface{}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(io.EOF, err) {
				break
			}

			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal(line, &data)
		if err != nil {
			return nil, err
		}

		rows = append(rows, data)
	}

	return rows, nil
}

func IsValidURL(s string) bool {
	if _, err := url.ParseRequestURI(s); err != nil {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return false
	}

	return true
}

func MaskSecret(config datatypes.JSON, secretKeyPath ...string) datatypes.JSON {
	maskedConfiguration := string(config)

	for _, key := range secretKeyPath {
		if gjson.Get(maskedConfiguration, key).Exists() {
			maskedConfiguration, _ = sjson.Set(maskedConfiguration, key, CredMask)
		}
	}

	return []byte(maskedConfiguration)
}
