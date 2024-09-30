package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UPnP device description template
const deviceDescription = `<?xml version="1.0"?>
<root xmlns="urn:schemas-upnp-org:device-1-0" configId="{{.ConfigId}}">
	<specVersion>
		<major>1</major>
		<minor>0</minor>
	</specVersion>
	<device>
		<deviceType>urn:schemas-upnp-org:device:MediaRenderer:3</deviceType>
		<friendlyName>{{.FriendlyName}}</friendlyName>
		<manufacturer>-</manufacturer>
		<modelDescription></modelDescription>
		<modelName>{{.ModelName}}</modelName>
		<modelNumber>{{.ModelNumber}}</modelNumber>
		<UDN>uuid:{{.DeviceUUID}}</UDN>
		<serviceList>
			<service>
				<serviceType>urn:schemas-upnp-org:service:RenderingControl:3</serviceType>
				<serviceId>urn:upnp-org:serviceId:RenderingControl</serviceId>
				<SCPDURL>/upnp/scpd/RenderingControl3.xml</SCPDURL>
				<controlURL>/upnp/MediaRenderer/RenderingControl3.xml</controlURL>
				<eventSubURL>/upnp/non-existing-eventsuburl</eventSubURL>
			</service>
			<service>
				<serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>
				<serviceId>urn:upnp-org:serviceId:ConnectionManager</serviceId>
				<SCPDURL>/upnp/scpd/ConnectionManager1.xml</SCPDURL>
				<controlURL>/upnp/MediaRenderer/ConnectionManager1.xml</controlURL>
				<eventSubURL>/upnp/non-existing-eventsuburl</eventSubURL>
			</service>
		</serviceList>
	</device>
</root>
`

const scpdRenderingControl = `<?xml version="1.0" encoding="utf-8"?>
<scpd xmlns="urn:schemas-upnp-org:service-1-0">
  <specVersion>
    <major>1</major>
    <minor>0</minor>
  </specVersion>
  <serviceStateTable>
    <stateVariable sendEvents="yes">
      <name>LastChange</name>
      <dataType>string</dataType>
    </stateVariable>
    <stateVariable sendEvents="no">
      <name>Volume</name>
      <dataType>ui2</dataType>
      <allowedValueRange>
        <minimum>0</minimum>
        <maximum>100</maximum>
        <step>1</step>
      </allowedValueRange>
    </stateVariable>
  </serviceStateTable>
  <actionList>
    <action>
      <name>GetVolume</name>
      <argumentList>
        <argument>
          <name>CurrentVolume</name>
          <direction>out</direction>
          <relatedStateVariable>Volume</relatedStateVariable>
        </argument>
      </argumentList>
    </action>
    <action>
      <name>SetVolume</name>
      <argumentList>
        <argument>
          <name>DesiredVolume</name>
          <direction>in</direction>
          <relatedStateVariable>Volume</relatedStateVariable>
        </argument>
      </argumentList>
    </action>
  </actionList>
</scpd>`

var (
	deviceUUID = uuid.New()
)

func (s *ApiServer) upnpServeDescription(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Application-URL", fmt.Sprintf("http://10.65.42.114/%d/upnp/apps/", s.port))

	tpl := template.Must(template.New("").Parse(deviceDescription))

	w.Header().Set("Content-Type", "text/xml; charset=utf-8")

	// TODO: XML escape some of these values
	tpl.Execute(w, map[string]interface{}{
		"ConfigId":     1,              // update on every description change
		"FriendlyName": "go-librespot", // TODO: make configurable
		"ModelName":    "go-librespot",
		"ModelNumber":  "0.0.1",    // TODO: make configurable
		"DeviceUUID":   deviceUUID, // TODO: base on the device_id?
	})
}

type xmlEnvelopeGetVolumeResponse struct {
	CurrentVolume int
}

type xmlEnvelopeSetVolume struct {
	DesiredVolume int
}

type xmlEnvelopeBody struct {
	GetVolume         *struct{}                     `xml:"urn:schemas-upnp-org:service:RenderingControl:3 GetVolume"`
	GetVolumeResponse *xmlEnvelopeGetVolumeResponse `xml:"urn:schemas-upnp-org:service:RenderingControl:3 GetVolumeResponse"`
	SetVolume         *xmlEnvelopeSetVolume         `xml:"urn:schemas-upnp-org:service:RenderingControl:3 SetVolume"`
	SetVolumeResponse *struct{}                     `xml:"urn:schemas-upnp-org:service:RenderingControl:3 SetVolumeResponse"`
}

type xmlEnvelope struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    xmlEnvelopeBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func (s *ApiServer) upnpRenderingControl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read XML SOAP request.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("could not read XML SOAP")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var result xmlEnvelope
	err = xml.Unmarshal(body, &result)
	if err != nil {
		log.WithError(err).Error("could not unmarshal XML SOAP")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process request.
	envelope := xmlEnvelope{
		Body: xmlEnvelopeBody{},
	}
	req := ApiRequest{
		resp: make(chan apiResponse, 1),
	}
	switch {
	case result.Body.SetVolume != nil:
		req.Type = ApiRequestTypeSetVolume
		req.Data = ApiRequestDataVolume{
			Volume: int32(result.Body.SetVolume.DesiredVolume),
		}
		s.requests <- req
		<-req.resp
		envelope.Body.SetVolumeResponse = &struct{}{}
	case result.Body.GetVolume != nil:
		req.Type = ApiRequestTypeGetVolume
		s.requests <- req
		resp := <-req.resp
		data := resp.data.(*ApiResponseVolume)
		envelope.Body.GetVolumeResponse = &xmlEnvelopeGetVolumeResponse{
			CurrentVolume: int(data.Value),
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	// Send the response to the client.
	resp, err := xml.Marshal(&envelope)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err).Error("could not marshal XML SOAP")
		return
	}
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>`))
	w.Write(resp)
}
