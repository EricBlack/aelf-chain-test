package api

import (
	"encoding/json"
	"ericshu.io/aelf-chain-test/dto"
	"ericshu.io/aelf-chain-test/utils"
)

type Net struct {
	HostUrl string
	Client  utils.HttpClient
}

func GetNet(hostUrl string) (net Net) {
	net = Net{
		HostUrl: hostUrl,
		Client:  utils.NewClient(hostUrl),
	}
	net.Client.SetDefaultHeaders()

	return
}

func (net Net) GetPeers() (result []dto.PeerDto) {
	response := net.Client.Get(GetPeers, nil)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = []dto.PeerDto{}
	}

	return
}

func (net Net) AddPeer(address string) (result bool) {
	params := map[string]interface{}{
		"address": address,
	}
	response := net.Client.Post(RemovePeer, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = false
	}

	return
}

func (net Net) RemovePeer(address string) (result bool) {
	params := map[string]interface{}{
		"address": address,
	}
	response := net.Client.Delete(RemovePeer, params)
	err := json.Unmarshal([]byte(response), &result)
	if err != nil {
		result = false
	}

	return
}
