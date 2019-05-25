package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.3ag.xyz/backend/common/fail"
	"gitlab.3ag.xyz/backend/common/mq"
	"time"
)

func main() {



	// 1. 最一開始需要先建立 mq adapter
	//
	mqAdp := mq.GenerateConnect("amqp://guest:guest@localhost:5672")
	defer mqAdp.Close()

	// ----------------------------------------------------------------------
	// |                          Basic Part                                |
	// ----------------------------------------------------------------------

	// 2. mq adapter 可以用 GetChannel 取得 Channel Adapter
	chAdp := mqAdp.GetChannel()
	defer chAdp.Close()

	// 3. Channel Adapter 有和原本官方 library 相同的 Queue Declare, 參數完全相同，不同的是回傳為 Queue Adapter
	_, err := chAdp.QueueDeclare("demo-q-1", true, false, false, false, nil)
	fail.FailOnError(err, "Failed to declear queue demo-q-1")

	// 4. Channel Adapter 有定義一個 Consume ，這是比較方使的 Consume 方式
	//    原本的 Consume 要用 Channel Consume 再寫 Queue 名稱，這個 Queue 的 consume 可以少寫一個 Queue 名稱，也比較清楚
	//    回傳和原本 Library 一樣是 chan amqp.Delivery
	//msg1 := q1.Consume("", false, false, false, false, nil)
	msg1, err0 := chAdp.Consume("demo-q-1", "", false, false, false, false, nil)
	fail.FailOnError(err0, "failed consume")

	// 5. 用 Channel Adapter Publish，參數和官方的完全相同
	err = chAdp.Publish("", "demo-q-1", false, false,
		// 這邊一樣是用原本 library 的  Publishing struct
		amqp.Publishing{
			Type: "string",
			Body: []byte("test"),
		})

	fail.FailOnError(err, "publish faled")

	// 定義 timeout，用 select block 取 amqp.Delivery，能確保沒收到東西一直卡在這邊
	timer := time.NewTimer(100 * time.Millisecond)
	select {
	case d := <-msg1:
		fmt.Println("[Log] get msg successful: ", string(d.Body))
		d.Ack(false) // Consume 沒有設定 autoDelete 的話就要手動 Ack
	case <-timer.C:
		fmt.Println("[Error] time out")
	}

	// ----------------------------------------------------------------------
	// |                          Exchange Part                             |
	// ----------------------------------------------------------------------
	// 1. 前面己有 chAdp，這邊再 Declare 兩個 channel
	_, err = chAdp.QueueDeclare("queue-2", true, false, false, false, nil)
	fail.FailOnError(err, "declare queue-2 failed")
	q3, err := chAdp.QueueDeclare("queue-3", true, false, false, false, nil)
	fail.FailOnError(err, "declare queue-3 failed")

	// 2. ExchangeDeclare 和原本的 library 相同，direct 為 route by key 的設定 (無萬用字元）
	chAdp.ExchangeDeclare("exg-demo", "direct", false, false, false, false, nil)

	// 3. Bind Queue 的方式和原本 library 相同
	chAdp.QueueBindEasy("queue-2", "route-q-2", "exg-demo")
	chAdp.QueueBindEasy("queue-3", "route-q-3", "exg-demo")


	// 這邊一樣可以用原本的 Channel consume 方式，不過 err 也會回傳出來
	msg2, err := chAdp.Consume("queue-2", "", false, false, false, false, nil)
	// msg2 := q2.Consume("", false, false, false, false, nil) // <- 用 queue consume 的方式
	fail.FailOnError(err, "Consume queue-2 failed")
	msg3 := q3.Consume("", false, false, false, false, nil)

	_ = chAdp.Publish("exg-demo", "route-q-3", false, false,
		amqp.Publishing{
			Type: "text",
			Body: []byte("推給 queue-3"),
		})

	// timer 可以 reset
	timer.Reset(1 * time.Second)
	select {
	case d := <-msg2:
		fmt.Println("因為 publish 到 q3 所以這行不可能印出", d.Body)
	case d := <- msg3:
		fmt.Print("收到資料 = ", string(d.Body))
	case <-timer.C:
		fmt.Println("[Error] Timeout")
	}

}


