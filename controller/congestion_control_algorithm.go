package controller

import "log"

type CongestionControlAlgorithm int

const (
	Reno CongestionControlAlgorithm = iota
	Cubic
	BBR
	SCReAM
	GCC
)

func CongestionControlAlgorithmFromString(a string) CongestionControlAlgorithm {
	switch a {
	case "reno":
		return Reno
	case "cubic":
		return Cubic
	case "bbr":
		return BBR
	case "scream":
		return SCReAM
	case "gcc":
		return GCC
	default:
		log.Printf("warning, unknown algorithm: %v, using default ('reno')", a)
		return Reno
	}
}

func (a CongestionControlAlgorithm) String() string {
	switch a {
	case Reno:
		return "reno"
	case Cubic:
		return "cubic"
	case BBR:
		return "bbr"
	case SCReAM:
		return "scream"
	case GCC:
		return "gcc"
	default:
		log.Printf("warning, undefined algorithm: %v", a)
		return "none"
	}
}
