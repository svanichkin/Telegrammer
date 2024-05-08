package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v3"
)

func formatAsINI(obj map[string]any) string {

	parts := []string{}
	for k, v := range obj {
		parts = append(parts, fmt.Sprintf("%s = %v", k, v))
	}
	return strings.Join(parts, "\n")

}

func responseError(c *gin.Context, err error, rt string) {

	if teleErr, ok := err.(*telebot.Error); ok {
		response(c, teleErr.Code, gin.H{"error": teleErr.Description}, rt)
	} else {
		response(c, http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)}, rt)
	}

}

func response(c *gin.Context, code int, obj map[string]any, t string) {

	switch strings.ToLower(t) {
	case "json":
		c.JSON(code, obj)
	case "jsonp":
		c.JSONP(code, obj)
	case "securejson":
		c.SecureJSON(code, obj)
	case "indentedjson":
		c.IndentedJSON(code, obj)
	case "asciijson":
		c.AsciiJSON(code, obj)
	case "purejson":
		c.PureJSON(code, obj)
	case "xml":
		c.XML(code, obj)
	case "html":
		c.HTML(code, "error", obj)
	case "yaml":
		c.YAML(code, obj)
	case "toml":
		c.TOML(code, obj)
	case "protobuf":
		c.ProtoBuf(code, obj)
	case "ini":
		c.String(code, formatAsINI(obj))
	}

}
