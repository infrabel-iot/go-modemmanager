// Code generated by "stringer -type=MMSmsDeliveryState -trimprefix=MmSmsDeliveryState"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsDeliveryStateCompletedReceived-0]
	_ = x[MmSmsDeliveryStateCompletedForwardedUnconfirmed-1]
	_ = x[MmSmsDeliveryStateCompletedReplacedBySc-2]
	_ = x[MmSmsDeliveryStateTemporaryErrorCongestion-32]
	_ = x[MmSmsDeliveryStateTemporaryErrorSmeBusy-33]
	_ = x[MmSmsDeliveryStateTemporaryErrorNoResponseFromSme-34]
	_ = x[MmSmsDeliveryStateTemporaryErrorServiceRejected-35]
	_ = x[MmSmsDeliveryStateTemporaryErrorQosNotAvailable-36]
	_ = x[MmSmsDeliveryStateTemporaryErrorInSme-37]
	_ = x[MmSmsDeliveryStateErrorRemoteProcedure-64]
	_ = x[MmSmsDeliveryStateErrorIncompatibleDestination-65]
	_ = x[MmSmsDeliveryStateErrorConnectionRejected-66]
	_ = x[MmSmsDeliveryStateErrorNotObtainable-67]
	_ = x[MmSmsDeliveryStateErrorQosNotAvailable-68]
	_ = x[MmSmsDeliveryStateErrorNoInterworkingAvailable-69]
	_ = x[MmSmsDeliveryStateErrorValidityPeriodExpired-70]
	_ = x[MmSmsDeliveryStateErrorDeletedByOriginatingSme-71]
	_ = x[MmSmsDeliveryStateErrorDeletedByScAdministration-72]
	_ = x[MmSmsDeliveryStateErrorMessageDoesNotExist-73]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorCongestion-96]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorSmeBusy-97]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorNoResponseFromSme-98]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorServiceRejected-99]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorQosNotAvailable-100]
	_ = x[MmSmsDeliveryStateTemporaryFatalErrorInSme-101]
	_ = x[MmSmsDeliveryStateUnknown-256]
	_ = x[MmSmsDeliveryStateNetworkProblemAddressVacant-512]
	_ = x[MmSmsDeliveryStateNetworkProblemAddressTranslationFailure-513]
	_ = x[MmSmsDeliveryStateNetworkProblemNetworkResourceOutage-514]
	_ = x[MmSmsDeliveryStateNetworkProblemNetworkFailure-515]
	_ = x[MmSmsDeliveryStateNetworkProblemInvalidTeleserviceId-516]
	_ = x[MmSmsDeliveryStateNetworkProblemOther-517]
	_ = x[MmSmsDeliveryStateTerminalProblemNoPageResponse-544]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationBusy-545]
	_ = x[MmSmsDeliveryStateTerminalProblemNoAcknowledgment-546]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationResourceShortage-547]
	_ = x[MmSmsDeliveryStateTerminalProblemSmsDeliveryPostponed-548]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationOutOfService-549]
	_ = x[MmSmsDeliveryStateTerminalProblemDestinationNoLongerAtThisAddress-550]
	_ = x[MmSmsDeliveryStateTerminalProblemOther-551]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemResourceShortage-576]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemIncompatibility-577]
	_ = x[MmSmsDeliveryStateRadioInterfaceProblemOther-578]
	_ = x[MmSmsDeliveryStateGeneralProblemEncoding-608]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsOriginationDenied-609]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsTerminationDenied-610]
	_ = x[MmSmsDeliveryStateGeneralProblemSupplementaryServiceNotSupported-611]
	_ = x[MmSmsDeliveryStateGeneralProblemSmsNotSupported-612]
	_ = x[MmSmsDeliveryStateGeneralProblemMissingExpectedParameter-614]
	_ = x[MmSmsDeliveryStateGeneralProblemMissingMandatoryParameter-615]
	_ = x[MmSmsDeliveryStateGeneralProblemUnrecognizedParameterValue-616]
	_ = x[MmSmsDeliveryStateGeneralProblemUnexpectedParameterValue-617]
	_ = x[MmSmsDeliveryStateGeneralProblemUserDataSizeError-618]
	_ = x[MmSmsDeliveryStateGeneralProblemOther-619]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemAddressVacant-768]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemAddressTranslationFailure-769]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemNetworkResourceOutage-770]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemNetworkFailure-771]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemInvalidTeleserviceId-772]
	_ = x[MmSmsDeliveryStateTemporaryNetworkProblemOther-773]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemNoPageResponse-800]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationBusy-801]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemNoAcknowledgment-802]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationResourceShortage-803]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemSmsDeliveryPostponed-804]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationOutOfService-805]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemDestinationNoLongerAtThisAddress-806]
	_ = x[MmSmsDeliveryStateTemporaryTerminalProblemOther-807]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemResourceShortage-832]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemIncompatibility-833]
	_ = x[MmSmsDeliveryStateTemporaryRadioInterfaceProblemOther-834]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemEncoding-864]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsOriginationDenied-865]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsTerminationDenied-866]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSupplementaryServiceNotSupported-867]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemSmsNotSupported-868]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemMissingExpectedParameter-870]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemMissingMandatoryParameter-871]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUnrecognizedParameterValue-872]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUnexpectedParameterValue-873]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemUserDataSizeError-874]
	_ = x[MmSmsDeliveryStateTemporaryGeneralProblemOther-875]
}

const _MMSmsDeliveryState_name = "CompletedReceivedCompletedForwardedUnconfirmedCompletedReplacedByScTemporaryErrorCongestionTemporaryErrorSmeBusyTemporaryErrorNoResponseFromSmeTemporaryErrorServiceRejectedTemporaryErrorQosNotAvailableTemporaryErrorInSmeErrorRemoteProcedureErrorIncompatibleDestinationErrorConnectionRejectedErrorNotObtainableErrorQosNotAvailableErrorNoInterworkingAvailableErrorValidityPeriodExpiredErrorDeletedByOriginatingSmeErrorDeletedByScAdministrationErrorMessageDoesNotExistTemporaryFatalErrorCongestionTemporaryFatalErrorSmeBusyTemporaryFatalErrorNoResponseFromSmeTemporaryFatalErrorServiceRejectedTemporaryFatalErrorQosNotAvailableTemporaryFatalErrorInSmeUnknownNetworkProblemAddressVacantNetworkProblemAddressTranslationFailureNetworkProblemNetworkResourceOutageNetworkProblemNetworkFailureNetworkProblemInvalidTeleserviceIdNetworkProblemOtherTerminalProblemNoPageResponseTerminalProblemDestinationBusyTerminalProblemNoAcknowledgmentTerminalProblemDestinationResourceShortageTerminalProblemSmsDeliveryPostponedTerminalProblemDestinationOutOfServiceTerminalProblemDestinationNoLongerAtThisAddressTerminalProblemOtherRadioInterfaceProblemResourceShortageRadioInterfaceProblemIncompatibilityRadioInterfaceProblemOtherGeneralProblemEncodingGeneralProblemSmsOriginationDeniedGeneralProblemSmsTerminationDeniedGeneralProblemSupplementaryServiceNotSupportedGeneralProblemSmsNotSupportedGeneralProblemMissingExpectedParameterGeneralProblemMissingMandatoryParameterGeneralProblemUnrecognizedParameterValueGeneralProblemUnexpectedParameterValueGeneralProblemUserDataSizeErrorGeneralProblemOtherTemporaryNetworkProblemAddressVacantTemporaryNetworkProblemAddressTranslationFailureTemporaryNetworkProblemNetworkResourceOutageTemporaryNetworkProblemNetworkFailureTemporaryNetworkProblemInvalidTeleserviceIdTemporaryNetworkProblemOtherTemporaryTerminalProblemNoPageResponseTemporaryTerminalProblemDestinationBusyTemporaryTerminalProblemNoAcknowledgmentTemporaryTerminalProblemDestinationResourceShortageTemporaryTerminalProblemSmsDeliveryPostponedTemporaryTerminalProblemDestinationOutOfServiceTemporaryTerminalProblemDestinationNoLongerAtThisAddressTemporaryTerminalProblemOtherTemporaryRadioInterfaceProblemResourceShortageTemporaryRadioInterfaceProblemIncompatibilityTemporaryRadioInterfaceProblemOtherTemporaryGeneralProblemEncodingTemporaryGeneralProblemSmsOriginationDeniedTemporaryGeneralProblemSmsTerminationDeniedTemporaryGeneralProblemSupplementaryServiceNotSupportedTemporaryGeneralProblemSmsNotSupportedTemporaryGeneralProblemMissingExpectedParameterTemporaryGeneralProblemMissingMandatoryParameterTemporaryGeneralProblemUnrecognizedParameterValueTemporaryGeneralProblemUnexpectedParameterValueTemporaryGeneralProblemUserDataSizeErrorTemporaryGeneralProblemOther"

var _MMSmsDeliveryState_map = map[MMSmsDeliveryState]string{
	0:   _MMSmsDeliveryState_name[0:17],
	1:   _MMSmsDeliveryState_name[17:46],
	2:   _MMSmsDeliveryState_name[46:67],
	32:  _MMSmsDeliveryState_name[67:91],
	33:  _MMSmsDeliveryState_name[91:112],
	34:  _MMSmsDeliveryState_name[112:143],
	35:  _MMSmsDeliveryState_name[143:172],
	36:  _MMSmsDeliveryState_name[172:201],
	37:  _MMSmsDeliveryState_name[201:220],
	64:  _MMSmsDeliveryState_name[220:240],
	65:  _MMSmsDeliveryState_name[240:268],
	66:  _MMSmsDeliveryState_name[268:291],
	67:  _MMSmsDeliveryState_name[291:309],
	68:  _MMSmsDeliveryState_name[309:329],
	69:  _MMSmsDeliveryState_name[329:357],
	70:  _MMSmsDeliveryState_name[357:383],
	71:  _MMSmsDeliveryState_name[383:411],
	72:  _MMSmsDeliveryState_name[411:441],
	73:  _MMSmsDeliveryState_name[441:465],
	96:  _MMSmsDeliveryState_name[465:494],
	97:  _MMSmsDeliveryState_name[494:520],
	98:  _MMSmsDeliveryState_name[520:556],
	99:  _MMSmsDeliveryState_name[556:590],
	100: _MMSmsDeliveryState_name[590:624],
	101: _MMSmsDeliveryState_name[624:648],
	256: _MMSmsDeliveryState_name[648:655],
	512: _MMSmsDeliveryState_name[655:682],
	513: _MMSmsDeliveryState_name[682:721],
	514: _MMSmsDeliveryState_name[721:756],
	515: _MMSmsDeliveryState_name[756:784],
	516: _MMSmsDeliveryState_name[784:818],
	517: _MMSmsDeliveryState_name[818:837],
	544: _MMSmsDeliveryState_name[837:866],
	545: _MMSmsDeliveryState_name[866:896],
	546: _MMSmsDeliveryState_name[896:927],
	547: _MMSmsDeliveryState_name[927:969],
	548: _MMSmsDeliveryState_name[969:1004],
	549: _MMSmsDeliveryState_name[1004:1042],
	550: _MMSmsDeliveryState_name[1042:1089],
	551: _MMSmsDeliveryState_name[1089:1109],
	576: _MMSmsDeliveryState_name[1109:1146],
	577: _MMSmsDeliveryState_name[1146:1182],
	578: _MMSmsDeliveryState_name[1182:1208],
	608: _MMSmsDeliveryState_name[1208:1230],
	609: _MMSmsDeliveryState_name[1230:1264],
	610: _MMSmsDeliveryState_name[1264:1298],
	611: _MMSmsDeliveryState_name[1298:1344],
	612: _MMSmsDeliveryState_name[1344:1373],
	614: _MMSmsDeliveryState_name[1373:1411],
	615: _MMSmsDeliveryState_name[1411:1450],
	616: _MMSmsDeliveryState_name[1450:1490],
	617: _MMSmsDeliveryState_name[1490:1528],
	618: _MMSmsDeliveryState_name[1528:1559],
	619: _MMSmsDeliveryState_name[1559:1578],
	768: _MMSmsDeliveryState_name[1578:1614],
	769: _MMSmsDeliveryState_name[1614:1662],
	770: _MMSmsDeliveryState_name[1662:1706],
	771: _MMSmsDeliveryState_name[1706:1743],
	772: _MMSmsDeliveryState_name[1743:1786],
	773: _MMSmsDeliveryState_name[1786:1814],
	800: _MMSmsDeliveryState_name[1814:1852],
	801: _MMSmsDeliveryState_name[1852:1891],
	802: _MMSmsDeliveryState_name[1891:1931],
	803: _MMSmsDeliveryState_name[1931:1982],
	804: _MMSmsDeliveryState_name[1982:2026],
	805: _MMSmsDeliveryState_name[2026:2073],
	806: _MMSmsDeliveryState_name[2073:2129],
	807: _MMSmsDeliveryState_name[2129:2158],
	832: _MMSmsDeliveryState_name[2158:2204],
	833: _MMSmsDeliveryState_name[2204:2249],
	834: _MMSmsDeliveryState_name[2249:2284],
	864: _MMSmsDeliveryState_name[2284:2315],
	865: _MMSmsDeliveryState_name[2315:2358],
	866: _MMSmsDeliveryState_name[2358:2401],
	867: _MMSmsDeliveryState_name[2401:2456],
	868: _MMSmsDeliveryState_name[2456:2494],
	870: _MMSmsDeliveryState_name[2494:2541],
	871: _MMSmsDeliveryState_name[2541:2589],
	872: _MMSmsDeliveryState_name[2589:2638],
	873: _MMSmsDeliveryState_name[2638:2685],
	874: _MMSmsDeliveryState_name[2685:2725],
	875: _MMSmsDeliveryState_name[2725:2753],
}

func (i MMSmsDeliveryState) String() string {
	if str, ok := _MMSmsDeliveryState_map[i]; ok {
		return str
	}
	return "MMSmsDeliveryState(" + strconv.FormatInt(int64(i), 10) + ")"
}
