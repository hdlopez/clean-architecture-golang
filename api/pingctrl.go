package api

type pingCtrl struct {
}

func newPingCtrl() *pingCtrl {
	return new(pingCtrl)
}

func (ctrl *pingCtrl) Ping(c Ctx) {
	c.String(200, "pong")
}
