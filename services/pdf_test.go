package services

import (
	"testing"
)

func TestAddWatermarks(t *testing.T) {

	for _, tt := range []struct {
		msg             string
		inFile, outFile string
		selectedPages   []string
		mode            string
		modeParm        string
		wmConf          string
	}{
		{"TestWatermarkPDF",
			"./(Everything I Do）Do It For You.pdf",
			"A Thousand Years.pdf",
			nil,
			"pdf",
			"../data/cover.pdf",
			"sc:1, pos:tr, off:-10 0, rot:0"},
	}{
		AddWatermarks( tt.msg, tt.inFile, tt.outFile, tt.selectedPages, tt.mode, tt.modeParm, tt.wmConf, false)
	}
}

