package go_modemmanager

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

// This interface provides access to extended signal quality information.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	ModemSignalInterface = ModemInterface + ".Signal"

	/* Methods */
	ModemSignalSetup = ModemSignalInterface + ".Setup"
	/* Property */
	ModemSignalPropertyRate = ModemSignalInterface + ".Rate"
	ModemSignalPropertyCdma = ModemSignalInterface + ".Cdma"
	ModemSignalPropertyEvdo = ModemSignalInterface + ".Evdo"
	ModemSignalPropertyGsm  = ModemSignalInterface + ".Gsm"
	ModemSignalPropertyUmts = ModemSignalInterface + ".Umts"
	ModemSignalPropertyLte  = ModemSignalInterface + ".Lte"
)

type ModemSignal interface {
	/* METHODS */

	// Returns object path
	GetObjectPath() dbus.ObjectPath

	MarshalJSON() ([]byte, error)

	// Setup extended signal quality information retrieval.
	// refresh rate to set, in seconds. 0 to disable retrieval.
	Setup(rate uint32) error

	/* PROPERTIES */
	//Refresh rate for the extended signal quality information updates, in seconds. A value of 0 disables the retrieval of the values.
	GetRate() (rate uint32, err error)

	// The CDMA1x access technology.
	GetCdma() (SignalProperty, error)

	// The CDMA EV-DO access technology.
	GetEvdo() (SignalProperty, error)

	// The GSM/GPRS access technology.
	GetGsm() (SignalProperty, error)

	// The UMTS (WCDMA)  access technology.
	GetUmts() (SignalProperty, error)

	// The LTE access technology.
	GetLte() (SignalProperty, error)
}

func NewModemSignal(objectPath dbus.ObjectPath) (ModemSignal, error) {
	var si modemSignal
	return &si, si.init(ModemManagerInterface, objectPath)
}

type modemSignal struct {
	dbusBase
}

type SignalProperty struct {
	Type MMSignalPropertyType `json:"property-type"` // define the Signal Property Type
	Rssi float64              `json:"rssi"`          // The CDMA1x / CDMA EV-DO / GSM / UMTS / LTE RSSI (Received Signal Strength Indication), in dBm, given as a floating point value (Applicable for all types).
	Ecio float64              `json:"sinr"`          // The CDMA1x Ec/Io / CDMA EV-DO Ec/Io / UMTS Ec/Io level in dBm, given as a floating point value (Only applicable for type Cdma, Evdo, Umts).
	Sinr float64              `json:"ecio"`          // CDMA EV-DO SINR level, in dB, given as a floating point value (Only applicable for type Evdo).
	Io   float64              `json:"io"`            // The CDMA EV-DO Io, in dBm, given as a floating point value (Only applicable for type Evdo).
	Rscp float64              `json:"rscp"`          // The UMTS RSCP (Received Signal Code Power), in dBm, given as a floating point value (Only applicable for type Umts).
	Rsrq float64              `json:"rsrq"`          // The LTE RSRP (Reference Signal Received Power), in dBm, given as a floating point value (Only applicable for type LTE).
	Rsrp float64              `json:"rsrp"`          // The LTE RSRP (Reference Signal Received Power), in dBm, given as a floating point value (Only applicable for type LTE).
	Snr  float64              `json:"snr"`           // The LTE S/R ratio, in dB, given as a floating point value (Only applicable for type LTE).
}

func (sp SignalProperty) String() string {
	return "Type: " + fmt.Sprint(sp.Type) +
		" Rssi: " + fmt.Sprint(sp.Rssi) +
		" Ecio: " + fmt.Sprint(sp.Ecio) +
		" Sinr: " + fmt.Sprint(sp.Sinr) +
		" Io: " + fmt.Sprint(sp.Io) +
		" Rscp: " + fmt.Sprint(sp.Rscp) +
		" Rsrq: " + fmt.Sprint(sp.Rsrq) +
		" Rsrp: " + fmt.Sprint(sp.Rsrp) +
		" Snr: " + fmt.Sprint(sp.Snr)
}

func (sp SignalProperty) ReadMap(inputMap map[string]dbus.Variant, signalType MMSignalPropertyType) () {
	sp.Type = signalType
	for key, element := range inputMap {
		switch key {
		case "rssi":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Rssi = tmpValue
			}

		case "sinr":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Sinr = tmpValue
			}

		case "ecio":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Ecio = tmpValue
			}

		case "io":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Io = tmpValue
			}

		case "rscp":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Rscp = tmpValue
			}

		case "rsrq":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Rsrq = tmpValue
			}

		case "rsrp":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Rsrp = tmpValue
			}

		case "snr":
			tmpValue, ok := element.Value().(float64)
			if ok {
				sp.Snr = tmpValue
			}

		}
	}
}

func (si modemSignal) GetObjectPath() dbus.ObjectPath {
	return si.obj.Path()
}

func (si modemSignal) Setup(rate uint32) error {
	return si.call(ModemSignalSetup, &rate)
}

func (si modemSignal) GetRate() (rate uint32, err error) {
	return si.getUint32Property(ModemSignalPropertyRate)
}

func (si modemSignal) GetCdma() (sp SignalProperty, err error) {
	res, err := si.getMapStringVariantProperty(ModemSignalPropertyCdma)
	if err != nil {
		return
	}
	sp.ReadMap(res, MMSignalPropertyTypeCdma)
	return
}

func (si modemSignal) GetEvdo() (sp SignalProperty, err error) {
	res, err := si.getMapStringVariantProperty(ModemSignalPropertyEvdo)
	if err != nil {
		return
	}
	sp.ReadMap(res, MMSignalPropertyTypeEvdo)
	return
}

func (si modemSignal) GetGsm() (sp SignalProperty, err error) {
	res, err := si.getMapStringVariantProperty(ModemSignalPropertyGsm)
	if err != nil {
		return
	}
	sp.ReadMap(res, MMSignalPropertyTypeGsm)
	return
}

func (si modemSignal) GetUmts() (sp SignalProperty, err error) {
	res, err := si.getMapStringVariantProperty(ModemSignalPropertyUmts)
	if err != nil {
		return
	}
	sp.ReadMap(res, MMSignalPropertyTypeUmts)
	return
}

func (si modemSignal) GetLte() (sp SignalProperty, err error) {
	res, err := si.getMapStringVariantProperty(ModemSignalPropertyLte)
	if err != nil {
		return
	}
	sp.ReadMap(res, MMSignalPropertyTypeLte)
	return
}

func (si modemSignal) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
