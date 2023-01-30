// Package credential defines interfaces shared by other packages that
// manage credentials for Boundary sessions.
package credential

import (
	"context"

	"github.com/hashicorp/boundary/internal/boundary"
	"google.golang.org/protobuf/types/known/structpb"
)

// Domain defines the domain for the credential package.
const Domain = "credential"

// A Store is a resource that can store, retrieve, and potentially generate
// credentials of differing types and access levels. It belongs to a project
// and must support the principle of least privilege by providing
// mechanisms to limit the credentials it can access to the minimum
// necessary for the project it is in.
type Store interface {
	boundary.Resource
	GetProjectId() string
}

// Type is the type of credential provided by a library.
type Type string

// Credential type values.
const (
	UnspecifiedType      Type = "unspecified"
	UsernamePasswordType Type = "username_password"
	SshPrivateKeyType    Type = "ssh_private_key"
	SshCertificateType   Type = "ssh_certificate"
	JsonType             Type = "json"
)

// A Library is a resource that provides credentials that are of the same
// type and access level from a single store.
type Library interface {
	boundary.Resource
	GetStoreId() string
	CredentialType() Type
}

// Purpose is the purpose of the credential.
type Purpose string

func (p Purpose) String() string {
	return string(p)
}

// Credential purpose values.
const (
	// BrokeredPurpose is a credential used for brokering specific
	// purposes. Brokered credentials are returned to the user.
	BrokeredPurpose Purpose = "brokered"

	// InjectedApplicationPurpose is a credential used by a boundary
	// worker to secure the connection between the worker and the endpoint.
	// Injected Application credentials are never returned to the user.
	InjectedApplicationPurpose Purpose = "injected_application"
)

// ValidPurposes are the set of all credential Purposes.
var ValidPurposes = []Purpose{
	BrokeredPurpose,
	InjectedApplicationPurpose,
}

// SecretData represents secret data.
type SecretData any

// Credential is an entity containing secret data.
type Credential interface {
	boundary.Entity
	Secret() SecretData
}

// Dynamic is a credential generated by a library for a specific session.
type Dynamic interface {
	Credential
	GetSessionId() string
	Library() Library
	Purpose() Purpose
}

// Static is a static credential that is stored directly in a credential store.
type Static interface {
	boundary.Resource
	GetStoreId() string
}

// A Request represents a request for a credential from the SourceId for
// the given purpose. For dynamic credentials, the SourceId is the PublicId
// of a credential library.
type Request struct {
	SourceId string
	Purpose  Purpose
}

// SshCertificate is a credential containing a client certificate, username,
// and SSH private key.
type SshCertificate interface {
	SshPrivateKey
	Certificate() []byte
}

// Issuer issues dynamic credentials.
type Issuer interface {
	// Issue issues dynamic credentials for a session from the requested
	// libraries and for the requested purposes. The sessionId must be a
	// valid sessionId. The SourceId in each request must be the public id
	// of a library the Issuer can issue credentials from.
	//
	// If Issue encounters an error, it returns no credentials and revokes
	// any credentials issued before encountering the error.
	//
	// Supported Options: WithTemplateData
	Issue(ctx context.Context, sessionId string, requests []Request, opt ...Option) ([]Dynamic, error)
}

// Revoker revokes dynamic credentials.
type Revoker interface {
	// Revoke revokes the dynamic credentials issued for sessionid.
	Revoke(ctx context.Context, sessionId string) error
}

// Password represents a secret password.
type Password string

// PrivateKey represents a secret private key.
type PrivateKey []byte

// JsonObject represents a JSON object that is serialized.
type JsonObject struct {
	structpb.Struct
}

// UsernamePassword is a credential containing a username and a password.
type UsernamePassword interface {
	Credential
	Username() string
	Password() Password
}

// SshPrivateKey is a credential containing a username an SSH private key and
// an optional private key passphrase.
type SshPrivateKey interface {
	Credential
	Username() string
	PrivateKey() PrivateKey
	PrivateKeyPassphrase() []byte
}
