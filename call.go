package rip

type call struct {
	SimpleCall
}

func UpgradeSimpleCall(simple SimpleCall) Call {
	return &call{
		SimpleCall: simple,
	}
}

func (c *call) Arg(arg interface{}) Call {
	c.SetArg(arg)
	return c
}

func (c *call) Return(into interface{}) Error {
	resp := c.Response()
	if e := resp.Err(); e != nil {
		return e
	}
	return resp.Unmarshal(into)
}

func (c *call) Err() Error {
	return c.Response().Err()
}
