package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WitchCancel()
	//WithTimeout()
	//WithValue()
	WithDeadline()
}

func WitchCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	cancel()
	go BuyFlour(ctx)
	go BuyEgg(ctx)
	fmt.Println("停电了，不做蛋挞了")
	time.Sleep(2 * time.Second)
}

//买面
func BuyFlour(ctx context.Context) {
	fmt.Println("去买面")
	//time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买面了")
		return
	default:
	}

	fmt.Println("买好面了")
}

//买蛋
func BuyEgg(ctx context.Context) {
	fmt.Println("去买蛋")
	//time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:
	}

	fmt.Println("买好蛋了")

	go BuyEgg(ctx)
	time.Sleep(1 * time.Second)
}

func WithTimeout() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	fmt.Println("开始部署")
	go Distribute(ctx)
	time.Sleep(10 * time.Second)
}

//部署
func Distribute(ctx context.Context) {
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("部署取消")
		return
	default:
	}
	fmt.Println("部署成功")
}

func WithValue() {
	ctx := context.WithValue(context.Background(), "1", "钱包")
	GotoPapa(ctx)
}
func GotoPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	GotoMama(ctx)
}
func GotoMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	GotoGrandma(ctx)
}
func GotoGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	GotoParty(ctx)
}
func GotoParty(ctx context.Context) {
	fmt.Println("1:", ctx.Value("1"))
	fmt.Println("2:", ctx.Value("2"))
	fmt.Println("3:", ctx.Value("3"))
	fmt.Println("4:", ctx.Value("4"))
}

func WithDeadline() {
	now := time.Now()
	t := now.Add(1 * time.Second)
	ctx, _ := context.WithDeadline(context.Background(), t)
	go WatchTv(ctx)
	go WatchPhone(ctx)
	time.Sleep(2 * time.Second)
}
func WatchTv(ctx context.Context) {
	fmt.Println("看电视")
	select {
	case <-ctx.Done():
		fmt.Println("关电视")
	}
}
func WatchPhone(ctx context.Context) {
	fmt.Println("看手机")
	select {
	case <-ctx.Done():
		fmt.Println("关手机")
	}
}
