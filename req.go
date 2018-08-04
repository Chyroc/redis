package redis

import (
	"bytes"
	"fmt"
	"strconv"
)

func (r Redis) write(p []byte) error {
	_, err := r.conn.Write(p)
	return err
}

func (r Redis) cmd(args ...string) error {
	if len(args) == 0 {
		return ErrEmptyCommand
	}

	buf := new(bytes.Buffer)

	buf.WriteString("*")
	buf.WriteString(strconv.Itoa(len(args)))
	buf.Write(crlf)

	for _, arg := range args {
		p := []byte(arg)

		buf.WriteString("$")
		buf.WriteString(strconv.Itoa(len(p)))
		buf.Write(crlf)

		buf.Write(p)
		buf.Write(crlf)
	}

	return r.write(buf.Bytes())
}

func (r Redis) run(args ...string) *Reply {
	if err := r.cmd(args...); err != nil {
		fmt.Printf("send %#v [error:%s]\n", args, err)
		return errToReply(err)
	}

	reply, err := r.read()
	if err != nil {
		fmt.Printf("send %#v [error:%s]\n", args, err)
		return errToReply(err)
	}

	fmt.Printf("send %#v | got %v\n", args, reply)
	return reply
}
