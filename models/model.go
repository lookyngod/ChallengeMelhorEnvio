package models

import (
	uuid "github.com/google/uuid"
	"gorm.io/datatypes"
)

type Log struct {
	ID                  uuid.UUID           `gorm:"primary_key;type:string"`
	Request             Request             `json:"request" gorm:"embedded;embedded_prefix:request_"`
	UpstreamURI         string              `json:"upstream_uri"`
	Response            Response            `json:"response" gorm:"embedded;embedded_prefix:response_"`
	AuthenticatedEntity AuthenticatedEntity `json:"authenticated_entity" gorm:"embedded;embedded_prefix:authenticated_entity_"`
	Route               Route               `json:"route" gorm:"embedded;embedded_prefix:route_"`
	Service             Service             `json:"service" gorm:"embedded;embedded_prefix:service_"`
	Latencies           Latencies           `json:"latencies" gorm:"embedded;embedded_prefix:latencies_"`
	ClientIP            string              `json:"client_ip"`
	StartedAt           int64               `json:"started_at"`
}

type Request struct {
	Method      string         `json:"method"`
	URI         string         `json:"uri"`
	URL         string         `json:"url"`
	Size        int64          `json:"size"`
	Querystring datatypes.JSON `json:"querystring" gorm:"type:json"`
	Headers     RequestHeaders `json:"headers" gorm:"embedded;embedded_prefix:headers_"`
}

type RequestHeaders struct {
	Accept    string `json:"accept"`
	Host      string `json:"host"`
	UserAgent string `json:"user-agent"`
}

type Headers struct {
	ContentLenght                 string `json:"Content_lenght"`
	Via                           string `json:"via"`
	Connection                    string `json:"Connection"`
	AccessControlAllowCredentials string `json:"access-control-allow-credentials"`
	ContentType                   string `json:"Content-type"`
	Server                        string `json:"server"`
	AccessControlAllowOrigin      string `json:"access-control-allow-origin"`
}

type Response struct {
	Status  int64   `json:"status"`
	Size    int64   `json:"size"`
	Headers Headers `json:"headers" gorm:"embedded;embedded_prefix:headers_"`
}

type AuthenticatedEntity struct {
	ConsumerID ConsumerID `json:"consumer_id" gorm:"embedded;embedded_prefix:consumer_id_"`
}
type ConsumerID struct {
	UUID string `json:"uuid"`
}

type Route struct {
	CreatedAt     int64          `json:"created_at"`
	Hosts         string         `json:"hosts"`
	ID            string         `json:"id"`
	Methods       datatypes.JSON `json:"methods" gorm:"type:json"`
	Paths         datatypes.JSON `json:"paths" gorm:"type:json"`
	PreserveHost  bool           `json:"preserve_host"`
	Protocols     datatypes.JSON `json:"protocols" gorm:"type:json"`
	RegexPriority int64          `json:"regex_priority"`
	Service       ServiceRoute   `json:"service" gorm:"embedded;embedded_prefix:service_"`
	StripPath     bool           `json:"strip_path"`
	UpdatedAt     int64          `json:"updated_at"`
}

type ServiceRoute struct {
	ID string `json:"id"`
}
type Service struct {
	ConnectTimeout int64  `json:"connect_timeout"`
	CreatedAt      int64  `json:"created_at"`
	Host           string `json:"host"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Port           int64  `json:"port"`
	Protocol       string `json:"protocol"`
	ReadTimeout    int64  `json:"read_timeout"`
	Retries        int64  `json:"retries"`
	UpdatedAt      int64  `json:"updated_at"`
	WriteTimeout   int64  `json:"write_timeout"`
}
type Latencies struct {
	Proxy   int64 `json:"proxy"`
	Kong    int64 `json:"kong"`
	Request int64 `json:"request"`
}
