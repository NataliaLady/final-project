package src

import (
	"example.com/types"
	"example.com/types/data"
)

func GetResultData(
	sliceSMSData []data.SMSData, sliceMMSData []data.MMSData,
	sliceVoiceData []data.VoiceData,
	sliceEmailData []data.EmailData,
	billingData data.BillingData,
	sliceSupportData []data.SupportData,
	sliceIncidentData []data.IncidentData,
	countries []types.Country) data.Result {

	sms := prepareSMSData(sliceSMSData, countries)
	mms := prepareMMSData(sliceMMSData, countries)
	email := prepareEmailData(sliceEmailData, countries)
	support := prepareSupportData(sliceSupportData)
	incident := prepareIncidentData(sliceIncidentData)

	resultSetT := data.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: sliceVoiceData,
		Email:     email,
		Billing:   billingData,
		Support:   support,
		Incidents: incident,
	}

	status := sms != nil && mms != nil && sliceVoiceData != nil && email != nil && resultSetT.Billing == billingData && support != nil && sliceIncidentData != nil
	resultT := data.Result{
		Status: status,
	}

	if !status {
		resultT.Error = "Ошибка сбора данных"
	} else {
		resultT.Data = resultSetT
	}

	return resultT
}
