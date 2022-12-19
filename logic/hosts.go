package logic

import (
	"encoding/json"
	"fmt"

	"github.com/gravitl/netmaker/database"
	"github.com/gravitl/netmaker/models"
)

// GetAllHosts - returns all hosts in flat list or error
func GetAllHosts() ([]models.Host, error) {
	currHostMap, err := GetHostsMap()
	if err != nil {
		return nil, err
	}
	var currentHosts = []models.Host{}
	for k := range currHostMap {
		var h = *currHostMap[k]
		currentHosts = append(currentHosts, h)
	}

	return currentHosts, nil
}

// GetHostsMap - gets all the current hosts on machine in a map
func GetHostsMap() (map[string]*models.Host, error) {
	records, err := database.FetchRecords(database.HOSTS_TABLE_NAME)
	if err != nil && !database.IsEmptyRecord(err) {
		return nil, err
	}
	currHostMap := make(map[string]*models.Host)
	for k := range records {
		var h models.Host
		err = json.Unmarshal([]byte(records[k]), &h)
		if err != nil {
			return nil, err
		}
		currHostMap[h.ID.String()] = &h
	}

	return currHostMap, nil
}

// GetHost - gets a host from db given id
func GetHost(hostid string) (*models.Host, error) {
	record, err := database.FetchRecord(database.HOSTS_TABLE_NAME, hostid)
	if err != nil {
		return nil, err
	}

	var h models.Host
	if err = json.Unmarshal([]byte(record), &h); err != nil {
		return nil, err
	}

	return &h, nil
}

// CreateHost - creates a host if not exist
func CreateHost(h *models.Host) error {
	_, err := GetHost(h.ID.String())
	if (err != nil && !database.IsEmptyRecord(err)) || (err == nil) {
		return fmt.Errorf("host already exists")
	}

	return UpsertHost(h)
}

// UpdateHost - updates host data by field
func UpdateHost(newHost, currentHost *models.Host) {
	// unchangeable fields via API here
	newHost.DaemonInstalled = currentHost.DaemonInstalled
	newHost.OS = currentHost.OS
	newHost.IPForwarding = currentHost.IPForwarding
	newHost.HostPass = currentHost.HostPass
	newHost.NodePassword = currentHost.NodePassword
	newHost.MacAddress = currentHost.MacAddress
	newHost.Debug = currentHost.Debug
	newHost.Nodes = currentHost.Nodes
	newHost.PublicKey = currentHost.PublicKey
	newHost.InternetGateway = currentHost.InternetGateway
	newHost.TrafficKeyPublic = currentHost.TrafficKeyPublic

	// changeable fields
	if len(newHost.Version) == 0 {
		newHost.Version = currentHost.Version
	}

	if len(newHost.Name) == 0 {
		newHost.Name = currentHost.Name
	}

	if newHost.LocalAddress.String() != currentHost.LocalAddress.String() {
		newHost.LocalAddress = currentHost.LocalAddress
	}

	if newHost.LocalRange.String() != currentHost.LocalRange.String() {
		newHost.LocalRange = currentHost.LocalRange
	}

	if newHost.MTU == 0 {
		newHost.MTU = currentHost.MTU
	}

	if newHost.ListenPort == 0 {
		newHost.ListenPort = currentHost.ListenPort
	}

	if newHost.ProxyListenPort == 0 {
		newHost.ProxyListenPort = currentHost.ProxyListenPort
	}
}

// UpsertHost - upserts into DB a given host model, does not check for existence*
func UpsertHost(h *models.Host) error {
	data, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return database.Insert(h.ID.String(), string(data), database.HOSTS_TABLE_NAME)
}

// RemoveHost - removes a given host from server
func RemoveHost(h *models.Host) error {
	if len(h.Nodes) > 0 {
		for i := range h.Nodes {
			id := h.Nodes[i]
			n, err := GetNodeByID(id)
			if err == nil {
				if err = DeleteNodeByID(&n); err != nil {
					return err // must remove associated nodes before removing a host
				}
			}
		}
	}
	return database.DeleteRecord(database.HOSTS_TABLE_NAME, h.ID.String())
}
