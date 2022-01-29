// MIT License

// Copyright (c) 2019 gonethopper

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// * @Author: ankye
// * @Date: 2019-06-24 11:07:19
// * @Last Modified by:   ankye
// * @Last Modified time: 2019-06-24 11:07:19

package quic

import (
	"time"

	"github.com/airkits/nethopper/examples/simple_client/protocol"
	"github.com/airkits/nethopper/mediator"
	"github.com/airkits/nethopper/network"
	"github.com/airkits/nethopper/network/quic"
	"github.com/airkits/nethopper/server"
)

// ModuleCreate  module create function
func ModuleCreate() (mediator.IModule, error) {
	return &Module{}, nil
}

// Module struct to define module
type Module struct {
	server.BaseContext
	quicClient *quic.Client
}

//Handlers set moudle handlers
func (s *Module) Handlers() map[string]interface{} {
	return map[string]interface{}{}
}

//ReflectHandlers set moudle reflect handlers
func (s *Module) ReflectHandlers() map[string]interface{} {
	return map[string]interface{}{
		protocol.QUICLogin: Login,
	}
}

// Setup init custom module and pass config map to module
// config
// m := map[string]interface{}{
//  "queueSize":1000,
// }
func (s *Module) Setup(conf server.IConfig) (mediator.IModule, error) {
	if err := s.ReadConfig(conf); err != nil {
		panic(err)
	}

	s.CreateWorkerPool(s, 128, 10*time.Second, true)

	s.quicClient = quic.NewClient(conf, func(conn network.IConn, uid uint64, token string) network.IAgent {
		a := network.NewAgent(NewAgentAdapter(conn), uid, token)
		network.GetInstance().AddAgent(a)
		return a
	}, func(agent network.IAgent) {
		network.GetInstance().RemoveAgent(agent)
	})
	s.quicClient.Run()

	return s, nil
}

// ReadConfig config map
// address default :80
func (s *Module) ReadConfig(conf server.IConfig) error {
	return nil
}

//Reload reload config
func (s *Module) Reload(conf server.IConfig) error {
	return nil
}

// OnRun goruntine run and call OnRun , always use ModuleRun to call this function
func (s *Module) OnRun(dt time.Duration) {
	server.RunSimpleFrame(s)
}

// Stop goruntine
func (s *Module) Stop() error {

	return nil
}

// // Call async send message to module
// func (s *Module) Call(option int32, obj *server.CallObject) error {
// 	return nil
// }

// PushBytes async send string or bytes to queue
func (s *Module) PushBytes(option int32, buf []byte) error {
	return nil
}
