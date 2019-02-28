package scte224v20180501

import (
	"encoding/xml"
	"testing"
)

const viewingpolicy string = `
<?xml version="1.0" encoding="utf-8"?>
<ViewingPolicy 
	xmlns:action="urn:scte:224:action" 
  xmlns:xlink="http://www.w3.org/1999/xlink" 
	xmlns:audience="urn:scte:224:audience" 
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
  xmlns:metadata="urn:scte:224:metadata" id="test.com/viewingpolicy/test" lastUpdated="2019-02-22T16:03:19.907Z" 
  xmlns="http://www.scte.org/schemas/224">
  <Audience xlink:href="test.com/audience/all" />
	<action:FastForward>false</action:FastForward>
	<action:Content>TARGETSTREAM</action:Content>
  <action:Capture>
    <action:StartWindow>
      <action:Percentage>0</action:Percentage>
    </action:StartWindow>
    <action:StopWindow>
      <action:Percentage>250</action:Percentage>
    </action:StopWindow>
    <action:MidrollDAI>false</action:MidrollDAI>
  </action:Capture>
</ViewingPolicy>`

type testViewingPolicy struct {
	ReusableType
	ActionProperty []*AnyProperty `xml:"http://www.scte.org/schemas/224 Any,omitempty" json:"anys,omitempty"`
	XMLName        xml.Name       `xml:"http://www.scte.org/schemas/224 ViewingPolicy" json:"-"`
	Audience       *Audience      `xml:"http://www.scte.org/schemas/224 Audience" json:"audience,omitempty"`
	FastForward    string         `xml:"FastForward,omitempty" json:"fastForward,omitempty"`
	Capture        []*testCapture `xml:"Capture,omitempty" "json:"capture,omitempty"`
	Content        string         `xml:"Content,omitempty" "json:"content,omitempty"`
}

//Capture struct
type testCapture struct {
	StartWindow *testStartWindow `xml:"StartWindow,omitempty" json:"startWindow"`
	StopWindow  *testStopWindow  `xml:"StopWindow,omitempty" json:"stopWindow"`
	MidrollDAI  string           `xml:"MidrollDAI,omitempty" json:"midrollDAI"`
}

//StartWindow struct
type testStartWindow struct {
	Percentage string `xml:"Percentage,omitempty" json:"percentage"`
}

//StopWindow struct
type testStopWindow struct {
	Offset     string `xml:"Offset,omitempty" json:"offset"`
	Percentage string `xml:"Percentage,omitempty" json:"percentage"`
}

func TestViewingPolicy2018(t *testing.T) {

	var vpol *ViewingPolicy
	err := xml.Unmarshal([]byte(viewingpolicy), &vpol)
	if err != nil {
		t.Errorf("Error unmarshalling viewingpolicys %v", err)
		t.FailNow()
	}

	vbyte, marshalErr := xml.Marshal(vpol)
	if marshalErr != nil {
		t.Errorf("Error Marshal %v", err)
		t.FailNow()
	}

	var tvpol *testViewingPolicy
	err = xml.Unmarshal(vbyte, &tvpol)
	if err != nil {
		t.Errorf("Error unmarshalling %v", err)
		t.FailNow()
	}

	if tvpol.FastForward == "" {
		t.Log("action:FastForward is empty")
		t.FailNow()
	}

	if tvpol.Content == "" {
		t.Log("action:Content is empty")
		t.FailNow()
	}

	// Capture
	if len(tvpol.Capture) == 0 {
		t.Log("action:Capture is empty and is supposed to have child nodes")
		t.FailNow()
	}

	for _, capture := range tvpol.Capture {
		// StartWindow
		if capture.StartWindow == nil {
			t.Log("action:Capture StartWindow is nil")
			t.FailNow()
		}
		if capture.StartWindow.Percentage == "" {
			t.Log("action:Capture StartWindow Percentage is empty")
			t.Fail()
		}

		// StopWindow
		if capture.StopWindow == nil {
			t.Log("action:Capture StopWindow is nil")
			t.FailNow()
		}

		if capture.StopWindow.Percentage == "" {
			t.Log("action:Capture StopWindow Percentage is empty")
			t.Fail()
		}

		// MidrollDAI
		if capture.MidrollDAI == "" {
			t.Log("action:Capture MidrollDAI is empty")
			t.Fail()
		}
	}
}
