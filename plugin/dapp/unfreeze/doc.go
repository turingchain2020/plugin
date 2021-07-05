// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package unfreeze 提供了定期解冻合约的实现
// 功能描述：定期解冻合约帮助用户锁定一定量的币， 按在指定的规制解冻给受益人，
// 适用于分期付款， 分期支付形式的员工激励等情景。
//
// 合约提供了3类操作
//  1. 创建定期解冻合约：创建时需要指定支付的资产和总量，以及定期解冻的形式。
//  1. 受益人提币：受益人提走解冻了的资产。
//  1. 发起人终止合约： 发起人可以终止合约的履行。
//
// 解冻的形式目前支持两种
//  1. 固定数额解冻：指定时间间隔，解冻固定的资产。
//  1. 按剩余量的固定比例解冻：指定时间间隔，按剩余量的固定比例解冻。 这种方式，越到后面解冻的越少。
// 说明：在合约创建时， 就可以解冻一次。
// 举例1， 一个固定数额解冻和合约， 总量为100, 一个月解冻10. 创建时可以由受益人提走10, 第一个月后又可以提走10.
//       在受益人没有及时提币的情况下， 受益人在一段时间之后可以一次性提走本该解冻的所有的币。 即解冻的币是按指定
//       形式解冻的，和受益人的提币时间和次数等都不会影响解冻的进程。
// 举例2， 一个按剩余量的固定比例解冻的合约， 总量为100, 一个月解冻剩余的10%. 创建时可以由受益人提走10 （100× 10%）, 第一个月后又可以提走9 （90 × 10%）.
//       在受益人没有及时提币的情况下， 受益人在一段时间之后可以一次性提走本该解冻的所有的币。 即解冻的币是按指定
//       形式解冻的，和受益人的提币时间和次数等都不会影响解冻的进程。

package unfreeze