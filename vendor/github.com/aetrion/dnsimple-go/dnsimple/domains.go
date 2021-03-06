package dnsimple

import (
	"fmt"
)

// DomainsService handles communication with the domain related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/domains/
type DomainsService struct {
	client *Client
}

// Domain represents a domain in DNSimple.
type Domain struct {
	ID           int    `json:"id,omitempty"`
	AccountID    int    `json:"account_id,omitempty"`
	RegistrantID int    `json:"registrant_id,omitempty"`
	Name         string `json:"name,omitempty"`
	UnicodeName  string `json:"unicode_name,omitempty"`
	Token        string `json:"token,omitempty"`
	State        string `json:"state,omitempty"`
	AutoRenew    bool   `json:"auto_renew,omitempty"`
	PrivateWhois bool   `json:"private_whois,omitempty"`
	ExpiresOn    string `json:"expires_on,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// DomainResponse represents a response from an API method that returns a Domain struct.
type DomainResponse struct {
	Response
	Data *Domain `json:"data"`
}

// DomainsResponse represents a response from an API method that returns a collection of Domain struct.
type DomainsResponse struct {
	Response
	Data []Domain `json:"data"`
}

// domainRequest represents a generic wrapper for a domain request,
// when domainWrapper cannot be used because of type constraint on Domain.
type domainRequest struct {
	Domain interface{} `json:"domain"`
}

func domainIdentifier(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return fmt.Sprintf("%d", value)
	}
	return ""
}

func domainPath(accountID string, domain interface{}) string {
	if domain != nil {
		return fmt.Sprintf("/%v/domains/%v", accountID, domainIdentifier(domain))
	}
	return fmt.Sprintf("/%v/domains", accountID)
}

// ListDomains lists the domains for an account.
//
// See https://developer.dnsimple.com/v2/domains/#list
func (s *DomainsService) ListDomains(accountID string, options *ListOptions) (*DomainsResponse, error) {
	path := versioned(domainPath(accountID, nil))
	domainsResponse := &DomainsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(path, domainsResponse)
	if err != nil {
		return nil, err
	}

	domainsResponse.HttpResponse = resp
	return domainsResponse, nil
}

// CreateDomain creates a new domain in the account.
//
// See https://developer.dnsimple.com/v2/domains/#create
func (s *DomainsService) CreateDomain(accountID string, domainAttributes Domain) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, nil))
	domainResponse := &DomainResponse{}

	resp, err := s.client.post(path, domainAttributes, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}

// GetDomain fetches a domain.
//
// See https://developer.dnsimple.com/v2/domains/#get
func (s *DomainsService) GetDomain(accountID string, domain interface{}) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, domain))
	domainResponse := &DomainResponse{}

	resp, err := s.client.get(path, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}

// DeleteDomain PERMANENTLY deletes a domain from the account.
//
// See https://developer.dnsimple.com/v2/domains/#delete
func (s *DomainsService) DeleteDomain(accountID string, domain interface{}) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, domain))
	domainResponse := &DomainResponse{}

	resp, err := s.client.delete(path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}

// ResetDomainToken resets the domain token.
//
// See https://developer.dnsimple.com/v2/domains/#reset-token
func (s *DomainsService) ResetDomainToken(accountID string, domain interface{}) (*DomainResponse, error) {
	path := versioned(domainPath(accountID, domain) + "/token")
	domainResponse := &DomainResponse{}

	resp, err := s.client.post(path, nil, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}
