package dnsimple

import (
	"fmt"
)

// ZoneRecord represents a DNS record in DNSimple.
type ZoneRecord struct {
	ID           int    `json:"id,omitempty"`
	ZoneID       string `json:"zone_id,omitempty"`
	ParentID     int    `json:"parent_id,omitempty"`
	Type         string `json:"type,omitempty"`
	Name         string `json:"name"`
	Content      string `json:"content,omitempty"`
	TTL          int    `json:"ttl,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	SystemRecord bool   `json:"system_record,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// ZoneRecordResponse represents a response from an API method that returns a ZoneRecord struct.
type ZoneRecordResponse struct {
	Response
	Data *ZoneRecord `json:"data"`
}

// ZoneRecordsResponse represents a response from an API method that returns a collection of ZoneRecord struct.
type ZoneRecordsResponse struct {
	Response
	Data []ZoneRecord `json:"data"`
}

func zoneRecordPath(accountID string, zoneID string, recordID int) string {
	path := fmt.Sprintf("/%v/zones/%v/records", accountID, zoneID)

	if recordID != 0 {
		path += fmt.Sprintf("/%d", recordID)
	}

	return path
}

// ListRecords lists the zone records for a zone.
//
// See https://developer.dnsimple.com/v2/zones/#list
func (s *ZonesService) ListRecords(accountID string, zoneID string, options *ListOptions) (*ZoneRecordsResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneID, 0))
	recordsResponse := &ZoneRecordsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(path, recordsResponse)
	if err != nil {
		return nil, err
	}

	recordsResponse.HttpResponse = resp
	return recordsResponse, nil
}

// CreateRecord creates a zone record.
//
// See https://developer.dnsimple.com/v2/zones/#create
func (s *ZonesService) CreateRecord(accountID string, zoneID string, recordAttributes ZoneRecord) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneID, 0))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.post(path, recordAttributes, recordResponse)
	if err != nil {
		return nil, err
	}

	recordResponse.HttpResponse = resp
	return recordResponse, nil
}

// GetRecord fetches a zone record.
//
// See https://developer.dnsimple.com/v2/zones/#get
func (s *ZonesService) GetRecord(accountID string, zoneID string, recordID int) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneID, recordID))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.get(path, recordResponse)
	if err != nil {
		return nil, err
	}

	recordResponse.HttpResponse = resp
	return recordResponse, nil
}

// UpdateRecord updates a zone record.
//
// See https://developer.dnsimple.com/v2/zones/#update
func (s *ZonesService) UpdateRecord(accountID string, zoneID string, recordID int, recordAttributes ZoneRecord) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneID, recordID))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.patch(path, recordAttributes, recordResponse)
	if err != nil {
		return nil, err
	}

	recordResponse.HttpResponse = resp
	return recordResponse, nil
}

// DeleteRecord PERMANENTLY deletes a zone record from the zone.
//
// See https://developer.dnsimple.com/v2/zones/#delete
func (s *ZonesService) DeleteRecord(accountID string, zoneID string, recordID int) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneID, recordID))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.delete(path, nil, nil)
	if err != nil {
		return nil, err
	}

	recordResponse.HttpResponse = resp
	return recordResponse, nil
}
