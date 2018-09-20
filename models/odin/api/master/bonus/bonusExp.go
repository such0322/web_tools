package bonus

type BonusExp struct {
}

func (m *BonusExp) GetRewardNames() {
}

func (m *BonusExp) GetType() string {
	return TypeEXP
}

func (m *BonusExp) ToJson() string {
	return ""
}
