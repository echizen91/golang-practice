package easy

// ParkingSystem : data struct to store parking spaces
type ParkingSystem struct {
	big    int
	medium int
	small  int
}

// ConstructorParking : constructor for initializing parking spaces
func ConstructorParking(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

// AddCar : function to add car to parking space
func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.big > 0 {
			p.big--
			return true
		}
	case 2:
		if p.medium > 0 {
			p.medium--
			return true
		}
	case 3:
		if p.small > 0 {
			p.small--
			return true
		}
	default:
		return false

	}
	return false
}
