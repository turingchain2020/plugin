package plugin

import (
	_ "github.com/turingchain2020/plugin/plugin/consensus/init" //consensus init
	_ "github.com/turingchain2020/plugin/plugin/crypto/init"    //crypto init
	_ "github.com/turingchain2020/plugin/plugin/dapp/init"      //dapp init
	_ "github.com/turingchain2020/plugin/plugin/mempool/init"   //mempool init
	_ "github.com/turingchain2020/plugin/plugin/p2p/init"       //p2p init
	_ "github.com/turingchain2020/plugin/plugin/store/init"     //store init
)
