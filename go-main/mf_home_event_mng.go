package main

type MfHomeEventMng struct {
}
type WebActionParams struct {
}
type WebCtrl struct {
}

type WebViewData struct {
}

func (m *MfHomeEventMng) UploadEventImg(rParams *WebActionParams, rCtrl *WebCtrl, rData *WebViewData) int {

	m.UploadEventImgUtils(rParams, rCtrl, rData)

	return 0
}
