package core

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

type DeviceUsage struct {
	ID            string      `json:"id"`
	ShortID       string      `json:"short_id"`
	Hostname      string      `json:"hostname"`
	Description   interface{} `json:"description"`
	Tags          []string    `json:"tags"`
	ImageURL      interface{} `json:"image_url"`
	BillingCycle  string      `json:"billing_cycle"`
	User          string      `json:"user"`
	Iqn           string      `json:"iqn"`
	Locked        bool        `json:"locked"`
	BondingMode   int         `json:"bonding_mode"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	IpxeScriptURL interface{} `json:"ipxe_script_url"`
	AlwaysPxe     bool        `json:"always_pxe"`
	Storage       struct {
	} `json:"storage"`
	Customdata struct {
	} `json:"customdata"`
	CreatedBy struct {
		ID             string    `json:"id"`
		ShortID        string    `json:"short_id"`
		FirstName      string    `json:"first_name"`
		LastName       string    `json:"last_name"`
		FullName       string    `json:"full_name"`
		Email          string    `json:"email"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Level          string    `json:"level"`
		AvatarThumbURL string    `json:"avatar_thumb_url"`
		Href           string    `json:"href"`
	} `json:"created_by"`
	OperatingSystem struct {
		ID             string `json:"id"`
		Slug           string `json:"slug"`
		Name           string `json:"name"`
		Distro         string `json:"distro"`
		Version        string `json:"version"`
		Preinstallable bool   `json:"preinstallable"`
		Pricing        struct {
		} `json:"pricing"`
		DistroLabel     string        `json:"distro_label"`
		ProvisionableOn []interface{} `json:"provisionable_on"`
		Licensed        bool          `json:"licensed"`
	} `json:"operating_system"`
	Facility struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Code     string   `json:"code"`
		Features []string `json:"features"`
		Metro    struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country string `json:"country"`
		} `json:"metro"`
		IPRanges []interface{} `json:"ip_ranges"`
		Address  struct {
			ID          string `json:"id"`
			Address     string `json:"address"`
			Address2    string `json:"address2"`
			City        string `json:"city"`
			State       string `json:"state"`
			ZipCode     string `json:"zip_code"`
			Country     string `json:"country"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
		} `json:"address"`
	} `json:"facility"`
	Project struct {
		Href string `json:"href"`
	} `json:"project"`
	SSHKeys []struct {
		Href string `json:"href"`
	} `json:"ssh_keys"`
	ProjectLite struct {
		Href string `json:"href"`
	} `json:"project_lite"`
	Volumes     []interface{} `json:"volumes"`
	IPAddresses []struct {
		ID            string        `json:"id"`
		AddressFamily int           `json:"address_family"`
		Netmask       string        `json:"netmask"`
		CreatedAt     time.Time     `json:"created_at"`
		Details       interface{}   `json:"details"`
		Tags          []interface{} `json:"tags"`
		Public        bool          `json:"public"`
		Cidr          int           `json:"cidr"`
		Management    bool          `json:"management"`
		Manageable    bool          `json:"manageable"`
		Enabled       bool          `json:"enabled"`
		GlobalIP      bool          `json:"global_ip"`
		Customdata    struct {
		} `json:"customdata"`
		Project struct {
			Href string `json:"href"`
		} `json:"project"`
		ProjectLite struct {
			Href string `json:"href"`
		} `json:"project_lite"`
		AssignedTo struct {
			Href string `json:"href"`
		} `json:"assigned_to"`
		Interface struct {
			Href string `json:"href"`
		} `json:"interface"`
		Network  string `json:"network"`
		Address  string `json:"address"`
		Gateway  string `json:"gateway"`
		Href     string `json:"href"`
		Facility struct {
			ID       string   `json:"id"`
			Name     string   `json:"name"`
			Code     string   `json:"code"`
			Features []string `json:"features"`
			Metro    struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Code    string `json:"code"`
				Country string `json:"country"`
			} `json:"metro"`
			IPRanges []interface{} `json:"ip_ranges"`
			Address  struct {
				ID          string `json:"id"`
				Address     string `json:"address"`
				Address2    string `json:"address2"`
				City        string `json:"city"`
				State       string `json:"state"`
				ZipCode     string `json:"zip_code"`
				Country     string `json:"country"`
				Coordinates struct {
					Latitude  string `json:"latitude"`
					Longitude string `json:"longitude"`
				} `json:"coordinates"`
			} `json:"address"`
		} `json:"facility"`
		Metro struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country string `json:"country"`
		} `json:"metro"`
	} `json:"ip_addresses"`
	Favorite bool `json:"favorite"`
	Metro    struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Code    string `json:"code"`
		Country string `json:"country"`
	} `json:"metro"`
	Plan struct {
		ID          string `json:"id"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Line        string `json:"line"`
		Specs       struct {
			Cpus []struct {
				Count int    `json:"count"`
				Type  string `json:"type"`
			} `json:"cpus"`
			Memory struct {
				Total string `json:"total"`
			} `json:"memory"`
			Drives []struct {
				Count    int    `json:"count"`
				Size     string `json:"size"`
				Type     string `json:"type"`
				Category string `json:"category"`
			} `json:"drives"`
			Nics []struct {
				Count int    `json:"count"`
				Type  string `json:"type"`
			} `json:"nics"`
			Features struct {
				Raid bool `json:"raid"`
				Txt  bool `json:"txt"`
			} `json:"features"`
		} `json:"specs"`
		Legacy          bool     `json:"legacy"`
		DeploymentTypes []string `json:"deployment_types"`
		Type            string   `json:"type"`
		Class           string   `json:"class"`
		Pricing         struct {
			Hour float64 `json:"hour"`
		} `json:"pricing"`
		ReservationPricing struct {
			OneMonth struct {
				Month float64 `json:"month"`
			} `json:"one_month"`
			OneYear struct {
				Month float64 `json:"month"`
			} `json:"one_year"`
			ThreeYear struct {
				Month float64 `json:"month"`
			} `json:"three_year"`
			Pa struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"pa"`
			Dc struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"dc"`
			Ty struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"ty"`
			Ny struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"ny"`
			Sp struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"sp"`
			Am struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"am"`
			Hk struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"hk"`
			Da struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"da"`
			La struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"la"`
			Sg struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"sg"`
			Se struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"se"`
			Ch struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"ch"`
			Tr struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"tr"`
			Sy struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"sy"`
			Fr struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"fr"`
			Md struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"md"`
			Sl struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"sl"`
			Sv struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"sv"`
			Ld struct {
				OneMonth struct {
					Month float64 `json:"month"`
				} `json:"one_month"`
				OneYear struct {
					Month float64 `json:"month"`
				} `json:"one_year"`
				ThreeYear struct {
					Month float64 `json:"month"`
				} `json:"three_year"`
			} `json:"ld"`
		} `json:"reservation_pricing"`
		AvailableIn []struct {
			Href  string `json:"href"`
			Price struct {
				Hour float64 `json:"hour"`
			} `json:"price"`
		} `json:"available_in"`
		AvailableInMetros []struct {
			Href  string `json:"href"`
			Price struct {
				Hour float64 `json:"hour"`
			} `json:"price"`
		} `json:"available_in_metros"`
	} `json:"plan"`
	DeviceType string `json:"device_type"`
	Actions    []struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"actions"`
	NetworkFrozen bool   `json:"network_frozen"`
	Userdata      string `json:"userdata"`
	SwitchUUID    string `json:"switch_uuid"`
	NetworkPorts  []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
		Data struct {
			Bonded bool   `json:"bonded"`
			Mac    string `json:"mac"`
		} `json:"data,omitempty"`
		Bond struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"bond,omitempty"`
		NativeVirtualNetwork      interface{}   `json:"native_virtual_network"`
		VirtualNetworks           []interface{} `json:"virtual_networks"`
		DisbondOperationSupported bool          `json:"disbond_operation_supported"`
		Href                      string        `json:"href"`
		NetworkType               string        `json:"network_type,omitempty"`
	} `json:"network_ports"`
	State                        string `json:"state"`
	AllowSameVlanOnMultiplePorts bool   `json:"allow_same_vlan_on_multiple_ports"`
	Href                         string `json:"href"`
}

func (m Metrics) GetDeviceInfo(deviceUUID string, createdDate string, createdBefore string) (*DeviceUsage, error) {

	deviceUsage := DeviceUsage{}
	url := EquinixEndpoint +
		deviceUUID +
		"?created%5Bafter%5D=" + createdDate + "&created%5Bbefore%5D=" + createdBefore

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return &DeviceUsage{}, err
	}
	req.Header.Add("X-Auth-Token", viper.Get("EQUINIX_AUTH_TOKEN").(string))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &DeviceUsage{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &DeviceUsage{}, err
	}
	fmt.Println(string(body))
	errMarshal := json.Unmarshal(body, &deviceUsage)
	if errMarshal != nil {
		return nil, errMarshal
	}
	fmt.Println(deviceUsage)
	return &deviceUsage, nil
}
