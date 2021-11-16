package main

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "maverick gun",
			power: 5,
		},
	}
}
