package bonus

type BonusMileage struct {
}

func (m *BonusMileage) GetRewardNames() {
}

func (m *BonusMileage) GetType() string {
	return TypeMileage
}

func (m *BonusMileage) ToJson() string {
	return ""
}
