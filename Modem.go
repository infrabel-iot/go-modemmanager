package go_modemmanager

import (
	"errors"
	"github.com/godbus/dbus/v5"
	"reflect"
)

// The Modem interface controls the status and actions in a given modem object.
// This interface will always be available as long a the modem is considered valid.

const (
	ModemInterface = ModemManagerInterface + ".Modem"

	ModemsObjectPath = modemManagerMainObjectPath + "Modems"

	/* Methods */
	ModemEnable  = ModemInterface + ".Enable"
	ModemDisable = ModemEnable
	// Deprecated ModemListBearers            = ModemInterface + ".ListBearers"
	ModemCreateBearer           = ModemInterface + ".CreateBearer"
	ModemDeleteBearer           = ModemInterface + ".DeleteBearer"
	ModemReset                  = ModemInterface + ".Reset"
	ModemFactoryReset           = ModemInterface + ".FactoryReset"
	ModemSetPowerState          = ModemInterface + ".SetPowerState"
	ModemSetCurrentCapabilities = ModemInterface + ".SetCurrentCapabilities"
	ModemSetCurrentModes        = ModemInterface + ".SetCurrentModes"
	ModemSetCurrentBands        = ModemInterface + ".SetCurrentBands"
	ModemCommand                = ModemInterface + ".Command"

	/* Property */

	ModemPropertySim                          = ModemInterface + ".Sim"                          //                           readable   o
	ModemPropertyBearers                      = ModemInterface + ".Bearers"                      //    readable   ao
	ModemPropertySupportedCapabilities        = ModemInterface + ".SupportedCapabilities"        //    readable   au
	ModemPropertyCurrentCapabilities          = ModemInterface + ".CurrentCapabilities"          //     readable   u
	ModemPropertyMaxBearers                   = ModemInterface + ".MaxBearers"                   //     readable   u
	ModemPropertyMaxActiveBearers             = ModemInterface + ".MaxActiveBearers"             //     readable   u
	ModemPropertyManufacturer                 = ModemInterface + ".Manufacturer"                 //     readable   s
	ModemPropertyModel                        = ModemInterface + ".Model"                        //     readable   s
	ModemPropertyRevision                     = ModemInterface + ".Revision"                     //     readable   s
	ModemPropertyCarrierConfiguration         = ModemInterface + ".CarrierConfiguration"         //    readable   s
	ModemPropertyCarrierConfigurationRevision = ModemInterface + ".CarrierConfigurationRevision" //  readable   s
	ModemPropertyHardwareRevision             = ModemInterface + ".HardwareRevision"             //       readable   s
	ModemPropertyDeviceIdentifier             = ModemInterface + ".DeviceIdentifier"             //         readable   s
	ModemPropertyDevice                       = ModemInterface + ".Device"                       //         readable   s
	ModemPropertyDrivers                      = ModemInterface + ".Drivers"                      //         readable   as
	ModemPropertyPlugin                       = ModemInterface + ".Plugin"                       //         readable   s
	ModemPropertyPrimaryPort                  = ModemInterface + ".PrimaryPort"                  //         readable   s
	ModemPropertyPorts                        = ModemInterface + ".Ports"                        //        readable   a(su)
	ModemPropertyEquipmentIdentifier          = ModemInterface + ".EquipmentIdentifier"          //         readable   s
	ModemPropertyUnlockRequired               = ModemInterface + ".UnlockRequired"               //         readable   u
	ModemPropertyUnlockRetries                = ModemInterface + ".UnlockRetries"                //          readable   a{uu}
	ModemPropertyState                        = ModemInterface + ".State"                        //         readable   i
	ModemPropertyStateFailedReason            = ModemInterface + ".StateFailedReason"            //         readable   u
	ModemPropertyAccessTechnologies           = ModemInterface + ".AccessTechnologies"           //         readable   u
	ModemPropertySignalQuality                = ModemInterface + ".SignalQuality"                //           readable   (ub)
	ModemPropertyOwnNumbers                   = ModemInterface + ".OwnNumbers"                   //           readable   as
	ModemPropertyPowerState                   = ModemInterface + ".PowerState"                   //         readable   u
	ModemPropertySupportedModes               = ModemInterface + ".SupportedModes"               //           readable   a(uu)
	ModemPropertyCurrentModes                 = ModemInterface + ".CurrentModes"                 //           readable   (uu)
	ModemPropertySupportedBands               = ModemInterface + ".SupportedBands"               //           readable   au
	ModemPropertyCurrentBands                 = ModemInterface + ".CurrentBands"                 //          readable   au
	ModemPropertySupportedIpFamilies          = ModemInterface + ".SupportedIpFamilies"          //         readable   u

)

type Modem interface {
	/* METHODS */

	// Returns object path
	GetObjectPath() dbus.ObjectPath

	// Return ModemSimple Object
	GetSimpleModem() (ModemSimple,error)
	// Enables the Modem: When enabled, the modem's radio is powered on and data sessions, voice calls,
	// location services, and Short Message Service may be available.
	Enable() error
	// Disable the Modem: When disabled, the modem enters low-power state and no network-related operations are available.
	Disable() error
	// Deprecated: List configured packet data bearers (EPS Bearers, PDP Contexts, or CDMA2000 Packet Data Sessions).
	// ListBearers() ([]Bearer, error)

	// Create a new packet data bearer using the given characteristics.
	// This request may fail if the modem does not support additional bearers,
	// if too many bearers are already defined, or if properties are invalid.
	// see BearerProperty struct
	//
	// Some properties are only applicable to a bearer of certain access technologies, for example the "apn" property is
	// not applicable to CDMA2000 Packet Data Session bearers.
	CreateBearer(BearerProperty) (Bearer, error)

	// If the bearer is currently active and providing packet data server, it will be disconnected and that packet data service will terminate.
	DeleteBearer(bearer Bearer) error

	// Clear non-persistent configuration and state, and return the device to a newly-powered-on state.
	// This command may power-cycle the device.
	Reset() error

	// Clear the modem's configuration (including persistent configuration and state), and return the device to a
	// factory-default state.
	// If not required by the modem, code may be ignored. This command may or may not power-cycle the device.
	FactoryReset(code string) error

	// Set the power state of the modem. This action can only be run when the modem is in MM_MODEM_STATE_DISABLED state.
	SetPowerState(MMModemPowerState) error

	// Set the capabilities of the device. A restart of the modem may be required. Bitmask of MMModemCapability values, to specify the capabilities to use.
	SetCurrentCapabilities([]MMModemCapability) error

	// Set the access technologies (e.g. 2G/3G/4G preference) the device is currently allowed to use when connecting to a network.
	// The given combination should be supported by the modem, as specified in the "SupportedModes" property.
	// A pair of MMModemMode values, where the first one is a bitmask of allowed modes, and the second one the preferred mode, if any.
	SetCurrentModes(SupportedModesProperty) error

	// Set the radio frequency and technology bands the device is currently allowed to use when connecting to a network.
	// List of MMModemBand values, to specify the bands to be used.
	SetCurrentBands([]MMModemBand) error

	// AT command for the Modem (operation only allowed in debug mode)
	Command(cmd string, timeout uint32) (string, error)

	// StateChanged (i old,	i new,	u reason);
	// The modem's state (see "State") changed.
	// i old: A MMModemState value, specifying the new state.
	// i new: A MMModemState value, specifying the new state.
	// u reason: A MMModemStateChangeReason value, specifying the reason for this state change.
	Subscribe() <-chan *dbus.Signal
	Unsubscribe()


	/* METHODS to get Properties */
	// The path of the SIM object available in this device, if any.
	GetSim() (Sim, error)

	// The list of bearer object paths (EPS Bearers, PDP Contexts, or CDMA2000 Packet Data Sessions) as requested by the user.
	// This list does not include the initial EPS bearer details (see "InitialEpsBearer").
	GetBearers() ([]Bearer, error)

	// List of MMModemCapability values, specifying the combinations of generic family of access technologies the modem supports.
	// If the modem doesn't allow changing the current capabilities, a single entry with MM_MODEM_CAPABILITY_ANY will be given.
	// It's an array of bitmasks because the modem may support different combinations
	// (E.g. "gsm/umts+lte" or "cdma/evdo+lte" or "gsm/umts+cdma/evdo+lte".
	// This property gives you the list of combinations supported, Then, you have CurrentCapabilities, which gives you the actual combination in use currently.
	GetSupportedCapabilities() ([][]MMModemCapability, error)

	// Bitmask of MMModemCapability values, specifying the generic family of access technologies the modem currently supports without a firmware reload or reinitialization.
	GetCurrentCapabilities() ([]MMModemCapability, error)

	// The maximum number of defined packet data bearers the modem supports.
	// This is not the number of active/connected bearers the modem supports,
	// but simply the number of bearers that may be defined at any given time.
	// For example, POTS and CDMA2000-only devices support only one bearer, while GSM/UMTS devices
	// typically support three or more, and any LTE-capable device (whether LTE-only, GSM/UMTS-capable,
	// and/or CDMA2000-capable) also typically support three or more.
	GetMaxBearers() (uint32, error)

	// The maximum number of active MM_BEARER_TYPE_DEFAULT bearers that may be explicitly enabled by the user.
	// POTS and CDMA2000-only devices support one active bearer, while GSM/UMTS and LTE-capable devices
	//(including LTE/CDMA devices) typically support at least two active bearers.
	GetMaxActiveBearers() (uint32, error)

	// The equipment manufacturer, as reported by the modem.
	GetManufacturer() (string, error)

	// The equipment model, as reported by the modem.
	GetModel() (string, error)

	// The revision identification of the software, as reported by the modem.
	GetRevision() (string, error)

	// The description of the carrier-specific configuration (MCFG) in use by the modem.
	GetCarrierConfiguration() (string, error)

	// The revision identification of the carrier-specific configuration (MCFG) in use by the modem.
	GetCarrierConfigurationRevision() (string, error)

	// The revision identification of the hardware, as reported by the modem.
	GetHardwareRevision() (string, error)

	// A best-effort device identifier based on various device information like model name, firmware
	// revision, USB/PCI/PCMCIA IDs, and other properties.
	// This ID is not guaranteed to be unique and may be shared between identical devices with the same firmware,
	// but is intended to be "unique enough" for use as a casual device identifier for various user experience operations.
	// his is not the device's IMEI or ESN since those may not be available before unlocking the device via a PIN.
	GetDeviceIdentifier() (string, error)

	// The physical modem device reference (ie, USB, PCI, PCMCIA device), which may be dependent upon the operating system.
	// In Linux for example, this points to a sysfs path of the usb_device object.
	// This value may also be set by the user using the MM_ID_PHYSDEV_UID udev tag (e.g. binding the tag to a
	// specific sysfs path).
	GetDevice() (string, error)

	// The Operating System device drivers handling communication with the modem hardware.
	GetDriver() ([]string, error)

	// The name of the plugin handling this modem.
	GetPlugin() (string, error)

	// The name of the primary port using to control the modem.
	GetPrimaryPort() (string, error)

	// The list of ports in the modem, given as an array of string and unsigned integer pairs.
	// The string is the port name or path, and the integer is the port type given as a MMModemPortType value.
	GetPorts() ([]Port, error)

	// The identity of the device. This will be the IMEI number for GSM devices and the hex-format ESN/MEID for CDMA devices.
	GetEquipmentIdentifier() (string, error)

	// Current lock state of the device, given as a MMModemLock value.
	GetUnlockRequired() (MMModemLock, error)

	// A dictionary in which the keys are MMModemLock flags, and the values are integers giving the number of PIN tries
	// remaining before the code becomes blocked (requiring a PUK) or permanently blocked. Dictionary entries exist
	// only for the codes for which the modem is able to report retry counts.
	GetUnlockRetries() ([]Pair, error)

	// Overall state of the modem, given as a MMModemState value.
	// If the device's state cannot be determined, MM_MODEM_STATE_UNKNOWN will be reported.
	GetState() (MMModemState, error)

	// Error specifying why the modem is in MM_MODEM_STATE_FAILED state, given as a MMModemStateFailedReason value.
	GetStateFailedReason() (MMModemStateFailedReason, error)

	// Bitmask of MMModemAccessTechnology values, specifying the current network access technologies used by the device
	// to communicate with the network.
	// If the device's access technology cannot be determined, MM_MODEM_ACCESS_TECHNOLOGY_UNKNOWN will be reported.
	GetAccessTechnologies() ([]MMModemAccessTechnology, error)

	// Signal quality in percent (0 - 100) of the dominant access technology the device is using to communicate with the network. Always 0 for POTS devices.
	// The additional boolean value indicates if the quality value given was recently taken.
	GetSignalQuality() (percent uint32, recent bool, err error)

	// List of numbers (e.g. MSISDN in 3GPP) being currently handled by this modem.
	GetOwnNumbers() ([]string, error)

	// A MMModemPowerState value specifying the current power state of the modem.
	GetPowerState() (MMModemPowerState, error)

	// This property exposes the supported mode combinations, given as an array of unsigned integer pairs, where:
	// The first integer is a bitmask of MMModemMode values, specifying the allowed modes.
	// The second integer is a single MMModemMode, which specifies the preferred access technology, among the ones defined in the allowed modes.
	GetSupportedModes() ([]Mode, error)

	// A pair of MMModemMode values, where the first one is a bitmask specifying the access technologies (eg 2G/3G/4G) the device is currently allowed to use when connecting to a network, and the second one is the preferred mode of those specified as allowed.
	// The pair must be one of those specified in "SupportedModes".
	GetCurrentModes() (Mode, error)

	// List of MMModemBand values, specifying the radio frequency and technology bands supported by the device.
	// For POTS devices, only the MM_MODEM_BAND_ANY mode will be returned.
	GetSupportedBands() ([]MMModemBand, error)

	// List of MMModemBand values, specifying the radio frequency and technology bands the device is currently using when connecting to a network.
	// It must be a subset of "SupportedBands".
	GetCurrentBands() ([]MMModemBand, error)

	// Bitmask of MMBearerIpFamily values, specifying the IP families supported by the device.
	GetSupportedIpFamilies() ([]MMBearerIpFamily, error)

	MarshalJSON() ([]byte, error)
}

func NewModem(objectPath dbus.ObjectPath) (Modem, error) {
	var m modem
	return &m, m.init(ModemManagerInterface, objectPath)
}

type modem struct {
	dbusBase
	sigChan chan *dbus.Signal
}

type Port struct {
	PortName string          // Port Name or Path
	PortType MMModemPortType // Modem Port Type
}
type Mode struct {
	AllowedModes  []MMModemMode
	PreferredMode MMModemMode
}
func (m modem) GetObjectPath() dbus.ObjectPath {
	return m.obj.Path()
}

func (m modem) GetSimpleModem()(ModemSimple,error){
	return NewModemSimple(m.obj.Path())
}

func (m modem) Enable() error {
	err := m.call(ModemEnable, true)
	return err
}

func (m modem) Disable() error {
	err := m.call(ModemDisable, false)
	return err
}

func (m modem) CreateBearer(property BearerProperty) (Bearer, error) {
	// untested
	v := reflect.ValueOf(property)
	st := reflect.TypeOf(property)
	type dynMap interface{}
	var myMap map[string]dynMap
	myMap = make(map[string]dynMap)
	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i).Interface()
		if v.Field(i).IsZero() {
			continue
		}
		myMap[tag] = value
	}
	var path dbus.ObjectPath
	err := m.callWithReturn(&path, ModemCreateBearer, &myMap)
	if err != nil {
		return nil, err
	}
	return NewBearer(path)
}

func (m modem) DeleteBearer(bearer Bearer) error {
	// todo: not implemented - ModemDeleteBearer
	panic("implement me")
}

func (m modem) Reset() error {
	return m.call(ModemReset)
}

func (m modem) FactoryReset(code string) error {
	// todo: untested
	return m.call(ModemFactoryReset, code)
}

func (m modem) SetPowerState(state MMModemPowerState) error {
	// handle with care ...
	return m.call(ModemSetPowerState, state)
}

func (m modem) SetCurrentCapabilities(capabilities []MMModemCapability) error {
	// todo: untested
	var caps MMModemCapability
	err := m.call(ModemSetCurrentCapabilities, caps.SliceToBitmask(capabilities))
	return err
}

func (m modem) SetCurrentModes(property SupportedModesProperty) error {
	// todo: untested
	var mode MMModemMode
	var resSlice = []uint32{mode.SliceToBitmask(property.AllowedModes),
		mode.SliceToBitmask([]MMModemMode{property.PreferredMode})}
	return m.call(ModemSetCurrentModes, resSlice)
}

func (m modem) SetCurrentBands(bands []MMModemBand) error {
	// todo: untested
	return m.call(ModemSetCurrentBands, bands)
}

func (m modem) Command(cmd string, timeout uint32) (response string, err error) {
	// untested - only works in debug mode
	err = m.callWithReturn(&response, ModemCommand, cmd, timeout)
	return
}

func (m modem) GetSim() (Sim, error) {
	simPath, err := m.getObjectProperty(ModemPropertySim)
	if err != nil {
		return nil, err
	}
	return NewSim(simPath)
}

func (m modem) GetBearers() ([]Bearer, error) {
	bearerPaths, err := m.getSliceObjectProperty(ModemPropertyBearers)
	if err != nil {
		return nil, err
	}
	var bearers []Bearer
	for idx := range bearerPaths {
		bearer, err := NewBearer(bearerPaths[idx])
		if err != nil {
			return nil, err
		}
		bearers = append(bearers, bearer)
	}
	return bearers, nil
}

func (m modem) GetSupportedCapabilities() (capabilities [][]MMModemCapability, err error) {
	caps, err := m.getSliceUint32Property(ModemPropertySupportedCapabilities)
	if err != nil {
		return nil, err
	}
	var capability MMModemCapability
	for _, c := range caps {
		capabilities = append(capabilities, capability.BitmaskToSlice(c))

	}
	return

}

func (m modem) GetCurrentCapabilities() (capabilities []MMModemCapability, err error) {
	res, err := m.getUint32Property(ModemPropertyCurrentCapabilities)
	if err != nil {
		return nil, err
	}
	var capability MMModemCapability
	return capability.BitmaskToSlice(res), nil
}

func (m modem) GetMaxBearers() (maxBearers uint32, err error) {

	return m.getUint32Property(ModemPropertyMaxBearers)
}

func (m modem) GetMaxActiveBearers() (uint32, error) {
	return m.getUint32Property(ModemPropertyMaxActiveBearers)
}

func (m modem) GetManufacturer() (string, error) {
	return m.getStringProperty(ModemPropertyManufacturer)
}

func (m modem) GetModel() (string, error) {
	return m.getStringProperty(ModemPropertyModel)
}

func (m modem) GetRevision() (string, error) {
	return m.getStringProperty(ModemPropertyRevision)
}

func (m modem) GetCarrierConfiguration() (string, error) {
	return m.getStringProperty(ModemPropertyCarrierConfiguration)
}

func (m modem) GetCarrierConfigurationRevision() (string, error) {
	return m.getStringProperty(ModemPropertyCarrierConfigurationRevision)
}

func (m modem) GetHardwareRevision() (string, error) {
	return m.getStringProperty(ModemPropertyHardwareRevision)
}

func (m modem) GetDeviceIdentifier() (string, error) {
	return m.getStringProperty(ModemPropertyDeviceIdentifier)
}

func (m modem) GetDevice() (string, error) {
	return m.getStringProperty(ModemPropertyDevice)
}

func (m modem) GetDriver() ([]string, error) {
	return m.getSliceStringProperty(ModemPropertyDrivers)
}

func (m modem) GetPlugin() (string, error) {
	return m.getStringProperty(ModemPropertyPlugin)
}

func (m modem) GetPrimaryPort() (string, error) {
	return m.getStringProperty(ModemPropertyPrimaryPort)
}

func (m modem) GetPorts() (ports []Port, err error) {
	res, err := m.getSliceSlicePairProperty(ModemPropertyPorts)
	if err != nil {
		return nil, err
	}
	for _, pair := range res {
		newA, ok := pair.a.(string)
		if !ok {
			return nil, errors.New("wrong type != string")
		}
		newB, ok := pair.b.(uint32)
		if !ok {

			return nil, errors.New("wrong type != uin32")
		}
		ports = append(ports, Port{PortName: newA, PortType: MMModemPortType(newB)})
	}
	return
}

func (m modem) GetEquipmentIdentifier() (string, error) {
	return m.getStringProperty(ModemPropertyEquipmentIdentifier)
}

func (m modem) GetUnlockRequired() (MMModemLock, error) {

	res, err := m.getUint32Property(ModemPropertyUnlockRequired)
	if err != nil {
		return MmModemLockUnknown, err
	}
	return MMModemLock(res), nil
}

func (m modem) GetUnlockRetries() (values []Pair, err error) {
	res, err := m.getMapUint32Uint32Property(ModemPropertyUnlockRetries)
	if err != nil {
		return nil, err
	}
	for key, element := range res {
		values = append(values, Pair{a: MMModemLock(key), b: element})
	}
	return values, nil
}

func (m modem) GetState() (MMModemState, error) {

	res, err := m.getInt32Property(ModemPropertyState)
	if err != nil {
		return MmModemStateUnknown, err
	}
	return MMModemState(res), nil
}

func (m modem) GetStateFailedReason() (MMModemStateFailedReason, error) {
	res, err := m.getUint32Property(ModemPropertyStateFailedReason)
	if err != nil {
		return MmModemStateFailedReasonUnknown, err
	}

	return MMModemStateFailedReason(res), nil
}

func (m modem) GetAccessTechnologies() (result []MMModemAccessTechnology, err error) {
	res, err := m.getUint32Property(ModemPropertyAccessTechnologies)
	if err != nil {
		return nil, err
	}
	var tmp MMModemAccessTechnology

	return tmp.BitmaskToSlice(res), nil
}

func (m modem) GetSignalQuality() (percent uint32, recent bool, err error) {
	res, err := m.getPairProperty(ModemPropertySignalQuality)
	if err != nil {
		return
	}
	return res.a.(uint32), res.b.(bool), nil
}

func (m modem) GetOwnNumbers() ([]string, error) {
	return m.getSliceStringProperty(ModemPropertyOwnNumbers)

}

func (m modem) GetPowerState() (MMModemPowerState, error) {
	res, err := m.getUint32Property(ModemPropertyPowerState)
	if err != nil {
		return MmModemPowerStateUnknown, err
	}
	return MMModemPowerState(res), nil
}

func (m modem) GetSupportedModes() (modes []Mode, err error) {

	res, err := m.getSliceSlicePairProperty(ModemPropertySupportedModes)

	if err != nil {
		return nil, err
	}
	var tmp MMModemMode
	for _, x := range res {
		modes = append(modes, Mode{AllowedModes: tmp.BitmaskToSlice(x.a.(uint32)), PreferredMode: MMModemMode(x.b.(uint32))})

	}
	return

}
func (m modem) GetCurrentModes() (mode Mode, err error) {
	res, err := m.getPairProperty(ModemPropertyCurrentModes)
	if err != nil {
		return mode, err
	}
	var tmp MMModemMode
	mode.PreferredMode = MMModemMode(res.b.(uint32))
	mode.AllowedModes = tmp.BitmaskToSlice(res.a.(uint32))

	return
}

func (m modem) GetSupportedBands() (bands []MMModemBand, err error) {
	res, err := m.getSliceUint32Property(ModemPropertySupportedBands)
	if err != nil {
		return nil, err
	}
	for _, x := range res {
		bands = append(bands, MMModemBand(x))

	}
	return

}

func (m modem) GetCurrentBands() (bands []MMModemBand, err error) {
	res, err := m.getSliceUint32Property(ModemPropertyCurrentBands)
	if err != nil {
		return nil, err
	}
	for _, x := range res {
		bands = append(bands, MMModemBand(x))

	}
	return
}

func (m modem) GetSupportedIpFamilies() ([]MMBearerIpFamily, error) {
	res, err := m.getUint32Property(ModemPropertySupportedIpFamilies)
	if err != nil {
		return nil, err
	}
	var ipFam MMBearerIpFamily

	return ipFam.BitmaskToSlice(res), nil
}

func (m modem) Subscribe() <-chan *dbus.Signal {
	if m.sigChan != nil {
		return m.sigChan
	}

	m.subscribeNamespace(ModemManagerObjectPath)
	m.sigChan = make(chan *dbus.Signal, 10)
	m.conn.Signal(m.sigChan)

	return m.sigChan
}

func (m modem) Unsubscribe() {
	m.conn.RemoveSignal(m.sigChan)
	m.sigChan = nil
}

func (m modem) MarshalJSON() ([]byte, error) {
	// todo: not implemented yet
	panic("implement me")
}

// todo duplicate
type SupportedModesProperty struct {
	AllowedModes  []MMModemMode // the first one is a bitmask of allowed modes
	PreferredMode MMModemMode   // the second one the preferred mode, if any
}