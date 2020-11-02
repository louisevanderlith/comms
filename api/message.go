package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"io/ioutil"
	"net/http"
)

func FetchMessage(web *http.Client, host string, k hsk.Key) (core.Message, error) {
	url := fmt.Sprintf("%s/message/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Message{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Message{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Message{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchLatestMessages(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/message/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Message{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
