package gonsx

import (
	"github.com/sky-uk/gonsx/api/securitytags"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClientGetAllSecurityTags(t *testing.T) {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<securityTags><securityTag><objectId>securitytag-12</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>Block_All_Ports</name><description>Security tag used to demonstrate all ports being blocked.</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-19</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>rr</name><description>rr</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-15</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>2</revision><type><typeName>SecurityTag</typeName></type><name>DEMO_WEB_SERVER</name><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>1</vmCount></securityTag><securityTag><objectId>securitytag-16</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>SecurityTag</typeName></type><name>DEMO_APP_SERVER</name><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-17</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>1</revision><type><typeName>SecurityTag</typeName></type><name>DEMO_DB_SERVER</name><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>false</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-1</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>VULNERABILITY_MGMT.VulnerabilityFound.threat=high</name><description>Tag indicates that the vulnerability found has a high threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-2</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>ANTI_VIRUS.VirusFound.threat=low</name><description>Tag indicates that the viruses found have a low threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-3</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>ANTI_VIRUS.VirusFound.threat=medium</name><description>Tag indicates that the viruses found have a medium threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-4</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>IDS_IPS.threat=high</name><description>Tag indicates that the data violation detected has a high threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-5</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>DATA_SECURITY.violationsFound</name><description>Tag indicates that data violation was detected</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-6</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>IDS_IPS.threat=low</name><description>Tag indicates that the data violation detected has a low threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-7</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>AntiVirus.virusFound</name><description>This tag is now deprecated. Please use ANTI_VIRUS.* tags</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-8</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>VULNERABILITY_MGMT.VulnerabilityFound.threat=low</name><description>Tag indicates that the vulnerability found has a low threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-9</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>VULNERABILITY_MGMT.VulnerabilityFound.threat=medium</name><description>Tag indicates that the vulnerability found has a medium threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-10</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>IDS_IPS.threat=medium</name><description>Tag indicates that the data violation detected has a medium threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag><securityTag><objectId>securitytag-11</objectId><objectTypeName>SecurityTag</objectTypeName><vsmUuid>4221A849-079E-D13E-6B36-068D4F1222A9</vsmUuid><nodeId>dd3b6a28-b778-4310-8803-b6eae482b2c0</nodeId><revision>0</revision><type><typeName>SecurityTag</typeName></type><name>ANTI_VIRUS.VirusFound.threat=high</name><description>Tag indicates that the viruses found have a high threat level</description><clientHandle></clientHandle><extendedAttributes/><isUniversal>false</isUniversal><universalRevision>0</universalRevision><systemResource>true</systemResource><vmCount>0</vmCount></securityTag></securityTags>`
	setup(http.StatusOK, xmlContent)
	defer server.Close()

	api := securitytags.NewGetAll()
	nsxClient.Do(api)

	assert.Len(t, api.GetResponse().SecurityTags, 16)
	assert.Equal(t, "securitytag-12", api.GetResponse().SecurityTags[0].ObjectID)
	assert.Equal(t, "Security tag used to demonstrate all ports being blocked.", api.GetResponse().SecurityTags[0].Description)
	assert.Equal(t, "Block_All_Ports", api.GetResponse().SecurityTags[0].Name)
	assert.Equal(t, "SecurityTag", api.GetResponse().SecurityTags[0].TypeName)
	assert.Equal(t, "securitytag-19", api.GetResponse().SecurityTags[1].ObjectID)
	assert.Equal(t, "rr", api.GetResponse().SecurityTags[1].Description)
	assert.Equal(t, "rr", api.GetResponse().SecurityTags[1].Name)
	assert.Equal(t, "SecurityTag", api.GetResponse().SecurityTags[1].TypeName)

}