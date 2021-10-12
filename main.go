package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
一次性密码本算法：
思路：
	1、然后定义62位字母表
	2、用户输入明文
	3、获取明文长度、字母表长度
	4、把长度值传入随机数函数，使其生成和明文长度一致的随机数密钥
	5、把密钥和明文转为uint[]类型，便于异或运算
	6、定义异或运算函数，
	6、编码函数：让明文和密钥key进行异或运算,返回加密后的密文
	7、解码函数：让密文和密钥key进行异或运算，返回解密后的明文
	8、输出明文 + 密文

（GenKey、EnCrypt、DeCrypt、ByteToUint、main、）
*/
var divAlphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
func main() {
	fmt.Println("输入你要加密的明文数据")
	//明文数据
	var str []byte
	//随机数
	var randunm []byte
	fmt.Scanln(&str)
	fmt.Println(str)
	fmt.Println(string(str))
	//明文长度
	strlen :=len(str)
	//字母表长度
	apslen := len(divAlphabets)

	//随机数函数，返回密钥
	key := GenKey(strlen, apslen, divAlphabets, randunm)//返回随机数密钥的10进制
	//fmt.Println(key)

	//byte类型转为uint类型
	ustr:= ByteToUint(str)
	ukey := ByteToUint(key)

	/*编码解码都需要使用uint类型*/

	//编码函数，返回密文
	yihuoEncode := EnCrypt(ustr, ukey)
	//解码函数，返回明文
	yihuoDecode := DeCrypt(yihuoEncode, ukey)

	//打印明文 + 密文
	fmt.Printf("一次性密码本加密后的密文： %X\n",yihuoEncode)
	fmt.Printf("一次性密码本解密后的明文： %q\n",yihuoDecode)



/*	fmt.Println(reflect.TypeOf(yihuoDecode)) //打印数据类型*/
}

//随机数函数
func GenKey(strlen,apslen int,apsbets string,randnum []byte)(r []byte ){//返回随机数
		rand.Seed(time.Now().Unix())
		//生成随机数下标
		for j := 1;j<=strlen;j++{
			//0~62字母表长度的随机数
			i2 := rand.Intn(apslen)
			fmt.Println("随机整数：",i2)
			randnum = append(randnum,apsbets[i2])
		}
		fmt.Println("字母表中对应的值：",string(randnum))
		fmt.Println(randnum)
		return randnum
}
//编码函数
func EnCrypt(str,key []uint) []uint{
	//fmt.Println("uint类型的str：" ,str)
	//fmt.Println("uint类型的key：" ,key)

	//明文和key异或运算
	yihuo := Yihuo(str, key)
	fmt.Println("密文：",yihuo)
	//返回密文
	return yihuo
}
//解码函数
func DeCrypt(str ,key []uint) []uint{
	//密文和key进行异或运算
	yihuo := Yihuo(str, key)
	fmt.Println("明文：",yihuo)
	//返回明文
	return yihuo
}

//异或运算
func Yihuo(str,key []uint) []uint{
	//把明文的每一位和key进行异或运算
	var yihuo []uint
	for i :=0;i<len(str);i++ {
		yihuo =append(yihuo, str[i] ^ key[i])
		//fmt.Println(yihuo)
	}
	return yihuo//加密后密码本中对应的值
}

//byte[]转uint[]
func ByteToUint(str []byte)[]uint{
	var  uslice []uint
	for i := 0;i < len(str);i++{
		//转为uint类型，放进新切片
		var suu = uint(str[i])
		uslice = append(uslice,suu)
	}
	fmt.Println(string(str)," 转为uint类型：" ,uslice)
	return uslice
}


//转为2进制函数
func ToBinary(str,key []byte) (b,b2 string){
	bstr := fmt.Sprintf("%b",str)
	bkey := fmt.Sprintf("%b",key)
	fmt.Println(bstr)
	fmt.Println(bkey)
	//返回2进制的明文和key
	return bstr,bkey
}