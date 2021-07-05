// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
区块链卡牌游戏：斗牛

一、玩法简介：
游戏可以由2-5人进行，总共52张牌（除大小王），系统将随机发给玩家每人5张牌，并根据5张牌进行排列组合，进行大小比较确定胜负。
支持两种游戏玩法，普通玩法和庄家玩法
普通玩法：没有庄家，所有玩家直接比大小，最大的玩家赢得所有筹码
庄家玩法：玩家只和庄家比大小，比庄家大则赢得庄家筹码，反之则输给庄家筹码，庄家有创建者开始按加入游戏的顺序轮换

二、发牌和洗牌
1、洗牌由游戏创建者发起交易的区块时间blocktime作为随机数因子，所以排序链上可见
2、发牌使用各玩家加入游戏所在交易的hash作为随机数因子，由于nonce的存在，txhash本身具有随机性，且每个玩家每一局的txhash都不一样，可以保证发牌具有随机性

三、制胜策略
将玩家a的5张牌分为两组（3+2）后，与玩家b进行大小比较。
1、第一组3张牌的比较规则：要求将5张牌中取出任意3张组成10、20、30的整数（加法运算）。数字A-10的扑克牌数字代表其大小，JQK统一以10计算。
若玩家a和b有那么三张牌能凑成10或20或30的整数，我们称之为有牛，那么则进行第2组两张牌的大小比较。若玩家a或b有某人无法使用3张牌凑成10或20
或30的整数，我们称之为没牛，同时该玩家判定为输。

2、第二组牌的比较则把剩下的两张牌按照加法计算，10的整数倍数最大，1最小，若大于10小于20则取个位数计算。数字越大则牌型越大，数字越小则牌型
越小。若第2组牌数字为1我们称之为牛一，若第2组数字为10或20我们称之为牛牛，其他以牛二、牛三等名称称呼。牌型从小到大排序为：没牛-牛一-牛二……牛八-牛九-牛牛。

3、若玩家a和b都无法使用3张牌凑成10或20或30的整数，即两家均无牛，则此时进行5张牌中最大一张牌的比较，大小次序为K-Q-J-10-9……A，若最大一
张牌也相同则根据花色进行比较，大小次序为黑桃、红桃、梅花、方片。

4、牌型翻倍
没牛、牛1-6： 1倍
牛7-9：      2倍
牛牛：       3倍
四花：       4倍
五花：       5倍

四、游戏过程和状态
1、充值，往pokerbull合约对应的合约账户充值
2、玩家指定游戏人数和筹码加入游戏，如果有等待状态的游戏，直接加入，如果没有，创建新游戏
3、等待指定游戏人数的玩家加入，等待期间，可发送quit交易退出
4、最后一名玩家加入，游戏开始、发牌，计算结果
5、根据结果，划转筹码
6、合约账户提现到真实账户
7、状态：start->continue->quit

*/
