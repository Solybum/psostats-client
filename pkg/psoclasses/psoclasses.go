package psoclasses

import "errors"

type PsoClass struct {
	Name      string
	MaxShifta int
	MinAtp    int
	MaxAtp    int
	Ata       int
}

var HUmar = PsoClass{Name: "HUmar", MaxShifta: 7, MinAtp: 1492, MaxAtp: 1497, Ata: 210}
var HUnewearl = PsoClass{Name: "HUnewearl", MaxShifta: 20, MinAtp: 1332, MaxAtp: 1337, Ata: 209}
var HUcast = PsoClass{Name: "HUcast", MaxShifta: 3, MinAtp: 1734, MaxAtp: 1739, Ata: 201}
var HUcaseal = PsoClass{Name: "HUcaseal", MaxShifta: 7, MinAtp: 1396, MaxAtp: 1401, Ata: 228}
var RAmar = PsoClass{Name: "RAmar", MaxShifta: 15, MinAtp: 1356, MaxAtp: 1360, Ata: 259}
var RAmarl = PsoClass{Name: "RAmarl", MaxShifta: 20, MinAtp: 1241, MaxAtp: 1245, Ata: 251}
var RAcast = PsoClass{Name: "RAcast", MaxShifta: 7, MinAtp: 1446, MaxAtp: 1450, Ata: 234}
var RAcaseal = PsoClass{Name: "RAcaseal", MaxShifta: 7, MinAtp: 1271, MaxAtp: 1275, Ata: 241}
var FOmar = PsoClass{Name: "FOmar", MaxShifta: 30, MinAtp: 1100, MaxAtp: 1102, Ata: 173}
var FOmarl = PsoClass{Name: "FOmarl", MaxShifta: 30, MinAtp: 970, MaxAtp: 972, Ata: 180}
var FOnewm = PsoClass{Name: "FOnewm", MaxShifta: 30, MinAtp: 912, MaxAtp: 914, Ata: 190}
var FOnewearl = PsoClass{Name: "FOnewearl", MaxShifta: 30, MinAtp: 681, MaxAtp: 683, Ata: 196}

func ForName(name string) (PsoClass, error) {
	switch name {
	case "HUmar":
		return HUmar, nil
	case "HUnewearl":
		return HUnewearl, nil
	case "HUcast":
		return HUcast, nil
	case "HUcaseal":
		return HUcaseal, nil
	case "RAmar":
		return RAmar, nil
	case "RAmarl":
		return RAmarl, nil
	case "RAcast":
		return RAcast, nil
	case "RAcaseal":
		return RAcaseal, nil
	case "FOmar":
		return FOmar, nil
	case "FOmarl":
		return FOmarl, nil
	case "FOnewm":
		return FOnewm, nil
	case "FOnewearl":
		return FOnewearl, nil
	}
	return PsoClass{}, errors.New("invalid class")
}

func GetAll() []PsoClass {
	return []PsoClass{
		HUmar,
		HUnewearl,
		HUcast,
		HUcaseal,
		RAmar,
		RAmarl,
		RAcast,
		RAcaseal,
		FOmar,
		FOmarl,
		FOnewm,
		FOnewearl,
	}
}
