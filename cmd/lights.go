package cmd


type DeconzConfiguration struct {
	Utc                      string      `json:"UTC"`
	Announceinterval         int         `json:"announceinterval"`
	Announceurl              string      `json:"announceurl"`
	Apiversion               string      `json:"apiversion"`
	Bridgeid                 string      `json:"bridgeid"`
	Datastoreversion         string      `json:"datastoreversion"`
	Devicename               string      `json:"devicename"`
	Dhcp                     bool        `json:"dhcp"`
	DisablePermitJoinAutoOff bool        `json:"disablePermitJoinAutoOff"`
	Discovery                bool        `json:"discovery"`
	Factorynew               bool        `json:"factorynew"`
	Fwneedupdate             bool        `json:"fwneedupdate"`
	Fwupdatestate            string      `json:"fwupdatestate"`
	Fwversion                string      `json:"fwversion"`
	Fwversionupdate          string      `json:"fwversionupdate"`
	Gateway                  string      `json:"gateway"`
	Groupdelay               int         `json:"groupdelay"`
	Homebridge               string      `json:"homebridge"`
	Homebridgepin            interface{} `json:"homebridgepin"`
	Homebridgeupdate         bool        `json:"homebridgeupdate"`
	Homebridgeupdateversion  interface{} `json:"homebridgeupdateversion"`
	Homebridgeversion        interface{} `json:"homebridgeversion"`
	Ipaddress                string      `json:"ipaddress"`
	Lightlastseeninterval    int         `json:"lightlastseeninterval"`
	Linkbutton               bool        `json:"linkbutton"`
	Localtime                string      `json:"localtime"`
	Mac                      string      `json:"mac"`
	Modelid                  string      `json:"modelid"`
	Name                     string      `json:"name"`
	Netmask                  string      `json:"netmask"`
	Networkopenduration      int         `json:"networkopenduration"`
	Otauactive               bool        `json:"otauactive"`
	Otaustate                string      `json:"otaustate"`
	Panid                    int         `json:"panid"`
	Permitjoin               int         `json:"permitjoin"`
	Permitjoinfull           int         `json:"permitjoinfull"`
	Port                     int         `json:"port"`
	Portalservices           bool        `json:"portalservices"`
	Proxyaddress             string      `json:"proxyaddress"`
	Proxyport                int         `json:"proxyport"`
	Replacesbridgeid         interface{} `json:"replacesbridgeid"`
	Rfconnected              bool        `json:"rfconnected"`
	Runmode                  string      `json:"runmode"`
	Starterkitid             string      `json:"starterkitid"`
	Swcommit                 string      `json:"swcommit"`
	Swupdate                 struct {
		Notify      bool   `json:"notify"`
		Text        string `json:"text"`
		Updatestate int    `json:"updatestate"`
		URL         string `json:"url"`
		Version     string `json:"version"`
	} `json:"swupdate"`
	Swupdate2 struct {
		Autoinstall struct {
			On         bool   `json:"on"`
			Updatetime string `json:"updatetime"`
		} `json:"autoinstall"`
		Bridge struct {
			Lastinstall string `json:"lastinstall"`
			State       string `json:"state"`
		} `json:"bridge"`
		Checkforupdate bool   `json:"checkforupdate"`
		Lastchange     string `json:"lastchange"`
		State          string `json:"state"`
	} `json:"swupdate2"`
	Swversion          string `json:"swversion"`
	System             string `json:"system"`
	Timeformat         string `json:"timeformat"`
	Timezone           string `json:"timezone"`
	Updatechannel      string `json:"updatechannel"`
	UUID               string `json:"uuid"`
	Websocketnotifyall bool   `json:"websocketnotifyall"`
	Websocketport      int    `json:"websocketport"`
	Whitelist          struct {
		Zero97E0464C4 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"097E0464C4"`
		ZeroC322DF78A struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"0C322DF78A"`
		Four0BCF69D5D struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"40BCF69D5D"`
		FourE6D332084 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"4E6D332084"`
		Five0E4F04085 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"50E4F04085"`
		FiveAFA341B5D struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"5AFA341B5D"`
		FiveD33EBED74 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"5D33EBED74"`
		Seven1E9811709 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"71E9811709"`
		Eight7207A9FC9 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"87207A9FC9"`
		Eight7B73BC800 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"87B73BC800"`
		EightDC179EA6B struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"8DC179EA6B"`
		EightFE2467DF1 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"8FE2467DF1"`
		Num9018417029 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"9018417029"`
		Nine71660D323 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"971660D323"`
		B801831F72 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"B801831F72"`
		Cc0D5B655A struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"CC0D5B655A"`
		D4241Cc77B struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"D4241CC77B"`
		D9Db61E3Aa struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"D9DB61E3AA"`
		Edf21Dda49 struct {
			CreateDate  string `json:"create date"`
			LastUseDate string `json:"last use date"`
			Name        string `json:"name"`
		} `json:"EDF21DDA49"`
	} `json:"whitelist"`
	Wifi           string        `json:"wifi"`
	Wifiavailable  []interface{} `json:"wifiavailable"`
	Wifichannel    string        `json:"wifichannel"`
	Wificlientname interface{}   `json:"wificlientname"`
	Wifiip         string        `json:"wifiip"`
	Wifimgmt       int           `json:"wifimgmt"`
	Wifiname       interface{}   `json:"wifiname"`
	Wifitype       string        `json:"wifitype"`
	Zigbeechannel  int           `json:"zigbeechannel"`
}