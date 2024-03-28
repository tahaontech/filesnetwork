package p2p

type HandShakeFunc func(any) error

func NopHandShakeFunc(any) error { return nil }
