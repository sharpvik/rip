package rip

type call struct {
	*Client
	function string
	argument interface{}
}

func (c *call) Arg(arg interface{}) *call {
	c.argument = arg
	return c
}

func (c *call) Response() *Response {
	req, err := NewRequest(c.function, c.argument)
	if err != nil {
		return ResponseError(err)
	}
	if err = c.Send(req); err != nil {
		return ResponseError(err)
	}
	return ReadResponse(c.conn)
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Unmarshal to decode response.
func (c *call) Return(v interface{}) Error {
	resp := c.Response()
	if err := resp.Err(); err != nil {
		return err
	}
	return resp.Unmarshal(v)
}

// Err ignores the return value and reports response error if present.
func (c *call) Err() Error {
	return c.Response().Err()
}
